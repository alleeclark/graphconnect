package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	filePath string
)

func init() {
	flag.StringVar(&filePath, "filepath", "./urls.csv", "Path of csv")
}
func main() {
	urls := readFile(filePath)
	var urlArray []string
	for _, val := range urls {
		newURL := ParseCleanURL(val)
		urlArray = append(urlArray, newURL)
	}
	writeCSV(urlArray)
}

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var urlArry []string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		urlArry = append(urlArry, scanner.Text())
	}
	return urlArry

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
		log.Fatalf("Cannot create file %v", err)
		defer file.Close()
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(cleanURLS)
	if err != nil {
		log.Fatalf("Cannot write to file %v", err)
	}

}

//parseCleanURL does xyz
func ParseCleanURL(url string) string {
	url = strings.Replace(url, "//", "/", -1)
	return strings.Replace(url, "/", "~", -1)
}
