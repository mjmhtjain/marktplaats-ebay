package services

import (
	"testing"

	"github.com/mjmhtjain/marktplaats-ebay/src/daos"
	"github.com/mjmhtjain/marktplaats-ebay/src/models"
	"github.com/stretchr/testify/assert"
)

func Test_ecgCreditService_UploadCreditorInfo(t *testing.T) {

	t.Run("IF daoService performs well, THEN expect persisted creditors", func(t *testing.T) {
		var expectedErr error = nil
		creditService := ecgCreditService{
			creditDAO: NewFakeCreditDAO(expectedErr),
		}
		creditors := []models.Creditor{}

		insertedCreditors, err := creditService.UploadCreditorInfo(creditors)

		assert.Nil(t, err)
		assert.NotNil(t, insertedCreditors)
		assert.Len(t, insertedCreditors, len(creditors))
	})

	t.Run("IF daoService returns an error, THEN expect error", func(t *testing.T) {

	})
}

func NewFakeCreditDAO(e error) daos.CreditDAO {
	return &fakeCreditDAO{
		expectedErr: e,
	}
}

type fakeCreditDAO struct {
	expectedErr error
}

func (dao *fakeCreditDAO) GetAll() ([]models.Creditor, error) {
	return nil, nil
}

func (dao *fakeCreditDAO) InsertAll(creditors []models.Creditor) ([]models.Creditor, error) {
	if dao.expectedErr != nil {
		return nil, dao.expectedErr
	}

	return creditors, nil
}
