package service

import (
	"backend/src/account-service/entity"
	"backend/src/account-service/repository"
	"backend/src/user-service-mock/dto"
	"log"
)

type IAddressService interface {
	CreateAddress(address *entity.Address) error
	UpdateAddress(address *entity.Address) (*entity.Address, error)
	GetAllAddresses(dto *dto.GetAddressDTO) (*[]entity.Address, error)
	GetAddressById(dto *dto.GetAddressByIdDTO) (*entity.Address, error)
	DeleteAddress(dto *dto.GetAddressByIdDTO) error
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

func (a *AddressService) UpdateAddress(address *entity.Address) (*entity.Address, error) {
	updateAddress, err := a.AddressRepository.UpdateAddress(address)
	if err != nil {
		log.Println("UpdateAddress: Error Update address in package service", err)
		return nil, err
	}
	return updateAddress, nil
}

func (a *AddressService) GetAllAddresses(dto *dto.GetAddressDTO) (*[]entity.Address, error) {
	address, err := a.AddressRepository.GetAllAddresses(dto)
	if err != nil {
		log.Println("GetAddress: Error GetAddress in package address-service", err)
		return nil, err
	}
	return address, nil
}

func (a *AddressService) DeleteAddress(dto *dto.GetAddressByIdDTO) error {
	err := a.AddressRepository.DeleteAddress(dto)
	if err != nil {
		log.Println("DeletedAddress: Error Delete Address in package service")
		return err
	}
	return nil
}

func (a *AddressService) GetAddressById(dto *dto.GetAddressByIdDTO) (*entity.Address, error) {
	address, err := a.AddressRepository.GetAddressById(dto)
	if err != nil {
		log.Println("GetAddressById: Error in get address by id in package uer-service", err)
		return nil, err
	}
	return address, nil
}
