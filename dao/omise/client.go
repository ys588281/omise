package omise

import (
	"fmt"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"

	config "../../config"
)

type OmiseClient struct {
	client *omise.Client
}

func NewClient() (*OmiseClient, error) {
	client, err := omise.NewClient(config.OmisePublicKey, config.OmiseSecretKey)
	if err != nil {
		return nil, err
	}
	return &OmiseClient{
		client: client,
	}, err
}

type CreateChargeInfo struct {
	Amount int64
	TokenId string
}

func (c *OmiseClient) CreateCharge(createChargeInfo *CreateChargeInfo) (*omise.Charge, error) {
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   createChargeInfo.Amount,
		Currency: "thb",
		Card:     createChargeInfo.TokenId,
	}
	if err := c.client.Do(charge, createCharge); err != nil {
		fmt.Printf("failed to call omise api to create charge, %s \n", err)
		return nil, err
	}

	// log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
	return charge, nil
}

type CreateCardInfo struct {
	Name string
	Number string
	ExpirationMonth int
	ExpirationYear int
}

func (c *OmiseClient) CreateCardToken(createCardInfo *CreateCardInfo) (*omise.Card, error) {
	card, createToken := &omise.Card{}, &operations.CreateToken{
		Name:            createCardInfo.Name,
		Number:          createCardInfo.Number,
		ExpirationYear:  createCardInfo.ExpirationYear,
		ExpirationMonth: time.Month(createCardInfo.ExpirationMonth),
	}
	
	if err := c.client.Do(card, createToken); err != nil {
		fmt.Printf("failed to call omise api to create card token, %s \n", err)
		return nil, err
	}
	
	// log.Printf("created card: %#v\n", card)

	return card, nil
}

