package easycass

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// GetSession creates and returns a gocql.Session that can be used to access
// the database.
func GetSession(username, password, pathToZip string) (*gocql.Session, error) {

	zi, err := readZip(pathToZip)
	if err != nil {
		return nil, err
	}

	cluster := gocql.NewCluster(zi.hostname)
	cluster.ConnectTimeout = time.Second * 5
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	cluster.Hosts = []string{fmt.Sprintf("%s:%d", zi.hostname, zi.port)}

	cluster.SslOpts = &gocql.SslOptions{
		Config:                 zi.tlsConfig,
		EnableHostVerification: false,
	}

	cluster.Keyspace = zi.keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	setKeyspace(session, zi.keyspace)

	return session, nil
}
