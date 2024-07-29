package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&requests, "requests", 0, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas")
}

func main() {
	flag.Parse()

	if url == "" || requests <= 0 || concurrency <= 0 {
		fmt.Println("Uso: --url=http://example.com --requests=1000 --concurrency=10")
		return
	}

	fmt.Printf("Iniciando teste de carga em %s com %d requests e %d concorrências...\n", url, requests, concurrency)

	start := time.Now()

	var wg sync.WaitGroup
	requestsChan := make(chan int, requests)
	statusCodes := make(map[int]int)
	mu := sync.Mutex{}

	for i := 0; i < requests; i++ {
		requestsChan <- i
	}
	close(requestsChan)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range requestsChan {
				resp, err := http.Get(url)
				mu.Lock()
				if err != nil {
					fmt.Println("Erro ao realizar request:", err)
					statusCodes[0]++ // Contar os erros como status code 0
				} else {
					statusCodes[resp.StatusCode]++
					resp.Body.Close()
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("\nTeste de carga concluído em %s\n", elapsed)
	fmt.Printf("Total de requests: %d\n", requests)

	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("Status HTTP 200: %d\n", statusCodes[http.StatusOK])

	for code, count := range statusCodes {
		if code != http.StatusOK {
			fmt.Printf("Status HTTP %d: %d\n", code, count)
		}
	}
}
