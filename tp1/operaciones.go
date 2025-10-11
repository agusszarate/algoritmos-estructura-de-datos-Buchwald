package main

import "math"

const (
	_OP_SUMA           = "+"
	_OP_RESTA          = "-"
	_OP_MULTIPLICACION = "*"
	_OP_DIVISION       = "/"
	_OP_POTENCIA       = "^"
	_OP_LOGARITMO      = "log"
	_OP_RAIZ           = "sqrt"
	_OP_TERNARIO       = "?"
)

type operacion struct {
	simbolo string
	aridad  int
	operar  func(operandos []int64) (int64, string)
}

func crearOperaciones() map[string]operacion {
	operaciones := make(map[string]operacion)

	operaciones[_OP_SUMA] = operacion{
		simbolo: _OP_SUMA,
		aridad:  2,
		operar: func(ops []int64) (int64, string) {
			return ops[0] + ops[1], ""
		},
	}

	operaciones[_OP_RESTA] = operacion{
		simbolo: _OP_RESTA,
		aridad:  2,
		operar: func(ops []int64) (int64, string) {
			return ops[0] - ops[1], ""
		},
	}

	operaciones[_OP_MULTIPLICACION] = operacion{
		simbolo: _OP_MULTIPLICACION,
		aridad:  2,
		operar: func(ops []int64) (int64, string) {
			return ops[0] * ops[1], ""
		},
	}

	operaciones[_OP_DIVISION] = operacion{
		simbolo: _OP_DIVISION,
		aridad:  2,
		operar: func(ops []int64) (int64, string) {
			if ops[1] == 0 {
				return 0, "division por cero"
			}
			return ops[0] / ops[1], ""
		},
	}

	operaciones[_OP_POTENCIA] = operacion{
		simbolo: _OP_POTENCIA,
		aridad:  2,
		operar: func(ops []int64) (int64, string) {
			if ops[1] < 0 {
				return 0, "exponente negativo"
			}
			return int64(math.Pow(float64(ops[0]), float64(ops[1]))), ""
		},
	}

	operaciones[_OP_LOGARITMO] = operacion{
		simbolo: _OP_LOGARITMO,
		aridad:  2,
		operar: func(ops []int64) (int64, string) {
			if ops[1] < 2 {
				return 0, "base invalida"
			}
			if ops[0] <= 0 {
				return 0, "argumento invalido"
			}
			return int64(math.Log(float64(ops[0])) / math.Log(float64(ops[1]))), ""
		},
	}

	operaciones[_OP_RAIZ] = operacion{
		simbolo: _OP_RAIZ,
		aridad:  1,
		operar: func(ops []int64) (int64, string) {
			if ops[0] < 0 {
				return 0, "raiz negativa"
			}
			return int64(math.Sqrt(float64(ops[0]))), ""
		},
	}

	operaciones[_OP_TERNARIO] = operacion{
		simbolo: _OP_TERNARIO,
		aridad:  3,
		operar: func(ops []int64) (int64, string) {
			if ops[0] != 0 {
				return ops[1], ""
			}
			return ops[2], ""
		},
	}

	return operaciones
}
