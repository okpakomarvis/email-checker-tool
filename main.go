package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spreRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan(){
		checker(scanner.Text())
	}

	if err:= scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input %v", err)
	}
}
func checker(domain string){
	var hasMX, hasSPF, hasDMARC bool
	var spreRecord, dmarcRecord string

	 mxRecord, err := net.LookupMX(domain)

	 if err != nil{
		fmt.Printf("Error: %v", err)
	 }

	 if len(mxRecord) >0 {
		hasMX = true
	 }
	 txtRecord, erro:= net.LookupTXT(domain)
	 if erro != nil {
		fmt.Printf("Error: %v", err)
	 }
	 for _, records:= range txtRecord {
		if strings.HasPrefix(records, "v=spf1"){
			hasSPF = true
			spreRecord = records
			break
		}
	 }

	dmRecords, err1:= net.LookupTXT("_dmarc."+domain)

	if err1 != nil{
		fmt.Printf("Error: %v", err)
	}
	for _, record:= range dmRecords {
		if strings.HasPrefix(record, "DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v",domain, hasMX, hasSPF, spreRecord, hasDMARC, dmarcRecord)

}