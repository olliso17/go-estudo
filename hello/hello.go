package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const dalay = 5

func main() {

	exibeIntroducao()
	for {
		exibreMenu()
		//retorno da função
		comando := lerComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("exibindo logs")
		case 3:
			fmt.Println("saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	//sempre tenho que usar uma variável que eu declaro
	nome := "Patricia"
	versao := 1.1
	fmt.Println("Olá ", nome)
	fmt.Println("Este programa está na versão", versao)

}
func lerComando() int {
	//para o usuario poder digitar usa abaixo, %d -> modificador e o &comando ->endereço da variável comando dentro do computador
	var comandoLido int

	fmt.Scan(&comandoLido)

	fmt.Println("o comandoLido escolhido foi", comandoLido)

	return comandoLido
}

func exibreMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("3- Sair do Programa")

}

func iniciarMonitoramento() {
	fmt.Println("monitorando")
	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("testando site", i, ":", site)
			//quando eu tiver uma função que retorna dois valores vc pode so retornar um com exemplo abaixo usando
			//resp,_:= ao inves de resp, err
			testaSite(site)

		}
		time.Sleep(dalay * time.Minute)
	}
	//posição i e o item da posição site

}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("o site", site, "foi carregado com sucesso")

	} else {
		fmt.Println("o site", site, "está com problemas", resp.StatusCode)
	}
}
