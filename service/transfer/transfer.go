package transferService

import (
	"fmt"
	"strconv"
	
	"github.com/omise/omise-go"

	parser "../../parser"
	omiseClient "../../dao/omise"
)

type transaferService struct {
	client *omiseClient.OmiseClient
}

func NewTransferService() (*transaferService, error) {
	client, err := omiseClient.NewClient()
	if err != nil {
		fmt.Println("failed to create omise client")
		return nil, err
	}
	return &transaferService{client: client}, nil
}


func (t *transaferService) TransferByCreditCard(donationInfo *parser.DonationRow) (*omise.Charge, error) {
	m, err := strconv.Atoi(donationInfo.ExpMonth)
	y, err := strconv.Atoi(donationInfo.ExpYear)

	card, err := t.client.CreateCardToken(&omiseClient.CreateCardInfo{
		Name: donationInfo.Name,
		Number: donationInfo.CreditCardNumber,
		ExpirationMonth: m,
		ExpirationYear: y,
	})
	if err != nil {
		fmt.Println("failed to create card token")
		return nil, err
	}
	charge, err := t.client.CreateCharge(&omiseClient.CreateChargeInfo{
		Amount: donationInfo.Price,
		TokenId: card.ID,
	})
	if err != nil {
		fmt.Println("failed to create charge")
		return nil, err
	}
	return charge, nil
}