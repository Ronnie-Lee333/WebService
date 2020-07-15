package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

// Post is table Post's struct
type Post struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	// operation with table posts
	http.HandleFunc("/post/", handleRequest)

	// table posts and other addition info
	http.HandleFunc("/japi/", handleJSON)
	server.ListenAndServe()
}

// main handler function
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// json handler function with goroutine
func handleJSON(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = getJSONWithID(w, r)
	default:
		fmt.Println("NG except GET.")
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve a post data from DB in a goroutine
// and retrieve other data in antoher goroutine.
// finally connect the two data to return to client.
// GET /japi/1
func getJSONWithID(w http.ResponseWriter, r *http.Request) (err error) {

	chStr1 := make(chan string, 10)
	chStr2 := make(chan string, 10)

	// get data from table posts in a goroutine
	go func() {
		id, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			return
		}
		post, err := retrieve(id)
		if err != nil {
			return
		}
		output, err := json.MarshalIndent(&post, "", "\t\t")
		if err != nil {
			return
		}
		// fmt.Println(string(output))
		chStr1 <- string(output)
	}()

	// get other data in another goroutine
	go func() {
		chStr2 <- `{
		"addition1" : "addition infomation 1",
		"addition2" : "addition infomation 2"
}`
	}()

	outStr1 := <-chStr1
	outStr2 := <-chStr2
	outALL := outStr1 + "\r\n" + outStr2

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outALL))
	return
}

// Retrieve a post
// GET /post/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Create a post
// POST /post/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	fmt.Println(string(body))
	json.Unmarshal(body, &post)

	// fmt.Println(post.Id)
	// fmt.Println(post.Content)
	// fmt.Println(post.Author)

	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// Update a post
// PUT /post/1
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// Delete a post
// DELETE /post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
