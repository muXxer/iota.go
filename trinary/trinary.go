// Package trinary provides functions for validating and converting Trits and Trytes.
package trinary

import (
	"bytes"
	"math"
	"strings"

	. "github.com/muxxer/iota.go/consts"
	iotaGoMath "github.com/muxxer/iota.go/math"
	"github.com/pkg/errors"
)

var (
	// TryteValueToTritsLUT is a lookup table to convert tryte values into trits.
	TryteValueToTritsLUT = [TryteRadix][TritsPerTryte]int8{
		{-1, -1, -1}, {0, -1, -1}, {1, -1, -1}, {-1, 0, -1}, {0, 0, -1}, {1, 0, -1},
		{-1, 1, -1}, {0, 1, -1}, {1, 1, -1}, {-1, -1, 0}, {0, -1, 0}, {1, -1, 0},
		{-1, 0, 0}, {0, 0, 0}, {1, 0, 0}, {-1, 1, 0}, {0, 1, 0}, {1, 1, 0},
		{-1, -1, 1}, {0, -1, 1}, {1, -1, 1}, {-1, 0, 1}, {0, 0, 1}, {1, 0, 1},
		{-1, 1, 1}, {0, 1, 1}, {1, 1, 1},
	}

	// TryteValueToTyteLUT is a lookup table to convert tryte values into trytes.
	TryteValueToTyteLUT = [TryteRadix]byte{'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M'}

	// TryteToTryteValueLUT is a lookup table to convert trytes into tryte values.
	TryteToTryteValueLUT = [...]int8{
		0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		-13, -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1,
	}

	// lookup table to unpack a byte into 5 trits.
	bytesToTritsLUT = [256][NumberOfTritsInAByte]int8{
		{0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {-1, 1, 0, 0, 0}, {0, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {-1, -1, 1, 0, 0},
		{0, -1, 1, 0, 0}, {1, -1, 1, 0, 0}, {-1, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {1, 0, 1, 0, 0}, {-1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0}, {1, 1, 1, 0, 0}, {-1, -1, -1, 1, 0}, {0, -1, -1, 1, 0}, {1, -1, -1, 1, 0}, {-1, 0, -1, 1, 0},
		{0, 0, -1, 1, 0}, {1, 0, -1, 1, 0}, {-1, 1, -1, 1, 0}, {0, 1, -1, 1, 0}, {1, 1, -1, 1, 0}, {-1, -1, 0, 1, 0},
		{0, -1, 0, 1, 0}, {1, -1, 0, 1, 0}, {-1, 0, 0, 1, 0}, {0, 0, 0, 1, 0}, {1, 0, 0, 1, 0}, {-1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0}, {1, 1, 0, 1, 0}, {-1, -1, 1, 1, 0}, {0, -1, 1, 1, 0}, {1, -1, 1, 1, 0}, {-1, 0, 1, 1, 0},
		{0, 0, 1, 1, 0}, {1, 0, 1, 1, 0}, {-1, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {1, 1, 1, 1, 0}, {-1, -1, -1, -1, 1},
		{0, -1, -1, -1, 1}, {1, -1, -1, -1, 1}, {-1, 0, -1, -1, 1}, {0, 0, -1, -1, 1}, {1, 0, -1, -1, 1}, {-1, 1, -1, -1, 1},
		{0, 1, -1, -1, 1}, {1, 1, -1, -1, 1}, {-1, -1, 0, -1, 1}, {0, -1, 0, -1, 1}, {1, -1, 0, -1, 1}, {-1, 0, 0, -1, 1},
		{0, 0, 0, -1, 1}, {1, 0, 0, -1, 1}, {-1, 1, 0, -1, 1}, {0, 1, 0, -1, 1}, {1, 1, 0, -1, 1}, {-1, -1, 1, -1, 1},
		{0, -1, 1, -1, 1}, {1, -1, 1, -1, 1}, {-1, 0, 1, -1, 1}, {0, 0, 1, -1, 1}, {1, 0, 1, -1, 1}, {-1, 1, 1, -1, 1},
		{0, 1, 1, -1, 1}, {1, 1, 1, -1, 1}, {-1, -1, -1, 0, 1}, {0, -1, -1, 0, 1}, {1, -1, -1, 0, 1}, {-1, 0, -1, 0, 1},
		{0, 0, -1, 0, 1}, {1, 0, -1, 0, 1}, {-1, 1, -1, 0, 1}, {0, 1, -1, 0, 1}, {1, 1, -1, 0, 1}, {-1, -1, 0, 0, 1},
		{0, -1, 0, 0, 1}, {1, -1, 0, 0, 1}, {-1, 0, 0, 0, 1}, {0, 0, 0, 0, 1}, {1, 0, 0, 0, 1}, {-1, 1, 0, 0, 1},
		{0, 1, 0, 0, 1}, {1, 1, 0, 0, 1}, {-1, -1, 1, 0, 1}, {0, -1, 1, 0, 1}, {1, -1, 1, 0, 1}, {-1, 0, 1, 0, 1},
		{0, 0, 1, 0, 1}, {1, 0, 1, 0, 1}, {-1, 1, 1, 0, 1}, {0, 1, 1, 0, 1}, {1, 1, 1, 0, 1}, {-1, -1, -1, 1, 1},
		{0, -1, -1, 1, 1}, {1, -1, -1, 1, 1}, {-1, 0, -1, 1, 1}, {0, 0, -1, 1, 1}, {1, 0, -1, 1, 1}, {-1, 1, -1, 1, 1},
		{0, 1, -1, 1, 1}, {1, 1, -1, 1, 1}, {-1, -1, 0, 1, 1}, {0, -1, 0, 1, 1}, {1, -1, 0, 1, 1}, {-1, 0, 0, 1, 1},
		{0, 0, 0, 1, 1}, {1, 0, 0, 1, 1}, {-1, 1, 0, 1, 1}, {0, 1, 0, 1, 1}, {1, 1, 0, 1, 1}, {-1, -1, 1, 1, 1},
		{0, -1, 1, 1, 1}, {1, -1, 1, 1, 1}, {-1, 0, 1, 1, 1}, {0, 0, 1, 1, 1}, {1, 0, 1, 1, 1}, {-1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {-1, -1, -1, -1, -1}, {0, -1, -1, -1, -1}, {1, -1, -1, -1, -1},
		{-1, 0, -1, -1, -1}, {0, 0, -1, -1, -1}, {1, 0, -1, -1, -1}, {-1, 1, -1, -1, -1}, {0, 1, -1, -1, -1}, {1, 1, -1, -1, -1},
		{-1, -1, 0, -1, -1}, {0, -1, 0, -1, -1}, {1, -1, 0, -1, -1}, {-1, 0, 0, -1, -1}, {0, 0, 0, -1, -1}, {1, 0, 0, -1, -1},
		{-1, 1, 0, -1, -1}, {0, 1, 0, -1, -1}, {1, 1, 0, -1, -1}, {-1, -1, 1, -1, -1}, {0, -1, 1, -1, -1}, {1, -1, 1, -1, -1},
		{-1, 0, 1, -1, -1}, {0, 0, 1, -1, -1}, {1, 0, 1, -1, -1}, {-1, 1, 1, -1, -1}, {0, 1, 1, -1, -1}, {1, 1, 1, -1, -1},
		{-1, -1, -1, 0, -1}, {0, -1, -1, 0, -1}, {1, -1, -1, 0, -1}, {-1, 0, -1, 0, -1}, {0, 0, -1, 0, -1}, {1, 0, -1, 0, -1},
		{-1, 1, -1, 0, -1}, {0, 1, -1, 0, -1}, {1, 1, -1, 0, -1}, {-1, -1, 0, 0, -1}, {0, -1, 0, 0, -1}, {1, -1, 0, 0, -1},
		{-1, 0, 0, 0, -1}, {0, 0, 0, 0, -1}, {1, 0, 0, 0, -1}, {-1, 1, 0, 0, -1}, {0, 1, 0, 0, -1}, {1, 1, 0, 0, -1},
		{-1, -1, 1, 0, -1}, {0, -1, 1, 0, -1}, {1, -1, 1, 0, -1}, {-1, 0, 1, 0, -1}, {0, 0, 1, 0, -1}, {1, 0, 1, 0, -1},
		{-1, 1, 1, 0, -1}, {0, 1, 1, 0, -1}, {1, 1, 1, 0, -1}, {-1, -1, -1, 1, -1}, {0, -1, -1, 1, -1}, {1, -1, -1, 1, -1},
		{-1, 0, -1, 1, -1}, {0, 0, -1, 1, -1}, {1, 0, -1, 1, -1}, {-1, 1, -1, 1, -1}, {0, 1, -1, 1, -1}, {1, 1, -1, 1, -1},
		{-1, -1, 0, 1, -1}, {0, -1, 0, 1, -1}, {1, -1, 0, 1, -1}, {-1, 0, 0, 1, -1}, {0, 0, 0, 1, -1}, {1, 0, 0, 1, -1},
		{-1, 1, 0, 1, -1}, {0, 1, 0, 1, -1}, {1, 1, 0, 1, -1}, {-1, -1, 1, 1, -1}, {0, -1, 1, 1, -1}, {1, -1, 1, 1, -1},
		{-1, 0, 1, 1, -1}, {0, 0, 1, 1, -1}, {1, 0, 1, 1, -1}, {-1, 1, 1, 1, -1}, {0, 1, 1, 1, -1}, {1, 1, 1, 1, -1},
		{-1, -1, -1, -1, 0}, {0, -1, -1, -1, 0}, {1, -1, -1, -1, 0}, {-1, 0, -1, -1, 0}, {0, 0, -1, -1, 0}, {1, 0, -1, -1, 0},
		{-1, 1, -1, -1, 0}, {0, 1, -1, -1, 0}, {1, 1, -1, -1, 0}, {-1, -1, 0, -1, 0}, {0, -1, 0, -1, 0}, {1, -1, 0, -1, 0},
		{-1, 0, 0, -1, 0}, {0, 0, 0, -1, 0}, {1, 0, 0, -1, 0}, {-1, 1, 0, -1, 0}, {0, 1, 0, -1, 0}, {1, 1, 0, -1, 0},
		{-1, -1, 1, -1, 0}, {0, -1, 1, -1, 0}, {1, -1, 1, -1, 0}, {-1, 0, 1, -1, 0}, {0, 0, 1, -1, 0}, {1, 0, 1, -1, 0},
		{-1, 1, 1, -1, 0}, {0, 1, 1, -1, 0}, {1, 1, 1, -1, 0}, {-1, -1, -1, 0, 0}, {0, -1, -1, 0, 0}, {1, -1, -1, 0, 0},
		{-1, 0, -1, 0, 0}, {0, 0, -1, 0, 0}, {1, 0, -1, 0, 0}, {-1, 1, -1, 0, 0}, {0, 1, -1, 0, 0}, {1, 1, -1, 0, 0},
		{-1, -1, 0, 0, 0}, {0, -1, 0, 0, 0}, {1, -1, 0, 0, 0}, {-1, 0, 0, 0, 0},
	}

	// Pow27LUT is a Look-up-table for Decoding Trits to int64
	Pow27LUT = []int64{1,
		27,
		729,
		19683,
		531441,
		14348907,
		387420489,
		10460353203,
		282429536481,
		7625597484987,
		205891132094649,
		5559060566555523,
		150094635296999136,
		4052555153018976256}

	encodedZero = []int8{1, 0, 0, -1}
)

// Trits is a slice of int8. You should not use cast, use NewTrits instead to ensure the validity.
type Trits = []int8

// Trytes is a string of trytes. Use NewTrytes() instead of typecasting.
type Trytes = string

// Hash represents a trinary hash
type Hash = Trytes

// Hashes is a slice of Hash.
type Hashes = []Hash

// ValidTrit returns true if t is a valid trit.
func ValidTrit(t int8) bool {
	return t >= -1 && t <= 1
}

// ValidTrits returns true if t is valid trits (non-empty and -1, 0 or 1).
func ValidTrits(trits Trits) error {
	if len(trits) == 0 {
		return errors.Wrap(ErrInvalidTrit, "trits slice is empty")
	}
	for i, trit := range trits {
		if !ValidTrit(trit) {
			return errors.Wrapf(ErrInvalidTrit, "at index %d", i)
		}
	}
	return nil
}

// NewTrits casts Trits and checks its validity.
func NewTrits(t []int8) (Trits, error) {
	err := ValidTrits(t)
	return t, err
}

// TritsEqual returns true if t and b are equal Trits.
func TritsEqual(a Trits, b Trits) (bool, error) {
	if err := ValidTrits(a); err != nil {
		return false, err
	}
	if err := ValidTrits(b); err != nil {
		return false, err
	}

	if len(a) != len(b) {
		return false, nil
	}

	for i := range a {
		if a[i] != b[i] {
			return false, nil
		}
	}
	return true, nil
}

// ReverseTrits reverses the given trits.
func ReverseTrits(trits Trits) Trits {
	for left, right := 0, len(trits)-1; left < right; left, right = left+1, right-1 {
		trits[left], trits[right] = trits[right], trits[left]
	}

	return trits
}

// TrailingZeros returns the number of trailing zeros of the given trits.
func TrailingZeros(trits Trits) int {
	var z int
	for i := len(trits) - 1; i >= 0 && trits[i] == 0; i-- {
		z++
	}
	return z
}

// roundUpToTryteMultiple rounds the given number up the the nearest multiple of 3 to make a valid tryte count.
func roundUpToTryteMultiple(n uint) uint {
	rem := n % TritsPerTryte
	if rem == 0 {
		return n
	}
	return n + TritsPerTryte - rem
}

// MinTrits returns the length of trits needed to encode the value.
func MinTrits(value int64) int {
	valueAbs := iotaGoMath.AbsInt64(value)

	var vp uint64
	var num int
	switch {
	case valueAbs >= 308836698141973:
		vp = 308836698141973
		num = 31
	case valueAbs >= 5230176601:
		vp = 5230176601
		num = 21
	case valueAbs >= 88573:
		vp = 88573
		num = 11
	default:
		vp = 1
		num = 1
	}

	for valueAbs > vp {
		vp = vp*TrinaryRadix + 1
		num++
	}
	return num
}

// IntToTrits converts int64 to a slice of trits.
func IntToTrits(value int64) Trits {
	numTrits := MinTrits(value)
	numTrytes := (numTrits + TritsPerTryte - 1) / TritsPerTryte
	trits := MustTrytesToTrits(IntToTrytes(value, numTrytes))
	return trits[:numTrits]
}

// TritsToInt converts a slice of trits into an integer and assumes little-endian notation.
func TritsToInt(t Trits) int64 {
	var val int64
	for i := len(t) - 1; i >= 0; i-- {
		val = val*TrinaryRadix + int64(t[i])
	}
	return val
}

// IntToTrytes converts int64 to a slice of trytes.
func IntToTrytes(value int64, trytesCnt int) Trytes {
	negative := value < 0
	v := iotaGoMath.AbsInt64(value)

	var trytes strings.Builder
	trytes.Grow(trytesCnt)

	for i := 0; i < trytesCnt; i++ {
		if v == 0 {
			trytes.WriteByte('9')
			continue
		}

		v += TryteRadix / 2
		tryte := int8(v%TryteRadix) - TryteRadix/2
		v /= TryteRadix
		if negative {
			tryte = -tryte
		}
		trytes.WriteByte(MustTryteValueToTryte(tryte))
	}
	return trytes.String()
}

// TrytesToInt converts a slice of trytes to int64.
func TrytesToInt(t Trytes) int64 {
	// ignore tailing 9s
	var i int
	for i = len(t) - 1; i >= 0; i-- {
		if t[i] != '9' {
			break
		}
	}

	var val int64
	for ; i >= 0; i-- {
		val = val*TryteRadix + int64(MustTryteToTryteValue(t[i]))
	}
	return val
}

// CanTritsToTrytes returns true if t can be converted to trytes.
func CanTritsToTrytes(trits Trits) bool {
	if len(trits) == 0 {
		return false
	}
	return len(trits)%TritsPerTryte == 0
}

// MustTryteValueToTryte converts the value of a tryte v in [-13,13] to a tryte char in [9A-Z].
// It panics when v is an invalid value.
func MustTryteValueToTryte(v int8) byte {
	idx := uint(v - MinTryteValue)
	if idx >= uint(len(TryteValueToTyteLUT)) {
		panic(ErrInvalidTrytes)
	}
	return TryteValueToTyteLUT[idx]
}

// MustTryteToTryteValue converts a tryte char t in [9A-Z] to a tryte value in [-13,13].
// It panics when t is an invalid tryte.
func MustTryteToTryteValue(t byte) int8 {
	idx := uint(t - '9')
	if idx >= uint(len(TryteToTryteValueLUT)) {
		panic(ErrInvalidTrytes)
	}
	return TryteToTryteValueLUT[idx]
}

// TritsToTrytes converts a slice of trits into trytes. Returns an error if len(t)%3!=0
func TritsToTrytes(trits Trits) (Trytes, error) {
	if err := ValidTrits(trits); err != nil {
		return "", err
	}
	if !CanTritsToTrytes(trits) {
		return "", errors.Wrap(ErrInvalidTritsLength, "trits slice size must be a multiple of 3")
	}
	return MustTritsToTrytes(trits), nil
}

// MustTritsToTrytes converts a slice of trits into trytes.
// Performs no validation on the input trits and might therefore return an invalid trytes representation
// (without a panic).
func MustTritsToTrytes(trits Trits) Trytes {
	trytes := make([]byte, len(trits)/TritsPerTryte)
	for i := range trytes {
		tryteTrits := trits[i*TritsPerTryte:]
		_ = tryteTrits[2] // bounds check hint to compiler
		v := tryteTrits[0] + tryteTrits[1]*3 + tryteTrits[2]*9

		trytes[i] = MustTryteValueToTryte(v)
	}
	return string(trytes)
}

func validTryte(t rune) bool {
	return (t >= 'A' && t <= 'Z') || t == '9'
}

// ValidTryte returns the validity of a tryte (must be rune A-Z or 9)
func ValidTryte(t rune) error {
	if !validTryte(t) {
		return ErrInvalidTrytes
	}
	return nil
}

// ValidTrytes returns true if t is made of valid trytes.
func ValidTrytes(trytes Trytes) error {
	if trytes == "" {
		return ErrInvalidTrytes
	}
	for _, tryte := range trytes {
		if !validTryte(tryte) {
			return ErrInvalidTrytes
		}
	}
	return nil
}

// NewTrytes casts to Trytes and checks its validity.
func NewTrytes(s string) (Trytes, error) {
	err := ValidTrytes(s)
	return s, err
}

// TrytesToTrits converts a slice of trytes into trits.
func TrytesToTrits(trytes Trytes) (Trits, error) {
	if err := ValidTrytes(trytes); err != nil {
		return nil, err
	}
	return MustTrytesToTrits(trytes), nil
}

// MustTrytesToTrits converts a slice of trytes into trits.
// Performs no validation on the provided inputs (therefore might return an invalid representation) and might panic.
func MustTrytesToTrits(trytes Trytes) Trits {
	trits := make(Trits, len(trytes)*TritsPerTryte)
	for i := 0; i < len(trytes); i++ {
		MustPutTryteTrits(trits[i*TritsPerTryte:], MustTryteToTryteValue(trytes[i]))
	}
	return trits
}

// MustPutTryteTrits converts v in [-13,13] to its corresponding 3-trit value and writes this to trits.
// It panics on invalid input.
func MustPutTryteTrits(trits []int8, v int8) {
	idx := v - MinTryteValue
	_ = trits[2] // early bounds check to guarantee safety of writes below
	trits[0] = TryteValueToTritsLUT[idx][0]
	trits[1] = TryteValueToTritsLUT[idx][1]
	trits[2] = TryteValueToTritsLUT[idx][2]
}

// CanBeHash returns the validity of the trit length.
func CanBeHash(trits Trits) bool {
	return len(trits) == HashTrinarySize
}

// TrytesToBytes packs trytes into a slice of bytes (5 packed trits in 1 byte).
func TrytesToBytes(trytes Trytes) ([]byte, error) {
	trits, err := TrytesToTrits(trytes)
	if err != nil {
		return nil, err
	}
	return MustTritsToBytes(trits), nil
}

// MustTrytesToBytes packs trytes into a slice of bytes (5 packed trits in 1 byte).
// Performs no validation on the provided inputs (therefore might return an invalid representation) and might panic.
func MustTrytesToBytes(trytes Trytes) []byte {
	return MustTritsToBytes(MustTrytesToTrits(trytes))
}

// BytesToTrytes unpacks a slice of bytes (5 packed trits in 1 byte) into trytes.
func BytesToTrytes(bytes []byte, numTrytes ...int) (Trytes, error) {
	var numTrits int
	if len(numTrytes) > 0 {
		numTrits = numTrytes[0] * TritsPerTryte
	} else {
		numTrits = int(roundUpToTryteMultiple(uint(len(bytes)) * NumberOfTritsInAByte))
	}

	trits, err := BytesToTrits(bytes, numTrits)
	if err != nil {
		return "", err
	}
	return MustTritsToTrytes(trits), nil
}

// MustBytesToTrytes unpacks a slice of bytes (5 packed trits in 1 byte) into trytes.
// Performs no validation on the provided inputs (therefore might return an invalid representation) and might panic.
func MustBytesToTrytes(bytes []byte, numTrytes ...int) Trytes {
	var numTrits int
	if len(numTrytes) > 0 {
		numTrits = numTrytes[0] * TritsPerTryte
	} else {
		numTrits = int(roundUpToTryteMultiple(uint(len(bytes)) * NumberOfTritsInAByte))
	}
	return MustTritsToTrytes(MustBytesToTrits(bytes, numTrits))
}

// TritsToBytes packs an array of trits into an array of bytes (5 packed trits in 1 byte).
func TritsToBytes(trits Trits) ([]byte, error) {
	if err := ValidTrits(trits); err != nil {
		return nil, err
	}
	return MustTritsToBytes(trits), nil
}

// MustTritsToBytes packs an array of trits into an array of bytes (5 packed trits in 1 byte).
// Performs no validation on the provided inputs (therefore might return an invalid representation) and might panic.
func MustTritsToBytes(trits Trits) (bytes []byte) {
	tritsLength := len(trits)
	bytesLength := (tritsLength + NumberOfTritsInAByte - 1) / NumberOfTritsInAByte

	bytes = make([]byte, bytesLength)
	for i := len(bytes) - 1; i >= 0; i-- {
		tmp := trits[i*NumberOfTritsInAByte:]
		if len(tmp) < NumberOfTritsInAByte {
			var v int
			for j := len(tmp) - 1; j >= 0; j-- {
				v = v*TrinaryRadix + int(tmp[j])
			}
			bytes[i] = byte(v)
			continue
		}

		v := int(tmp[0]) + int(tmp[1])*3 + int(tmp[2])*9 + int(tmp[3])*27 + int(tmp[4])*81
		bytes[i] = byte(v)
	}
	return bytes
}

// ValidBytesForTrits checks whether the given bytes are valid for bytes to trits conversion (5 packed trits in 1 byte).
func ValidBytesForTrits(bytes []byte) error {
	for i, b := range bytes {
		c := int8(b)
		if c > 121 || c < -121 {
			return errors.Wrapf(ErrInvalidByte, "at index %d (byte value: %d)", i, c)
		}
	}
	return nil
}

// BytesToTrits unpacks an array of bytes (5 packed trits in 1 byte) into an array of trits.
func BytesToTrits(bytes []byte, numTrits ...int) (Trits, error) {
	if err := ValidBytesForTrits(bytes); err != nil {
		return nil, err
	}
	if len(numTrits) > 0 {
		tritsLength := numTrits[0]

		minTritLength := (len(bytes)-1)*NumberOfTritsInAByte + 1
		if tritsLength < minTritLength {
			return nil, errors.Wrapf(ErrInvalidTritsLength, "must be at least %d in size", minTritLength)
		}
	}
	return MustBytesToTrits(bytes, numTrits...), nil
}

// MustBytesToTrits unpacks an array of bytes (5 packed trits in 1 byte) into an array of trits.
// Performs no validation on the provided inputs (therefore might return an invalid representation) and might panic.
func MustBytesToTrits(bytes []byte, numTrits ...int) Trits {
	tritsLength := len(bytes) * NumberOfTritsInAByte
	resultLength := tritsLength

	if len(numTrits) > 0 {
		// if the number of trits is specified this becomes the target length
		resultLength = numTrits[0]

		// if the specified number of trits decreases, we need to adapt the number of bytes to be converted
		if resultLength < tritsLength {
			bytesLength := (resultLength + NumberOfTritsInAByte - 1) / NumberOfTritsInAByte
			tritsLength = bytesLength * NumberOfTritsInAByte
			bytes = bytes[:bytesLength]
		} else {
			tritsLength = resultLength
		}
	}

	trits := make(Trits, tritsLength)
	for i := range bytes {
		tmp := trits[i*NumberOfTritsInAByte:]
		_ = tmp[4] // bounds check hint to compiler
		tmp[0] = bytesToTritsLUT[bytes[i]][0]
		tmp[1] = bytesToTritsLUT[bytes[i]][1]
		tmp[2] = bytesToTritsLUT[bytes[i]][2]
		tmp[3] = bytesToTritsLUT[bytes[i]][3]
		tmp[4] = bytesToTritsLUT[bytes[i]][4]
	}
	return trits[:resultLength]
}

// Pad pads the given trytes with 9s up to the given size.
func Pad(trytes Trytes, n int) (Trytes, error) {
	if len(trytes) > 0 {
		if err := ValidTrytes(trytes); err != nil {
			return "", err
		}
	}
	return MustPad(trytes, n), nil
}

// MustPad pads the given trytes with 9s up to the given size.
// Performs no validation on the provided inputs (therefore might return an invalid representation) and might panic.
func MustPad(trytes Trytes, n int) Trytes {
	if len(trytes) >= n {
		return trytes
	}

	var result strings.Builder
	result.Grow(n)

	result.WriteString(trytes)
	result.Write(bytes.Repeat([]byte{'9'}, n-len(trytes)))

	return result.String()
}

// PadTrits pads the given trits with 0 up to the given size.
func PadTrits(trits Trits, n int) (Trits, error) {
	if len(trits) > 0 {
		if err := ValidTrits(trits); err != nil {
			return nil, err
		}
	}
	return MustPadTrits(trits, n), nil
}

// MustPadTrits pads the given trits with 0 up to the given size.
// Performs no validation on the provided inputs (therefore might return an invalid representation) and might panic.
func MustPadTrits(trits Trits, n int) Trits {
	if len(trits) >= n {
		return trits
	}

	result := make(Trits, n)
	copy(result, trits)
	return result
}

// EncodedLength returns the length of trits needed to encode the value + encoding information.
func EncodedLength(value int64) uint64 {
	if value == 0 {
		return uint64(len(encodedZero))
	}
	length := uint64(roundUpToTryteMultiple(uint(MinTrits(value))))

	// trits length + encoding length
	return length + uint64(MinTrits((1<<(length/uint64(TrinaryRadix)))-1))
}

// EncodeInt64 encodes an int64 as a slice of trits with encoding information.
func EncodeInt64(value int64) (t Trits, size uint64, err error) {
	size = EncodedLength(value)

	if value == 0 {
		return encodedZero, size, nil
	}

	var encoding int64 = 0
	index := 0
	length := roundUpToTryteMultiple(uint(MinTrits(value)))
	t = make(Trits, size)
	copy(t, IntToTrits(value))

	for i := 0; i < int(length)-TrinaryRadix; i += TrinaryRadix {
		if TritsToInt(t[i:i+TrinaryRadix]) >= 0 {
			encoding |= 1 << uint(index)
			for j := 0; j < TrinaryRadix; j++ {
				t[i+j] = -t[i+j]
			}
		}
		index++
	}

	if TritsToInt(t[length-TrinaryRadix:length]) <= 0 {
		encoding |= 1 << uint(index)
		for i := 1; i < TrinaryRadix+1; i++ {
			t[int(length)-i] = -t[int(length)-i]
		}
	}

	copy(t[length:], IntToTrits(encoding))
	return t, size, nil
}

// DecodeInt64 decodes a slice of trits with encoding information as an int64.
func DecodeInt64(t Trits) (value int64, size uint64, err error) {
	numTrits := uint64(len(t))

	equal, err := TritsEqual(t[0:4], encodedZero)
	if err != nil {
		return 0, 0, err
	}

	if equal {
		return 0, EncodedLength(0), nil
	}

	value = 0
	var encodingStart uint64 = 0

	for (encodingStart < numTrits) && (TritsToInt(t[encodingStart:encodingStart+TrinaryRadix]) <= 0) {
		encodingStart += TrinaryRadix
	}

	if encodingStart >= numTrits {
		return 0, 0, errors.New("encodingStart > numTrits")
	}

	encodingStart += TrinaryRadix
	encodingLength := MinTrits((1 << (encodingStart / TrinaryRadix)) - 1)
	encoding := TritsToInt(t[encodingStart : encodingStart+uint64(encodingLength)])

	// Bound checking for the lookup table
	if encodingStart/TrinaryRadix > 13 {
		return 0, 0, errors.New("encodingStart/TrinaryRadix > 13")
	}

	for i := 0; i < int(encodingStart/TrinaryRadix); i++ {
		tryteValue := TritsToInt(t[i*TrinaryRadix : (i*TrinaryRadix)+TrinaryRadix])

		if ((encoding >> uint(i)) & 1) == 1 {
			tryteValue = -tryteValue
		}
		value += Pow27LUT[i] * tryteValue
	}

	return value, encodingStart + uint64(encodingLength), nil
}

// Sum returns the sum of two trits.
func Sum(a int8, b int8) int8 {
	s := a + b

	switch s {
	case 2:
		return -1
	case -2:
		return 1
	default:
		return s
	}
}

func cons(a int8, b int8) int8 {
	if a == b {
		return a
	}

	return 0
}

func any(a int8, b int8) int8 {
	s := a + b

	if s > 0 {
		return 1
	}

	if s < 0 {
		return -1
	}

	return 0
}

func fullAdd(a int8, b int8, c int8) [2]int8 {
	sA := Sum(a, b)
	cA := cons(a, b)
	cB := cons(sA, c)
	cOut := any(cA, cB)
	sOut := Sum(sA, c)
	return [2]int8{sOut, cOut}
}

// AddTrits adds a to b.
func AddTrits(a Trits, b Trits) Trits {
	maxLen := int64(math.Max(float64(len(a)), float64(len(b))))
	if maxLen == 0 {
		return Trits{0}
	}
	out := make(Trits, maxLen)
	var aI, bI, carry int8

	for i := 0; i < len(out); i++ {
		if i < len(a) {
			aI = a[i]
		} else {
			aI = 0
		}
		if i < len(b) {
			bI = b[i]
		} else {
			bI = 0
		}

		fA := fullAdd(aI, bI, carry)
		out[i] = fA[0]
		carry = fA[1]
	}
	return out
}
