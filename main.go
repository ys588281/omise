package main

import (
	donationApplication "./application"
)

func main() {
	donationApplication.DonationByFile("data/fng.1000.csv.rot128")
}