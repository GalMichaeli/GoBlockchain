package curve

import (
	"math/big"
	//"fmt"
	"GoBlockchain/field"
	"GoBlockchain/point"
	//"errors"
)

type Curve struct {
	P big.Int
	A field.FieldElement
	B field.FieldElement
	G point.Point
	N field.FieldElement
}

SECP256K1 := curve.New(
	"fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f",
	"0",
	"7",
	"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
	"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
	"fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141",
	16,
)

func New(p, a, b, Gx, Gy, n string, base int) *Curve {
	c := new(Curve)

	c.P.SetString(p, base)
	c.A.SetString(a, p, base)
	c.B.SetString(b, p, base)
	c.G.SetString(Gx, Gy, a, b, p, base)
	c.N.SetString(n, p, base)
	return c
}

func (c *Curve) ScalarMul(z *point.Point, coef *field.FieldElement) *point.Point {
	var r field.FieldElement
	res := new(point.Point)

	r.Mod(coef,&c.N)
	return res.ScalarMul(z, &r)
}
