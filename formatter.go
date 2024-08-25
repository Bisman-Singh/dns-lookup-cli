package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// FormatTable prints query results in a human-readable table format.
func FormatTable(results []QueryResult) {
	for _, result := range results {
		fmt.Printf("\n--- %s (%s) ---\n", result.Domain, result.RecordType)

		if result.Error != "" {
			fmt.Printf("  Error: %s\n", result.Error)
			continue
		}

		if len(result.Records) == 0 {
			fmt.Println("  No records found.")
			continue
		}

		switch result.RecordType {
		case "MX":
			fmt.Printf("  %-40s %s\n", "VALUE", "PRIORITY")
			fmt.Printf("  %s\n", strings.Repeat("-", 55))
			for _, r := range result.Records {
				fmt.Printf("  %-40s %d\n", r.Value, r.Priority)
			}
		default:
			fmt.Printf("  %s\n", "VALUE")
			fmt.Printf("  %s\n", strings.Repeat("-", 50))
			for _, r := range result.Records {
				fmt.Printf("  %s\n", r.Value)
			}
		}
	}
	fmt.Println()
}

// FormatJSON prints query results as a JSON array.
func FormatJSON(results []QueryResult) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
	}
}
