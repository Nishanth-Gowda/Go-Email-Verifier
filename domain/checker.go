package domain

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func CheckDomain(domain string) {
	startTime := time.Now()
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mx, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal("Sorry No such Domain Exists\n", err)
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

	fmt.Println("==============")
	fmt.Printf(" domain: %v,\n hasMX: %v,\n hasSPF: %v,\n hasDMARC: %v,\n spfRecord: %v,\n dmarcRecord:%v,\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time: %s", elapsedTime)
}
