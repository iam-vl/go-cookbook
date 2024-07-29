package main

import (
	"fmt"

	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/filemanager"
	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// cmdm := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("cannot process this job:", err)
		}
	}
}
