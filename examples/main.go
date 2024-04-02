package main

import (
	"fmt"
	"iso8583"
)

func main() {

	i := iso8583.New()
	// Adding fields with example values that correspond to
	// typical ISO 8583 usage

	msg := i.CreateISO("0200")

	msg.AddField(2, "1234567890123456")      // Primary Account Number
	msg.AddField(3, "000000")                // Processing Code
	msg.AddField(4, "000000010000")          // Amount, Transaction
	msg.AddField(11, "000001")               // Systems Trace Audit Number
	msg.AddField(12, "235959")               // Time, Local Transaction
	msg.AddField(17, "0225")                 // Date, Capture
	msg.AddField(18, "5999")                 // Merchant Type
	msg.AddField(19, "840")                  // Acquiring Institution Country Code
	msg.AddField(20, "840")                  // PAN Extended, Country Code
	msg.AddField(21, "840")                  // Forwarding Institution Country Code
	msg.AddField(22, "012")                  // Point of Service Entry Mode
	msg.AddField(23, "001")                  // Application PAN Sequence Number
	msg.AddField(25, "00")                   // Point of Service Condition Code
	msg.AddField(26, "12")                   // Point of Service Capture Code
	msg.AddField(27, "1")                    // Authorizing Identification Response Length
	msg.AddField(28, "C00000000")            // Amount, Transaction Fee
	msg.AddField(29, "D00000000")            // Amount, Settlement Fee
	msg.AddField(30, "D00000000")            // Amount, Transaction Processing Fee
	msg.AddField(31, "C00000000")            // Amount, Settlement Processing Fee
	msg.AddField(32, "123456")               // Acquiring Institution Identification Code
	msg.AddField(34, "12345678901234567890") // Primary Account Number, Extended
	msg.AddField(38, "123456")               // Authorization Identification Response
	msg.AddField(48, "XPTO")                 // Additional Data - Private
	msg.AddField(49, "840")                  // Currency Code, Transaction
	msg.AddField(51, "840")                  // Currency Code, Cardholder Billing
	msg.AddField(54, "CSHB840D0000002000")   // Cashback amount              // Currency Code, Cardholder Billing
	msg.AddField(53, "2600000000000000")     // Security Related Control Information
	msg.AddField(57, "2500")                 // Amount, Cash
	msg.AddField(58, "1234")                 // Authorizing Agent Institution ID
	msg.AddField(59, "XXXXXXYYYYYYYYYZ")     // Echo Data
	msg.AddField(64, "IFFFXPTO")             // Network Management Information Code

	msg.AddField(65, "IFFFXPTO") // Network Management Information Code
	msg.AddField(66, "IFFFXPTO") // Network Management Information Code
	msg.AddField(67, "IFFFXPTO") // Network Management Information Code

	isoMessage, err := msg.Build()
	if err != nil {
		fmt.Printf("Error building ISO message: %v\n", err)
		return
	}

	fmt.Printf("ISO Message: %s\n", isoMessage)

	_, err = i.Parse(isoMessage)
	if err != nil {
		fmt.Printf("Error Parse: %v\n", err)
		return
	}

}
