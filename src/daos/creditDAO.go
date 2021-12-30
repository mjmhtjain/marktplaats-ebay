package daos

import "github.com/mjmhtjain/marktplaats-ebay/src/models"

type CreditDAO interface {
	GetAll() ([]models.Creditor, error)
	InsertAll([]models.Creditor) ([]models.Creditor, error)
}

type ecgCreditDAO struct {
}

func NewECGCreditDAO() CreditDAO {
	return &ecgCreditDAO{}
}

func (dao *ecgCreditDAO) GetAll() ([]models.Creditor, error) {
	return nil, nil
}

func (dao *ecgCreditDAO) InsertAll([]models.Creditor) ([]models.Creditor, error) {
	return nil, nil
}
