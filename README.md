# DNS Lookup CLI

A command-line DNS query tool written in Go. Supports multiple record types and can query multiple domains at once with both table and JSON output formats.

## Features

- Support for record types: A, AAAA, MX, CNAME, TXT, NS
- Query multiple domains in a single command
- JSON output mode for scripting and automation
- Human-readable table output by default
- MX records displayed with priority values
- Exit code 1 if any lookup fails

## Build

```bash
go build -o dns-lookup-cli .
```

## Usage

```bash
# Look up A records for a domain
./dns-lookup-cli example.com

# Look up MX records
./dns-lookup-cli -type MX example.com

# Query multiple domains
./dns-lookup-cli -type NS example.com example.org

# Get JSON output
./dns-lookup-cli -json -type TXT example.com

# Look up AAAA (IPv6) records
./dns-lookup-cli -type AAAA example.com
```

### Flags

| Flag    | Default | Description                              |
|---------|---------|------------------------------------------|
| `-type` | `A`     | DNS record type (A, AAAA, MX, CNAME, TXT, NS) |
| `-json` | `false` | Output results as JSON                   |

### Example Output

**Table format:**
```
--- example.com (A) ---
  VALUE
  --------------------------------------------------
  93.184.216.34

--- example.com (MX) ---
  VALUE                                    PRIORITY
  -------------------------------------------------------
  mail.example.com.                        10
  mail2.example.com.                       20
```

**JSON format:**
```json
[
  {
    "domain": "example.com",
    "record_type": "A",
    "records": [
      {
        "domain": "example.com",
        "record_type": "A",
        "value": "93.184.216.34"
      }
    ]
  }
]
```


