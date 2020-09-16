// Package address provides primitives for generating and validating addresses (with and without checksum).
package address

import (
	"github.com/muxxer/iota.go/checksum"
	. "github.com/muxxer/iota.go/consts"
	"github.com/muxxer/iota.go/kerl"
	. "github.com/muxxer/iota.go/signing"
	"github.com/muxxer/iota.go/signing/key"
	. "github.com/muxxer/iota.go/trinary"
)

// Checksum returns the checksum of the given address.
func Checksum(address Hash) (Trytes, error) {
	if len(address) < 81 {
		return "", ErrInvalidAddress
	}

	addressWithChecksum, err := checksum.AddChecksum(address[:81], true, 9)
	if err != nil {
		return "", err
	}
	return addressWithChecksum[AddressWithChecksumTrytesSize-AddressChecksumTrytesSize : 90], nil
}

// ValidAddress checks whether the given address is valid.
func ValidAddress(address Hash) error {
	switch len(address) {
	case 90:
		if err := ValidChecksum(address[:81], address[81:]); err != nil {
			return err
		}
	case 81:
	default:
		return ErrInvalidAddress
	}
	return ValidTrytes(address)
}

// ValidChecksum checks whether the given checksum corresponds to the given address.
func ValidChecksum(address Hash, checksum Trytes) error {
	actualChecksum, err := Checksum(address)
	if err != nil {
		return err
	}
	if checksum != actualChecksum {
		return ErrInvalidChecksum
	}
	return nil
}

// GenerateAddress generates an address deterministically, according to the given seed, index and security level.
func GenerateAddress(seed Trytes, index uint64, secLvl SecurityLevel, addChecksum ...bool) (Hash, error) {
	for len(seed)%81 != 0 {
		seed += "9"
	}

	if secLvl == 0 {
		secLvl = SecurityLevelMedium
	}

	// use Kerl for the entire address generation
	h := kerl.NewKerl()

	subseed, err := Subseed(seed, index, h)
	if err != nil {
		return "", err
	}

	prvKey, err := key.Sponge(subseed, secLvl, h)
	if err != nil {
		return "", err
	}

	digests, err := Digests(prvKey, h)
	if err != nil {
		return "", err
	}

	addressTrits, err := Address(digests, h)
	if err != nil {
		return "", err
	}

	address := MustTritsToTrytes(addressTrits)

	if len(addChecksum) > 0 && addChecksum[0] {
		return checksum.AddChecksum(address, true, 9)
	}

	return address, nil
}

// GenerateAddresses generates N new addresses from the given seed, indices and security level.
func GenerateAddresses(seed Trytes, start uint64, count uint64, secLvl SecurityLevel, addChecksum ...bool) (Hashes, error) {
	addresses := make(Hashes, count)

	var withChecksum bool
	if len(addChecksum) > 0 && addChecksum[0] {
		withChecksum = true
	}

	var err error
	for i := 0; i < int(count); i++ {
		addresses[i], err = GenerateAddress(seed, start+uint64(i), secLvl, withChecksum)
		if err != nil {
			return nil, err
		}
	}
	return addresses, nil
}
