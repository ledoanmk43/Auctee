package repository

import (
	"backend/pkg/utils"
	auction "backend/src/auction-service/config"
	"backend/src/auction-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type IAuctionRepository interface {
	CreateAuction(auction *entity.Auction) error
	UpdateAuction(product *entity.Auction) error
	DeleteAuction(auctionId, userIdId uint) error
	GetAuctionById(auctionId uint) (*entity.Auction, error)
	GetActiveAuctionById(auctionId uint) (*entity.Auction, error)
	GetAllAuctions(page int) (*[]entity.Auction, error)
	GetAllAuctionsByUserId(userId uint) (*[]entity.Auction, error)
	UpdateCurrentBidByAuctionId(bid *entity.Bid, maxPrice float32) error
	CheckIfUserIsWinner(userId, auctionId uint) bool
	GetAllAuctionsByProductName(nameList []string) (*[]entity.Auction, error)
}

type AuctionRepositoryDefault struct {
	connection *gorm.DB
}

func NewAuctionRepositoryDefault(dbConn *gorm.DB) *AuctionRepositoryDefault {
	return &AuctionRepositoryDefault{connection: dbConn}
}

func (a *AuctionRepositoryDefault) GetAllAuctionsByProductName(nameList []string) (*[]entity.Auction, error) {
	var auctions *[]entity.Auction
	var count int64
	_ = a.connection.Where("product_name IN ? AND end_time >= ?", nameList, time.Now()).Find(&auctions).Count(&count)
	if count == 0 {
		return nil, errors.New("no auction found")
	}

	return auctions, nil
}

func (a *AuctionRepositoryDefault) GetAllAuctionsByUserId(userId uint) (*[]entity.Auction, error) {
	var auctions *[]entity.Auction
	var count int64
	record := a.connection.Where("user_id = ?", userId).Order("created_at desc").Find(&auctions).Count(&count)
	if record.Error != nil {
		log.Println("Get auctions: Error get all auctions in repo", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, errors.New("product not found")
	}

	return auctions, nil
}

func (a *AuctionRepositoryDefault) GetAllAuctions(page int) (*[]entity.Auction, error) {
	var auctions *[]entity.Auction
	//Maybe lazy load will require about 20 auctions at a time
	//Or search about lazy load API
	record := a.connection.Limit(auction.PerPage).Offset((page-1)*auction.PerPage).Order("end_time asc").Where("end_time >= ? and is_active = ?", time.Now(), true).Find(&auctions)
	if record.Error != nil {
		log.Println("Get auctions: Error get all auctions in repo", record.Error)
		return nil, record.Error
	}

	return auctions, nil
}

func (a *AuctionRepositoryDefault) CreateAuction(auction *entity.Auction) error {
	record := a.connection.Create(&auction)
	if record.Error != nil {
		log.Println("Error to create auction in repo: ", record.Error)
		return record.Error
	}
	return nil
}

func (a *AuctionRepositoryDefault) UpdateAuction(updateBody *entity.Auction) error {
	var auctionToUpdate *entity.Auction
	var count int64
	record := a.connection.Where("id = ? AND user_id = ?", updateBody.ID, updateBody.UserId).Find(&auctionToUpdate).Count(&count)

	if record.Error != nil {
		log.Println("Error to find auction in repo", record.Error)
		return record.Error
	}
	if count == 0 {
		return errors.New("auction not found")
	}

	auctionToUpdate.StartTime = updateBody.StartTime
	auctionToUpdate.EndTime = updateBody.EndTime
	auctionToUpdate.CurrentBid = updateBody.CurrentBid
	auctionToUpdate.IsActive = utils.BoolAddr(*updateBody.IsActive)
	auctionToUpdate.Quantity = updateBody.Quantity
	auctionToUpdate.PricePerStep = updateBody.PricePerStep
	auctionToUpdate.ImagePath = updateBody.ImagePath
	recordSave := a.connection.Updates(&auctionToUpdate)
	if recordSave.Error != nil {
		log.Println("Error to update auction repo", recordSave.Error)
		return recordSave.Error
	}
	return nil
}

func (a *AuctionRepositoryDefault) DeleteAuction(auctionId, userId uint) error {
	var auction *entity.Auction
	var count int64
	result := a.connection.Where("id = ? AND user_id = ?", auctionId, userId).Find(&auction).Count(&count)
	if result.Error != nil {
		log.Println("Delete auction: Error in find auction to delete in package repository", result.Error)
		return result.Error
	}
	if count == 0 {
		return errors.New("auction not found")
	}
	a.connection.Delete(&auction)
	return nil
}

func (a *AuctionRepositoryDefault) GetAuctionById(auctionId uint) (*entity.Auction, error) {
	var auction *entity.Auction
	var count int64
	record := a.connection.Where("id = ?", auctionId).Find(&auction).Count(&count)
	if record.Error != nil {
		log.Println("Error to find auction in repo")
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetAuctionById: auction not found", count)
		return nil, errors.New("error: auction not found")
	}
	return auction, nil
}

func (a *AuctionRepositoryDefault) GetActiveAuctionById(auctionId uint) (*entity.Auction, error) {
	var auction *entity.Auction
	var count int64
	record := a.connection.Where("id = ? and is_active = ?", auctionId, true).Find(&auction).Count(&count)
	if record.Error != nil {
		log.Println("Error to find auction in repo")
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetAuctionById: auction not found", count)
		return nil, errors.New("error: auction not found")
	}
	return auction, nil
}

func (a *AuctionRepositoryDefault) UpdateCurrentBidByAuctionId(newBid *entity.Bid, maxPrice float32) error {
	var auction *entity.Auction
	record := a.connection.Where("id = ?", newBid.AuctionId).Find(&auction)

	if record.Error != nil {
		log.Println(record.Error)
		return errors.New("error to find auction when update current winner in repo")
	}

	if newBid.BidValue >= float64(maxPrice) {
		auction.IsActive = utils.BoolAddr(false)
		auction.EndTime = time.Now().Format("2006-01-02 15:04:05")
	}
	auction.StartTime = strings.ReplaceAll(auction.StartTime, "Z", "")
	auction.StartTime = strings.ReplaceAll(auction.StartTime, "T", " ")
	auction.EndTime = strings.ReplaceAll(auction.EndTime, "Z", "")
	auction.EndTime = strings.ReplaceAll(auction.EndTime, "T", " ")
	auction.CurrentBid = newBid.BidValue
	auction.WinnerId = newBid.UserId
	recordSave := a.connection.Updates(&auction)
	if recordSave.Error != nil {
		log.Println("Error to update auction repo", recordSave.Error)
		return recordSave.Error
	}
	return record.Error
}

func (a *AuctionRepositoryDefault) CheckIfUserIsWinner(userId, auctionId uint) bool {
	var auction *entity.Auction
	var count int64
	res := a.connection.Where("winner_id = ? AND id = ?", userId, auctionId).Find(&auction).Count(&count)
	if res.Error != nil {
		log.Println("Error: ", res.Error)
		return false
	}
	if count == 0 {
		return false
	}
	return true
}
