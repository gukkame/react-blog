package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type post struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Content  string `json:"content"`
	Created  string `json:"created"`
	Image    string `json:"image"`
}

func home(w http.ResponseWriter, req *http.Request) {

	json.NewDecoder(req.Body)
	// var title string

	//decoded recieved data
	// decoder.Decode(&title)
	// fmt.Println("blog posts title", title)

	var postdata []post
	db, err := sql.Open("sqlite3", "./database/database.db")
	checkErr(err)

	postdata = getAllPosts(db)

	// fmt.Println(postdata)

	defer db.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(postdata); err != nil {
		panic(err)
	}
}

func getAllPosts(db *sql.DB) []post {
	rows, err := db.Query("SELECT id, username, title, content, created, image FROM post")
	checkErr(err)
	postinfo := make([]post, 0)
	for rows.Next() { //for loop through database table
		onePost := post{}
		err = rows.Scan(&onePost.ID, &onePost.Username, &onePost.Title, &onePost.Content, &onePost.Created, &onePost.Image)
		checkErr(err)
		time := ""
		time = onePost.Created[:10]
		time += " " + onePost.Created[11:19]
		onePost.Created = time
		onePost.Slug = strings.Replace(onePost.Title, " ", "-", -1)
		postinfo = append(postinfo, onePost)
	}
	return postinfo
}
