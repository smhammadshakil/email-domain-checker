package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain:\nhasMX:\nhasSPF:\nspfRecord:\nhasDMARC:\ndmarcRecord\n\n")
	fmt.Print("Input domain:")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error: cound not read from input: %v\n\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}
	fmt.Printf("=> hasMX: %v\n", hasMX)
	fmt.Printf("=> hasSPF: %v\n", hasSPF)
	fmt.Printf("=> spfRecord: %v\n", spfRecord)
	fmt.Printf("=> hasDMARC: %v\n", hasDMARC)
	fmt.Printf("=> dmarcRecord: %v\n", dmarcRecord)

}
