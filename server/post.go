package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type incomingData struct {
	Title  string `json:"title"`
	
}

func singlePost(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data incomingData

	//decoded recieved data
	decoder.Decode(&data)
	fmt.Println("blog posts title", data.Title)

	var postdata []post
	db, err := sql.Open("sqlite3", "./database/database.db")
	checkErr(err)

	postdata = getPost(db, data)

	// fmt.Println(postdata)

	defer db.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(postdata); err != nil {
		panic(err)
	}
}

func getPost(db *sql.DB, data incomingData) []post {
	rows, err := db.Query("SELECT id, username, title, content, image, created FROM post WHERE title = ?", (data.Title))
	checkErr(err)
	postinfo := make([]post, 0)
	for rows.Next() { //for loop through database table
		onePost := post{}
		err = rows.Scan(&onePost.ID, &onePost.Username, &onePost.Title, &onePost.Content, &onePost.Image, &onePost.Created)
		checkErr(err)
		time := ""
		time = onePost.Created[:10]
		time += " " + onePost.Created[11:19]
		onePost.Created = time
		postinfo = append(postinfo, onePost)
	}
	return postinfo
}
