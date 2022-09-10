package service

import (
	"backend/src/auction-service/entity"
	"backend/src/auction-service/repository"
	"log"
)

type IAuctionService interface {
	CreateAuction(auction *entity.Auction) (*entity.Auction, error)
}

type AuctionServiceDefault struct {
	AuctionRepository repository.IAuctionRepository
}

func NewAuctionServiceDefault(auctionRepository repository.IAuctionRepository) *AuctionServiceDefault {
	return &AuctionServiceDefault{AuctionRepository: auctionRepository}
}

func (a *AuctionServiceDefault) CreateAuction(auction *entity.Auction) (*entity.Auction, error) {
	createdAuction, err := a.AuctionRepository.CreateAuction(auction)
	if err != nil {
		log.Println("CreateAuction: Error Create Auction in package service", err)
		return nil, err
	}
	return createdAuction, nil
}
