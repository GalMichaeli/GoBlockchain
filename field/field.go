package field

import (
	"log"
	"github.com/holiman/uint256"
)

type FieldElement struct {
	Number uint256.Int
	Prime  uint256.Int
}

func NewFieldElement(numStr string, primeStr string, base string) *FieldElement {

	var number, prime *uint256.Int
	var err error
	
	if base == "d" {
		number, err = uint256.FromDecimal(numStr)
		if err != nil {
			log.Fatal(err)
		}
		prime, err = uint256.FromDecimal(primeStr)
		if err != nil {
			log.Fatal(err)
		}
	} else if base == "x" {
		number, err = uint256.FromHex(numStr)
		if err != nil {
			log.Fatal(err)
		}
		prime, err = uint256.FromHex(primeStr)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Unknown base. Third argument must be \"d\" or \"x\"")
	}
	
	fe := new(FieldElement)
	fe.Number = *number
	fe.Prime = *prime
	return fe
}
