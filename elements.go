package iso8583

// LenType represents the length type of an ISO 8583 element.
type LenType int

// List of length types.
const (
	Fixed LenType = iota
	LLVAR
	LLLVAR
)

// Element represents an ISO 8583 element.
type Element struct {
	ContentType string
	Label       string
	LenType     LenType
	MaxLen      int
	MinLen      int
}

// String returns the string representation of the length type.
func (l LenType) String() string {
	return [...]string{"Fixed", "LLVAR", "LLLVAR"}[l]
}

// NewElement initializes a new ISO 8583 element.
// have a chance to customize the elements for your own use case
// need to be careful when changing the elements and change
// the parser and builder accordingly
var dataElem = map[int]Element{
	0:   {ContentType: "n", Label: "Message Type Indicator", LenType: Fixed, MaxLen: 4},
	1:   {ContentType: "b", Label: "Bitmap", LenType: Fixed, MaxLen: 8},
	2:   {ContentType: "n", Label: "Primary account number (PAN)", LenType: LLVAR, MaxLen: 19, MinLen: 12},
	3:   {ContentType: "n", Label: "Processing code", LenType: Fixed, MaxLen: 6},
	4:   {ContentType: "n", Label: "Amount, transaction", LenType: Fixed, MaxLen: 12},
	5:   {ContentType: "n", Label: "Amount, settlement", LenType: Fixed, MaxLen: 12},
	6:   {ContentType: "n", Label: "Amount, cardholder billing", LenType: Fixed, MaxLen: 12},
	7:   {ContentType: "n", Label: "Transmission date & time", LenType: Fixed, MaxLen: 10},
	8:   {ContentType: "n", Label: "Amount, cardholder billing fee", LenType: Fixed, MaxLen: 8},
	9:   {ContentType: "n", Label: "Conversion rate, settlement", LenType: Fixed, MaxLen: 8},
	10:  {ContentType: "n", Label: "Conversion rate, cardholder billing", LenType: Fixed, MaxLen: 8},
	11:  {ContentType: "n", Label: "System trace audit number", LenType: Fixed, MaxLen: 6},
	12:  {ContentType: "n", Label: "Time, local transaction (hhmmss)", LenType: Fixed, MaxLen: 6},
	13:  {ContentType: "n", Label: "Date, local transaction (MMDD)", LenType: Fixed, MaxLen: 4},
	14:  {ContentType: "n", Label: "Date, expiration", LenType: Fixed, MaxLen: 4},
	15:  {ContentType: "n", Label: "Date, settlement", LenType: Fixed, MaxLen: 4},
	16:  {ContentType: "n", Label: "Date, conversion", LenType: Fixed, MaxLen: 4},
	17:  {ContentType: "n", Label: "Date, capture", LenType: Fixed, MaxLen: 4},
	18:  {ContentType: "n", Label: "Merchant type", LenType: Fixed, MaxLen: 4},
	19:  {ContentType: "n", Label: "Acquiring institution country code", LenType: Fixed, MaxLen: 3},
	20:  {ContentType: "n", Label: "PAN extended, country code", LenType: Fixed, MaxLen: 3},
	21:  {ContentType: "n", Label: "Forwarding institution. country code", LenType: Fixed, MaxLen: 3},
	22:  {ContentType: "n", Label: "Point of service entry mode", LenType: Fixed, MaxLen: 3},
	23:  {ContentType: "n", Label: "Application PAN sequence number", LenType: Fixed, MaxLen: 3},
	24:  {ContentType: "n", Label: "Network International identifier (NII)", LenType: Fixed, MaxLen: 3},
	25:  {ContentType: "n", Label: "Point of service condition code", LenType: Fixed, MaxLen: 2},
	26:  {ContentType: "n", Label: "Point of service capture code", LenType: Fixed, MaxLen: 2},
	27:  {ContentType: "n", Label: "Authorizing identification response length", LenType: Fixed, MaxLen: 1},
	28:  {ContentType: "an", Label: "Amount, transaction fee", LenType: Fixed, MaxLen: 9},
	29:  {ContentType: "an", Label: "Amount, settlement fee", LenType: Fixed, MaxLen: 9},
	30:  {ContentType: "an", Label: "Amount, transaction processing fee", LenType: Fixed, MaxLen: 9},
	31:  {ContentType: "an", Label: "Amount, settlement processing fee", LenType: Fixed, MaxLen: 9},
	32:  {ContentType: "n", Label: "Acquiring institution identification code", LenType: LLVAR, MaxLen: 11},
	33:  {ContentType: "n", Label: "Forwarding institution identification code", LenType: LLVAR, MaxLen: 11},
	34:  {ContentType: "ns", Label: "Primary account number, extended", LenType: LLVAR, MaxLen: 28},
	35:  {ContentType: "z", Label: "Track 2 data", LenType: LLVAR, MaxLen: 37},
	36:  {ContentType: "n", Label: "Track 3 data", LenType: LLLVAR, MaxLen: 104},
	37:  {ContentType: "an", Label: "Retrieval reference number", LenType: Fixed, MaxLen: 12},
	38:  {ContentType: "an", Label: "Authorization identification response", LenType: Fixed, MaxLen: 6},
	39:  {ContentType: "an", Label: "Response code", LenType: Fixed, MaxLen: 2},
	40:  {ContentType: "an", Label: "Service restriction code", LenType: Fixed, MaxLen: 3},
	41:  {ContentType: "ans", Label: "Card acceptor terminal identification", LenType: Fixed, MaxLen: 8},
	42:  {ContentType: "ans", Label: "Card acceptor identification code", LenType: Fixed, MaxLen: 15},
	43:  {ContentType: "ans", Label: "Card acceptor name/location", LenType: Fixed, MaxLen: 40},
	44:  {ContentType: "an", Label: "Additional response data", LenType: LLVAR, MaxLen: 25},
	45:  {ContentType: "an", Label: "Track 1 data", LenType: LLVAR, MaxLen: 76},
	46:  {ContentType: "an", Label: "Additional data - ISO", LenType: LLLVAR, MaxLen: 999},
	47:  {ContentType: "an", Label: "Additional data - national", LenType: LLLVAR, MaxLen: 999},
	48:  {ContentType: "an", Label: "Additional data - private", LenType: LLLVAR, MaxLen: 999},
	49:  {ContentType: "an", Label: "Currency code, transaction", LenType: Fixed, MaxLen: 3},
	50:  {ContentType: "an", Label: "Currency code, settlement", LenType: Fixed, MaxLen: 3},
	51:  {ContentType: "an", Label: "Currency code, cardholder billing", LenType: Fixed, MaxLen: 3},
	52:  {ContentType: "b", Label: "Personal identification number data", LenType: Fixed, MaxLen: 8},
	53:  {ContentType: "n", Label: "Security related control information", LenType: Fixed, MaxLen: 16},
	54:  {ContentType: "an", Label: "Additional amounts", LenType: LLLVAR, MaxLen: 120},
	55:  {ContentType: "ans", Label: "Reserved ISO", LenType: LLLVAR, MaxLen: 999},
	56:  {ContentType: "ans", Label: "Reserved ISO", LenType: LLLVAR, MaxLen: 999},
	57:  {ContentType: "ans", Label: "Reserved national", LenType: LLLVAR, MaxLen: 999},
	58:  {ContentType: "ans", Label: "Reserved national", LenType: LLLVAR, MaxLen: 999},
	59:  {ContentType: "ans", Label: "Reserved national", LenType: LLLVAR, MaxLen: 999},
	60:  {ContentType: "ans", Label: "Reserved national", LenType: LLLVAR, MaxLen: 999},
	61:  {ContentType: "ans", Label: "Reserved private", LenType: LLLVAR, MaxLen: 999},
	62:  {ContentType: "ans", Label: "Reserved private", LenType: LLLVAR, MaxLen: 999},
	63:  {ContentType: "ans", Label: "Reserved private", LenType: LLLVAR, MaxLen: 999},
	64:  {ContentType: "b", Label: "Message authentication code (MAC)", LenType: Fixed, MaxLen: 8},
	65:  {ContentType: "b", Label: "Bitmap, extended", LenType: Fixed, MaxLen: 1},
	66:  {ContentType: "n", Label: "Settlement code", LenType: Fixed, MaxLen: 1},
	67:  {ContentType: "n", Label: "Extended payment code", LenType: Fixed, MaxLen: 2},
	68:  {ContentType: "n", Label: "Receiving institution country code", LenType: Fixed, MaxLen: 3},
	69:  {ContentType: "n", Label: "Settlement institution country code", LenType: Fixed, MaxLen: 3},
	70:  {ContentType: "n", Label: "Network management information code", LenType: Fixed, MaxLen: 3},
	71:  {ContentType: "n", Label: "Message number", LenType: Fixed, MaxLen: 4},
	72:  {ContentType: "n", Label: "Message number, last", LenType: Fixed, MaxLen: 4},
	73:  {ContentType: "n", Label: "Date, action (YYMMDD)", LenType: Fixed, MaxLen: 6},
	74:  {ContentType: "n", Label: "Credits, number", LenType: Fixed, MaxLen: 10},
	75:  {ContentType: "n", Label: "Credits, reversal number", LenType: Fixed, MaxLen: 10},
	76:  {ContentType: "n", Label: "Debits, number", LenType: Fixed, MaxLen: 10},
	77:  {ContentType: "n", Label: "Debits, reversal number", LenType: Fixed, MaxLen: 10},
	78:  {ContentType: "n", Label: "Transfer number", LenType: Fixed, MaxLen: 10},
	79:  {ContentType: "n", Label: "Transfer, reversal number", LenType: Fixed, MaxLen: 10},
	80:  {ContentType: "n", Label: "Inquiries number", LenType: Fixed, MaxLen: 10},
	81:  {ContentType: "n", Label: "Authorizations, number", LenType: Fixed, MaxLen: 10},
	82:  {ContentType: "n", Label: "Credits, processing fee amount", LenType: Fixed, MaxLen: 12},
	83:  {ContentType: "n", Label: "Credits, transaction fee amount", LenType: Fixed, MaxLen: 12},
	84:  {ContentType: "n", Label: "Debits, processing fee amount", LenType: Fixed, MaxLen: 12},
	85:  {ContentType: "n", Label: "Debits, transaction fee amount", LenType: Fixed, MaxLen: 12},
	86:  {ContentType: "n", Label: "Credits, amount", LenType: Fixed, MaxLen: 16},
	87:  {ContentType: "n", Label: "Credits, reversal amount", LenType: Fixed, MaxLen: 16},
	88:  {ContentType: "n", Label: "Debits, amount", LenType: Fixed, MaxLen: 16},
	89:  {ContentType: "n", Label: "Debits, reversal amount", LenType: Fixed, MaxLen: 16},
	90:  {ContentType: "n", Label: "Original data elements", LenType: Fixed, MaxLen: 42},
	91:  {ContentType: "an", Label: "File update code", LenType: Fixed, MaxLen: 1},
	92:  {ContentType: "an", Label: "File security code", LenType: Fixed, MaxLen: 2},
	93:  {ContentType: "an", Label: "Response indicator", LenType: Fixed, MaxLen: 5},
	94:  {ContentType: "an", Label: "Service indicator", LenType: Fixed, MaxLen: 7},
	95:  {ContentType: "an", Label: "Replacement amounts", LenType: Fixed, MaxLen: 42},
	96:  {ContentType: "b", Label: "Message security code", LenType: Fixed, MaxLen: 8},
	97:  {ContentType: "an", Label: "Amount, net settlement", LenType: Fixed, MaxLen: 17},
	98:  {ContentType: "ans", Label: "Payee", LenType: Fixed, MaxLen: 25},
	99:  {ContentType: "n", Label: "Settlement institution identification code", LenType: LLVAR, MaxLen: 11},
	100: {ContentType: "n", Label: "Receiving institution identification code", LenType: LLVAR, MaxLen: 11},
	101: {ContentType: "ans", Label: "File name", LenType: LLVAR, MaxLen: 17},
	102: {ContentType: "ans", Label: "Account identification 1", LenType: LLVAR, MaxLen: 28},
	103: {ContentType: "ans", Label: "Account identification 2", LenType: LLVAR, MaxLen: 28},
	104: {ContentType: "ans", Label: "Transaction description", LenType: LLLVAR, MaxLen: 100},
	105: {ContentType: "ans", Label: "Reserved for ISO use", LenType: LLLVAR, MaxLen: 999},
	106: {ContentType: "ans", Label: "Reserved for ISO use", LenType: LLLVAR, MaxLen: 999},
	107: {ContentType: "ans", Label: "Reserved for ISO use", LenType: LLLVAR, MaxLen: 999},
	108: {ContentType: "ans", Label: "Reserved for ISO use", LenType: LLLVAR, MaxLen: 999},
	109: {ContentType: "ans", Label: "Reserved for ISO use", LenType: LLLVAR, MaxLen: 999},
	110: {ContentType: "ans", Label: "Reserved for ISO use", LenType: LLLVAR, MaxLen: 999},
	111: {ContentType: "ans", Label: "Reserved for ISO use", LenType: LLLVAR, MaxLen: 999},
	112: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	113: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	114: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	115: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	116: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	117: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	118: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	119: {ContentType: "ans", Label: "Reserved for national use", LenType: LLLVAR, MaxLen: 999},
	120: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	121: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	122: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	123: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	124: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	125: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	126: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	127: {ContentType: "ans", Label: "Reserved for private use", LenType: LLLVAR, MaxLen: 999},
	128: {ContentType: "b", Label: "Message authentication code", LenType: Fixed, MaxLen: 8},
}
