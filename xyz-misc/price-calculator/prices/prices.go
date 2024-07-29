package prices

import (
	"fmt"

	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/conversion"
	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("Failed to read the lines:", err)
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Failed to convert to float:", err)
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInPrice)
	}
	// fmt.Printf("Result: %+v\n", result)
	job.TaxIncludedPrices = result
	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)

	// filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
