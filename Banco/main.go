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

	fmt.Println(contaCorrente == contaCorrente2) //true, pq compara o valor

	minhaConta2 := new(ContaCorrente)
	minhaConta2.titular = "outra"
	fmt.Println(minhaConta2 == minhaConta)   //false, pq compara pelo endereço de memoria
	fmt.Println(*minhaConta2 == *minhaConta) //true, pq compara o valor
	fmt.Println(&minhaConta2, &minhaConta)

}
