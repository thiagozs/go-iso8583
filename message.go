package iso8583

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// LenType represents the length type of an ISO 8583 field.
type Parser struct {
	MTI          string
	Bitmap       string
	Fields       map[int]string
	ActiveFields []int
	HasSecBitmap bool
	LastField    int
}

// NewParser initializes a new ISO 8583 message parser.
func NewParser() *Parser {
	return &Parser{
		Fields: make(map[int]string),
	}
}

// Parse decodes an ISO 8583 message.
func (m *Parser) Parse(raw string) (*Parser, error) {
	if len(raw) < 4 {
		return m, fmt.Errorf("raw data too short to contain MTI")
	}

	// Reset fields if they are not empty
	if !m.fieldsAreEmpty() {
		m.resetFields()
	}

	// Parse the MTI and bitmap
	m.MTI = raw[:4]

	// Parse the bitmap and fields
	if err := m.ParseBitmap(raw); err != nil {
		return m, err
	}

	// Parse the fields
	if m.HasSecBitmap {
		return m, m.ParseFields(raw[36:])
	}

	return m, m.ParseFields(raw[20:])
}

// ParseBitmap converts a hexadecimal bitmap string to a binary string and identifies active fields.
func (m *Parser) ParseBitmap(rawBitmap string) error {
	m.Bitmap = ""
	bitmapHex := rawBitmap[4:20] // Primary bitmap
	fmt.Println("Bitmap1: ", bitmapHex)

	bitmap, err := hex.DecodeString(bitmapHex)
	if err != nil {
		return fmt.Errorf("failed to decode primary bitmap: %v", err)
	}

	for _, b := range bitmap {
		m.Bitmap += fmt.Sprintf("%08b", b)
	}

	if m.Bitmap[0] == '1' {
		// Secondary bitmap is present
		m.HasSecBitmap = true
		secondaryBitmapHex := rawBitmap[20:36]
		fmt.Println("Bitmap2: ", secondaryBitmapHex)

		secondaryBitmap, err := hex.DecodeString(secondaryBitmapHex)
		if err != nil {
			return fmt.Errorf("failed to decode secondary bitmap: %v", err)
		}

		for _, b := range secondaryBitmap {
			m.Bitmap += fmt.Sprintf("%08b", b)
		}
	}

	for i, bit := range m.Bitmap {
		if bit == '1' {
			m.ActiveFields = append(m.ActiveFields, i+1)
		}
	}

	m.LastField = m.ActiveFields[len(m.ActiveFields)-1]

	return nil
}

// ParseFields parses all fields indicated by the bitmap.
func (m *Parser) ParseFields(rawData string) error {

	fmt.Println("ActiveFields: ", m.ActiveFields)

	for _, fieldNum := range m.ActiveFields {
		if fieldNum == 1 {
			continue // Skip the bitmap field itself
		}

		elem, exists := dataElem[fieldNum]
		if !exists {
			return fmt.Errorf("no parser for field %d", fieldNum)
		}

		var fieldValue, remaining string
		var err error

		switch elem.LenType {
		case Fixed:
			if len(rawData) < elem.MaxLen {
				return fmt.Errorf("not enough data for fixed-length field %d", fieldNum)
			}
			fieldValue = rawData[:elem.MaxLen]
			remaining = rawData[elem.MaxLen:]

		case LLVAR:
			fieldValue, remaining, err = m.parseLLVAR(rawData, fieldNum)
			if err != nil {
				return fmt.Errorf("failed to parse LLVAR field %d: %v", fieldNum, err)
			}

		case LLLVAR:
			fieldValue, remaining, err = m.parseLLLVAR(rawData, fieldNum)
			if err != nil {
				return fmt.Errorf("failed to parse LLLVAR field %d: %v", fieldNum, err)
			}
		}

		// Update the field value in the message
		m.Fields[fieldNum] = fieldValue

		// Update rawData to the remaining part for the next field parsing
		rawData = remaining
		fmt.Printf("Element%d - %v\n", fieldNum, elem)
		fmt.Printf("Field%d - %s\n", fieldNum, fieldValue)
		fmt.Println("Remaining for next:", remaining)
		fmt.Println("-----------------")
	}

	return nil
}

func (m *Parser) parseLLVAR(input string, fieldNum int) (value string, remaining string, err error) {
	if len(input) < 2 {
		return "", "", fmt.Errorf("input too short for LLVAR length indicator")
	}

	lengthIndicator := input[:2]
	length, err := strconv.Atoi(lengthIndicator)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse LLVAR length indicator: %v", err)
	}

	fmt.Println("parseLLVAR Length: ", length)

	// Adjust length if it exceeds the input's remaining length
	if length > len(input)-2 {
		length = len(input) - 2
	}

	value = input[2 : 2+length]
	remaining = input[2+length:]

	// If it's the last field or the remaining string is very short, capture the rest
	if fieldNum == m.LastField || len(remaining) < 2 {
		value = input[2:]
		remaining = ""
	}

	return value, remaining, nil
}

func (m *Parser) parseLLLVAR(input string, fieldNum int) (value string, remaining string, err error) {
	if len(input) < 3 {
		return "", "", fmt.Errorf("input too short for LLLVAR length indicator")
	}

	lengthIndicator := input[:3]
	length, err := strconv.Atoi(lengthIndicator)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse LLLVAR length indicator: %v", err)
	}

	fmt.Println("parseLLLVAR Length: ", length)

	// Adjust length if it exceeds the input's remaining length
	if length > len(input)-3 {
		length = len(input) - 3
	}

	value = input[3 : 3+length]
	remaining = input[3+length:]

	// If it's the last field or the remaining string is very short, capture the rest
	if fieldNum == m.LastField || len(remaining) < 3 {
		value = input[3:]
		remaining = ""
	}

	return value, remaining, nil
}

func (m *Parser) resetFields() {
	m.MTI = ""
	m.Bitmap = ""
	m.Fields = make(map[int]string)
	m.ActiveFields = []int{}
	m.HasSecBitmap = false
	m.LastField = 0
}

func (m *Parser) fieldsAreEmpty() bool {
	return m.MTI == "" && m.Bitmap == "" && len(m.Fields) == 0 && len(m.ActiveFields) == 0 && !m.HasSecBitmap
}

func (m *Parser) LogFields() {
	fmt.Printf("MTI: %s\n", m.MTI)
	for _, fieldNum := range m.ActiveFields {
		value, ok := m.Fields[fieldNum]
		if ok {
			fmt.Printf("Field %d: %s\n", fieldNum, value)
		}
	}
}
