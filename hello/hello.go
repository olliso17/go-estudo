package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
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
	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("o site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("o site", site, "está com problemas", resp.StatusCode)
		registraLog(site, false)

	}
}

// <nil> é igual a null
func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n') //no caso ele so ler uma linha do arquivo
		linha = strings.TrimSpace(linha)      //para tirar o qubra linha do arquivo
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //ler e criar arquivo com permissões
	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- online:" + strconv.FormatBool(status) + "\n") //passando uma string com tempo, site, formatando o bool em string e dando espaço

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
	//não precisa abrir e fechar
}
