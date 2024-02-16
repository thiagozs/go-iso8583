# ISO8583 Message Parser and Builder

This project provides a Go package for parsing and building ISO8583 financial transaction messages. It's designed to make working with ISO8583 messages straightforward by abstracting the complexities involved in encoding and decoding various field types.

* **Working in progress**

## Installation

To use this package in your Go project, you can install it by running:

```bash
go get -u github.com/thiagozs/go-iso8583
```

## Usage

### Parsing an ISO8583 Message

To parse an ISO8583 message received over a network (e.g., TCP data), you can use the `Parse` method from the `Parser` struct:

```go
package main

import (
	"fmt"
	"github.com/thiagozs/go-iso8583"
)

func main() {
	dataTcp := "Your_ISO8583_Message_Here"

	parser := iso8583.NewParser()
	parsedMsg, err := parser.Parse(dataTcp)
	if err != nil {
		fmt.Printf("Error parsing ISO8583 message: %v\n", err)
		return
	}

	parsedMsg.LogFields()
}
```

### Building an ISO8583 Message

To build an ISO8583 message, you can use the `MessageBuilder` struct and its methods to set the MTI, add fields, and then build the message:

```go
msg := parser.CreateISO("0200")
msg.AddField(2, "4000001234567890")
msg.AddField(3, "000000")
// Add more fields as needed...

isoMsg, err := msg.Build()
if err != nil {
	fmt.Printf("Error building ISO8583 message: %v\n", err)
	return
}

fmt.Println("ISO8583 Message:", isoMsg)
```

### Example

The following example demonstrates parsing an ISO8583 message, logging its fields, and then building a new ISO8583 message:

```go
func main() {
	// Parse example ISO8583 message
	dataTcp := "08002038000000200002810000000001084909052253415630305A303537363331202020205341564E47583130303131303032303030302020202020200011010008B9F3F723CA3CD2F8"
	parser := iso8583.NewParser()

	fmt.Println("Data TCP =", dataTcp)
	_, err := parser.Parse(dataTcp)
	if err != nil {
		fmt.Printf("Error parsing ISO8583 message: %v\n", err)
		return
	}

	parser.LogFields()

	// Build a new ISO8583 message
	msg := parser.CreateISO("0200")
	msg.AddField(2, "4000001234567890")
	// Add more fields as needed...

	iso, err := msg.Build()
	if err != nil {
		fmt.Printf("Error building ISO8583 message: %v\n", err)
		return
	}

	fmt.Println("ISO Generated =", iso)
}
```

Replace `Your_ISO8583_Message_Here` with an actual ISO8583 message string you want to parse. Make sure to adjust field numbers and values according to the specifications you are working with.

## Contributing

We welcome contributions to this project! If you have improvements or bug fixes, please open a pull request or an issue.

-----

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-iso8583/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2024**, thiagozs
