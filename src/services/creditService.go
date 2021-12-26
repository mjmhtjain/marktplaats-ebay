package services

import "github.com/mjmhtjain/marktplaats-ebay/src/models"

type CreditService interface {
	uploadCreditorInfo(creditors []models.Creditor) (models.Creditor, error)
	getCreditors()
}

type ECGCreditService struct {
}
