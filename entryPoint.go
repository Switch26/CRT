package main

//The package name can be an arbitrary identifier, though if you want a package to serve as an entry point for an executable program, it needs to be named “main” and have a function main() with no arguments and no return type.

import (
	"fmt"
	"github.com/gocql/gocql"
)

func main() {

	cluster := gocql.NewCluster("cassandra_some_network")
	cluster.Keyspace = "links_keyspace"
	cluster.Consistency = gocql.Quorum

	if _, err := cluster.CreateSession(); err != nil {
		panic(err)
	}

	fmt.Println("Connection successfull...")
	fmt.Println("hahaha")
}
