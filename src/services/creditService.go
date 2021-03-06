package services

import (
	"github.com/mjmhtjain/marktplaats-ebay/src/daos"
	"github.com/mjmhtjain/marktplaats-ebay/src/models"
)

type CreditService interface {
	UploadCreditorInfo(creditors []models.Creditor) ([]models.Creditor, error)
	GetCreditors() ([]models.Creditor, error)
}

type ecgCreditService struct {
	creditDAO daos.CreditDAO
}

func NewCreditService() CreditService {
	return &ecgCreditService{
		creditDAO: daos.NewECGCreditDAO(),
	}
}

func (ecg *ecgCreditService) GetCreditors() ([]models.Creditor, error) {
	creditors, err := ecg.creditDAO.GetAll()
	if err != nil {
		return nil, err
	}

	return creditors, nil
}

func (ecg *ecgCreditService) UploadCreditorInfo(creditors []models.Creditor) ([]models.Creditor, error) {
	if len(creditors) == 0 {
		return creditors, nil
	}

	insertedCreditors, err := ecg.creditDAO.InsertAll(creditors)
	if err != nil {
		return nil, err
	}

	return insertedCreditors, nil
}
