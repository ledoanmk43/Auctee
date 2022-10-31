package repository

import (
	"backend/pkg/utils"
	"backend/src/account-service/dto"
	"backend/src/account-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(address *entity.Address) error
	UpdateAddress(userId uint, address *dto.UpdateAddressDTO) error
	GetAllAddresses(userId uint) (*[]entity.Address, error)
	GetAddressByAddressId(addressId, userId uint) (*entity.Address, error)
	DeleteAddress(userId, addressId uint) error
}

type AddressRepositoryDefault struct {
	db *gorm.DB
}

func NewAddressRepositoryDefault(db *gorm.DB) *AddressRepositoryDefault {
	return &AddressRepositoryDefault{db: db}
}

func (a *AddressRepositoryDefault) CreateAddress(address *entity.Address) error {
	if errCheckEmptyField := address.Validate(); errCheckEmptyField != nil {
		log.Println("CreateAddress: Error empty field in package repository", errCheckEmptyField)
		return errCheckEmptyField
	}

	if *address.IsDefault {
		//Check if any address of an user has default tag and switch to false
		var addressCheckDefault *entity.Address
		var countDefault int64
		_ = a.db.Where("user_id = ? AND is_default = ?", address.UserId, true).Find(&addressCheckDefault).Count(&countDefault)
		if addressCheckDefault != nil || countDefault != 0 {
			addressCheckDefault.IsDefault = utils.BoolAddr(false)
			_ = a.db.Updates(&addressCheckDefault)
		}
	}
	result := a.db.Create(&address)
	if result.Error != nil {
		log.Println("CreateAddress: Error Create in package repository", result)
		return result.Error
	}
	return nil
}

func (a *AddressRepositoryDefault) UpdateAddress(userId uint, updateBody *dto.UpdateAddressDTO) error {
	if errCheckEmptyField := updateBody.Validate(); errCheckEmptyField != nil {
		log.Println("CreateAddress: Error empty field in package repository", errCheckEmptyField)
		return errCheckEmptyField
	}

	var addressToUpdate *entity.Address
	var count int64
	record := a.db.Where("user_id = ? AND id = ?", userId, updateBody.Id).Find(&addressToUpdate).Count(&count)
	if record.Error != nil {
		log.Println("Error update address in package repository")
		return record.Error
	}
	if count == 0 {
		return errors.New("address not found")
	}

	addressToUpdate.Firstname = updateBody.Firstname
	addressToUpdate.Lastname = updateBody.Lastname
	addressToUpdate.Phone = updateBody.Phone
	addressToUpdate.Email = updateBody.Email
	addressToUpdate.Province = updateBody.Province
	addressToUpdate.District = updateBody.District
	addressToUpdate.SubDistrict = updateBody.SubDistrict
	addressToUpdate.Address = updateBody.Address
	addressToUpdate.TypeAddress = updateBody.TypeAddress

	if addressToUpdate.IsDefault != nil {
		addressToUpdate.IsDefault = utils.BoolAddr(*updateBody.IsDefault)
	}
	if *addressToUpdate.IsDefault == true {
		var addressCheckDefault *entity.Address
		var countDefault int64
		_ = a.db.Where("id != ? AND is_default = ?", updateBody.Id, true).Find(&addressCheckDefault).Count(&countDefault)
		if addressCheckDefault != nil || countDefault != 0 {
			addressCheckDefault.IsDefault = utils.BoolAddr(false)
			_ = a.db.Updates(&addressCheckDefault)
		}
	}

	res := a.db.Updates(&addressToUpdate)
	if res.Error != nil {
		log.Println("Update Address: Error in package repository", res.Error)
		return res.Error
	}
	return nil
}

func (a *AddressRepositoryDefault) GetAllAddresses(userId uint) (*[]entity.Address, error) {
	var addresses *[]entity.Address
	result := a.db.Where("user_id = ?", userId).Order("is_default desc").Order("created_at desc").Find(&addresses)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	return addresses, nil
}

func (a *AddressRepositoryDefault) DeleteAddress(userId, addressId uint) error {
	var addressToDelete *entity.Address
	var count int64
	resultFind := a.db.Where("user_id = ? AND id = ?", userId, addressId).Find(&addressToDelete).Count(&count)
	if resultFind.Error != nil {
		log.Println("Delete Address: Error to find Address  in package repository", resultFind)
		return resultFind.Error
	}
	if count == 0 {
		return errors.New("address not found")
	}

	resultDelete := a.db.Where(" id = ?", addressId).Delete(&addressToDelete)
	if resultDelete.Error != nil {
		log.Println("Delete Address: Error to Deleted Address  in package repository", resultDelete)
		return resultDelete.Error
	}
	return nil
}

func (a *AddressRepositoryDefault) GetAddressByAddressId(addressId, userId uint) (*entity.Address, error) {
	var address *entity.Address
	var count int64
	result := a.db.Where("id = ? AND user_id =?", addressId, userId).Find(&address).Count(&count)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	if count == 0 {
		return nil, errors.New("address not found")
	}
	return address, nil
}
