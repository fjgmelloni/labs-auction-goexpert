# Full Cycle - Desafio de Leilão com Fechamento Automático

Este projeto é um dos desafios da pós-graduação **Full Cycle**. O objetivo é implementar a funcionalidade de **fechamento automático de leilões** após um tempo determinado, utilizando **Go**, **MongoDB**, **goroutines** e **Docker**.

---

## 📋 Descrição do Desafio

Ao criar um novo leilão, ele deverá ser **automaticamente encerrado** após um intervalo de tempo definido pela variável de ambiente `AUCTION_DURATION_SECONDS`. O processo de encerramento deve ocorrer de forma **assíncrona** usando goroutines.

O projeto já contempla:

- Criação de leilões
- Lançamento de lances
- Validação de leilão ativo
- Fechamento automático do leilão após o tempo expirar ⏱️

---

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/) 1.22.5
- [MongoDB](https://www.mongodb.com/)
- [Docker & Docker Compose](https://docs.docker.com/)
- [Testify](https://github.com/stretchr/testify) para testes

---

## 📦 Como rodar o projeto

### Pré-requisitos

- [Docker](https://www.docker.com/products/docker-desktop) instalado
- [Go](https://golang.org/dl/) instalado (caso queira rodar localmente sem container)

---

### 📁 Clone o projeto

```bash
git clone https://github.com/seu-usuario/seu-repo.git
cd seu-repo
```

---

### 🐳 Subindo com Docker Compose

```bash
docker compose up --build
```

Isso irá:

- Subir o serviço de MongoDB
- Construir e rodar o binário do Go localizado em `cmd/auction/main.go`
- Utilizar as variáveis de ambiente definidas em `cmd/auction/.env`

---

### ⚙️ Variáveis de Ambiente

O arquivo `.env` define a duração dos leilões e configurações de banco:

```env
BATCH_INSERT_INTERVAL=20s
MAX_BATCH_SIZE=4
AUCTION_INTERVAL=20s
AUCTION_DURATION_SECONDS=60

MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=admin
MONGODB_URL=mongodb://admin:admin@mongodb:27017/auctions?authSource=admin
MONGODB_DB=auctions
```

---

### 🧪 Rodando os Testes

Para rodar os testes unitários (ex: validação do fechamento automático de leilões):

```bash
# Ative as variáveis do .env
export $(cat cmd/auction/.env | xargs)

# Execute os testes
go test ./internal/infra/database/auction -v -run ^TestAuctionAutoClose$
```

Se preferir rodar dentro do container, use:

```bash
docker exec -it <nome-do-container> go test ./internal/infra/database/auction
```

---

## 📂 Estrutura Principal

```
.
├── cmd
│   └── auction          # Entrypoint do app
├── internal
│   ├── entity           # Entidades de domínio
│   ├── infra
│   │   └── database
│   │       └── auction  # Persistência e lógica do leilão
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## 📚 O que foi implementado

- ✅ Encerramento automático do leilão com goroutine
- ✅ Tempo de vida do leilão controlado por variável de ambiente
- ✅ Teste automatizado para validar fechamento automático (`TestAuctionAutoClose`)
- ✅ Suporte a Docker e MongoDB via `docker-compose`

---

## 🎓 Sobre o Desafio

Este repositório foi desenvolvido como parte da **pós-graduação Full Cycle** da [Code.education](https://fullcycle.com.br/). O objetivo é explorar **concorrência em Go**, **persistência com MongoDB** e boas práticas de arquitetura em projetos reais.

---


> 📌 Este projeto foi desenvolvido como parte do desafio da pós-graduação [Full Cycle](https://fullcycle.com.br), com base no repositório oficial: [https://github.com/devfullcycle/labs-auction-goexpert](https://github.com/devfullcycle/labs-auction-goexpert).