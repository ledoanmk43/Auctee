package service

import (
	"backend/internal/account-service/dto"
	"backend/internal/account-service/entity"
	"backend/internal/account-service/repository"
	"log"
)

type IAddressService interface {
	CreateAddress(address *entity.Address) error
	UpdateAddress(userId uint, address *dto.UpdateAddressDTO) error
	GetAllAddresses(userId uint) (*[]entity.Address, error)
	GetAddressByAddressId(addressId, userId uint) (*entity.Address, error)
	DeleteAddress(userId, addressId uint) error
}

type AddressService struct {
	AddressRepository repository.IAddressRepository
}

func NewAddressServiceDefault(addressRepository repository.IAddressRepository) *AddressService {
	return &AddressService{AddressRepository: addressRepository}
}

func (a *AddressService) CreateAddress(address *entity.Address) error {
	err := a.AddressRepository.CreateAddress(address)
	if err != nil {
		log.Println("CreateAddress: Error Create address in package service", err)
		return err
	}
	return nil
}

func (a *AddressService) UpdateAddress(userId uint, address *dto.UpdateAddressDTO) error {
	err := a.AddressRepository.UpdateAddress(userId, address)
	if err != nil {
		log.Println("UpdateAddress: Error Update address in package service", err)
		return err
	}
	return nil
}

func (a *AddressService) GetAllAddresses(userId uint) (*[]entity.Address, error) {
	address, err := a.AddressRepository.GetAllAddresses(userId)
	if err != nil {
		log.Println("GetAddress: Error GetAddress in package address-service", err)
		return nil, err
	}
	return address, nil
}

func (a *AddressService) DeleteAddress(userId, addressId uint) error {
	err := a.AddressRepository.DeleteAddress(userId, addressId)
	if err != nil {
		log.Println("DeletedAddress: Error Delete Address in package service")
		return err
	}
	return nil
}

func (a *AddressService) GetAddressByAddressId(addressId, userId uint) (*entity.Address, error) {
	address, err := a.AddressRepository.GetAddressByAddressId(addressId, userId)
	if err != nil {
		log.Println("GetAddressById: Error in get address by id in package service", err)
		return nil, err
	}
	return address, nil
}
