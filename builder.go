package iso8583

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// MessageBuilder provides methods to construct
// an ISO 8583 message.
type MessageBuilder struct {
	MTI    string
	Fields map[int]string
}

// NewMessage initializes a new MessageBuilder with an MTI.
func NewISO() *MessageBuilder {
	return &MessageBuilder{
		Fields: make(map[int]string),
	}
}

// SetMTI sets the Message Type Indicator (MTI) for the ISO 8583 message.
func (mb *MessageBuilder) SetMTI(mti string) *MessageBuilder {
	mb.MTI = mti
	return mb
}

// AddField adds or updates a field in the ISO 8583 message.
func (mb *MessageBuilder) AddField(fieldNum int, value string) *MessageBuilder {
	mb.Fields[fieldNum] = value
	return mb
}

// Build constructs the ISO 8583 message based on the MTI and fields.
func (mb *MessageBuilder) Build() (string, error) {
	if mb.MTI == "" {
		return "", fmt.Errorf("MTI is required")
	}

	// Initialize bitmaps
	primaryBitmap := make([]int, 64)
	secondaryBitmap := make([]int, 64)
	needSecondaryBitmap := false

	// Initialize the field builder
	var fields strings.Builder

	// Sort the field numbers
	fieldNumbers := make([]int, 0, len(mb.Fields))
	for fieldNum := range mb.Fields {
		fieldNumbers = append(fieldNumbers, fieldNum)
	}
	sort.Ints(fieldNumbers)

	// Construct the field values and update bitmap
	for _, fieldNum := range fieldNumbers {
		value := mb.Fields[fieldNum]

		if fieldNum >= 1 && fieldNum <= 64 {
			primaryBitmap[fieldNum-1] = 1
		} else if fieldNum > 64 && fieldNum <= 128 {
			secondaryBitmap[fieldNum-65] = 1
			needSecondaryBitmap = true
		} else {
			continue // Skip unsupported field numbers
		}

		elem, exists := dataElem[fieldNum]
		if !exists {
			return "", fmt.Errorf("unsupported field %d", fieldNum)
		}

		fieldValue, err := constructFieldValue(fieldNum, value, elem)
		if err != nil {
			return "", fmt.Errorf("error constructing field %d: %v", fieldNum, err)
		}

		fields.WriteString(fieldValue)
	}

	// If a secondary bitmap is needed, set the first bit of the primary bitmap
	if needSecondaryBitmap {
		primaryBitmap[0] = 1
	}

	// Convert bitmaps to hexadecimal strings
	hexPrimaryBitmap := bitmapToHex(primaryBitmap)
	var hexSecondaryBitmap string
	if needSecondaryBitmap {
		hexSecondaryBitmap = bitmapToHex(secondaryBitmap)
	}

	// Combine MTI, bitmaps, and fields
	finalMessage := mb.MTI + hexPrimaryBitmap + hexSecondaryBitmap + fields.String()

	return finalMessage, nil
}

// bitmapToHex converts a binary bitmap slice to a hexadecimal string.
func bitmapToHex(bitmap []int) string {
	var binaryBitmap strings.Builder
	for _, bit := range bitmap {
		binaryBitmap.WriteString(strconv.Itoa(bit))
	}

	// Convert binary string to hexadecimal
	return binaryStringToHex(binaryBitmap.String())
}

// binaryStringToHex converts a binary string to a hexadecimal string.
func binaryStringToHex(binary string) string {
	var hex strings.Builder
	for i := 0; i < len(binary); i += 4 {
		nibble := binary[i : i+4]
		val, _ := strconv.ParseUint(nibble, 2, 8)
		hex.WriteString(fmt.Sprintf("%X", val))
	}
	return hex.String()
}

// constructFieldValue formats the field value based on ISO 8583 standards (fixed, LLVAR, LLLVAR).
func constructFieldValue(fieldNum int, value string, elem Element) (string, error) {
	var fieldBuilder strings.Builder

	switch elem.LenType {
	case Fixed:
		paddedValue := padOrTruncate(value, elem.MaxLen, elem.ContentType)
		fieldBuilder.WriteString(paddedValue)

	case LLVAR:
		lengthIndicator := fmt.Sprintf("%02d", len(value))
		fieldBuilder.WriteString(lengthIndicator + value)

	case LLLVAR:
		lengthIndicator := fmt.Sprintf("%03d", len(value))
		fieldBuilder.WriteString(lengthIndicator + value)

	default:
		return "", fmt.Errorf("unsupported length type for field %d", fieldNum)
	}

	return fieldBuilder.String(), nil
}

// padOrTruncate ensures the value fits the specified length for fixed fields.
func padOrTruncate(value string, length int, contentType string) string {
	if len(value) > length {
		return value[:length] // Truncate if longer
	}

	// Pad if shorter
	switch contentType {
	case "n": // Numeric fields are zero-padded on the left
		return fmt.Sprintf("%0*s", length, value)
	case "an", "ans": // Alphanumeric fields are space-padded on the right
		return fmt.Sprintf("%-*s", length, value)
	case "z": // Tracks are space-padded on the right
		return fmt.Sprintf("%-*s", length, value)
	case "b": // Binary fields are zero-padded on the left
		return fmt.Sprintf("%0*s", length, value)
	default:
		return value
	}
}
