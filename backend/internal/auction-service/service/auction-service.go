package service

import (
	"backend/internal/auction-service/entity"
	"backend/internal/auction-service/repository"
	"backend/pkg/utils"
	"errors"
	"log"
)

type IAuctionService interface {
	CreateAuction(auction *entity.Auction) error
	Update(b *entity.Auction) error
	Delete(auctionId, userId uint) error
	GetAuctionById(auctionId uint) (*entity.Auction, error)
	GetAllAuctions(page int) (*[]entity.Auction, error)
	GetAllAuctionsByUserId(userId uint) (*[]entity.Auction, error)
	GetAllAuctionsByProductName(nameList []string) (*[]entity.Auction, error)
}

type AuctionServiceDefault struct {
	AuctionRepository repository.IAuctionRepository
}

func NewAuctionServiceDefault(auctionRepository repository.IAuctionRepository) *AuctionServiceDefault {
	return &AuctionServiceDefault{AuctionRepository: auctionRepository}
}

func (a *AuctionServiceDefault) GetAllAuctionsByProductName(nameList []string) (*[]entity.Auction, error) {
	auctions, err := a.AuctionRepository.GetAllAuctionsByProductName(nameList)
	if err != nil {
		log.Println("Get auctions : Error get auctions in package service", err)
		return nil, err
	}
	return auctions, nil
}

func (a *AuctionServiceDefault) GetAllAuctionsByUserId(userId uint) (*[]entity.Auction, error) {
	auctions, err := a.AuctionRepository.GetAllAuctionsByUserId(userId)
	if err != nil {
		log.Println("Get auctions : Error get auctions in package service", err)
	}
	return auctions, nil
}

func (a *AuctionServiceDefault) GetAllAuctions(page int) (*[]entity.Auction, error) {
	auctions, err := a.AuctionRepository.GetAllAuctions(page)
	if err != nil {
		log.Println("Get auctions : Error get auctions in package service", err)
	}
	return auctions, nil
}

func (a *AuctionServiceDefault) CreateAuction(auction *entity.Auction) error {
	startTime, err := utils.StringToTime(auction.StartTime)
	if err != nil {
		log.Println("CreateAuction: Error Create Auction in package service", err)
		return err
	}
	endTime, err := utils.StringToTime(auction.EndTime)
	if err != nil {
		log.Println("CreateAuction: Error Create Auction in package service", err)
		return err
	}
	now, err := utils.GetMoment()
	if err != nil {
		log.Println("CreateAuction: Error Create Auction in package service", err)
		return err
	}

	//check if start/end time is < present
	if startTime.Before(now) || endTime.Before(now) || startTime == endTime || endTime.Before(startTime) {
		return errors.New("invalid start time or end time")
	}

	errCreate := a.AuctionRepository.CreateAuction(auction)
	if err != nil {
		log.Println("CreateAuction: Error Create Auction in package service", errCreate)
		return errCreate
	}
	return nil
}

func (a *AuctionServiceDefault) Update(auction *entity.Auction) error {
	startTime, err := utils.StringToTime(auction.StartTime)
	if err != nil && len(auction.StartTime) != 0 {
		return err
	}
	endTime, err := utils.StringToTime(auction.EndTime)
	if err != nil && len(auction.EndTime) != 0 {
		return err
	}
	now, err := utils.GetMoment()
	if err != nil {
		return err
	}

	//check if start/end time is < present
	if len(auction.StartTime) != 0 {
		if startTime.Before(now) {
			return errors.New("invalid start time")
		}
		if startTime == endTime {
			return errors.New("equal start time and end time")
		}
		if endTime.Before(startTime) {
			return errors.New("invalid start time and end time")
		}
	}
	if len(auction.EndTime) != 0 {
		if endTime.Before(now) {
			return errors.New("invalid end time")
		}
		if startTime == endTime {
			return errors.New("equal start time and end time")
		}
		if endTime.Before(startTime) {
			return errors.New("invalid start time and end time")
		}
	}

	err = a.AuctionRepository.UpdateAuction(auction)
	if err != nil {
		log.Println("Error in package service", err)
		return err
	}
	return nil
}
func (a *AuctionServiceDefault) Delete(auctionId, userId uint) error {
	err := a.AuctionRepository.DeleteAuction(auctionId, userId)
	if err != nil {
		log.Println("CreateAuction: Error Create Auction in package service", err)
		return err
	}
	return nil
}
func (a *AuctionServiceDefault) GetAuctionById(auctionId uint) (*entity.Auction, error) {
	auctionDetail, err := a.AuctionRepository.GetAuctionById(auctionId)
	if err != nil {
		log.Println("GetAuctionById: Error in get auction by Id", err)
		return nil, err
	}
	return auctionDetail, nil
}
