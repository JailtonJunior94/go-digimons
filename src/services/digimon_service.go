package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jailtonjunior94/go-digimons/src/structs"
)

// GetNames is a function
func GetNames() []string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "src", "files", "digimon.txt")

	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

// Request is a function
func Request(url string, wg *sync.WaitGroup) {
	var digimons []structs.Digimon

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&digimons); err != nil {
		if err != nil {
			panic(err)
		}
	}

	defer wg.Done()
	fmt.Println(digimons)
}
