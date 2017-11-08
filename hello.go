package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	showIntroduction()
	for {
		showMenu()
		comando := readInput()

		switch comando {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	nome := "Junior"
	fmt.Println("Ola sr:", nome)
}

func readInput() int {
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	const monitoramento = 5
	const delay = 5
	sites := []string{"https://www.alura.com.br", "https://www.google.com.br"}
	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso", resp.StatusCode)
	} else {
		fmt.Println("Site", site, "esta com problemas", resp.StatusCode)
	}
}
