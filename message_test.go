package iso8583

import (
	"testing"
)

func TestParse(t *testing.T) {
	rawMessage := "08002038000000200002810000000001084909052253415630305A303537363331202020205341564E47583130303131303032303030302020202020200011010008B9F3F723CA3CD2F8"

	parser := NewParser()
	parsedMessage, err := parser.Parse(rawMessage)
	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}

	// Check the MTI
	if parsedMessage.MTI != "0800" {
		t.Errorf("Expected MTI = 0800, got %s", parsedMessage.MTI)
	}

	// Define a map of expected field values for comparison
	expectedFields := map[int]string{
		3:  "810000",
		11: "000001",
		12: "084909",
		13: "0522",
		43: "53415630305A303537363331202020205341564E",
		63: "83130303131303032303030302020202020200011010008B9F3F723CA3CD2F8",
	}

	for fieldNum, expectedValue := range expectedFields {
		if value, ok := parsedMessage.Fields[fieldNum]; ok {
			if value != expectedValue {
				t.Errorf("Field %d: expected %s, got %s", fieldNum, expectedValue, value)
			}
		} else {
			t.Errorf("Field %d: expected %s, but was not present", fieldNum, expectedValue)
		}
	}
}

func TestCreateISO(t *testing.T) {
	msg := NewISO()

	// Set the MTI
	msg.SetMTI("0200")

	// Add fields
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
	msg.AddField(65, "0400600D000000000000")
	msg.AddField(66, "0400600D000000000000")
	msg.AddField(67, "0400600D000000000000")

	// Build the ISO message
	isoMessage, err := msg.Build()
	if err != nil {
		t.Errorf("Build() error = %v", err)
		return
	}

	// Expected ISO message
	expectedISO := "0200F23C040128C08400E00000000000000016400000123456789000000000000000600002091234560000011234560209240202206123456264000001234567890=240212345123456789012TERM12341234567890123458400200400600D0000000000000004"

	if isoMessage != expectedISO {
		t.Errorf("Expected ISO message = %s, got %s", expectedISO, isoMessage)
	}
}
