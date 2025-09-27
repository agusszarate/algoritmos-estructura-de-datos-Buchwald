package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"tdas/pila"
)

const _MENSAJE_ERROR = "ERROR"

type calculadoraRPN struct {
	pila pila.Pila[int64]
}

func CrearCalculadoraRPN() *calculadoraRPN {
	return &calculadoraRPN{
		pila: pila.CrearPilaDinamica[int64](),
	}
}

func (calc *calculadoraRPN) EvaluarExpresion(linea string) string {
	calc.vaciarPila()

	tokens := strings.Fields(linea)
	if len(tokens) == 0 {
		return _MENSAJE_ERROR
	}

	for _, token := range tokens {
		err := calc.procesarToken(token)
		if err != "" {
			calc.vaciarPila()
			return _MENSAJE_ERROR
		}
	}

	if calc.pila.EstaVacia() {
		return _MENSAJE_ERROR
	}

	resultado := calc.pila.Desapilar()

	if !calc.pila.EstaVacia() {
		calc.vaciarPila()
		return _MENSAJE_ERROR
	}

	return fmt.Sprintf("%d", resultado)
}

func (calc *calculadoraRPN) vaciarPila() {
	for !calc.pila.EstaVacia() {
		calc.pila.Desapilar()
	}
}

func (calc *calculadoraRPN) procesarToken(token string) string {
	numero, err := strconv.ParseInt(token, 10, 64)

	if err == nil {
		calc.pila.Apilar(numero)
		return ""
	}

	return calc.procesarOperador(token)
}

func (calc *calculadoraRPN) procesarOperador(operador string) string {
	switch operador {
	case "+":
		return calc.ejecutarOperacion(2, func(ops []int64) (int64, string) {
			return ops[0] + ops[1], ""
		})
	case "-":
		return calc.ejecutarOperacion(2, func(ops []int64) (int64, string) {
			return ops[0] - ops[1], ""
		})
	case "*":
		return calc.ejecutarOperacion(2, func(ops []int64) (int64, string) {
			return ops[0] * ops[1], ""
		})
	case "/":
		return calc.ejecutarOperacion(2, func(ops []int64) (int64, string) {
			if ops[1] == 0 {
				return 0, "division por cero"
			}
			return ops[0] / ops[1], ""
		})
	case "^":
		return calc.ejecutarOperacion(2, func(ops []int64) (int64, string) {
			if ops[1] < 0 {
				return 0, "exponente negativo"
			}
			resultado := int64(math.Pow(float64(ops[0]), float64(ops[1])))
			return resultado, ""
		})
	case "log":
		return calc.ejecutarOperacion(2, func(ops []int64) (int64, string) {
			if ops[1] < 2 {
				return 0, "base invalida"
			}
			if ops[0] <= 0 {
				return 0, "argumento invalido"
			}
			resultado := int64(math.Log(float64(ops[0])) / math.Log(float64(ops[1])))
			return resultado, ""
		})
	case "sqrt":
		return calc.ejecutarOperacion(1, func(ops []int64) (int64, string) {
			if ops[0] < 0 {
				return 0, "raiz negativa"
			}
			resultado := int64(math.Sqrt(float64(ops[0])))
			return resultado, ""
		})
	case "?":
		return calc.ejecutarOperacion(3, func(ops []int64) (int64, string) {
			if ops[0] != 0 {
				return ops[1], ""
			}
			return ops[2], ""
		})
	default:
		return "operador invalido"
	}
}

func (calc *calculadoraRPN) ejecutarOperacion(numOperandos int, operacion func([]int64) (int64, string)) string {
	operandos := make([]int64, numOperandos)
	for i := 0; i < numOperandos; i++ {
		if calc.pila.EstaVacia() {
			return "faltan operandos"
		}
		operandos[numOperandos-1-i] = calc.pila.Desapilar()
	}

	resultado, err := operacion(operandos)
	if err != "" {
		return err
	}

	calc.pila.Apilar(resultado)
	return ""
}

func main() {
	calc := CrearCalculadoraRPN()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		resultado := calc.EvaluarExpresion(linea)
		fmt.Println(resultado)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error al leer entrada: %v\n", err)
	}
}
