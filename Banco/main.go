package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	saldo         float64
	numeroConta   int
}

func main() {
	contaCorrente := ContaCorrente{titular: "jouse", numeroAgencia: 2565, saldo: 25.5, numeroConta: 123}
	fmt.Println(contaCorrente)

	minhaConta := new(ContaCorrente)
	minhaConta.titular = "outra"

	fmt.Println(*minhaConta)

	contaCorrente2 := ContaCorrente{titular: "jouse", numeroAgencia: 2565, saldo: 25.5, numeroConta: 123}

	fmt.Println(contaCorrente == contaCorrente2)

}
