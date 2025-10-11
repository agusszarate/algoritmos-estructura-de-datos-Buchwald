package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/pila"
)

const _MENSAJE_ERROR = "ERROR"

type calculadoraRPN struct {
	pila        pila.Pila[int64]
	operaciones map[string]operacion
}

func CrearCalculadoraRPN() *calculadoraRPN {
	return &calculadoraRPN{
		pila:        pila.CrearPilaDinamica[int64](),
		operaciones: crearOperaciones(),
	}
}

func (calc *calculadoraRPN) EvaluarExpresion(linea string) string {
	calc.vaciarPila()

	tokens := strings.Fields(linea)
	if len(tokens) == 0 {
		return _MENSAJE_ERROR
	}

	for _, token := range tokens {
		if err := calc.procesarToken(token); err != "" {
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

func (calc *calculadoraRPN) procesarOperador(simbolo string) string {
	op, existe := calc.operaciones[simbolo]
	if !existe {
		return "operador invalido"
	}

	operandos := make([]int64, op.aridad)
	for i := 0; i < op.aridad; i++ {
		if calc.pila.EstaVacia() {
			return "faltan operandos"
		}
		operandos[op.aridad-1-i] = calc.pila.Desapilar()
	}

	resultado, err := op.operar(operandos)
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
