package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func imprimirLentoln(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
		time.Sleep(18 * time.Millisecond)
	}
	fmt.Printf("\n")
}

func imprimirRapidoln(str string) {
	fmt.Printf("\n")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
		time.Sleep(3 * time.Millisecond)
	}
}

func generar_numero_random(numero_inicial int, numero_final int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	numero_aleatorio := random.Intn((numero_final-numero_inicial)+1) + numero_inicial
	time.Sleep(1 * time.Nanosecond)
	return numero_aleatorio
}

func generar_matriz_cuadrada_aleatoria(tamanio int, numero_aleatorio_inicial int, numero_aleatorio_final int) ([][]int, error) {
	var err error
	if tamanio < 2 {
		err = errors.New("no se pueden generar matrices de dimensiones menores 2 por 2")
		return nil, err
	}
	err = nil

	matriz := make([][]int, tamanio)

	for i := 0; i < tamanio; i++ {
		matriz[i] = make([]int, tamanio)
		for j := 0; j < tamanio; j++ {
			matriz[i][j] = generar_numero_random(numero_aleatorio_inicial, numero_aleatorio_final)
		}
	}
	return matriz, nil
}

func obtener_filas_columnas(matriz [][]int) (int, int) {
	filas := len(matriz)
	columnas := len(matriz[0])

	return filas, columnas
}

func sumar_matrices(matriz_1 [][]int, matriz_2 [][]int) ([][]int, error) {
	var err error
	filas_1, columnas_1 := obtener_filas_columnas(matriz_1)
	filas_2, columnas_2 := obtener_filas_columnas(matriz_2)
	if filas_1 != filas_2 || columnas_1 != columnas_2 || filas_1 != columnas_1 {
		err = errors.New("las matrices deben ser de la misma dimension")
		return nil, err
	}
	err = nil

	matriz_resultado := make([][]int, filas_1)
	for i := 0; i < filas_1; i++ {
		matriz_resultado[i] = make([]int, columnas_1)
		for j := 0; j < columnas_1; j++ {
			matriz_resultado[i][j] = matriz_1[i][j] + matriz_2[i][j]
		}
	}
	return matriz_resultado, err
}

func restar_matrices(matriz_1 [][]int, matriz_2 [][]int) ([][]int, error) {
	var err error
	filas_1, columnas_1 := obtener_filas_columnas(matriz_1)
	filas_2, columnas_2 := obtener_filas_columnas(matriz_2)
	if filas_1 != filas_2 || columnas_1 != columnas_2 || filas_1 != columnas_1 {
		err = errors.New("las matrices deben ser de la misma dimension")
		return nil, err
	}
	err = nil

	matriz_resultado := make([][]int, filas_1)
	for i := 0; i < filas_1; i++ {
		matriz_resultado[i] = make([]int, columnas_1)
		for j := 0; j < columnas_1; j++ {
			matriz_resultado[i][j] = matriz_1[i][j] - matriz_2[i][j]
		}
	}
	return matriz_resultado, err
}

func escalar_por_matriz(escalar int, matriz [][]int) [][]int {
	filas, columnas := obtener_filas_columnas(matriz)
	matriz_resultado := make([][]int, filas)

	for i := 0; i < filas; i++ {
		matriz_resultado[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			matriz_resultado[i][j] = matriz[i][j] * escalar
		}
	}
	return matriz_resultado
}

func generar_matriz_nxm_aleatoria(filas int, columnas int, numero_aleatorio_inicial int, numero_aleatorio_final int) ([][]int, error) {
	var err error
	if filas < 2 || columnas < 2 {
		err = errors.New("no se pueden generar matrices de dimensiones menores 2 por 2")
		return nil, err
	}
	err = nil

	matriz := make([][]int, filas)

	for i := 0; i < filas; i++ {
		matriz[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			matriz[i][j] = generar_numero_random(numero_aleatorio_inicial, numero_aleatorio_final)
		}
	}
	return matriz, err
}

