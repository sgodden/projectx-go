package main

import (
	"fmt"
	"io/ioutil"
	"projectx/persistence"
	"net/http"
	"encoding/json"
)

type Project struct {
	Self string
	Id string
	Key string
	Name string
}

func main() {
	client := &http.Client {}
	req, _ := http.NewRequest("GET", "https://blackcatsolutions.atlassian.net/rest/api/2/project", nil)
	req.Header.Add("Authorization", "Basic YWRtaW5pc3RyYXRvcjo0bXkgaGFtbSNyaGVhZA==")
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	
	projectSet := make([]Project, 1)
	
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &projectSet)
	
	fmt.Printf("There are %v projects:\n", len(projectSet))
	
	for _, proj := range projectSet {
		fmt.Println("\t", proj.Name)
	}
	
	return
	
	repo := persistence.NewCustomerOrderRepository()

	fmt.Println(repo.FindById(1))
}
