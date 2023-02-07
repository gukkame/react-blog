package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func database() {
	os.Remove("./database/database.db")

	log.Println("Creating sqlite-database.db...")

	file, err := os.Create("./database/database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("database.db created")

	dbConn, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	createPost(dbConn)

	defer dbConn.Close() // Defer Closing the database
}

// todo add image path to database
func createPost(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
	CREATE TABLE  post  (
		id  INTEGER not null PRIMARY KEY AUTOINCREMENT,
		username varchar(255) not null,
		title varchar(255) not null,
		content  text not null,
		image varchar(255) not null,
		created  datetime not null DEFAULT CURRENT_TIMESTAMP
	 )
		`)
	statement.Exec()
	stmt, err := dbConn.Prepare(`INSERT INTO post(id, username, title, content, image, created) VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec("1", "Gunta", "CSS Grid and Flexbox", "CSS Grid and Flexbox are modern CSS layout models that allow for creating flexible and responsive designs for websites and web applications. \n CSS Grid is a two-dimensional layout model that enables you to create rows and columns for arranging elements on a web page. You can define the size and behavior of these rows and columns and easily place elements within them. CSS Grid is particularly useful for creating complex and custom layouts that would otherwise be difficult to achieve with traditional CSS layout methods.\n Flexbox, on the other hand, is a one-dimensional layout model that allows you to control the alignment, direction, and order of elements within a container. With Flexbox, you can easily create flexible and responsive layouts that adapt to different screen sizes and device types. \n Both CSS Grid and Flexbox offer numerous benefits over traditional CSS layout methods, including easier implementation of responsive and adaptive designs, improved alignment and spacing control, and more efficient use of screen real estate. Understanding these layout models is crucial for modern front-end development, and can greatly simplify the process of creating high-quality, user-friendly web designs.", "./resources/css-logo.jpg", "2022-05-01 11:30:22")
	stmt.Exec("2", "Gunta", "Building a responsive website with CSS media queries", "Responsive websites are amazing!", "/image/path", "2022-09-01 10:30:22")
}
