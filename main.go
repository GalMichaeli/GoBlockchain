package main

import (
	"fmt"
	"log"
	"GoBlockchain/field"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	fe := field.NewFieldElement("12", "13", "d")
	fmt.Println(fe)
	
}
