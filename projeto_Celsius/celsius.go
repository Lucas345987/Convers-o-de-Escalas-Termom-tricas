package main

import "fmt"

const tempK = 351.65

func main() {
	tempC := tempK - 273.15
	fmt.Printf("A temperatura de Kelvin para Celsus é : %.2f graus.", tempC)
}
