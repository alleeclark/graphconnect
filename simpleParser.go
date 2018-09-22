package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	filePath string
)

const filecsv = ""

func init() {
	flag.StringVar(&filePath, "filepath", "./urls.csv", "Path of csv")
}

func main() {
	urls := readCSV(filePath)
	var urlArray []string
	for _, val := range urls {
		newURL := ParseCleanURL(val)
		urlArray = append(urlArray, newURL)
	}
	writeCSV(urlArray)
}

func readCSV(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("Unable to open file, %v", err)
	}
	urls, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalf("Unable to read tilda sv file %v", err)
	}
	var urlArry []string
	for _, url := range urls {
		for _, val := range url {
			urlArry = append(urlArry, val)
		}
	}
	return urlArry
}
func writeCSV(cleanURLS []string) {
	file, err := os.Create("./urlsclean.csv")
	if err != nil {
		log.Fatalf("Cannot create file", err)
		defer file.Close()
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(cleanURLS)
	if err != nil {
		log.Fatalf("Cannot write to file", err)
	}

}

//parseCleanURL does xyz
func ParseCleanURL(url string) string {
	url = strings.Replace(url, "//", "/", -1)
	return strings.Replace(url, "/", "~", -1)
}
