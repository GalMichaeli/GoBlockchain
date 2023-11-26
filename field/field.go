package field

import (
	"fmt"
	//"errors"
	"math/big"
)

type FieldElement struct {
	Number big.Int
	Prime big.Int
}

func samePrime(x, y *FieldElement) bool {
	if x.Prime.Cmp(&y.Prime) != 0 {
		return false
	}

	return true
}

/*******************************************************************************/
/*******************************************************************************/

func Zero(fe *FieldElement) *FieldElement {
	primeStr := fe.Prime.Text(10)
	return New("0", primeStr, 10)
}

func New(number, prime string, base int) *FieldElement {
	fe := new(FieldElement)
	var sucNum, sucPrime bool
	
	switch base {
	case 2:
		_, sucNum = fe.Number.SetString("0b" + number, 0)
		_, sucPrime = fe.Prime.SetString("0b" + prime, 0)
	case 16:
		_, sucNum = fe.Number.SetString("0x" + number, 0)
		_, sucPrime = fe.Prime.SetString("0x" + prime, 0)
	case 10:
		_, sucNum = fe.Number.SetString(number, 0)
		_, sucPrime = fe.Prime.SetString(prime, 0)
	default:
		sucNum, sucPrime = false, false 
	}

	if !sucNum || !sucPrime || (fe.Number.Cmp(&fe.Prime) > -1) {
		return nil
	}

	return fe
}

func (fe *FieldElement) String(base int) string {
	var str string
	
	switch base {
	case 2:
		str = fmt.Sprintf("Number: %s\nPrime: %s", fe.Number.Text(2), fe.Prime.Text(2))
	case 16:
		str = fmt.Sprintf("Number: %s\nPrime: %s", fe.Number.Text(16), fe.Prime.Text(16))
	case 10:
		str = fmt.Sprintf("Number: %s\nPrime: %s", fe.Number.Text(10), fe.Prime.Text(10))
	default:
		str = "<nil>"
	}

	return str	
}

func (fe *FieldElement) Set(x *FieldElement) *FieldElement {
	fe.Number.Set(&x.Number)
	fe.Prime.Set(&x.Prime)

	return fe
}

func (fe *FieldElement) SetString(n, p string, base int) (*FieldElement, bool) {
	_, nOk := fe.Number.SetString(n, base)
	_, pOk := fe.Prime.SetString(p, base)

	if !nOk || !pOk || (fe.Number.Cmp(&fe.Prime) > -1) {
		return nil, false
	}

	return fe, true
}

func (fe *FieldElement) Cmp(x *FieldElement) int {
	if !samePrime(fe, x) {
		return 666
	}

	return fe.Number.Cmp(&x.Number)
}

func (fe *FieldElement) Eq(x *FieldElement) bool {
	if !samePrime(fe, x) {
		return false
	}

	return fe.Cmp(x) == 0
}

func (fe *FieldElement) Add(x, y *FieldElement) *FieldElement {
	if !samePrime(x, y) {
		return nil
	}

	fe.Prime.Set(&x.Prime)
	fe.Number.Add(&x.Number, &y.Number)
	fe.Number.Mod(&fe.Number, &fe.Prime)

	return fe
}

func (fe *FieldElement) Sub(x, y *FieldElement) *FieldElement {
	if !samePrime(x, y) {
		return nil
	}

	fe.Prime.Set(&x.Prime)
	fe.Number.Sub(&x.Number, &y.Number)
	fe.Number.Mod(&fe.Number, &fe.Prime)

	return fe
}

func (fe *FieldElement) Mul(x, y *FieldElement) *FieldElement {
	if !samePrime(x, y) {
		return nil
	}

	fe.Prime.Set(&x.Prime)
	fe.Number.Mul(&x.Number, &y.Number)
	fe.Number.Mod(&fe.Number, &fe.Prime)

	return fe
}

func (fe *FieldElement) Div(x, y *FieldElement) *FieldElement {
	if !samePrime(x, y) {
		return nil
	}

	fe.Prime.Set(&x.Prime)
	fe.Number.Set(&y.Number)
	fe.Number.ModInverse(&fe.Number, &fe.Prime)
	fe.Number.Mul(&fe.Number, &x.Number)

	return fe
}

func (fe *FieldElement) Exp(base *FieldElement, exponent *big.Int) *FieldElement {
	fe.Prime.Set(&base.Prime)
	fe.Number.Exp(&fe.Number, exponent, &fe.Prime)
	return fe
}

func (fe *FieldElement) Inverse(x *FieldElement) *FieldElement {
	fe.Prime.Set(&x.Prime)
	fe.Number.ModInverse(&fe.Number, &fe.Prime)
	return fe
}

func (fe *FieldElement) Mod(x, y *FieldElement) *FieldElement {
	if x.Prime.Cmp(&y.Prime) != 0 {
		return nil
	}

	fe.Prime.Set(&x.Prime)
	fe.Number.Mod(&x.Number, &y.Number)
	return fe
}
