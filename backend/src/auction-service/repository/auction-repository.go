package repository

import (
	"chilindo/src/auction-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IAuctionRepository interface {
	CreateAuction(auction *entity.Auction) (*entity.Auction, error)
	GetAuctionById(auctionId uint) (*entity.Auction, error)
	UpdateCurrentBidByAuctionId(bid *entity.Bid) error
	CheckIfUserIsWinner(userId, auctionId uint) bool
}

type AuctionRepositoryDefault struct {
	connection *gorm.DB
}

func NewAuctionRepositoryDefault(connection *gorm.DB) *AuctionRepositoryDefault {
	return &AuctionRepositoryDefault{connection: connection}
}

func (a *AuctionRepositoryDefault) CreateAuction(auction *entity.Auction) (*entity.Auction, error) {

	record := a.connection.Create(&auction)
	if record.Error != nil {
		log.Println("Error to create auction in repo")
		return nil, record.Error
	}
	return auction, nil
}

func (a *AuctionRepositoryDefault) GetAuctionById(auctionId uint) (*entity.Auction, error) {
	var auction *entity.Auction
	record := a.connection.Where("id = ?", auctionId).Find(&auction)
	if record.Error != nil {
		log.Println("Error to find auction in repo")
		return nil, record.Error
	}
	return auction, nil
}

func (a *AuctionRepositoryDefault) UpdateCurrentBidByAuctionId(newBid *entity.Bid) error {
	var auction *entity.Auction
	record := a.connection.Where("id = ?", newBid.AuctionId).Find(&auction)

	if record.Error != nil {
		log.Println(record.Error)
		return errors.New("error to find auction when update current winner in repo")
	}
	auction.CurrentBid = newBid.BidValue
	auction.WinnerId = newBid.UserId
	a.connection.Save(&auction)
	return record.Error
}

func (a *AuctionRepositoryDefault) CheckIfUserIsWinner(userId, auctionId uint) bool {
	var newBid *entity.Bid
	res := a.connection.Where("winner_id = ? AND id = ?", userId, auctionId).First(&newBid)

	if res.Error != nil {
		log.Println("Error: ", res.Error)
		return false
	}
	return true
}
