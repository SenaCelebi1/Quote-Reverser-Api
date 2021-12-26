package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	models "quoteapi/model"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Quotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	url := "https://type.fit/api/quotes"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var c []models.Quotes
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string][]string)

	for i := 0; i < len(c); i++ {
		m[c[i].Author] = append(m[c[i].Author], Reverse(c[i].Text))
	}

	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}

	fizzbuzz := []models.QuotesNew{}

	for k, v := range m {
		fizzbuzz = append(fizzbuzz, models.QuotesNew{Author: k, Quotes: v})

	}

	_ = json.NewDecoder(r.Body).Decode(&fizzbuzz)
	json.NewEncoder(w).Encode(fizzbuzz)
}
