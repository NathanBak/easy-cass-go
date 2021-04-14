package main

import (
	"fmt"
	"log"
	"time"

	"github.com/NathanBak/easy-cass-go/pkg/easycass"
	"github.com/gocql/gocql"
)

// This example creates a cluster config that can be modified (changing the timeout) and then
// connects to the database and prints out the tables in the default keyspace.  To use, the
// Client ID, Client Secret and path to the Secure Connect Bundle zip  must be specified.  The
// Token associated with the Client ID and Client Secret is not needed.

const (
	clientID     = "PUT_CLIENT_ID_HERE"
	clientSecret = "PUT_CLIENT_SECRET_HERE"
	pathToZip    = "/home/me/Downloads/secure-connect-databasename.zip"
)

func main() {
	var err error

	// Create the new cluster config
	var cluster *gocql.ClusterConfig
	cluster, err = easycass.NewCluster(clientID, clientSecret, pathToZip)
	if err != nil {
		log.Fatal(err)
	}

	// Perform additional cluster configuration as desired (in this case set the Timeout)
	cluster.Timeout = 10 * time.Second

	// Create new session from customized ClusterConfig
	var session *gocql.Session
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	// This is the default keyspace for the session
	keyspace := cluster.Keyspace

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
