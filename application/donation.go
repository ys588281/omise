package application

import (
	"fmt"
	"strings"
	"strconv"
	"sync"
	"time"

	"go.uber.org/ratelimit"

	cipher "../cipher"
	parser "../parser"
	transferService "../service/transfer"
	statisticsService "../service/statistics"
)

func DonationByFile(filePath string) {
	rl := ratelimit.New(3)
	prev := time.Now()
	var wg sync.WaitGroup
	fmt.Println("performing donations...")
	buf, err := cipher.DecryptFile(filePath)
	if err != nil {
		fmt.Println("decrypt err: ", err)
	}
	summary := statisticsService.NewSummary()
	trnasferServiceImpl, err := transferService.NewTransferService()
	if err != nil {
		return
	}

	s := strings.Split(string(buf), "\n")
	for i := 0; i < len(s); i++ {
		now := rl.Take()
        fmt.Println(i, now.Sub(prev))
        prev = now
		fmt.Println(s[i])
		line := strings.TrimSpace(s[i])
		if line == "" {
			continue
		}
		ss := strings.Split(line, ",")
		price, err := strconv.ParseInt(ss[1], 10, 64)
		if err != nil {
			continue
		}
		donationRow := &parser.DonationRow{
			Name: ss[0],
			Price: price,
			CreditCardNumber: ss[2],
			CVV: ss[3],
			ExpMonth: ss[4],
			ExpYear: ss[5],
		}
		wg.Add(1)
		go func(donationRow *parser.DonationRow, summary *statisticsService.Summary) {
			_, err := trnasferServiceImpl.TransferByCreditCard(donationRow)
			if err != nil {
				summary.AddFailed(donationRow)
			} else {
				summary.AddSuccess(donationRow)
			}
			wg.Done()
		}(donationRow, summary)
		
	}
	wg.Wait()
	fmt.Printf("%+v", summary)
	fmt.Println("done.")
	fmt.Printf("total received: THB  %d\n", summary.TotalValue)
	fmt.Printf("successfully donated: THB  %d\n", summary.SuccessValue)
	fmt.Printf("faulty donation: THB  %d\n", summary.FailValue)

	fmt.Printf("average per person: THB  %d\n", fmt.Sprintf("%.2f", float64(summary.TotalValue)/float64(summary.TotalCount)))
	fmt.Println("top donors: ")
	for _, v := range summary.TopDonators {
		fmt.Printf("\t %s\n", v)
	}
}