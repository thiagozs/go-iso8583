package main

import (
	"fmt"
	"iso8583/iso8583"
	"strings"
)

func main() {

	dataTcp := "08002038000000200002810000000001084909052253415630305A303537363331202020205341564E47583130303131303032303030302020202020200011010008B9F3F723CA3CD2F8"

	i := iso8583.New()

	fmt.Println("Data TCP =", dataTcp)

	_, err := i.Parse(dataTcp)
	if err != nil {
		fmt.Printf("Error Parse: %v\n", err)
		return
	}

	i.LogFields()

	fmt.Println()
	fmt.Println("Testing ISO8583 Builder")
	fmt.Printf("%s\n", strings.Repeat("-", 80))

	index := `MTI: 0200
Field 2 (Primary Account Number): 4000001234567890
Field 3 (Processing Code): 000000 (Assuming purchase)
Field 4 (Amount, Transaction): 000000006000 (60 USD, 12 digits, no decimal point)
Field 7 (Transmission Date & Time): 0209123456 (Assuming February 9th, 12:34:56)
Field 11 (Systems Trace Audit Number): 000001
Field 12 (Time, Local Transaction): 123456 (12:34:56)
Field 13 (Date, Local Transaction): 0209 (February 9th)
Field 14 (Date, Expiration): 2402 (February 2024)
Field 22 (Point of Service Entry Mode): 022 (Magnetic stripe, including PIN capability)
Field 32 (Acquiring Institution Identification Code): 123456
Field 35 (Track 2 Data): 4000001234567890=240212345
Field 37 (Retrieval Reference Number): 123456789012
Field 41 (Card Acceptor Terminal ID): TERM1234
Field 42 (Card Acceptor ID Code): 123456789012345
Field 49 (Currency Code, Transaction): 840 (USD)
Field 54 (Additional Amounts): 0400600D000000000000 (4 installments of 60 USD)
	`
	fmt.Println(index)

	msg := i.CreateISO("0200")
	msg.AddField(2, "4000001234567890")
	msg.AddField(3, "000000")
	msg.AddField(4, "000000006000")
	msg.AddField(7, "0209123456")
	msg.AddField(11, "000001")
	msg.AddField(12, "123456")
	msg.AddField(13, "0209")
	msg.AddField(14, "2402")
	msg.AddField(22, "022")
	msg.AddField(32, "123456")
	msg.AddField(35, "4000001234567890=240212345")
	msg.AddField(37, "123456789012")
	msg.AddField(41, "TERM1234")
	msg.AddField(42, "123456789012345")
	msg.AddField(49, "840")
	msg.AddField(54, "0400600D000000000000")

	iso, err := msg.Build()
	if err != nil {
		fmt.Printf("Error Build: %v\n", err)
		return
	}

	fmt.Println("ISO Generated =", iso)

	_, err = i.Parse(iso)
	if err != nil {
		fmt.Printf("Error Parse: %v\n", err)
		return
	}

	i.LogFields()

	fmt.Println()
	fmt.Printf("Testing2 ISO8583 Builder, adding fields with example \nvalues that correspond to typical ISO 8583 usage\n")
	fmt.Printf("%s\n", strings.Repeat("-", 80))

	// Adding fields with example values that correspond to
	// typical ISO 8583 usage
	msg = i.CreateISO("0200")
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
