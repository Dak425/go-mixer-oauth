package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/gocolly/colly"
)

type scrapedScopeCollection []scrapedScope

type scrapedScope struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var matchColon = regexp.MustCompile(":")

func toCamelCase(text string) string {
	var camel strings.Builder

	words := matchColon.Split(text, -1)

	for _, word := range words {
		camel.WriteString(strings.Title(word))
	}

	return camel.String()
}

func main() {
	url := "https://dev.mixer.com/reference/oauth/scopes"

	rowSelector := "#body__padding > table > tbody > tr"
	cellSelector := "td"

	scopes := scrapedScopeCollection{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnHTML(rowSelector, func(e *colly.HTMLElement) {
		scope := scrapedScope{}

		e.ForEach(cellSelector, func(index int, td *colly.HTMLElement) {
			if index == 0 {
				scope.Name = td.Text
			} else {
				scope.Description = td.Text
			}
		})

		scopes = append(scopes, scope)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Complete")
		log.Printf("Compiled Scope Map: %+v \n", scopes)
		data, err := json.MarshalIndent(scopes, "", "	")
		if err != nil {
			log.Fatalf("error when marshalling scope collection to JSON: %v", err)
		}

		ioutil.WriteFile("gen/scopeMap.json", data, 0644)

		funcMap := template.FuncMap{
			"ToTitle":     strings.Title,
			"ToCamelCase": toCamelCase,
		}

		t := template.Must(template.New("scope.tmpl").Funcs(funcMap).ParseFiles("templates/scope.tmpl"))
		file, err := os.OpenFile("pkg/oauth/scope.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalf("unable to open file for writing scopes to: %v", err)
		}
		err = t.Execute(file, scopes)
		if err != nil {
			log.Fatalf("error when processing scraped scope to template: %v", err)
		}
		file.Close()
	})

	c.Visit(url)
}
