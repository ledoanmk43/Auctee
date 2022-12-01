package repository

import (
	"backend/src/auction-service/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

type IBidRepository interface {
	CreateBid(bid *entity.Bid, maxPrice float32) error
	GetAllBidsByAuctionId(auctionId uint) (*[]entity.Bid, error)
	GetAllBidsByUserId(userId uint) (*[]entity.Bid, error)
}

type BidRepositoryDefault struct {
	connection  *gorm.DB
	AuctionRepo IAuctionRepository
}

func NewBidRepositoryDefault(connection *gorm.DB, auctionRepo IAuctionRepository) *BidRepositoryDefault {
	return &BidRepositoryDefault{connection: connection, AuctionRepo: auctionRepo}
}

func (b *BidRepositoryDefault) GetAllBidsByAuctionId(auctionId uint) (*[]entity.Bid, error) {
	var bids *[]entity.Bid
	record := b.connection.Where("auction_id = ? ", auctionId).Order("bid_value desc").Find(&bids)
	if record.Error != nil {
		log.Println("Get auctions: Error get all bid in repo", record.Error)
		return nil, record.Error
	}

	return bids, nil
}

func (b *BidRepositoryDefault) GetAllBidsByUserId(userId uint) (*[]entity.Bid, error) {
	var bids *[]entity.Bid
	record := b.connection.Where("user_id = ? ", userId).Distinct("auction_id").Find(&bids)
	if record.Error != nil {
		log.Println("Get auctions: Error get all bid in repo", record.Error)
		return nil, record.Error
	}
	return bids, nil
}

func (b *BidRepositoryDefault) CreateBid(bid *entity.Bid, maxPrice float32) error {
	bid.BidTime = time.Now()
	record := b.connection.Create(&bid)
	if record.Error != nil {
		log.Println("Error to create bid in repo")
		return record.Error
	}

	err := b.AuctionRepo.UpdateCurrentBidByAuctionId(bid, maxPrice)
	if err != nil {
		return err
	}

	return nil
}
