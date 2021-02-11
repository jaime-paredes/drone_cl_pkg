package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jaime-paredes/drone_cl_pkg/models"
	"github.com/joho/godotenv"
)

var (
	env string
)

func init() {
	if env == "production" {
		godotenv.Load("../../.env")
	} else {
		env = "development"
		godotenv.Load("../../.env.development")
	}
	fmt.Println(fmt.Sprintf("Starting %s mode: ", env))
}

func main() {

	cpnURL := os.Getenv("trackingAPI")
	cpnURL = cpnURL + "/v1/companies/production.json"
	resp, err := http.Get(cpnURL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var companies []models.Company
	if err = json.Unmarshal(body, &companies); err != nil {
		fmt.Println(err)
	}
	internal.CertCheck(companies)
}
