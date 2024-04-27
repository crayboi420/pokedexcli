package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func testThis(){
	url := "https://pokeapi.co/api/v2/location-area/{id or name}/"
	res,err := http.Get(url)
	if err!=nil {log.Fatal(err)}
	fmt.Println(io.ReadAll(res.Body))
}