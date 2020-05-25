package main

import (
	"fmt"
	"log"

	"github.com/NathanBak/easy-cass-go/pkg/easycass"
)

// This example connects to the database and prints out the tables in the
// keyspace.  To use, the username, password, and pathToZip must be specified.

const (
	username  = "dbuser"
	password  = "dbpassword"
	pathToZip = "/home/me/Downloads/secure-connect-databasename.zip"
)

func main() {

	// This creates and returns the gocql.Session
	session, err := easycass.GetSession(username, password, pathToZip)
	if err != nil {
		log.Fatal(err)
	}

	// This is the default keyspace for the session
	keyspace := easycass.GetKeyspace(session)

	// Print the tables for the keyspace
	tableNames, err := easycass.GetKeyspaceTableNames(session, keyspace)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nKeyspace %q contains the following tables:\n", keyspace)
	for _, tableName := range tableNames {
		fmt.Printf("\t%s\n", tableName)
	}

	session.Close()
}
