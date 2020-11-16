package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jailtonjunior94/go-digimons/src/services"
)

var wg sync.WaitGroup

func main() {
	now := time.Now()

	for _, name := range services.GetNames() {
		go services.Request("https://digimon-api.vercel.app/api/digimon/name/"+name, &wg)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println("Duration: ", time.Since(now))
}
