package bigendian

import (
	"encoding/binary"
	"testing"
)

func TestBigEndianUint64(t *testing.T) {

	// convert ot bytes
	input := uint64(2984983220)
	inputBytes := ToUint64(input)

	// convert from bytes back
	result := FromUint64(inputBytes)
	if result != input {
		t.Fatal("Big endian conversion failed")
	}

	goResult := binary.BigEndian.Uint64(inputBytes[:])

	if goResult != input {
		t.Fatal("It's not a big endian representation")
	}

	input = uint64(18446744073709551615)
	inputBytes = ToUint64(input)

	// convert from bytes back
	result = FromUint64(inputBytes)
	if result != input {
		t.Fatal("Big endian conversion failed")
	}

	goResult = binary.BigEndian.Uint64(inputBytes[:])

	if goResult != input {
		t.Fatal("It's not a big endian representation")
	}

}

func TestBigEndianInt(t *testing.T) {

	// convert ot bytes
	input := int(2984983220)
	inputBytes := ToInt(input)

	// convert from bytes back
	result := FromInt(inputBytes)
	if result != input {
		t.Fatal("Big endian conversion failed")
	}

	goResult := binary.BigEndian.Uint32(inputBytes[:])

	if int(goResult) != input {
		t.Fatal("It's not a big endian representation")
	}

}
