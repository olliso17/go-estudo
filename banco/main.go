// toda a aplicação começa com esse package
package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// podemos utilizar ou não todas as variaveis
func main() {
	contaDaPatricia := ContaCorrente{
		titular:       "Patriciar",
		numeroAgencia: 558,
		numeroConta:   1234,
		saldo:         12.5}

	contaDaMari := ContaCorrente{"Mariana", 523, 12342, 122.5}

	fmt.Println(contaDaPatricia, contaDaMari)

}
