package repository

import (
	payment "backend/src/payment-service/config"
	"backend/src/payment-service/entity"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strings"
	"time"
)

type IPaymentRepository interface {
	CreatePayment(payment *entity.Payment) (string, error)
	UpdateAddressPayment(payment *entity.Payment) (*entity.Payment, error)
	CancelPayment(paymentId string, userId uint) error
	GetPaymentByPaymentId(paymentId string, userId uint) (*entity.Payment, error)
	GetAllPaymentsForWinner(page int, winnerId uint) (*[]entity.Payment, error)
	GetAllPaymentsForOwner(page int, ownerId uint) (*[]entity.Payment, error)
}

type PaymentRepositoryDefault struct {
	connection *gorm.DB
}

func NewPaymentRepositoryDefault(dbConn *gorm.DB) *PaymentRepositoryDefault {
	return &PaymentRepositoryDefault{connection: dbConn}
}

func (p *PaymentRepositoryDefault) CreatePayment(payment *entity.Payment) (string, error) {
	payment.CheckoutTime = time.Now()
	payment.EndTime = strings.ReplaceAll(payment.EndTime, "T", " ")
	payment.EndTime = strings.ReplaceAll(payment.EndTime, "Z", "")

	var count int64
	resIsDuplicated := p.connection.Where("auction_id = ? and winner_id = ?", payment.AuctionId, payment.WinnerId).Find(&payment).Count(&count)
	if resIsDuplicated.Error != nil {
		log.Println("Error to create payment in repo: ", resIsDuplicated.Error)
		return "", resIsDuplicated.Error
	}
	if count != 0 {
		return payment.Id, errors.New("order is pending")
	}
	rand.Seed(time.Now().UnixNano())
	transID := rand.Intn(1000)
	now := time.Now()
	appTransID := fmt.Sprintf("%02d%02d%02d_%v", now.Year()%100, int(now.Month()), now.Day(), transID)
	payment.Id = appTransID
	record := p.connection.Create(&payment)
	if record.Error != nil {
		log.Println("Error to create payment in repo: ", record.Error)
		return "", record.Error
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

func (p *PaymentRepositoryDefault) UpdateAddressPayment(updateBody *entity.Payment) (*entity.Payment, error) {
	var paymentToUpdate *entity.Payment

	var count int64
	record := p.connection.Where("id = ?", updateBody.Id).Find(&paymentToUpdate).Count(&count)

	//if paymentToUpdate.CreatedAt.Before(time.Now().Add(-time.Hour * 8)) {
	//	return errors.New("can not update after 12 hours")
	//}
	if record.Error != nil {
		log.Println("Error to find payment to update in repo", record.Error)
		return nil, record.Error
	}

	if count == 0 {
		return nil, errors.New("payment not found")
	}

	if updateBody.ShippingStatus == 2 || updateBody.CheckoutStatus == 5 {
		if updateBody.OwnerId != paymentToUpdate.OwnerId {
			return nil, errors.New("Unauthorized")
		}
	} else {
		if updateBody.WinnerId != paymentToUpdate.WinnerId {
			return nil, errors.New("Unauthorized")
		}
	}

	if paymentToUpdate.CheckoutStatus == 2 {
		return nil, errors.New("your order has been confirm")
	}

	//Address
	paymentToUpdate.EndTime = strings.ReplaceAll(paymentToUpdate.EndTime, "Z", "")
	paymentToUpdate.EndTime = strings.ReplaceAll(paymentToUpdate.EndTime, "T", " ")
	if updateBody.CheckoutStatus != 5 {
		paymentToUpdate.AddressId = updateBody.AddressId
		paymentToUpdate.Firstname = updateBody.Firstname
		paymentToUpdate.Lastname = updateBody.Lastname
		paymentToUpdate.Phone = updateBody.Phone
		paymentToUpdate.Email = updateBody.Email
		paymentToUpdate.Province = updateBody.Province
		paymentToUpdate.District = updateBody.District
		paymentToUpdate.SubDistrict = updateBody.SubDistrict
		paymentToUpdate.CheckoutStatus = updateBody.CheckoutStatus
		paymentToUpdate.Address = updateBody.Address
		paymentToUpdate.TypeAddress = updateBody.TypeAddress
		paymentToUpdate.PaymentMethod = updateBody.PaymentMethod
		paymentToUpdate.ShippingValue = updateBody.ShippingValue
		paymentToUpdate.Note = updateBody.Note
		paymentToUpdate.ShippingStatus = updateBody.ShippingStatus
		paymentToUpdate.Total = updateBody.Total
	} else {
		paymentToUpdate.CheckoutStatus = 5
	}
	recordSave := p.connection.Updates(&paymentToUpdate)
	if recordSave.Error != nil {
		log.Println("Error to update payment repo", recordSave.Error)
		return nil, recordSave.Error
	}
	return paymentToUpdate, nil
}

func (p *PaymentRepositoryDefault) CancelPayment(paymentId string, userId uint) error {
	var payment *entity.Payment
	var count int64
	result := p.connection.Where("id = ? AND winner_id = ?", paymentId, userId).Find(&payment).Count(&count)
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

	payment.EndTime = strings.ReplaceAll(payment.EndTime, "Z", "")
	payment.EndTime = strings.ReplaceAll(payment.EndTime, "T", " ")
	payment.CheckoutStatus = 4 // xác nhận huỷ đơn hàng
	recordSave := p.connection.Updates(&payment)
	if recordSave.Error != nil {
		log.Println("Error to cancel payment repo", recordSave.Error)
		return recordSave.Error
	}
	return nil
}

func (p *PaymentRepositoryDefault) GetPaymentByPaymentId(paymentId string, userId uint) (*entity.Payment, error) {
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
