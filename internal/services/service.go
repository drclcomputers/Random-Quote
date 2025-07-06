package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"quotes/internal/logger"
	"quotes/internal/model"
)

func Start() {
	CheckArgs()
}

func CheckArgs() {
	switch len(os.Args) {
	case 1:
		fmt.Println(GetQuote())
	case 2:
		switch os.Args[1] {
		case "-z":
			data := ContactZenquotesAPI()
			if data == "error" {
				fmt.Println(Fallback())
				return
			}
			fmt.Println(data)
		case "-q":
			data := ContactQuotableAPI()
			if data == "error" {
				fmt.Println(Fallback())
				return
			}
			fmt.Println(data)
		case "-s":
			SaveQuoteFile(GetQuote())
		case "-h":
			fmt.Print("Random Quote Generator -ver 0.3\n\n-s - Save the quote to a txt file.\n-z - Use the Zenquotes API.\n-q - Use the Quotable API.\n-h - Show help.\n\n")
		default:
			fmt.Println("Unknown argument passed!\n- This app")
		}
	default:
		fmt.Println("Too many arguments passed!\n- This app")
	}
}

func GetQuote() string {
	quote := ContactZenquotesAPI()
	if quote == "error" {
		quote = ContactQuotableAPI()
		if quote == "error" {
			return Fallback()
		}
		return quote
	}
	return quote
}

// ContactZenquotesAPIWithClient allows injecting a custom HTTP client (for testing)
func ContactZenquotesAPIWithClient(client *http.Client) string {
	resp, err := client.Get(model.ZENQUOTES_URL)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	logger.CheckError(err)
	var data []model.ZenquoteResponse
	err = json.Unmarshal(body, &data)
	logger.CheckError(err)
	return fmt.Sprintf("%s\n- %s", data[0].Content, data[0].Author)
}

func ContactZenquotesAPI() string {
	return ContactZenquotesAPIWithClient(http.DefaultClient)
}

// ContactQuotableAPIWithClient allows injecting a custom HTTP client (for testing)
func ContactQuotableAPIWithClient(client *http.Client) string {
	log.Printf("Warning! Using the less secure API, Quotable API.")
	resp, err := client.Get(model.QUOTABLE_URL)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	logger.CheckError(err)
	var data model.QuotableResponse
	err = json.Unmarshal(body, &data)
	logger.CheckError(err)
	return fmt.Sprintf("%s\n- %s", data.Content, data.Author)
}

func ContactQuotableAPI() string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return ContactQuotableAPIWithClient(client)
}

func Fallback() string {
	return "Unfortunately, our APIs aren't available right now. Try again later!\n- Every single app in existence"
}

func SaveQuoteFile(quote string) error {
	num := 1
	var filename string
	for {
		filename = fmt.Sprintf("quote_%d.txt", num)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			break
		}
		num++
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(quote)
	if err != nil {
		log.Fatalf("Failed to write to file: %v\n", err)
		return err
	}

	log.Printf("Quote saved to %s\n- This app", filename)
	return nil
}
