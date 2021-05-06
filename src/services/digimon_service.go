package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/jailtonjunior94/go-digimons/src/structs"
)

// GetNames is a function
func GetNames() []string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(dir, "src", "files", "digimon.txt")

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var digimons []string

	lines := bufio.NewScanner(file)
	for lines.Scan() {
		digimons = append(digimons, lines.Text())
	}

	return digimons
}

// Request is a function
func Request(url string, wg *sync.WaitGroup) {
	var digimons []structs.Digimon

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&digimons); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

	defer wg.Done()
	fmt.Println(digimons)
}
