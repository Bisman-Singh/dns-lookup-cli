package main

import (
	"fmt"
	"net"
	"strings"
)

// DNSRecord represents a single DNS record.
type DNSRecord struct {
	Domain     string `json:"domain"`
	RecordType string `json:"record_type"`
	Value      string `json:"value"`
	Priority   uint16 `json:"priority,omitempty"`
}

// QueryResult holds all records for a domain query.
type QueryResult struct {
	Domain     string      `json:"domain"`
	RecordType string      `json:"record_type"`
	Records    []DNSRecord `json:"records"`
	Error      string      `json:"error,omitempty"`
}

// Resolve performs a DNS lookup for the given domain and record type.
func Resolve(domain string, recordType string) QueryResult {
	result := QueryResult{
		Domain:     domain,
		RecordType: strings.ToUpper(recordType),
	}

	switch strings.ToUpper(recordType) {
	case "A":
		result.Records, result.Error = resolveA(domain)
	case "AAAA":
		result.Records, result.Error = resolveAAAA(domain)
	case "MX":
		result.Records, result.Error = resolveMX(domain)
	case "CNAME":
		result.Records, result.Error = resolveCNAME(domain)
	case "TXT":
		result.Records, result.Error = resolveTXT(domain)
	case "NS":
		result.Records, result.Error = resolveNS(domain)
	default:
		result.Error = fmt.Sprintf("unsupported record type: %s", recordType)
	}

	return result
}

func resolveA(domain string) ([]DNSRecord, string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err.Error()
	}

	var records []DNSRecord
	for _, ip := range ips {
		if ip.To4() != nil {
			records = append(records, DNSRecord{
				Domain:     domain,
				RecordType: "A",
				Value:      ip.String(),
			})
		}
	}

	if len(records) == 0 {
		return nil, "no A records found"
	}
	return records, ""
}

func resolveAAAA(domain string) ([]DNSRecord, string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err.Error()
	}

	var records []DNSRecord
	for _, ip := range ips {
		if ip.To4() == nil && ip.To16() != nil {
			records = append(records, DNSRecord{
				Domain:     domain,
				RecordType: "AAAA",
				Value:      ip.String(),
			})
		}
	}

	if len(records) == 0 {
		return nil, "no AAAA records found"
	}
	return records, ""
}

func resolveMX(domain string) ([]DNSRecord, string) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return nil, err.Error()
	}

	var records []DNSRecord
	for _, mx := range mxRecords {
		records = append(records, DNSRecord{
			Domain:     domain,
			RecordType: "MX",
			Value:      mx.Host,
			Priority:   mx.Pref,
		})
	}

	if len(records) == 0 {
		return nil, "no MX records found"
	}
	return records, ""
}

func resolveCNAME(domain string) ([]DNSRecord, string) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		return nil, err.Error()
	}

	records := []DNSRecord{
		{
			Domain:     domain,
			RecordType: "CNAME",
			Value:      cname,
		},
	}
	return records, ""
}

func resolveTXT(domain string) ([]DNSRecord, string) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		return nil, err.Error()
	}

	var records []DNSRecord
	for _, txt := range txtRecords {
		records = append(records, DNSRecord{
			Domain:     domain,
			RecordType: "TXT",
			Value:      txt,
		})
	}

	if len(records) == 0 {
		return nil, "no TXT records found"
	}
	return records, ""
}

func resolveNS(domain string) ([]DNSRecord, string) {
	nsRecords, err := net.LookupNS(domain)
	if err != nil {
		return nil, err.Error()
	}

	var records []DNSRecord
	for _, ns := range nsRecords {
		records = append(records, DNSRecord{
			Domain:     domain,
			RecordType: "NS",
			Value:      ns.Host,
		})
	}

	if len(records) == 0 {
		return nil, "no NS records found"
	}
	return records, ""
}
