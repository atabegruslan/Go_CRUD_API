package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "io/ioutil"
	"reflect"
)

type Place struct {
	id    int64  
	title string 
	body  string   
}
	
func main() {
    response, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
    data, _ := ioutil.ReadAll(response.Body)

	var places []Place
	json.Unmarshal(data, &places)
	fmt.Printf("%+v", places)
	fmt.Println(reflect.TypeOf(places))
}