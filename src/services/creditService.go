package services

import "github.com/mjmhtjain/marktplaats-ebay/src/models"

type CreditService interface {
	UploadCreditorInfo(creditors []models.Creditor) ([]models.Creditor, error)
	GetCreditors()
}

func NewCreditService() CreditService {
	return &ecgCreditService{}
}

type ecgCreditService struct {
}

func (ecg *ecgCreditService) UploadCreditorInfo(creditors []models.Creditor) ([]models.Creditor, error) {
	return nil, nil
}

func (ecg *ecgCreditService) GetCreditors() {

}
