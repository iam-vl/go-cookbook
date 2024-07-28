package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iam-vl/go-cookbook/xyz-misc/price-calculator/convert"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // *bufio.Scanner

	var lines []string
	// func (s *bufio.Scanner) Scan() bool
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	prices, err := convert.StringsToFloats(lines)
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
	fmt.Printf("Result: %+v\n", result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
