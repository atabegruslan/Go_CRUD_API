package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
    "github.com/gorilla/mux"
	"io/ioutil"
	"strconv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Article struct {
	gorm.Model // id and timestamps
    Title string `json:"Title"`
    Desc string `json:"Desc"`
    Content string `json:"Content"`
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
	
    myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	var articles []Article

	db.Find(&articles)
	fmt.Println("{}", articles)

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
	
	articles := []Article{}
	db.Find(&articles)
	
	for _, article := range articles {
		// string to int
		k, err:= strconv.Atoi(key)
		if err == nil{
			if article.ID == uint(k) {
				json.NewEncoder(w).Encode(article)
			}
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
	
	db.Create(&Article{Title: article.Title, Desc: article.Desc, Content: article.Content})
	fmt.Fprintf(w, "New Article Successfully Created")
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
    var toUpdate Article 
    json.Unmarshal(reqBody, &toUpdate)
	
	var article Article
	db.Where("id = ?", key).Find(&article)
	
	article.Title = toUpdate.Title
	article.Desc = toUpdate.Desc
	article.Content = toUpdate.Content

	db.Save(&article)
	fmt.Fprintf(w, "Successfully Updated Article")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var article Article
	db.Where("id = ?", key).Find(&article)
	db.Delete(&article)

	fmt.Fprintf(w, "Successfully Deleted Article")
}

func main() {
    db, err = gorm.Open("mysql", "ruslan:ruslan@tcp(127.0.0.1:3306)/go_articles?parseTime=true")

    if err != nil {
        panic(err.Error())
    }
	
	db.AutoMigrate(&Article{})

    handleRequests()
}

//https://tutorialedge.net/golang/golang-orm-tutorial/
//https://levelup.gitconnected.com/build-a-rest-api-using-go-mysql-gorm-and-mux-a02e9a2865ee
//go get github.com/gorilla/mux
//go get github.com/jinzhu/gorm
