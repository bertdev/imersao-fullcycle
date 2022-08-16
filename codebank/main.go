package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bertdev/imersao-fullcycle/domain"

	"github.com/bertdev/imersao-fullcycle/infrastructure/repository"
	"github.com/bertdev/imersao-fullcycle/usecase"

	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Herbert"
	cc.ExpirationMonth = 7
	cc.ExpirationYear = 2024
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 200

	repo := repository.NewTransactionRepositoryDb(db)
	repo.CreateCreditCard(*cc)
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	usecase := usecase.NewUseCaseTransaction(transactionRepository)
	return usecase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}
