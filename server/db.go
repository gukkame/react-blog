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
	stmt.Exec("1", "Gunta", "CSS Grid and Flexbox", "CSS Grid and Flexbox are modern CSS layout models that allow for creating flexible and responsive designs for websites and web applications. \n \nCSS Grid is a two-dimensional layout model that enables you to create rows and columns for arranging elements on a web page. You can define the size and behavior of these rows and columns and easily place elements within them. CSS Grid is particularly useful for creating complex and custom layouts that would otherwise be difficult to achieve with traditional CSS layout methods.\n \n Flexbox, on the other hand, is a one-dimensional layout model that allows you to control the alignment, direction, and order of elements within a container. With Flexbox, you can easily create flexible and responsive layouts that adapt to different screen sizes and device types. \n \nBoth CSS Grid and Flexbox offer numerous benefits over traditional CSS layout methods, including easier implementation of responsive and adaptive designs, improved alignment and spacing control, and more efficient use of screen real estate. Understanding these layout models is crucial for modern front-end development, and can greatly simplify the process of creating high-quality, user-friendly web designs.", "/css-logo.jpg", "2022-05-01 11:30:22")
	stmt.Exec("2", "Gunta", "Building a responsive website with CSS media queries", "Widgets deliver cutting-edge, reintermediate engage envisioneer orchestrate weblogs scale e-business users back-end e-tailers, reinvent technologies; implement; leading-edge mesh tagclouds magnetic. Proactive, mashups streamline, supply-chains web services; engage: infomediaries functionalities addelivery schemas integrateAJAX-enabled weblogs users monetize; optimize reinvent synthesize. Distributed. Disintermediate one-to-one, partnerships ubiquitous; supply-chains best-of-breed cutting-edge plug-and-play integrated aggregate e-markets, integrated orchestrate! Empower dynamic holistic design exploit plug-and-play integrated schemas; out-of-the-box implement web-enabled, magnetic Cluetrain. ROI convergence interfaces, ecologies blogospheres leading-edge extensible, e-commerce, communities; expedite; semantic, utilize virtual B2C tagclouds iterate paradigms. Orchestrate redefine, cultivate, applications networking harness aggregate end-to-end engage plug-and-play rich enterprise long-tail compelling enable proactive. Networking wikis frictionless front-end e-business podcasts bleeding-edge next-generation Cluetrain--wireless platforms integrate wireless extend tag solutions B2B, harness 24/7 synthesize folksonomies. \n \n 24/7, convergence beta-test supply-chains empower implement deliverables, enterprise optimize drive cross-platform mesh interactive: platforms transparent, frictionless widgets ecologies strategize. Viral social, e-business envisioneer widgets, mesh post utilize widgets. \n \n	Enable sticky; rich-clientAPIs cutting-edge: enhance networks embrace incentivize enable infrastructures. Seamless mashups, niches, remix generate envisioneer seamless e-commerce out-of-the-box. Real-time viral mindshare, wikis killer aggregate enable innovative. \n	Mesh authentic methodologies integrate user-centred iterate e-enable syndicate. Revolutionize web-readiness 24/365 cross-media, sticky, user-contributed reinvent one-to-one channels deliverables. Virtual--markets enhance innovate exploit viral methodologies aggregate ubiquitous engineer aggregate architectures revolutionary user-centred interactive harness deploy.\n \nBricks-and-clicks transition enhance vertical plug-and-play enable remix create 24/365 portals architectures ROI whiteboard maximize synergize blogging strategize, efficient addelivery vortals. Aggregate web services unleash enhance sticky, tag: viral holistic strategize initiatives killer user-centric rss-capable sexy disintermediate cutting-edge whiteboard.", "/info-tech.jpg", "2022-09-01 10:30:22")
	stmt.Exec("3", "Emily", "Testing and debugging front-end code with Jest", "Responsive websites are amazing!", "/responsive-web.png", "2022-09-01 10:30:22")
	stmt.Exec("4", "Gunta", "The benefits of using CSS preprocessors such as SASS or LESS", "Responsive websites are amazing!", "/website-layout.jpg", "2022-09-01 10:30:22")
	stmt.Exec("5", "Gunta", "How to implement lazy loading for images and videos", "Responsive websites are amazing!", "/css-image.jpg", "2022-09-01 10:30:22")
	stmt.Exec("6", "Rob", "Optimizing website performance with lazy loading, compression and minification", "Responsive websites are amazing!", "/frontend.jpg", "2022-09-01 10:30:22")
	stmt.Exec("7", "Emily", "Building interactive user interfaces with JavaScript and jQuery", "Responsive websites are amazing!", "/css-image.jpg", "2022-09-01 10:30:22")
	stmt.Exec("8", "Gunta", "A comprehensive guide to React and its features", "Responsive websites are amazing!", "/react-logo.png", "2022-09-01 10:30:22")
	stmt.Exec("9", "Rob", "State management in React using Redux and its alternatives", "Responsive websites are amazing!", "/react-logo.png", "2022-09-01 10:30:22")
	stmt.Exec("10", "Rob", "Creating animations and transitions with CSS and JavaScript", "Responsive websites are amazing!", "/css-image.jpg", "2022-09-01 10:30:22")
}
