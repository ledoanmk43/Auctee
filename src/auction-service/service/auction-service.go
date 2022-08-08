package service

import (
	"chilindo/src/auction-service/entity"
	"chilindo/src/auction-service/repository"
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

func (a AuctionServiceDefault) CreateAuction(auction *entity.Auction) (*entity.Auction, error) {
	newAuction, err := a.AuctionRepository.CreateAuction(auction)
	if err != nil {
		log.Println("CreateAuction: Error Create address in package service", err)
		return nil, err
	}
	return newAuction, nil
}
