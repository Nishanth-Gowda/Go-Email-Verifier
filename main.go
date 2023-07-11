package main

import (
	"bufio"
	"fmt"
	"github.com/nishanth-gowda/email-verifier/domain"
	"log"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan() {
		domain.CheckDomain(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}

}
