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

func generar_matriz_cuadrada_aleatoria(tamanio int, numero_aleatorio_inicial int, numero_aleatorio_final int) [][]int {

	matriz := make([][]int, tamanio)

	for i := 0; i < tamanio; i++ {
		matriz[i] = make([]int, tamanio)
		for j := 0; j < tamanio; j++ {
			matriz[i][j] = generar_numero_random(numero_aleatorio_inicial, numero_aleatorio_final)
		}
	}
	return matriz
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
	if filas == 0 || columnas == 0 {
		if filas == 0 {
			err = errors.New("no se pueden generar matrices de dimensiones n por 0")
		}
		if columnas == 0 {
			err = errors.New("no se pueden generar matrices de dimensiones 0 por m")
		}
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

	matriz_2 := generar_matriz_cuadrada_aleatoria(tamanio, numero_aleatorio_inicial, numero_aleatorio_final)
	matriz_1 := generar_matriz_cuadrada_aleatoria(tamanio, numero_aleatorio_inicial, numero_aleatorio_final)
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

func programa_2() {
	escalar := 5
	tamanio := 4
	matriz := generar_matriz_cuadrada_aleatoria(tamanio, 1, 10)
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
	matriz := generar_matriz_cuadrada_aleatoria(tamanio, 1, 10)
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

func test() {
	println("Seccioón para hacer pruebas.")
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
