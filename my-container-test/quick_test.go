package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttp(t *testing.T) {
	resp, err := http.Get("http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	assert.Equal(t, 200, resp.StatusCode)
}
