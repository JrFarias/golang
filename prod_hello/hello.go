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

func main() {
	showIntroduction()
	for {
		showMenu()
		comando := readInput()

		switch comando {
		case 1:
			initMonitoring()
		case 2:
			imprimeLogs()
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
	sites := readFile()
	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso", resp.StatusCode)
		saveLog(site, true)
	} else {
		fmt.Println("Site", site, "esta com problemas", resp.StatusCode)
		saveLog(site, false)
	}
}

func readFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func saveLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	fmt.Println(site, status, file)
	file.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}

	fmt.Println(string(file))
}
