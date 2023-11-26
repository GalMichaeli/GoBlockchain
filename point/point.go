package point

import (
	"fmt"
	//"errors"
	"GoBlockchain/field"
	"math/big"
)

type Point struct {
	X field.FieldElement
	Y field.FieldElement
	A field.FieldElement
	B field.FieldElement
}

var Infinity Point

func init() {
	Infinity.X.SetString("0", "0", 10)
	Infinity.Y.SetString("0", "0", 10)
	Infinity.A.SetString("0", "0", 10)
	Infinity.B.SetString("0", "0", 10)
}

func onCurve(z *Point) bool {
	var xCubed, ySquared, rhs, ax field.FieldElement

	xCubed.Set(&z.X).Exp(&xCubed, big.NewInt(3))
	ySquared.Set(&z.Y).Exp(&ySquared, big.NewInt(2))

	ax.Set(&z.X).Mul(&ax, &z.A)
	rhs.Set(&xCubed).Add(&rhs, &ax)
	rhs.Add(&rhs, &z.B)

	return ySquared.Cmp(&rhs) == 0
}

func sameCurve(p, q *Point) bool {
	if p.Eq(&Infinity) || q.Eq(&Infinity) {
		return true
	}
	return p.A.Eq(&q.A) && p.B.Eq(&q.B)
}

func New(x, y, a, b, p string, base int) *Point {
	z := new(Point)
	_, xOk := z.X.SetString(x, p, base)
	_, yOk := z.Y.SetString(y, p, base)
	_, aOk := z.A.SetString(a, p, base)
	_, bOk := z.B.SetString(b, p, base)

	if !onCurve(z) || !xOk || !yOk || !aOk || !bOk {
		return nil
	}

	return z
}

func (z *Point) SetString(x, y, a, b, p string, base int) *Point {
	z.X.SetString(x, p, base)
	z.Y.SetString(y, p, base)
	z.A.SetString(a, p, base)
	z.B.SetString(b, p, base) 
	return z
}

func (z *Point) String(base int) string {
	var str string
	if z.Eq(&Infinity) {
		return "Infinity"
	}
	str = fmt.Sprintf("x: %s\ny: %s\na: %s\nb: %s\nprime: %s",
		z.X.Number.Text(base),
		z.Y.Number.Text(base),
		z.A.Number.Text(base),
		z.B.Number.Text(base),
		z.X.Prime.Text(base),
	)

	return str
}



func (z *Point) Set(p *Point) *Point {
	z.X.Set(&p.X)
	z.Y.Set(&p.Y)
	z.A.Set(&p.A)
	z.B.Set(&p.B)
	return z
}

func (z *Point) Eq(p *Point) bool {
	return z.X.Eq(&p.X) && z.Y.Eq(&p.Y) && z.A.Eq(&p.A) && z.B.Eq(&p.B)
}

func (z *Point) Add(p, q *Point) *Point {
	if !sameCurve(p, q) {
		return nil
	}
	
	if p.Eq(&Infinity) {
		return z.Set(q)
	}

	if q.Eq(&Infinity) {
		return z.Set(p)
	}

	if p.X.Eq(&q.X) && (!p.Y.Eq(&q.Y) || p.Y.Eq(field.Zero(&p.Y))) {
		return z.Set(&Infinity)
	}

	var n, nn, d, s, sSquared, x, y field.FieldElement

	if p.X.Eq(&q.X) {
		nn.Mul(&p.X, &p.X) // nn = p.X^2
		n.Add(&nn, &nn).Add(&n, &nn) // n = 3 * p.X^2
		n.Add(&n, &p.A) // n = 3 * p.X^2 + p.A
		d.Add(&p.Y, &p.Y) // d = 2 * p.Y		
	} else {
		n.Sub(&q.Y, &p.Y)
		d.Sub(&q.X, &p.X)
	}

	s.Div(&n, &d)
	sSquared.Mul(&s, &s)
	x.Sub(&sSquared, &p.X).Sub(&x, &q.X)
	y.Sub(&p.X, &x).Mul(&y, &s).Sub(&y, &p.Y)
	z.X.Set(&x)
	z.Y.Set(&y)
	z.A.Set(&p.A)
	z.B.Set(&p.B)

	return z	
}

func (z *Point) ScalarMul(p *Point, c *field.FieldElement) *Point {
	byteSlice := c.Number.Bytes()
	bytesNo := len(byteSlice)
	bitsNo := c.Number.BitLen()
	it, acc := *p, Infinity
	var b byte

	for i := bytesNo - 1; i >= 0; i-- {
		b = byteSlice[i]
		for j := 0; j < 8 && bitsNo > 0; j, bitsNo = j+1, bitsNo-1 {
			if b & 1 == 1 {
				acc.Add(&acc, &it)
			}
			it.Add(&it, &it)
			b >>= 1
		}
	}
	
	return z.Set(&acc)
}
