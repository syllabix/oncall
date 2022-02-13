package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"text/template"
)

//go:embed manifest.template.yml
var static embed.FS

var (
	scheme string
	host   string
)

type Data struct {
	BaseURL string
}

func main() {

loop:
	for {
		fmt.Println("Will the bot use https? [y/n]")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("An error occurred while reading input: %v", err)
		}
		input = strings.ToLower(strings.TrimSuffix(input, "\n"))
		switch input {
		case "y", "yes":
			scheme = "https"
			break loop
		case "n", "no":
			scheme = "http"
			break loop
		default:
			fmt.Println("Please answer with y or n")
		}
	}

	fmt.Println("What will be the host name for the bot? (example: www.myhost.com)")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("An error occurred while reading input: %v", err)
	}
	host := strings.ToLower(strings.TrimSuffix(input, "\n"))
	url := url.URL{Scheme: scheme, Host: host}

	file, err := os.Create("manifest.yml")
	if err != nil {
		log.Fatalf("failed to create application Slack bot manifest.yml file\nfailure reason: %w\n", err)
	}

	tmpl, err := template.ParseFS(static, "manifest.template.yml")
	if err != nil {
		log.Fatalf("failed to parse service page html: %w", err)
	}

	err = tmpl.Execute(file, Data{BaseURL: url.String()})
	if err != nil {
		log.Fatalf("Bummer...\nfailed to generate applicaiton Slack bot manifest\nreason: %w", err)
	}

	fmt.Println("Nice! Your Slack bot app manifest.yml is ready to go!")
}
