# easy-cass-go

The easy-cass-go project makes it easy for go code to connect to a Datastax Astra Cassandra database.

[gocql](https://gocql.github.io/) is the defacto golang library for connecting to Apache Cassandra databases.  It works great, but gets a little confusing when trying to connect a [Datastax Astra database](https://www.datastax.com/cloud/datastax-astra).  The easycass package simplifies the necessary steps to connect to an Astra database.

## Database Setup

This package requires a [Datastax Astra database](https://www.datastax.com/cloud/datastax-astra) which is built on top of Cassandra.  The free tier of the database will be more than adequate for many use cases.  To create an Astra database, go to [https://astra.datastax.com/](https://astra.datastax.com/) and follow the instructions to register and create a database.

When creating the database, you will specify a database username and password--you will need these in order to connect to the database (note that the database username and password are different than the username and password used to login to the Astra console).

In addition to the username and password, you will need a secure connect bundle file.  You can download this zip file from the Astra console.  Save it to a known location as you will need the full path to the file.

## Basic Usage

The easiest way to get up and running is to call the `easycass.GetSession()` function and pass it the username, password, and path to the secure connect bundle zip.  It will return a `*gocql.Session` that can be used normally.

```golang
package main

import (
	"log"

    "github.com/NathanBak/easy-cass-go/pkg/easycass"
)

func main() {

    // Specify the database username and password 
    // and the path to the secure connect bundle
	username  := "dbuser"
	password  := "dbpassword"
	pathToZip := "/home/me/Downloads/secure-connect-databasename.zip"

	// This creates and returns the gocql.Session
	session, err := easycass.GetSession(username, password, pathToZip)
	if err != nil {
		log.Fatal(err)
    }
    
    // Do something neat

    // All done
    session.Close()
```

## Examples

- The [simplesession](examples/simplesession/main.go) example connects to an Astra Database and then lists the tables in the default keyspace.
- The [configuredsession](examples/configuredsession/main.go) allows modification of the cluster configuration before creating the session to connect to the Astra Database and list the tables in the default keyspace.