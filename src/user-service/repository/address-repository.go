package repository

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
	UpdateAddress(address *entity.Address) (*entity.Address, error)
	GetAddress(dto *dto.GetAddressDTO) (*[]entity.Address, error)
	GetAddressById(dto *dto.GetAddressByIdDTO) (*entity.Address, error)
	DeleteAddress(dto *dto.GetAddressByIdDTO) error
}

type AddressRepositoryDefault struct {
	db *gorm.DB
}

func NewAddressRepositoryDefault(db *gorm.DB) *AddressRepositoryDefault {
	return &AddressRepositoryDefault{db: db}
}

func (a AddressRepositoryDefault) CreateAddress(address *entity.Address) (*entity.Address, error) {
	if errCheckEmptyField := address.Validate(); errCheckEmptyField != nil {
		log.Println("CreateAddress: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}
	result := a.db.Create(&address)
	if result.Error != nil {
		log.Println("CreateAddress: Error Create in package repository", result)
		return nil, result.Error
	}
	return address, nil
}

func (a *AddressRepositoryDefault) UpdateAddress(address *entity.Address) (*entity.Address, error) {
	if errCheckEmptyField := address.Validate(); errCheckEmptyField != nil {
		log.Println("CreateAddress: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	var matchedAddress *entity.Address
	var count int64

	record := a.db.Where("user_id = ? AND id = ?", address.UserId, address.ID).Find(&matchedAddress).Count(&count)
	if record.Error != nil {
		log.Println("Error update serive in package repository")
		return nil, record.Error
	}
	if count == 0 {
		log.Println("=0")
		return nil, nil
	}
	matchedAddress = address
	recordUpdate := a.db.Updates(&matchedAddress)
	if recordUpdate.Error != nil {
		log.Println("Error ne thang lol")
		return nil, recordUpdate.Error
	}
	return matchedAddress, nil
}

func (a AddressRepositoryDefault) GetAddress(dto *dto.GetAddressDTO) (*[]entity.Address, error) {
	var address *[]entity.Address
	result := a.db.Where("user_id = ?", dto.UserId).Find(&address)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	return address, nil
}

func (a AddressRepositoryDefault) DeleteAddress(dto *dto.GetAddressByIdDTO) error {
	var deleteAddress *entity.Address
	resultFind := a.db.Where("user_id = ? AND id = ?", dto.UserId, dto.AddressId).Find(&deleteAddress)
	if resultFind.Error != nil {
		log.Println("DeleteAddress: Error to find Address  in package repository", resultFind)
		return resultFind.Error
	}
	resultDelete := a.db.Delete(&deleteAddress)
	if resultDelete.Error != nil {
		log.Println("DeleteAddress: Error to Deleted Address  in package repository", resultDelete)
		return resultDelete.Error
	}
	return nil
}

func (a AddressRepositoryDefault) GetAddressById(dto *dto.GetAddressByIdDTO) (*entity.Address, error) {
	var address *entity.Address
	var count int64
	result := a.db.Where("id = ? And user_id =?", dto.AddressId, dto.UserId).Find(&address).Count(&count)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	if count == 0 {
		return nil, errors.New("not found address")
	}
	return address, nil
}
