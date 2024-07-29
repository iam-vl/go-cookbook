package prices

import (
	"fmt"

	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/conversion"
	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	// add struct tags
	IOManager         iomanager.IOManager `json:"-"` // ignore in the output
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
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
	// filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
