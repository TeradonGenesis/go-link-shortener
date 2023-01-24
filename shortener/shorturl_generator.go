package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func Sha256Enc(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func Base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(iniLink string, userId string) string {
	urlHashBytes := Sha256Enc(iniLink + userId)
	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := Base58Encoded([]byte(fmt.Sprintf("%d", generatedNum)))
	return finalString[:8]
}
