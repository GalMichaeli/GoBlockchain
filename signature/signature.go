package signature

import (
	"fmt"
	"crypto/rand"
	"math/big"
	"GoBlockchain/field"
	"GoBlockchain/point"
	"GoBlockchain/curve"
)

type Signature struct {
	R FieldElement
	S FieldElement
}

type Key struct {
	privateKey FieldElement
	PublicKey FieldElement
}

func (s *Signature) String(base int) {
	str := fmt.Sprintf("r: %s\ns: %s", s.R.String(base), s.S.String(base))
}

func (s *Signature) Verify(c *curve.Curve, q *point.Point, z *field.FieldElement) bool {
	var u, v field.FieldElement

	u.Div(z, &s.S)
	v.Div(&s.R, &s.S)
	resPoint := c.ScalarMul(&c.G, &u)
	auxPoint := c.ScalarMul(q, &v)
	resPoint.Add(&resPoint, &auxPoint)

	return resPoint.X == s.R
}

func (l *Key) Sign(c *curve.Curve, z *field.FieldElement) *Signature {
	sig := new(Signature)
	var r, s field.FieldElement
	var q point.Point
	
	k, _ := rand.Int(rand.Reader, &c.P)
	kFe := field.FieldElement{*k, c.P}
	r = c.ScalarMul(&c.G, fe).X
	s.Mul(&r, &l.privateKey)
	s.Add(&s, z)
	s.Div(&s, &kFe)
}


