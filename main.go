package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	recordType := flag.String("type", "A", "DNS record type: A, AAAA, MX, CNAME, TXT, NS")
	jsonOutput := flag.Bool("json", false, "output results as JSON")
	flag.Parse()

	domains := flag.Args()
	if len(domains) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: dns-lookup-cli [flags] domain1 [domain2 ...]")
		fmt.Fprintln(os.Stderr, "\nFlags:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nSupported record types: A, AAAA, MX, CNAME, TXT, NS")
		os.Exit(1)
	}

	validTypes := map[string]bool{
		"A": true, "AAAA": true, "MX": true,
		"CNAME": true, "TXT": true, "NS": true,
	}

	rt := strings.ToUpper(*recordType)
	if !validTypes[rt] {
		fmt.Fprintf(os.Stderr, "Error: unsupported record type %q\n", *recordType)
		fmt.Fprintln(os.Stderr, "Supported types: A, AAAA, MX, CNAME, TXT, NS")
		os.Exit(1)
	}

	var results []QueryResult
	for _, domain := range domains {
		result := Resolve(domain, rt)
		results = append(results, result)
	}

	if *jsonOutput {
		FormatJSON(results)
	} else {
		FormatTable(results)
	}

	// Exit with error if any lookup failed
	for _, r := range results {
		if r.Error != "" {
			os.Exit(1)
		}
	}
}
