package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Padroes de Concorrencia em Go
// <-chan - Canal somente de leitura

func titulo (urls ...string) <-chan string {
	c := make (chan string)
	// Criando uma go routine para cada busca de url
	for _, url := range urls {
		go func (url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r,_  := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
			}(url)
	}
	return c
}

func main(){
	fmt.Println("Buscando titulos...")
	t1 := titulo("https://www.cod3r.com.br","https://www.google.com")
	t2 := titulo("https://www.youtube.com", "https://www.apple.com")
	fmt.Println("Primeiros :",<-t1 ,"|" ,<-t2)
	fmt.Println("Segundos :",<-t1 ,"|" ,<-t2)
}
