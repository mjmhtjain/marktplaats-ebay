package services

import (
	"errors"
	"os"
	"testing"

	"github.com/gocarina/gocsv"
	"github.com/mjmhtjain/marktplaats-ebay/src/daos"
	"github.com/mjmhtjain/marktplaats-ebay/src/models"
	"github.com/stretchr/testify/assert"
)

func Test_ecgCreditService_UploadCreditorInfo(t *testing.T) {
	fake := fakeCreditDAO{}
	creditService := newCreditService(&fake)

	t.Run("IF all valid arguments passed, THEN expect persisted creditors", func(t *testing.T) {
		creditors := readCSVFile(t, "/../../resources/Workbook2.csv")
		insertedCreditors, err := creditService.UploadCreditorInfo(creditors)

		assert.Nil(t, err)
		assert.NotNil(t, insertedCreditors)
		assert.Equal(t, len(creditors), len(insertedCreditors))
	})

	t.Run("IF empty creditors are passed, THEN expect empty creditors returned", func(t *testing.T) {
		creditors := []models.Creditor{}

		insertedCreditors, err := creditService.UploadCreditorInfo(creditors)

		assert.Nil(t, err)
		assert.NotNil(t, insertedCreditors)
		assert.Equal(t, len(creditors), len(insertedCreditors))
	})

	t.Run("IF daoService returns an error, THEN expect error", func(t *testing.T) {
		creditors := readCSVFile(t, "/../../resources/Workbook2.csv")
		expErr := errors.New("Unexpected error")
		fake.expectedErr = expErr

		_, err := creditService.UploadCreditorInfo(creditors)

		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})
}

func newCreditService(creditDAO daos.CreditDAO) CreditService {
	return &ecgCreditService{
		creditDAO: creditDAO,
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

func readCSVFile(t *testing.T, relativePath string) []models.Creditor {
	wd, _ := os.Getwd()
	filePath := wd + relativePath
	file, err := os.Open(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	creditors := []models.Creditor{}
	if err := gocsv.UnmarshalFile(file, &creditors); err != nil {
		t.Error(err)
	}

	return creditors
}
