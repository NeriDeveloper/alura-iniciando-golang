package main

import "fmt"
import "os"
import "net/http"
import "time"

const STATUS_CODE_OK = 200
const MENU_INICIAR_MONITORAMENT = 1
const MENU_EXIBIR_LOGS = 2
const MENU_SAIR_DO_PROGRAMA = 0
const RODADA_MONITORAMENTOS = 3
const TIME_DELAY_MONITORAMENTOS = 5

func main() {
	exibirIntroducao()
	
	for {
		exibirMenu()
		executarComandoDigitado(lerComandoDigitado())
	}
}

func exibirIntroducao() {
	nome := "Malaquias"
	versao := 0.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)
	
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComandoDigitado() int {
	var comando int
	
	fmt.Scan(&comando)
	fmt.Println("O comando digitado foi: ", comando)

	return comando
}

func executarComandoDigitado(comando int) {
switch comando {
	case MENU_INICIAR_MONITORAMENT:
		iniciarMonitoramento()
	case MENU_EXIBIR_LOGS:
		fmt.Println("Exibindo logs...")
	case MENU_SAIR_DO_PROGRAMA:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"http://random-status-code.herokuapp.com/"}

    for i := 0; i < RODADA_MONITORAMENTOS; i++ {
		for i, site := range sites {
	        fmt.Println("Testando site", i, ":", site)
	        testaSite(site)
		}
		
		time.Sleep(TIME_DELAY_MONITORAMENTOS * time.Second)
		imprimirLinhaEmBranco()
    }

	imprimirLinhaEmBranco()
}

func testaSite(site string) {
	res, _ := http.Get(site)
	
	if res.StatusCode == STATUS_CODE_OK {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "está com problemas. Status Code: ", res.StatusCode)
	}
}

func imprimirLinhaEmBranco() {
	fmt.Println("")
}