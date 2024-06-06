package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var cep string

func init() {
	if len(os.Args) < 2 || len(os.Args[1]) != 8 {
		fmt.Println("Favor inserir um CEP válido no arg da linha de comando")
		os.Exit(1)
	}
	cep = os.Args[1]

	if _, err := strconv.Atoi(cep); err != nil {
		fmt.Println("Cep inválido, favor inserir um CEP numérico.")
		os.Exit(1)
	}
}

func buscaBrasilApi(c chan<- string) {
	res, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	c <- string(body)
}

func buscaViaCep(c chan<- string) {
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	c <- string(body)
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go buscaBrasilApi(c1)
	go buscaViaCep(c2)

	select {
	case res := <-c1:
		fmt.Printf("%s\nRecebido por Brasil API\n", res)

	case res := <-c2:
		fmt.Printf("%s\nRecebido por ViaCep\n", res)

	case <-time.After(time.Second):
		fmt.Println("Timeout: nenhuma das apis responderam em tempo hábil")
	}

}
