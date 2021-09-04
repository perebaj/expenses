package connectors

import (
	"encoding/json"
	"expenses-control/pkg/structures"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Fetch() {
	NUBANK_EXPENSES_KEY := goDotEnvVariable("NUBANK_EXPENSES_KEY")
	url := "https://prod-global-webapp-proxy.nubank.com.br/api/proxy/" + NUBANK_EXPENSES_KEY
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	godotenv.Load("src/.env")
	NUBANK_KEY := goDotEnvVariable("NUBANK_KEY")

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("X-Correlation-Id", "WEB-APP.utcSk")
	req.Header.Add("Authorization", NUBANK_KEY)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	req.Header.Add("Origin", "https://app.nubank.com.br")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,pt;q=0.8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var jsonResponse structures.Expense
	json.Unmarshal([]byte(string(body)), &jsonResponse)

	expenseList := jsonResponse.Bill.LineItems

	for _, val := range expenseList {
		fmt.Printf(val.Title)
	}

}
