package main

import (
	"fmt"
	"log"
	//"strings"
	//"GoBlockchain/field"
	"GoBlockchain/point"
	"GoBlockchain/curve"
	//"math/big"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	var p1 point.Point
	p1.ScalarMul(&SECP256K1.G, &SECP256K1.N)
	fmt.Println(p1.String(10))

	
}
