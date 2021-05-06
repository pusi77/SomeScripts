package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pusi77/SomeScripts/mtf/api"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Language string `yaml:"lang"`
	ApiKey   string `yaml:"apiKey"`
	Logging  bool   `yaml:"log"`
}

func main() {

	var config conf
	config.readConfig()

	if !config.Logging {
		log.SetOutput(ioutil.Discard)
	}

	filepath := os.Args[1]

	inputFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(filepath + ".out")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		// check if 1st letter is uppercase or line contains *
		if line[0] < 'a' || strings.Contains(line, "*") {
			_, err := fmt.Fprintln(outputFile, line)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		possibleTitles := api.Search(line, config.ApiKey, config.Language)
		correctTitle := match(line, possibleTitles)
		if correctTitle == "" {
			correctTitle = askUser(line, possibleTitles)
			if correctTitle == "" {
				correctTitle = line
			}
		} else {
			fmt.Printf("Correct title found for %q: %q\n", line, correctTitle)
		}
		fmt.Println("=========================================================")
		_, err := fmt.Fprintln(outputFile, correctTitle)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (c *conf) readConfig() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println("found: " + c.Language)
	if c.ApiKey == "" {
		fmt.Println("Api Key not found!")
		os.Exit(1)
	}
	return c
}

// Check if the title is correct, returns the correct title if it's not, or "".
func match(title string, titles []string) string {
	for _, t := range titles {
		if strings.EqualFold(title, t) {
			return t
		}
	}
	return ""
}

func askUser(title string, possibleTitles []string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("No direct match found for %q, select one of the following: \n", title)
	fmt.Println("Press Enter without input to skip.")
	for i, t := range possibleTitles {
		fmt.Printf("%d) %s\n", i, t)
	}

	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\n", "")
	log.Println("selected: " + input)

	if input == "" {
		return ""
	}
	respNum, err := strconv.Atoi(input)
	if err != nil {
		// you had your chance
		log.Println("input not valid, skipping")
	}
	if respNum < 0 || (respNum > len(possibleTitles)) {
		log.Fatal("choice not valid")
	}

	fmt.Println("Selected: " + possibleTitles[respNum])
	return possibleTitles[respNum]
}
