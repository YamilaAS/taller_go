package main

import "fmt"

func main() {
	fmt.Println("Versiòn 1 de contador de votos")
	contadorDeVotos1()

	fmt.Println("Versiòn 2 con map de contador de votos")
	contadorDeVotos2()
}

func contadorDeVotos1() {
	puntajes := []int{5, 3, 4, 2, 1, 5, 4, 3, 2, 5}

	var c1, c2, c3, c4, c5 int

	for i := 0; i < len(puntajes); i++ {
		if puntajes[i] == 1 {
			c1++
		} else if puntajes[i] == 2 {
			c2++
		} else if puntajes[i] == 3 {
			c3++
		} else if puntajes[i] == 4 {
			c4++
		} else if puntajes[i] == 5 {
			c5++
		}
	}

	fmt.Println("Cantidad de votos:")
	fmt.Println("Puntaje 1:", c1)
	fmt.Println("Puntaje 2:", c2)
	fmt.Println("Puntaje 3:", c3)
	fmt.Println("Puntaje 4:", c4)
	fmt.Println("Puntaje 5:", c5)

	positivos := c4 + c5
	negativos := c1 + c2

	if positivos > negativos {
		fmt.Println("¡Buen resultado!")
	} else {
		fmt.Println("Resultado mejorable")
	}
}

func contadorDeVotos2() {
	puntajes := []int{5, 3, 4, 2, 1, 5, 4, 3, 2, 5}

	contador := make(map[int]int)

	for _, puntaje := range puntajes {
		contador[puntaje]++
	}

	fmt.Println("Cantidad de votos por puntaje:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Puntaje %d: %d votos\n", i, contador[i])
	}

	votosPositivos := contador[4] + contador[5]
	votosNegativos := contador[1] + contador[2]

	if votosPositivos > votosNegativos {
		fmt.Println("¡Buen resultado!")
	} else {
		fmt.Println("Resultado mejorable")
	}
}
