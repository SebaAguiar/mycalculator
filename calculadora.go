package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(entrada string, operador string) int {
	clean := strings.Split(entrada, operador)
	op1 := parser(clean[0])
	op2 := parser(clean[1])
	switch operador {
	case "+":
		fmt.Println(op1 + op2)
		return op1 + op2
	case "-":
		fmt.Println(op1 - op2)
		return op1 - op2
	case "*":
		fmt.Println(op1 * op2)
		return op1 * op2
	case "/":
		fmt.Println(op1 / op2)
		return op1 / op2
	default:
		fmt.Println(operador, "No esta soportado")
		return 0
	}
}

func parser(entrada string) int {
	op, _ := strconv.Atoi(entrada)
	return op
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
