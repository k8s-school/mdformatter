package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/piprate/json-gold/ld"
)

func jsonld() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// expanding remote document

	expanded, err := proc.Expand("/home/fjammes/src/mdformatter/micro-services.json", options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}

	ld.PrintDocument("JSON-LD expansion succeeded", expanded)

	doc := map[string]interface{}{
		"@context":  "http://schema.org/",
		"@type":     "Person",
		"name":      "Jane Doe",
		"jobTitle":  "Professor",
		"telephone": "(425) 123-4567",
		"url":       "http://www.janedoe.com",
	}

	expanded, err = proc.Expand(doc, options)
	if err != nil {
		panic(err)
	}

	ld.PrintDocument("JSON-LD expansion succeeded", expanded)
}

func main() {

	var jsonFilename string

	flag.StringVar(&jsonFilename, "json", "", "path to json-ld file")
	// md := flag.String("md", "", "path to md file")
	flag.Parse()

	// Open our jsonFile
	jsonFile, err := os.Open(jsonFilename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s\n", jsonFilename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	data := result["data"]
	jsonlist := data.([]interface{})
	product := jsonlist[0].(map[string]interface{})
	reviews := product["review"].([]interface{})
	// <div class="Stars" style="--rating: 5;" aria-label="La note de ce produit est 5 sur 5.">Denis C. de Hewlett-Packard: Formateur très professionnel et très compétent sur le sujet K8s.</div>
	for _, r := range reviews {
		review := r.(map[string]interface{})
		reviewRating := review["reviewRating"].(map[string]interface{})
		author := review["author"].(map[string]interface{})
		fmt.Printf("<div class=\"Stars\" style=\"--rating: 5;\" aria-label=\"La note de ce produit est %v sur 5.\">", reviewRating["ratingValue"])
		fmt.Printf("%s, %s</div>\n", author["name"], review["reviewBody"])
		// fmt.Println(review["datePublished"])
	}

}
