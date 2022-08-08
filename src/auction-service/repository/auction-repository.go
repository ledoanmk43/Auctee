package repository

import (
	"chilindo/src/auction-service/entity"
	"gorm.io/gorm"
	"log"
)

type IAuctionRepository interface {
	CreateAuction(auction *entity.Auction) (*entity.Auction, error)
}

type AuctionRepositoryDefault struct {
	connection *gorm.DB
}

func NewAuctionRepositoryDefault(connection *gorm.DB) *AuctionRepositoryDefault {
	return &AuctionRepositoryDefault{connection: connection}
}

func (a AuctionRepositoryDefault) CreateAuction(auction *entity.Auction) (*entity.Auction, error) {
	record := a.connection.Create(&auction)
	if record.Error != nil {
		log.Println("Error to create product repo")
		return nil, record.Error
	}
	return auction, nil
}
