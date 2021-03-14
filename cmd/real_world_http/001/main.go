package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.oreilly.co.jp/catalog/soon.xml")
	body, _ := ioutil.ReadAll(res.Body)
	log.Print(string(body))
	res.Body.Close()
}
