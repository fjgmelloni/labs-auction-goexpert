package auction_test

import (
	"context"
	"os"
	"testing"
	"time"

	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionAutoClose(t *testing.T) {

	if os.Getenv("MONGODB_URL") == "" {
		_ = os.Setenv("MONGODB_URL", "mongodb://admin:admin@localhost:27017/auctions?authSource=admin")
	}
	if os.Getenv("MONGODB_DB") == "" {
		_ = os.Setenv("MONGODB_DB", "auctions")
	}
	if os.Getenv("AUCTION_DURATION_SECONDS") == "" {
		_ = os.Setenv("AUCTION_DURATION_SECONDS", "2")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	assert.NoError(t, err)
	defer client.Disconnect(context.Background())

	db := client.Database(os.Getenv("MONGODB_DB"))
	repo := auction.NewAuctionRepository(db)

	newAuction, errCreate := auction_entity.CreateAuction(
		"Produto Teste", "Eletrônicos", "Leilão teste automático",
		auction_entity.New,
	)
	assert.Nil(t, errCreate)

	errCreateRepo := repo.CreateAuction(context.Background(), newAuction)
	assert.Nil(t, errCreateRepo)

	time.Sleep(3 * time.Second)

	result := db.Collection("auctions").FindOne(context.Background(), map[string]interface{}{
		"_id":    newAuction.Id,
		"status": auction_entity.Completed,
	})
	assert.NoError(t, result.Err())
}
