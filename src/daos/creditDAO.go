package daos

import (
	"context"
	"errors"
	"time"

	"github.com/mjmhtjain/marktplaats-ebay/src/models"
	"github.com/mjmhtjain/marktplaats-ebay/src/setup"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreditDAO interface {
	GetAll() ([]models.Creditor, error)
	InsertAll([]models.Creditor) ([]models.Creditor, error)
}

type ecgCreditDAO struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewECGCreditDAO() CreditDAO {
	c := setup.NewMongo().Client
	creditorsCollection := c.Database("ecg").Collection("creditors")

	return &ecgCreditDAO{
		client:     c,
		collection: creditorsCollection,
	}
}

func (dao *ecgCreditDAO) GetAll() ([]models.Creditor, error) {
	return nil, nil
}

func (dao *ecgCreditDAO) InsertAll(creditors []models.Creditor) ([]models.Creditor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	arr := []interface{}{}
	for _, c := range creditors {
		arr = append(arr, c)
	}

	res, err := dao.collection.InsertMany(ctx, arr)
	if err != nil {
		return nil, err
	}

	if len(res.InsertedIDs) > 0 {
		return creditors, nil
	}

	return nil, errors.New("Something")
}
