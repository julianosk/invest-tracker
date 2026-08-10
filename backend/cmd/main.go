package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type InvestmentRequest struct {
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`
	AssetClass    string  `json:"asset_class"`
	AssetSubClass string  `json:"asset_sub_class"`
}

func main() {
	file, err := os.Open("data/Degiro-Portfolio.csv")
	if err != nil {
		log.Fatal("failed to open CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("failed to read CSV:", err)
	}

	apiURL := "http://localhost:8080/api/investments"

	for i, record := range records {
		if i == 0 || len(record) < 7 || record[0] == "" {
			continue // skip header or empty rows
		}

		product := record[0]
		valueEurStr := record[6]
		if valueEurStr == "" {
			continue
		}

		valueEur, err := strconv.ParseFloat(valueEurStr, 64)
		if err != nil {
			log.Printf("failed to parse value for %s: %v", product, err)
			continue
		}

		assetClass := categorizeAssetClass(product)

		investment := InvestmentRequest{
			Name:          product,
			Amount:        valueEur,
			AssetClass:    assetClass,
			AssetSubClass: "",
		}

		jsonData, err := json.Marshal(investment)
		if err != nil {
			log.Printf("failed to marshal investment for %s: %v", product, err)
			continue
		}

		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("failed to post investment for %s: %v", product, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			log.Printf("failed to create investment for %s: status %d", product, resp.StatusCode)
			continue
		}

		log.Printf("Imported: %s - %.2f EUR (%s)", product, valueEur, assetClass)
	}
}

func categorizeAssetClass(product string) string {
	productUpper := strings.ToUpper(product)
	if strings.Contains(productUpper, "ETF") || strings.Contains(productUpper, "UCITS") {
		return "ETF"
	}
	if strings.Contains(productUpper, "CORP") || strings.Contains(productUpper, "INC") || strings.Contains(productUpper, "ADR") {
		return "Stock"
	}
	if strings.Contains(productUpper, "CASH") || strings.Contains(productUpper, "FUND") {
		return "Cash"
	}
	if strings.Contains(productUpper, "BOND") || strings.Contains(productUpper, "TREASURY") {
		return "Bond"
	}
	return "Other"
}