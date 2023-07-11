package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}

}

func checkDomain(domain string) {
	startTime := time.Now()
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mx, err := net.LookupMX(domain)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if len(mx) > 0 {
		hasMX = true
	}

	txt, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	for _, record := range txt {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	lookupTXT, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	for _, record := range lookupTXT {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time: %s", elapsedTime)
}