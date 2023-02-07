package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Created  string `json:"created"`
	Category string `json:"category"`
	Username string `json:"username"`
}

func home(w http.ResponseWriter, req *http.Request) {
	json.NewDecoder(req.Body)
	var postdata []post
	// var cookieData cookie
	// decoder.Decode(&cookieData)
	db, err := sql.Open("sqlite3", "./database/database.db")
	checkErr(err)
	// check := userCheck(cookieData.Username, cookieData.Cookie)
	// fmt.Println("cookie check ", check)
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
	rows, err := db.Query("SELECT id, user_name, title, content, created FROM post")
	checkErr(err)
	postinfo := make([]post, 0)
	for rows.Next() { //for loop through database table
		onePost := post{}
		err = rows.Scan(&onePost.ID, &onePost.Title, &onePost.Content, &onePost.Created, &onePost.Category, &onePost.Username)
		checkErr(err)
		time := ""
		time = onePost.Created[:10]
		time += " " + onePost.Created[11:19]
		onePost.Created = time
		postinfo = append(postinfo, onePost)
	}
	return postinfo
}
