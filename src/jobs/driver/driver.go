package driver

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

func createDBConnection() {
	// Connect to mongo
	session, err := mgo.Dial("mongodb://localhost:27017/IBDB")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("mongo err")
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	fmt.println("Database is up!!")
	// Get posts collection
	// posts = session.DB("app").C("posts")
}
