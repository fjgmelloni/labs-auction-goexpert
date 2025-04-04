package auction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction,
) *internal_error.InternalError {

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction: ", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	// ✅ Inicia goroutine segura para fechamento automático
	go func(auctionID string) {
		durationStr := os.Getenv("AUCTION_DURATION_SECONDS")
		durationSeconds, err := strconv.Atoi(durationStr)
		if err != nil || durationSeconds <= 0 {
			logger.Error("Invalid AUCTION_DURATION_SECONDS, using default 60s", err)
			durationSeconds = 60
		}

		logger.Info(fmt.Sprintf("Auction %s scheduled to close in %d seconds", auctionID, durationSeconds))
		time.Sleep(time.Duration(durationSeconds) * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{"_id": auctionID, "status": auction_entity.Active}
		update := bson.M{"$set": bson.M{"status": auction_entity.Completed}}

		result, err := ar.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.Error("Failed to close auction:", err)
			return
		}

		if result.ModifiedCount > 0 {
			logger.Info(fmt.Sprintf("Auction %s closed automatically.", auctionID))
		} else {
			logger.Info(fmt.Sprintf("Auction %s was not closed because it was already finalized or not found.", auctionID))
		}
	}(auctionEntity.Id)

	return nil
}