func multiplicar_matrices(matriz_1 [][]int, matriz_2 [][]int) ([][]int, error) {
	var err error
	filas_1, columnas_1 := obtener_filas_columnas(matriz_1) //3/3
	filas_2, columnas_2 := obtener_filas_columnas(matriz_2) //2/2

	if columnas_1 != filas_2 {
		err = errors.New("estas matrices no son operables, tienen que ser axn * nxb")
		return nil, err
	}
	err = nil

	matriz_resultado := make([][]int, filas_1)
	for i := 0; i < filas_1; i++ {
		matriz_resultado[i] = make([]int, columnas_2)
		for j := 0; j < columnas_2; j++ {
			for k := 0; k < columnas_1; k++ {
				matriz_resultado[i][j] += matriz_1[i][k] * matriz_2[k][j]
			}
		}
	}

	return matriz_resultado, err
}

func obtener_traza(matriz [][]int) (int, error) {
	var err error
	filas, columnas := obtener_filas_columnas(matriz)
	if filas != columnas {
		err = errors.New("las matrices deben ser de la misma dimension")
		return 0, err
	}
	err = nil

	traza := 0
	for i := 0; i < filas; i++ {
		traza += matriz[i][i]
	}

	return traza, err
}

func obtener_transpuesta(matriz [][]int) [][]int {
	filas, columnas := obtener_filas_columnas(matriz)

	matriz_transpuesta := make([][]int, columnas)
	for i := 0; i < columnas; i++ {
		matriz_transpuesta[i] = make([]int, filas)
		for j := 0; j < filas; j++ {
			matriz_transpuesta[i][j] = matriz[j][i]
		}
	}

	return matriz_transpuesta
}

