package service

import (
	"backend/src/auction-service/entity"
	"backend/src/auction-service/repository"
	"errors"
	"log"
)

type IBidService interface {
	CreateBid(bid *entity.Bid, maxPrice float32) error
	GetAllBidsByAuctionId(auctionId uint) (*[]entity.Bid, error)
	GetAllBidsByUserId(userId uint) (*[]entity.Auction, error)
}

type BidServiceDefault struct {
	BidRepository     repository.IBidRepository
	AuctionRepository repository.IAuctionRepository
}

func NewBidServiceDefault(bidRepository repository.IBidRepository, auctionRepository repository.IAuctionRepository) *BidServiceDefault {
	return &BidServiceDefault{BidRepository: bidRepository, AuctionRepository: auctionRepository}
}

func (b *BidServiceDefault) GetAllBidsByAuctionId(auctionId uint) (*[]entity.Bid, error) {
	bids, err := b.BidRepository.GetAllBidsByAuctionId(auctionId)
	if err != nil {
		log.Println("Get auctions : Error get auctions in package service", err)
	}
	return bids, nil
}

func (b *BidServiceDefault) GetAllBidsByUserId(userId uint) (*[]entity.Auction, error) {
	bids, err := b.BidRepository.GetAllBidsByUserId(userId)
	if err != nil {
		log.Println("Get auctions : Error get auctions in package service", err)
	}
	var auctions []entity.Auction
	for _, value := range *bids {
		auction, _ := b.AuctionRepository.GetActiveAuctionById(value.AuctionId)
		if auction != nil {
			auctions = append(auctions, *auction)
		}
	}

	return &auctions, nil
}

func (b *BidServiceDefault) CreateBid(newBid *entity.Bid, maxPrice float32) error {
	auction, errGetAuction := b.AuctionRepository.GetAuctionById(newBid.AuctionId)
	if errGetAuction != nil {
		log.Println("GetAuction: Error Get auction in bid package service", errGetAuction)
		return errGetAuction
	}

	//change this func into CheckIfBidder is winner
	if b.AuctionRepository.CheckIfUserIsWinner(newBid.UserId, auction.Id) {
		log.Println("CreateBid: Error Create Bid in package service: user is winner")
		return errors.New("Bạn đang là người đặt tiền cao nhất")
	}

	//if bidValue is smaller than currentBid
	if newBid.BidValue <= auction.CurrentBid {
		return errors.New("Số tiền nhỏ hơn giá hiện tại của sản phẩm")
	}

	//after all create Bid
	err := b.BidRepository.CreateBid(newBid, maxPrice)
	if err != nil {
		log.Println("CreateBid: Error Create Bid in package service: ", err)
		return err
	}
	return nil
}
