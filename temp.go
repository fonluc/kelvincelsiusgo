package main

import "fmt"

const ebulicaoK = 373

func main() {

	tempK := ebulicaoK
	tempC := (tempK - 273)

	fmt.Println("A temperatura de ebulição da água em ºK = ", tempK)
	fmt.Println("A temperatura de ebulição da água em ºC = ", tempC)
}
