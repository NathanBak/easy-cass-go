package easycass

import (
	"github.com/gocql/gocql"
)

// GetSession creates and returns a gocql.Session that can be used to access
// the database.
func GetSession(username, password, pathToZip string) (*gocql.Session, error) {

	cluster, err := NewCluster(username, password, pathToZip)
	if err != nil {
		return nil, err
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	setKeyspace(session, cluster.Keyspace)

	return session, nil
}
