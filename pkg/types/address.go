package types

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	goTron "github.com/kattana-io/go-tron/address"
	"github.com/shengdoushi/base58"
)

type Address struct {
	address string
	bytes   []byte
	ok      bool
}

func NewFromBase58(input string) *Address {
	bytes, err := base58.Decode(input, base58.BitcoinAlphabet)

	return &Address{
		bytes: bytes,
		ok:    err == nil,
	}
}

// NewEmptyAddress - create empty address
func NewEmptyAddress() *Address {
	return &Address{
		bytes: []byte{0, 0, 0, 0},
		ok:    false,
	}
}

// Credits to https://gist.github.com/motopig/c680f53897429fd15f5b3ca9aa6f6ed2
func NewFromHex(input string) *Address {
	// Check for empty string
	if input == "" {
		return &Address{
			bytes: []byte{},
			ok:    false,
		}
	}
	input = removeZeroX(input)
	if len(input)%2 == 1 {
		input = "0" + input
	}
	address := addPrefix(input)
	addb, _ := hex.DecodeString(address)
	hash1 := s256(s256(addb))
	secret := hash1[:4]
	addb = append(addb, secret...)

	return &Address{
		bytes: addb,
		ok:    true,
	}
}

func (a *Address) ToBase58() string {
	return base58.Encode(a.bytes, base58.BitcoinAlphabet)
}

// ToHex - we need to receive a string with prefix 41
func (a *Address) ToHex() string {
	return hex.EncodeToString(a.bytes)[:42]
}

func (a *Address) ToGoTronAddr() (goTron.Address, error) {
	return goTron.FromBase58(a.ToBase58())
}

// ToGoEthHex returns go-ethereum hex address
func (a *Address) ToGoEthHex() string {
	return toGoEthHex(a.bytes)
}

// ToGoEthAddr returns go-ethereum address
func (a *Address) ToGoEthAddr() common.Address {
	return common.HexToAddress(toGoEthHex(a.bytes))
}

func toGoEthHex(addr []byte) string {
	return "0x" + hex.EncodeToString(addr)[2:42]
}

func (a *Address) GetBytes() []byte {
	return a.bytes
}

func (a *Address) GetStatus() bool {
	return a.ok
}

func removeZeroX(input string) string {
	if input[0] == '0' && input[1] == 'x' {
		return input[2:]
	}
	return input
}

func addPrefix(input string) string {
	if input[0] == '4' && input[1] == '1' {
		return input
	}
	return "41" + input
}

func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}
