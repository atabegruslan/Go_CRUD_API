package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml"
		"strings")

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}
	
type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}
	
type NewsMap struct {
	Keyword string
	Location string
}
	
// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
// 	fmt.Fprintf(w, `<p>multi</p>
// 					<p>line</p>`)
// 	fmt.Fprintf(w, "aaa %s ccc %d", "bbb", 4)
// }

func main() {
	// http.HandleFunc("/", index_handler)
	// http.ListenAndServe(":8000", nil)

	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	var n News

	news_map := make(map[string]NewsMap)

	xml.Unmarshal(bytes, &s)

	//fmt.Printf("%+v\n", s.Locations) // Dump structs

	for _, location := range s.Locations {
		//fmt.Println("\n%s", location)

		resp, _ := http.Get(strings.Trim(location, " \n"))
		bytes, _ := ioutil.ReadAll(resp.Body)
		string_body := string(bytes)
		fmt.Println(string_body)
		xml.Unmarshal(bytes, &n)

		for index, _ := range n.Titles {
			news_map[n.Titles[index]] = NewsMap{n.Keywords[index], n.Locations[index]}
		}
	}

	// for index, data := range news_map {
	// 	fmt.Println("\n\n\n", index)
	// 	fmt.Println("\n", data.Keyword)
	// 	fmt.Println("\n", data.Location)
	// }
}
