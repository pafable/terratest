package main

// EXAMPLE OF MAKING HTTP GET REQUEST
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stdout.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
