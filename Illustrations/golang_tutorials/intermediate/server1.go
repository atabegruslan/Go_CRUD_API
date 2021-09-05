package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
    "github.com/gorilla/mux"
	"io/ioutil"
)

type Article struct {
    Id string `json:"Id"`
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
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
	
    //http.HandleFunc("/", homePage)
    myRouter.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles", returnAllArticles)
	
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST") // Gorilla Mux allow you to define POST, PUT ...
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

    //log.Fatal(http.ListenAndServe(":10000", nil))
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    //fmt.Fprintf(w, "Key: " + key)
	
	// Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(r.Body)
	
	//If post raw JSON from POSTMAN
	//{ "Id": "3", "Title": "Newly Created Post", "desc": "The description for my new post", "content": "my articles content" }
	//This will regurgitate the JSON
    //fmt.Fprintf(w, "%+v", string(reqBody))
	
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
    var toUpdate Article 
    json.Unmarshal(reqBody, &toUpdate)
	
	//fmt.Fprintf(w, "%+v", string(reqBody))
	//fmt.Fprintf(w, "Key: " + key)
	
	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
			Articles = append(Articles, toUpdate)
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the article
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

func main() {
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	
    handleRequests()
}

//https://tutorialedge.net/golang/creating-restful-api-with-golang/
//go get github.com/gorilla/mux

//http://localhost:10000/article/1
//{"Id":"1","Title":"Hello","desc":"Article Description","content":"Article Content"}

//POST http://localhost:10000/article RAW BODY { "Id": "3", "Title": "Newly Created Post", "desc": "The description for my new post", "content": "my articles content" }
//http://localhost:10000/articles
//[{"Id":"1","Title":"Hello","desc":"Article Description","content":"Article Content"},{"Id":"2","Title":"Hello 2","desc":"Article Description","content":"Article Content"},{"Id":"3","Title":"Newly Created Post","desc":"The description for my new post","content":"my articles content"}]

//http://localhost:10000/article/3
//http://localhost:10000/articles
//[{"Id":"1","Title":"Hello","desc":"Article Description","content":"Article Content"},{"Id":"2","Title":"Hello 2","desc":"Article Description","content":"Article Content"}]

