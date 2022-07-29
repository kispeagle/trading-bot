package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GetPairSymbolList(url string) []string {

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal()
	}

	var symbolList []string
	r, _ := regexp.Compile("\"symbol\":\"([A-Z]+)\"")
	rs := r.FindAllStringSubmatch(string(data), -1)
	for _, m := range rs {
		symbolList = append(symbolList, m[1])
	}

	return symbolList

}
