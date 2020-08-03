package statisticsService

import (
	"sync"
	parser "../../parser"
)

type Summary struct {
	mux sync.Mutex
	TotalValue int64
	TotalCount int64
	SuccessValue int64
	SuccessCount int64
	FailValue int64
	FailCount int64
	TopValue int64
	TopDonators []string
}

func NewSummary() *Summary {
	return &Summary{}
}


func (s *Summary) AddSuccess(donationRow *parser.DonationRow) {
	s.mux.Lock()
	s.TotalValue = s.TotalValue + donationRow.Price
	s.TotalCount = s.TotalCount + 1
	s.SuccessValue = s.SuccessValue + donationRow.Price
	s.SuccessCount = s.SuccessCount + 1
	s.updateTop(donationRow)
	defer s.mux.Unlock()
}

func (s *Summary) AddFailed(donationRow *parser.DonationRow) {
	s.mux.Lock()
	s.TotalValue = s.TotalValue + donationRow.Price
	s.TotalCount = s.TotalCount + 1
	s.FailValue = s.FailValue + donationRow.Price
	s.FailCount = s.FailCount + 1
	defer s.mux.Unlock()
}

func (s *Summary) updateTop(donationRow *parser.DonationRow) {
	if donationRow.Price > s.TopValue {
		s.TopValue = donationRow.Price
		s.TopDonators = []string{donationRow.Name}
	} else if donationRow.Price == s.TopValue {
		s.TopDonators = append(s.TopDonators, donationRow.Name)
	}
}

