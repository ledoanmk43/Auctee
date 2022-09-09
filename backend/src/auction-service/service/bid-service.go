package service

import (
	"chilindo/src/auction-service/entity"
	"chilindo/src/auction-service/repository"
	"errors"
	"log"
)

type IBidService interface {
	CreateBid(bid *entity.Bid) (*entity.Bid, error)
}

type BidServiceDefault struct {
	BidRepository     repository.IBidRepository
	AuctionRepository repository.IAuctionRepository
}

func NewBidServiceDefault(bidRepository repository.IBidRepository, auctionRepository repository.IAuctionRepository) *BidServiceDefault {
	return &BidServiceDefault{BidRepository: bidRepository, AuctionRepository: auctionRepository}
}

func (b *BidServiceDefault) CreateBid(newBid *entity.Bid) (*entity.Bid, error) {
	auction, errGetAuction := b.AuctionRepository.GetAuctionById(newBid.AuctionId)
	if errGetAuction != nil {
		log.Println("GetAuction: Error Get auction in bid package service", errGetAuction)
		return nil, errGetAuction
	}

	//also check if that user has already bid

	//change this func into CheckIfBidder is winner
	if b.AuctionRepository.CheckIfUserIsWinner(newBid.UserId, auction.ID) {
		log.Println("CreateBid: Error Create Bid in package service: user is winner")
		return nil, errors.New("you are currently the winner with highest bid")
	}

	//if bidValue is smaller than currentBid
	if newBid.BidValue <= auction.CurrentBid {
		return nil, errors.New("new bid must be greater than current bid")
	}

	//after all create Bid
	createdBid, err := b.BidRepository.CreateBid(newBid)
	if err != nil {
		log.Println("CreateBid: Error Create Bid in package service: ", err)
		return nil, err
	}
	return createdBid, nil
}
