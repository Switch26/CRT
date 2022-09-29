package main

//The package name can be an arbitrary identifier, though if you want a package to serve as an entry point for an executable program, it needs to be named “main” and have a function main() with no arguments and no return type.

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

var Global_keyspace = "links_keyspace"

func main() {

	cluster := gocql.NewCluster("cassandra_some_network")
	cluster.Keyspace = Global_keyspace
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	fmt.Println("Connection successfull...")

	// Check for links
	//Qme2RGkCMaj1ajYBR9oPcyu8M4PSL7zLAwow6RbANMgCCk
	scanner := session.Query(`SELECT * FROM links_keyspace.links`).Iter().Scanner()
	for scanner.Next() {
		var (
			hash  string
			links []string
		)
		err = scanner.Scan(&hash, &links)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hash:", hash)
		fmt.Println("links:", links)
	}

	fmt.Println("Query successfull?")

}
