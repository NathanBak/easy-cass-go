package easycass

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gocql/gocql"
)

var sessionKeyspace map[*gocql.Session]string
var sessionKeyspaceLock sync.RWMutex

func init() {
	sessionKeyspace = make(map[*gocql.Session]string)
}

// GetKeyspace returns the default keyspace which was specified when the
// database was created.  This will return an empty string if the session has
// not yet been properly initialized.
func GetKeyspace(session *gocql.Session) string {
	if session == nil {
		return ""
	}
	sessionKeyspaceLock.RLock()
	defer sessionKeyspaceLock.RUnlock()
	return sessionKeyspace[session]
}

// setKeyspace associates the specified keyspace with the provided session
func setKeyspace(session *gocql.Session, keyspace string) {
	sessionKeyspaceLock.Lock()
	defer sessionKeyspaceLock.Unlock()
	sessionKeyspace[session] = keyspace
}

// GetKeyspaceTableNames returns a slice containing the names of all the tables
// found associated with the specified keyspace.  If there are problems finding
// the information, an error is returned.  If the keyspace does not exist or
// does not contain any tables, an empty slice will be returned.
func GetKeyspaceTableNames(session *gocql.Session, keyspace string) ([]string, error) {
	var tableNames []string

	if session == nil {
		return tableNames, errors.New("nil session")
	}

	query := fmt.Sprintf("SELECT table_name FROM system_schema.tables WHERE keyspace_name = '%s';", keyspace)

	q := session.Query(query)
	iter := q.Iter()

	var tableName string
	scanner := iter.Scanner()
	for scanner.Next() {
		err := scanner.Scan(&tableName)
		if err != nil {
			return tableNames, err
		}
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}
