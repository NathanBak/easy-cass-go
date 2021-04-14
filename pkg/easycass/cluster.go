package easycass

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// NewCluster creates a returns a Cluster Config that and be modified and used to create a session.
func NewCluster(username, password, pathToZip string) (*gocql.ClusterConfig, error) {

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
	cluster.Port = zi.port
	cluster.Hosts = []string{fmt.Sprintf("%s:%d", zi.hostname, zi.port)}

	cluster.SslOpts = &gocql.SslOptions{Config: zi.tlsConfig}

	cluster.Keyspace = zi.keyspace

	return cluster, nil
}
