package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

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
	for _, r := range reviews {
		review := r.(map[string]interface{})
		fmt.Println(review["author"])
		fmt.Println(review["reviewBody"])
		fmt.Println(review["reviewRating"])
		fmt.Println(review["datePublished"])
	}

}
