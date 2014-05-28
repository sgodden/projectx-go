package main

import (
	"fmt"
	"projectx/persistence"
	"net/http"
)

func main() {
	client := &http.Client {}
	req, _ := http.NewRequest("GET", "https://blackcatsolutions.atlassian.net/rest/api/2/project", nil)
	req.Header.Add("Authorization", "Basic YWRtaW5pc3RyYXRvcjo0bXkgaGFtbSNyaGVhZA==")
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	
	return
	
	repo := persistence.NewCustomerOrderRepository()

	fmt.Println(repo.FindById(1))
}
