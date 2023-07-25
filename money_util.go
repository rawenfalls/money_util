package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Date   string   `xml:"Date,attr"`
	Name   string   `xml:"name,attr"`
	Valute []Valute `xml:"Valute"`
}

func getMoneyInfo(code, date string) {
	url := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily.asp?date_req=%s", date)

	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// проблема с 403 forbidden
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var currentQuota ValCurs
	// решение проблемы xml: encoding "windows-1251" declared but Decoder.CharsetReader is nil
	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&currentQuota)

	for _, val := range currentQuota.Valute {
		if strings.ToLower(val.CharCode) == strings.ToLower(code) {
			fmt.Printf("%s (%s): %s\n", val.CharCode, val.Name, val.Value)
		}
	}
}

func main() {
	code := flag.String("code", "", "код валюты в формате ISO 4217")
	currentDate := flag.String("date", "", "дата в формате YYYY-MM-DD")
	flag.Parse()

	if *code == "" {
		log.Fatalln("Необходимо указать код валюты")
	}

	if *currentDate != "" {
		date, err := time.Parse("2006-01-02", *currentDate)
		if err != nil {
			log.Fatalln("Неверный формат даты", err)
		}
		getMoneyInfo(*code, date.Format("02/01/2006"))
	} else {
		getMoneyInfo(*code, "")
	}
}
