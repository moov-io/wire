package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

// mockOriginatorOptionF creates a OriginatorOptionF
func mockOriginatorOptionF() *OriginatorOptionF {
	oof := NewOriginatorOptionF()
	oof.PartyIdentifier = "TXID/123-45-6789"
	oof.Name = "1/Name"
	oof.LineOne = "1/1234"
	oof.LineTwo = "2/1000 Colonial Farm Rd"
	oof.LineThree = "5/Pottstown"
	return oof
}

// TestMockOriginatorOptionF validates mockOriginatorOptionF
func TestMockOriginatorOptionF(t *testing.T) {
	oof := mockOriginatorOptionF()

	require.NoError(t, oof.Validate(), "mockOriginatorOptionF does not validate and will break other tests")
}

// TestOriginatorOptionFPartyIdentifier validates OriginatorOptionF PartyIdentifier is valid
func TestOriginatorOptionFPartyIdentifier(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "®®sdaasd"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("PartyIdentifier", ErrPartyIdentifier, oof.PartyIdentifier).Error())
}

// TestOriginatorOptionFPartyIdentifierNull validates OriginatorOptionF PartyIdentifier is not null
func TestOriginatorOptionFPartyIdentifierNull(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = ""

	err := oof.Validate()

	require.EqualError(t, err, fieldError("PartyIdentifier", ErrPartyIdentifier, oof.PartyIdentifier).Error())
}

// TestOriginatorOptionFPartyIdentifierCount validates OriginatorOptionF PartyIdentifier count is > 2
func TestOriginatorOptionFPartyIdentifierCount(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "X"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("PartyIdentifier", ErrPartyIdentifier, oof.PartyIdentifier).Error())
}

func TestOriginatorOptionF_PartyIdentifierField(t *testing.T) {
	tests := []struct {
		desc       string
		identifier string
	}{
		{"invalid empty string in piece 6", "CCPT/ BDF"},
		{"invalid empty string in piece 2 for an Account", "/ 1"},
		{"unique ID is valid", "ZZZZFGH/"},
		{"'/' is not in the correct spot", "CCPTFGH/"},
		{"unique ID has the correct count", "B/C"},
	}

	for _, test := range tests {
		oof := mockOriginatorOptionF()
		oof.PartyIdentifier = test.identifier

		assert.EqualErrorf(t, oof.Validate(), fieldError("PartyIdentifier", ErrPartyIdentifier,
			test.identifier).Error(), "Test PartyIdentifier: %s", test.desc)
	}
}

// TestOriginatorOptionFName validates OriginatorOptionF Name is valid
func TestOriginatorOptionFName(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.Name = "®"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("Name", ErrOptionFName, oof.Name).Error())
}

// TestOriginatorOptionFNameNull validates OriginatorOptionF Name is not null
func TestOriginatorOptionFNameNull(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.Name = ""

	err := oof.Validate()

	require.EqualError(t, err, fieldError("Name", ErrOptionFName, oof.Name).Error())
}

// TestOriginatorOptionFLineValid validates OriginatorOptionF Line* is valid
func TestOriginatorOptionFLineValid(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineOne = "Z/123"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrOptionFLine, oof.LineOne).Error())
}

// TestOriginatorOptionFLineSlash validates OriginatorOptionF Line* slash is valid
func TestOriginatorOptionFLineSlash(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineOne = "11123"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrOptionFLine, oof.LineOne).Error())
}

// TestOriginatorOptionFLineOne validates OriginatorOptionF LineOne is valid
func TestOriginatorOptionFLineOne(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineOne = "®"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrOptionFLine, oof.LineOne).Error())
}

// TestOriginatorOptionFLineTwo validates OriginatorOptionF LineTwo is valid
func TestOriginatorOptionFLineTwo(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineTwo = "®"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrOptionFLine, oof.LineTwo).Error())
}

// TestOriginatorOptionFLineThree validates OriginatorOptionF LineThree is valid
func TestOriginatorOptionFLineThree(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineThree = "®"

	err := oof.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrOptionFLine, oof.LineThree).Error())
}

// TestParseOriginatorOptionFWrongLength parses a wrong OriginatorOptionF record length
func TestParseOriginatorOptionFWrongLength(t *testing.T) {
	var line = "{5010}TXID/123-45-6789                   Name                               LineOne                            LineTwo                            LineThree                        "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorOptionF()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(181, len(r.line))).Error())
}

// TestParseOriginatorOptionFReaderParseError parses a wrong OriginatorOptionF reader parse error
func TestParseOriginatorOptionFReaderParseError(t *testing.T) {
	var line = "{5010}TXID/123-45-6789                   ®ame                               LineOne                            LineTwo                            LineThree                          "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorOptionF()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrOptionFName, "®ame")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrOptionFName, "®ame")).Error())
}
