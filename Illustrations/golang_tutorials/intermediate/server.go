package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func main() {
    Articles = []Article{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	
    handleRequests()
}

//https://tutorialedge.net/golang/creating-restful-api-with-golang/
//http://localhost:10000/articles
//[{"Title":"Hello","desc":"Article Description","content":"Article Content"},{"Title":"Hello 2","desc":"Article Description","content":"Article Content"}]