package main

import (
	"fmt"
	"log"

	lib "github.com/Kroning/zigzag/lib"	
)

func main() {
	candles := []*lib.Candle{}

	// Valley -> peak
	for i:=0; i<10; i++ {
		candle := &lib.Candle{ Low: 100 + float64(i), High: (101.0 + float64(i)) }
		candles = append(candles, candle)
	}
	// Peak -> valley
    for i:=0; i<9; i++ {
        candle := &lib.Candle{ Low: 109 - float64(i), High: (110.0 - float64(i)) }
        candles = append(candles, candle)
    } 

	swings, err := lib.GetSwings(candles, 5, 2)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	for _, swing := range swings {
	    fmt.Println(swing)
	}
}
