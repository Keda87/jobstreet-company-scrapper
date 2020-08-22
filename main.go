package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	urlTemplate = "https://api-js.prod.companyreview.co/companies?page=%d&company_name=&per_page=%d&sort=-reviews_count&api_key=jwt_jsIdBrowserKey"
	totalData   = 131_604
	itemPerPage = 20
)

func main() {
	totalJob := int(math.Floor(totalData / itemPerPage))
	jobs := make(chan Payload, totalJob)
	results := make(chan Company, totalJob)

	for i := 1; i <= 20; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= totalJob; i++ {
		url := fmt.Sprintf(urlTemplate, i, itemPerPage)

		payload := Payload{
			URL: url,
			UserAgent: GetRandomUserAgent(),
		}

		jobs <- payload
	}
	close(jobs)

	for r := range results {
		// TODO: insert to database
		fmt.Println(r.CompanyName)
		fmt.Println(r.Industry)
		fmt.Println("---")
	}

}

func worker(id int, jobs <-chan Payload, result chan<- Company) {
	for job := range jobs {
		fmt.Println("[FETCH]", job.URL)

		fmt.Println(job.UserAgent)

		random := rand.Intn(5)
		time.Sleep(time.Second * time.Duration(random))

		var resp Response
		jsonResponse := Fetch(job.URL, job.UserAgent)
		json.Unmarshal(jsonResponse, &resp)

		for _, company := range resp.Data {
			result <- company
		}
	}
}
