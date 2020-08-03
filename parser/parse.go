package parser

import (
	"fmt"
	"strings"
	"strconv"
)

func ParseString(input string) (map[string]*DonationRow, error) {
	m := make(map[string]*DonationRow)
	s := strings.Split(input, "\n")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		line := strings.TrimSpace(s[i])
		if line == "" {
			continue
		}
		ss := strings.Split(line, ",")
		price, _ := strconv.ParseInt(ss[1], 10, 64)
		m[ss[0]] = &DonationRow{
			Name: ss[0],
			Price: price,
			CreditCardNumber: ss[2],
			CVV: ss[3],
			ExpMonth: ss[4],
			ExpYear: ss[5],
		}
	}
	return m, nil
}

type DonationRow struct {
	Name string
	Price int64
	CreditCardNumber string
	CVV string
	ExpMonth string
	ExpYear string
}