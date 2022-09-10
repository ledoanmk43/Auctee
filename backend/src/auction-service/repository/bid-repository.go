package repository

import (
	"backend/src/auction-service/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

type IBidRepository interface {
	CreateBid(bid *entity.Bid) (*entity.Bid, error)
}

type BidRepositoryDefault struct {
	connection  *gorm.DB
	AuctionRepo IAuctionRepository
}

func NewBidRepositoryDefault(connection *gorm.DB, auctionRepo IAuctionRepository) *BidRepositoryDefault {
	return &BidRepositoryDefault{connection: connection, AuctionRepo: auctionRepo}
}

func (b *BidRepositoryDefault) CreateBid(bid *entity.Bid) (*entity.Bid, error) {
	bid.BidTime = time.Now()
	record := b.connection.Create(&bid)
	if record.Error != nil {
		log.Println("Error to create bid in repo")
		return nil, record.Error
	}

	//auction, err := b.AuctionRepo.GetAuctionById(bid.AuctionId)
	//if err != nil {
	//	return nil, err
	//}
	//log.Println("current bid: ", auction.CurrentBid)
	//log.Println("new bid value: ", bid.BidValue)

	err := b.AuctionRepo.UpdateCurrentBidByAuctionId(bid)
	if err != nil {
		return nil, err
	}

	return bid, nil
}

//func (b *BidRepositoryDefault) CheckIfUserIsWinner(userId, auctionId uint) bool {
//	var newBid *entity.Bid
//	res := b.connection.Where("user_id = ? AND auction_id = ?", userId, auctionId).First(&newBid)
//
//	if res.Error != nil {
//		log.Println("Error: ", res.Error)
//		return false
//	}
//	return true
//}
