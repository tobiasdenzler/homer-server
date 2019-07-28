package main

import (
	"log"

	"github.com/tobiasdenzler/homer-server/pkg/dss"
)

func main() {

	result, err := dss.Call("/json/property/query", map[string]string{"query": "/apartment/zones/*(*)"})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(result)
	}
}
