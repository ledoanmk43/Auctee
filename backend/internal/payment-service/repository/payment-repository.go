package repository

import (
	payment "backend/internal/payment-service/config"
	"backend/internal/payment-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type IPaymentRepository interface {
	CreatePayment(payment *entity.Payment) (uint, error)
	UpdateAddressPayment(b *entity.Payment) error
	DeletePayment(paymentId, userId uint) error
	GetPaymentByPaymentId(paymentId, userId uint) (*entity.Payment, error)
	GetAllPaymentsForWinner(page int, winnerId uint) (*[]entity.Payment, error)
	GetAllPaymentsForOwner(page int, ownerId uint) (*[]entity.Payment, error)
}

type PaymentRepositoryDefault struct {
	connection *gorm.DB
}

func NewPaymentRepositoryDefault(dbConn *gorm.DB) *PaymentRepositoryDefault {
	return &PaymentRepositoryDefault{connection: dbConn}
}

func (p *PaymentRepositoryDefault) CreatePayment(payment *entity.Payment) (uint, error) {
	payment.CheckoutTime = time.Now()
	var count int64
	resIsDuplicated := p.connection.Where("auction_id = ? and winner_id = ?", payment.AuctionId, payment.WinnerId).Find(&payment).Count(&count)
	if resIsDuplicated.Error != nil {
		log.Println("Error to create payment in repo: ", resIsDuplicated.Error)
		return 0, resIsDuplicated.Error
	}
	if count != 0 {
		return payment.Id, errors.New("order is pending")
	}

	record := p.connection.Create(&payment)
	if record.Error != nil {
		log.Println("Error to create payment in repo: ", record.Error)
		return 0, record.Error
	}

	return payment.Id, nil
}

func (p *PaymentRepositoryDefault) GetAllPaymentsForWinner(page int, winnerId uint) (*[]entity.Payment, error) {
	var payments *[]entity.Payment
	record := p.connection.Limit(payment.PerPage).Offset((page-1)*payment.PerPage).Where("winner_id = ? ", winnerId).Order("created_at desc").Find(&payments)
	if record.Error != nil {
		log.Println("Get auctions: Error get all auctions in repo", record.Error)
		return nil, record.Error
	}

	return payments, nil
}

func (p *PaymentRepositoryDefault) GetAllPaymentsForOwner(page int, ownerId uint) (*[]entity.Payment, error) {
	var payments *[]entity.Payment
	record := p.connection.Limit(payment.PerPage).Offset((page-1)*payment.PerPage).Where("owner_id = ? ", ownerId).Find(&payments)
	if record.Error != nil {
		log.Println("Get auctions: Error get all auctions in repo", record.Error)
		return nil, record.Error
	}

	return payments, nil
}

func (p *PaymentRepositoryDefault) GetAllBills(page int, userId uint) (*[]entity.Payment, error) {
	var payments *[]entity.Payment
	record := p.connection.Limit(payment.PerPage).Offset((page-1)*payment.PerPage).Where("winner_id = ? ", userId).Find(&payments)
	if record.Error != nil {
		log.Println("Get auctions: Error get all auctions in repo", record.Error)
		return nil, record.Error
	}

	return payments, nil
}

func (p *PaymentRepositoryDefault) UpdateAddressPayment(updateBody *entity.Payment) error {
	var paymentToUpdate *entity.Payment
	var count int64
	record := p.connection.Where("id = ?", updateBody.ID).Find(&paymentToUpdate).Count(&count)

	//if paymentToUpdate.CreatedAt.Before(time.Now().Add(-time.Hour * 8)) {
	//	return errors.New("can not update after 12 hours")
	//}
	if record.Error != nil {
		log.Println("Error to find payment to update in repo", record.Error)
		return record.Error
	}

	if count == 0 {
		return errors.New("payment not found")
	}

	if updateBody.WinnerId != paymentToUpdate.WinnerId {
		return errors.New("Unauthorized")
	}

	if paymentToUpdate.CheckoutStatus == 2 {
		return errors.New("your order has been confirm")
	}

	//Address
	paymentToUpdate.Firstname = updateBody.Firstname
	paymentToUpdate.Lastname = updateBody.Lastname
	paymentToUpdate.Phone = updateBody.Phone
	paymentToUpdate.Email = updateBody.Email
	paymentToUpdate.Province = updateBody.Province
	paymentToUpdate.District = updateBody.District
	paymentToUpdate.SubDistrict = updateBody.SubDistrict
	paymentToUpdate.Address = updateBody.Address
	paymentToUpdate.TypeAddress = updateBody.TypeAddress

	recordSave := p.connection.Updates(&paymentToUpdate)
	if recordSave.Error != nil {
		log.Println("Error to update payment repo", recordSave.Error)
		return recordSave.Error
	}
	return nil
}

func (a *PaymentRepositoryDefault) DeletePayment(paymentId, userId uint) error {
	var payment *entity.Payment
	var count int64
	result := a.connection.Where("id = ? AND winner_id = ?", paymentId, userId).Find(&payment).Count(&count)
	if result.Error != nil {
		log.Println("Delete payment: Error in find payment to delete in package repository", result.Error)
		return result.Error
	}
	if count == 0 {
		return errors.New("payment not found")
	}

	if payment.CheckoutStatus == 2 {
		return errors.New("your order has been confirm")
	}

	a.connection.Delete(&payment)
	return nil
}

func (p *PaymentRepositoryDefault) GetPaymentByPaymentId(paymentId, userId uint) (*entity.Payment, error) {
	var payment *entity.Payment
	var count int64
	record := p.connection.Where("id = ? and winner_id = ?", paymentId, userId).Find(&payment).Count(&count)
	if record.Error != nil {
		log.Println("Error to find payment in repo")
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetPaymentById: payment not found", count)
		return nil, errors.New("error: payment not found")
	}
	return payment, nil
}
