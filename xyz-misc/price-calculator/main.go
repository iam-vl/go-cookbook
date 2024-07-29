package main

import (
	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/cmdmanager"
	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		cmdm := cmdmanager.New()
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}