func imprimir_matriz(matriz [][]int) {
	filas, columnas := obtener_filas_columnas(matriz)
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func programa_1() bool {
	color.Set(color.FgHiMagenta)
	println("\nResultado de sumar 2 matrices 4x4:\n")
	color.Unset()
	tamanio := 4
	numero_aleatorio_inicial := 0
	numero_aleatorio_final := 10

	matriz_2, err := generar_matriz_cuadrada_aleatoria(tamanio, numero_aleatorio_inicial, numero_aleatorio_final)
	if err != nil {
		fmt.Println(err)
		return false
	}
	matriz_1, err := generar_matriz_cuadrada_aleatoria(tamanio, numero_aleatorio_inicial, numero_aleatorio_final)
	if err != nil {
		fmt.Println(err)
		return false
	}
	matriz_resultado, err := sumar_matrices(matriz_1, matriz_2)
	if err != nil {
		fmt.Println(err)
		return false
	}
	matriz_resultado_resta, err := restar_matrices(matriz_1, matriz_2)
	if err != nil {
		fmt.Println(err)
		return false
	}

	color.Set(color.FgHiMagenta)
	println("Matríz 1:")
	color.Unset()
	imprimir_matriz(matriz_1)

	color.Set(color.FgHiMagenta)
	println("Matríz 2:")
	color.Unset()
	imprimir_matriz(matriz_2)

	color.Set(color.FgHiMagenta)
	println("Matríz resultado suma:")
	color.Unset()
	imprimir_matriz(matriz_resultado)
	color.Set(color.FgHiMagenta)
	println("Matríz resultado resta:")
	color.Unset()
	imprimir_matriz(matriz_resultado_resta)

	return true
}

func programa_2() bool {
	escalar := 5
	tamanio := 4
	matriz, err := generar_matriz_cuadrada_aleatoria(tamanio, 1, 10)
	if err != nil {
		fmt.Println(err)
		return false
	}
	matriz_resultado := escalar_por_matriz(escalar, matriz)

	color.Set(color.FgHiMagenta)
	println("Escalar:")
	color.Unset()
	println(escalar)
	color.Set(color.FgHiMagenta)
	println("Matríz Original:")
	color.Unset()
	imprimir_matriz(matriz)
	color.Set(color.FgHiMagenta)
	println("Matríz resultado:")
	color.Unset()
	imprimir_matriz(matriz_resultado)
	return true
}

func programa_3() bool {
	filas_1 := 3
	columnas_1 := 2
	filas_2 := columnas_1
	columnas_2 := filas_1
	matriz_1, err := generar_matriz_nxm_aleatoria(filas_1, columnas_1, 1, 10)
	if err != nil {
		fmt.Print(err)
		return false
	}
	matriz_2, err := generar_matriz_nxm_aleatoria(filas_2, columnas_2, 1, 10)
	if err != nil {
		fmt.Print(err)
		return false
	}

	matriz_resultado, err := multiplicar_matrices(matriz_1, matriz_2)
	if err != nil {
		fmt.Print(err)
		return false
	}

	color.Set(color.FgHiMagenta)
	println("Matríz 1:")
	color.Unset()
	imprimir_matriz(matriz_1)

	color.Set(color.FgHiMagenta)
	println("Matríz 2:")
	color.Unset()
	imprimir_matriz(matriz_2)

	color.Set(color.FgHiMagenta)
	println("Matríz resultado:")
	color.Unset()
	imprimir_matriz(matriz_resultado)

	return true
}

func programa_4() bool {
	tamanio := 4
	matriz, err := generar_matriz_cuadrada_aleatoria(tamanio, 1, 10)
	if err != nil {
		fmt.Println(err)
		return false
	}
	traza, err := obtener_traza(matriz)
	if err != nil {
		println(err)
		return false
	}

	color.Set(color.FgHiMagenta)
	println("Matríz:")
	color.Unset()
	imprimir_matriz(matriz)
	color.Set(color.FgHiMagenta)
	println("Traza:")
	color.Unset()
	println(traza)
	return true
}

func programa_5() bool {
	tamanio := 3
	matriz, err := generar_matriz_nxm_aleatoria(tamanio, 2, 1, 10)
	if err != nil {
		println(err)
		return false
	}
	matriz_transpuesta := obtener_transpuesta(matriz)

	color.Set(color.FgHiMagenta)
	println("Matriz original:")
	color.Unset()
	imprimir_matriz(matriz)

	color.Set(color.FgHiMagenta)
	println("Matriz transpuesta:")
	color.Unset()
	imprimir_matriz(matriz_transpuesta)
	return true
}

func test() bool {
	color.Set(color.FgYellow)
	println("Probar creación de matrices y operaciones.")
	println("Genere la primera matriz para comenzar.")
	fmt.Printf("\nSeleccione la matriz que desea:\n")
	fmt.Printf("\n\t(1)Matriz cuadrada aleatoria. (2)Matriz NxM aleatoria.\n> ")
	color.Unset()
	opcion := 0
	fmt.Scan(&opcion)

	matriz := [][]int{}

	switch opcion {
	case 1:
		color.Set(color.FgYellow)
		fmt.Printf("\nIngrese el tamaño de la matriz cuadrada:\n> ")
		color.Unset()
		tamanio := 0
		fmt.Scan(&tamanio)
		color.Set(color.FgYellow)
		fmt.Printf("\nGenerar numeros aleatorios desde:\n> ")
		color.Unset()
		var desde int
		fmt.Scan(&desde)
		color.Set(color.FgYellow)
		fmt.Printf("\nHasta:\n> ")
		color.Unset()
		var hasta int
		fmt.Scan(&hasta)
		matriz_1, err := generar_matriz_cuadrada_aleatoria(tamanio, desde, hasta)
		matriz = matriz_1
		if err != nil {
			fmt.Println(err)
			return false
		}
		color.Set(color.FgMagenta)
		fmt.Printf("\nMatriz inicial:\n")
		color.Unset()
		imprimir_matriz(matriz_1)
	case 2:
		color.Set(color.FgYellow)
		fmt.Printf("\nIngrese las filas de la matriz:\n> ")
		color.Unset()
		filas := 0
		fmt.Scan(&filas)
		color.Set(color.FgYellow)
		fmt.Printf("\nIngrese las columnas de la matriz:\n> ")
		color.Unset()
		columnas := 0
		fmt.Scan(&columnas)
		color.Set(color.FgYellow)
		fmt.Printf("\nGenerar numeros aleatorios desde:\n> ")
		color.Unset()
		var desde int
		fmt.Scan(&desde)
		color.Set(color.FgYellow)
		fmt.Printf("\nHasta:\n> ")
		color.Unset()
		var hasta int
		fmt.Scan(&hasta)
		matriz_1, err := generar_matriz_nxm_aleatoria(filas, columnas, desde, hasta)
		matriz = matriz_1
		if err != nil {
			fmt.Println(err)
			return false
		}
		color.Set(color.FgMagenta)
		fmt.Printf("\nMatriz inicial:\n")
		color.Unset()
		imprimir_matriz(matriz_1)
	default:
		println("No existe la opción.")
	}

	for {
		bandera := 0
		color.Set(color.FgYellow)
		println("¿Qué operación desea ejecutar?")
		fmt.Printf("\n\t(1)Transpuesta.		   	(2)Traza.")
		fmt.Printf("\n\t(3)Sumar otra matriz.   (4)Restar otra matriz.")
		fmt.Printf("\n\t(5)Multiplicar escalar.	(6)Multiplicar por otra matriz.\n> ")
		color.Unset()
		fmt.Scan(&opcion)
		switch opcion {
		case 1:
			color.Set(color.FgMagenta)
			println("Matriz transpuesta:")
			color.Unset()
			imprimir_matriz(obtener_transpuesta(matriz))
		case 2:
			traza, err := obtener_traza(matriz)
			if err != nil {
				fmt.Println(err)
				return false
			}
			color.Set(color.FgHiMagenta)
			println("Traza:")
			color.Unset()
			println(traza)
		case 3:
			matriz_resultado, err := sumar_matrices(matriz, matriz)
			if err != nil {
				fmt.Println(err)
				return false
			}
			color.Set(color.FgMagenta)
			println("Matriz resultado:")
			color.Unset()
			imprimir_matriz(obtener_transpuesta(matriz_resultado))
		case 4:
			matriz_resultado, err := restar_matrices(matriz, matriz)
			if err != nil {
				fmt.Println(err)
				return false
			}
			color.Set(color.FgMagenta)
			println("Matriz resultado:")
			color.Unset()
			imprimir_matriz(obtener_transpuesta(matriz_resultado))
		case 5:
			color.Set(color.FgYellow)
			fmt.Printf("\nIngrese el escalar:\n> ")
			color.Unset()
			var escalar int
			fmt.Scan(&escalar)
			color.Set(color.FgMagenta)
			println("Matriz resultado:")
			color.Unset()
			imprimir_matriz(escalar_por_matriz(escalar, matriz))
		case 6:
			matriz_resultado, err := multiplicar_matrices(matriz, matriz)
			if err != nil {
				fmt.Println(err)
				return false
			}
			color.Set(color.FgMagenta)
			println("Matriz resultado:")
			color.Unset()
			imprimir_matriz(obtener_transpuesta(matriz_resultado))
		default:
			bandera = 1
		}
		if bandera != 0 {
			break
		}
	}

	return true
}

func main() {
	opcion := 0

	for {
		color.Set(color.FgGreen)
		fmt.Printf("\nIngrese la opción a probar: ")
		fmt.Printf("\n\n(1)Suma de matrices.\t\t(2)Escalar por matríz.\n")
		fmt.Printf("(3)Producto de matrices.\t(4)Traza de una matríz.\n")
		fmt.Printf("(5)Transpuesta de una matríz.\t(6)Probar.\n\n> ")
		color.Unset()

		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			programa_1()
		case 2:
			programa_2()
		case 3:
			programa_3()
		case 4:
			programa_4()
		case 5:
			programa_5()
		case 6:
			test()
		default:

		}

		if opcion <= 0 || opcion >= 7 {
			color.Set(color.FgCyan)
			imprimirLentoln("\n\nTen un buen dia! UwU")
			color.Unset()
			break
		} else {
			color.Set(color.FgGreen)
			imprimirRapidoln("Probar otro programa?")
			color.Unset()
		}

	}

	color.Set(color.FgCyan)
	imprimirLentoln("\n\n< By LuisFlahan4051  ;) />\n\n")
	color.Unset()
	time.Sleep(2 * time.Second)
}
