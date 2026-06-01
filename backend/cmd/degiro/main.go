package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	colProduct     = 0
	colISIN        = 1
	colValueInEUR  = 6
	minColumnCount = 7
)

type investment struct {
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`
	AssetClass    string  `json:"asset_class"`
	AssetSubClass string  `json:"asset_sub_class"`
}

func main() {
	csvPath := flag.String("csv", defaultCSVPath(), "path to Degiro Portfolio.csv export")
	apiURL := flag.String("api", envOrDefault("API_URL", "http://localhost:8080"), "backend API base URL")
	flag.Parse()

	records, err := readPortfolioCSV(*csvPath)
	if err != nil {
		log.Fatalf("read portfolio CSV: %v", err)
	}

	client := &http.Client{}
	existing, err := fetchExistingNames(client, *apiURL)
	if err != nil {
		log.Fatalf("fetch existing investments: %v", err)
	}

	var imported, skipped int
	for _, record := range records {
		inv, err := recordToInvestment(record)
		if err != nil {
			log.Printf("skip row %q: %v", record[colProduct], err)
			skipped++
			continue
		}

		if existing[inv.Name] {
			log.Printf("skip duplicate: %s", inv.Name)
			skipped++
			continue
		}

		if err := createInvestment(client, *apiURL, inv); err != nil {
			log.Fatalf("create investment %q: %v", inv.Name, err)
		}

		log.Printf("imported: %s (%.2f EUR, %s / %s)", inv.Name, inv.Amount, inv.AssetClass, inv.AssetSubClass)
		imported++
		existing[inv.Name] = true
	}

	log.Printf("done: imported %d, skipped %d", imported, skipped)
}

func defaultCSVPath() string {
	if path := os.Getenv("PORTFOLIO_CSV"); path != "" {
		return path
	}
	return filepath.Join("data", "Portfolio.csv")
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func readPortfolioCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}

	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read row: %w", err)
		}
		if len(record) < minColumnCount {
			continue
		}
		if strings.TrimSpace(record[colProduct]) == "" {
			continue
		}
		records = append(records, record)
	}

	return records, nil
}

func recordToInvestment(record []string) (investment, error) {
	product := strings.TrimSpace(record[colProduct])
	isin := strings.TrimSpace(record[colISIN])
	valueEUR, err := strconv.ParseFloat(strings.TrimSpace(record[colValueInEUR]), 64)
	if err != nil {
		return investment{}, fmt.Errorf("parse value in EUR %q: %w", record[colValueInEUR], err)
	}

	name := product
	if isin != "" {
		name = fmt.Sprintf("%s (%s)", product, isin)
	}

	return investment{
		Name:          name,
		Amount:        valueEUR,
		AssetClass:    inferAssetClass(product),
		AssetSubClass: inferAssetSubClass(product),
	}, nil
}

func inferAssetClass(product string) string {
	upper := strings.ToUpper(product)

	switch {
	case strings.Contains(upper, "CASH"):
		return "Cash"
	case strings.Contains(upper, "GOLD"):
		return "Commodity"
	case strings.Contains(upper, "BOND"),
		strings.Contains(upper, "GOVT"),
		strings.Contains(upper, "TREASURY"),
		strings.Contains(upper, "CORP BD"),
		strings.Contains(upper, "CRP BD"):
		return "Bond"
	case strings.Contains(upper, "PROPERTY"),
		strings.Contains(upper, "REIT"),
		strings.Contains(upper, "REALTY"),
		strings.Contains(upper, "EPRA"):
		return "Real Estate"
	default:
		return "Equity"
	}
}

func inferAssetSubClass(product string) string {
	upper := strings.ToUpper(product)

	switch {
	case strings.Contains(upper, "CASH"):
		return "Cash"
	case strings.Contains(upper, "ETF"), strings.Contains(upper, "UCITS"):
		return "ETF"
	default:
		return "Stock"
	}
}

func fetchExistingNames(client *http.Client, apiURL string) (map[string]bool, error) {
	resp, err := client.Get(strings.TrimRight(apiURL, "/") + "/api/investments")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GET /api/investments returned %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}

	var investments []investment
	if err := json.NewDecoder(resp.Body).Decode(&investments); err != nil {
		return nil, err
	}

	names := make(map[string]bool, len(investments))
	for _, inv := range investments {
		names[inv.Name] = true
	}
	return names, nil
}

func createInvestment(client *http.Client, apiURL string, inv investment) error {
	payload, err := json.Marshal(inv)
	if err != nil {
		return err
	}

	resp, err := client.Post(
		strings.TrimRight(apiURL, "/")+"/api/investments",
		"application/json",
		bytes.NewReader(payload),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("POST /api/investments returned %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}

	return nil
}
