package service

import (
	"backend/src/payment-service/entity"
	"backend/src/payment-service/repository"
	"log"
)

type IPaymentService interface {
	CreatePayment(payment *entity.Payment) error
	UpdateAddressPayment(payment *entity.Payment) error
	DeletePayment(paymentId, userId uint) error
	GetPaymentByPaymentId(paymentId uint) (*entity.Payment, error)
	GetAllPaymentsForWinner(page int, winnerId uint) (*[]entity.Payment, error)
	GetAllPaymentsForOwner(page int, winnerId uint) (*[]entity.Payment, error)
}

type PaymentServiceDefault struct {
	PaymentRepository repository.IPaymentRepository
}

func NewPaymentServiceDefault(paymentRepository repository.IPaymentRepository) *PaymentServiceDefault {
	return &PaymentServiceDefault{PaymentRepository: paymentRepository}
}

func (p *PaymentServiceDefault) GetAllPaymentsForWinner(page int, winnerId uint) (*[]entity.Payment, error) {
	payments, err := p.PaymentRepository.GetAllPaymentsForWinner(page, winnerId)
	if err != nil {
		log.Println("Get auctions : Error get auctions in package service", err)
		return nil, err
	}
	return payments, nil
}

func (p *PaymentServiceDefault) GetAllPaymentsForOwner(page int, ownerId uint) (*[]entity.Payment, error) {
	payments, err := p.PaymentRepository.GetAllPaymentsForOwner(page, ownerId)
	if err != nil {
		log.Println("Get auctions : Error get auctions in package service", err)
		return nil, err
	}
	return payments, nil
}

func (p *PaymentServiceDefault) CreatePayment(payment *entity.Payment) error {

	errCreate := p.PaymentRepository.CreatePayment(payment)
	if errCreate != nil {
		log.Println("CreateAuction: Error Create Auction in package service", errCreate)
		return errCreate
	}
	return nil
}

func (p *PaymentServiceDefault) UpdateAddressPayment(auction *entity.Payment) error {
	err := p.PaymentRepository.UpdateAddressPayment(auction)
	if err != nil {
		log.Println("Error in package service", err)
		return err
	}
	return nil
}

func (p *PaymentServiceDefault) DeletePayment(paymentId, userId uint) error {
	err := p.PaymentRepository.DeletePayment(paymentId, userId)
	if err != nil {
		log.Println("CreateAuction: Error Create Auction in package service", err)
		return err
	}
	return nil
}

func (p *PaymentServiceDefault) GetPaymentByPaymentId(paymentId uint) (*entity.Payment, error) {
	paymentDetail, err := p.PaymentRepository.GetPaymentByPaymentId(paymentId)
	if err != nil {
		log.Println("GetPaymentById: Error in get paymentDetail by Id", err)
		return nil, err
	}
	return paymentDetail, nil
}
