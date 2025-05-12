package main

import (
	"fmt"
	"goDasar/calculation" //Kalau beda package harus import dulu
)

func main() {
	fmt.Println("Hello World")
	kalimat := TestDulu()            //Bisa langsung panggil fungsi lain (ASALKAN SATU PACKAGE)
	result := calculation.Math(5, 2) //Fungsi yang di panggil dari package Calculation
	multiple := calculation.Multiple(2, 3)

	fmt.Println(kalimat)
	fmt.Println(result)
	fmt.Println(multiple)
}
