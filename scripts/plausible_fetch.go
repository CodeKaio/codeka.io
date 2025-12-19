package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// APIRequest defines the body we send to Plausible.
type APIRequest struct {
	SiteID     string   `json:"site_id"`
	Metrics    []string `json:"metrics"`
	DateRange  string   `json:"date_range"`
	Dimensions []string `json:"dimensions"`
}

// APIResult matches items under response.results
// Example item: {"metrics": [270], "dimensions": ["/path/"]}
type APIResult struct {
	Metrics    []int64  `json:"metrics"`
	Dimensions []string `json:"dimensions"`
}

// APIResponse is the top-level structure returned by Plausible.
type APIResponse struct {
	Results []APIResult `json:"results"`
}

// OutputEntry is the simplified shape we want to write: page/visits.
type OutputEntry struct {
	Page   string `json:"page"`
	Visits int64  `json:"visits"`
}

func main() {
	var (
		apiURL  = flag.String("url", "https://plausible.io/api/v2/query", "Plausible API endpoint")
		siteID  = flag.String("site", "dummy.site", "site_id parameter")
		date    = flag.String("range", "7d", "date_range parameter (e.g., 7d, 30d, month")
		outFile = flag.String("out", "plausible-pages.json", "output JSON file path")
	)
	flag.Parse()

	apiKey := os.Getenv("PLAUSIBLE_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "error: PLAUSIBLE_API_KEY environment variable is not set")
		os.Exit(1)
	}

	reqBody := APIRequest{
		SiteID:    *siteID,
		Metrics:   []string{
			"pageviews",
		},
		DateRange: *date,
		Dimensions: []string{
			"event:page",
		},
	}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal request: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, *apiURL, bytes.NewReader(payload))
	if err != nil {
		fmt.Fprintf(os.Stderr, "build request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(os.Stderr, "http error: %s\n%s\n", resp.Status, string(b))
		os.Exit(1)
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Fprintf(os.Stderr, "decode response: %v\n", err)
		os.Exit(1)
	}

	// Transform: results -> [{page, visits}]
	out := make([]OutputEntry, 0, len(apiResp.Results))
	for _, r := range apiResp.Results {
		var page string
		if len(r.Dimensions) > 0 {
			page = r.Dimensions[0]
		}
		var visits int64
		if len(r.Metrics) > 0 {
			visits = r.Metrics[0]
		}
		out = append(out, OutputEntry{Page: page, Visits: visits})
	}

	f, err := os.Create(*outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(out); err != nil {
		fmt.Fprintf(os.Stderr, "write output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("wrote %d entries to %s\n", len(out), *outFile)
}
