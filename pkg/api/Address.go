package api

import (
	"crypto/sha256"
	"encoding/hex"
	goTron "github.com/0x10f/go-tron/address"
	"github.com/shengdoushi/base58"
	"strings"
)

type Address struct {
	address string
	bytes   []byte
	ok      bool
}

func FromBase58(input string) *Address {
	bytes, err := base58.Decode(input, base58.BitcoinAlphabet)

	return &Address{
		bytes: bytes,
		ok:    err == nil,
	}
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

// EmptyAddress - create empty address
func EmptyAddress() *Address {
	return &Address{
		bytes: []byte{0, 0, 0, 0},
		ok:    false,
	}
}

// Credits to https://gist.github.com/motopig/c680f53897429fd15f5b3ca9aa6f6ed2
func FromHex(input string) *Address {
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

// ToHex - we need to receive string with 41 prefix
func (a *Address) ToHex() string {
	return hex.EncodeToString(a.bytes)[:42]
}

func (a *Address) PackIntoEthWord() (result string) {
	str := strings.ReplaceAll(a.ToHex(), "41", "00")
	content := len(str)
	pre := 64 - content
	for i := 0; i < pre; i++ {
		result += "0"
	}
	result += str
	return result
}

func (a *Address) PackIntoEthBytes() []byte {
	str := strings.ReplaceAll(a.ToHex(), "41", "00")
	return []byte(str)
}

func (a *Address) GetBytes() []byte {
	return a.bytes
}

func (a *Address) GetStatus() bool {
	return a.ok
}

func (a *Address) ToGoTronAddress() (goTron.Address, error) {
	return goTron.FromBase58(a.ToBase58())
}
