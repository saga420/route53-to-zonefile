package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func convertRoute53ToZoneFile(zoneID, output string) error {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		return err
	}

	svc := route53.New(sess)

	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	resp, err := svc.ListResourceRecordSets(params)
	if err != nil {
		return err
	}

	var outputWriter *os.File
	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			return err
		}
		defer file.Close()
		outputWriter = file
	} else {
		outputWriter = os.Stdout
	}

	for _, rrset := range resp.ResourceRecordSets {
		if *rrset.Type == "NS" || *rrset.Type == "SOA" {
			continue
		}

		fmt.Fprintf(outputWriter, "%s\t%d\t%s\t%s\n", *rrset.Name, *rrset.TTL, *rrset.Type, *rrset.ResourceRecords[0].Value)
	}

	return nil
}

func main() {
	zoneID := flag.String("zone-id", "", "The hosted zone ID")
	output := flag.String("output", "", "Output file name")

	flag.Parse()

	if *zoneID == "" {
		fmt.Println("Please provide the hosted zone ID.")
		os.Exit(1)
	}

	err := convertRoute53ToZoneFile(*zoneID, *output)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
