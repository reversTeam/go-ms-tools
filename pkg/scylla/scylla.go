package scylla

import (
	"github.com/gocql/gocql"
)

const (
	consistency = gocql.LocalOne
)

type ScyllaDBManager struct {
	session *gocql.Session
}

type ScyllaConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Scylla map[string]ScyllaConfig `mapstructure:"scylla"`
}

func NewScyllaDBManager(keyspace string, cnf ScyllaConfig) (*ScyllaDBManager, error) {
	cluster := gocql.NewCluster(cnf.Host)
	cluster.Keyspace = keyspace
	cluster.Consistency = consistency

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &ScyllaDBManager{session: session}, nil
}

func (manager *ScyllaDBManager) Close() {
	manager.session.Close()
}

func (manager *ScyllaDBManager) ExecuteQuery(query string, values ...interface{}) error {
	return manager.session.Query(query, values...).Exec()
}

func (manager *ScyllaDBManager) GetOne(query string, values ...interface{}) *gocql.Query {
	return manager.session.Query(query, values...).Consistency(gocql.One)
}

func (manager *ScyllaDBManager) Get(query string, values ...interface{}) *gocql.Iter {
	return manager.session.Query(query, values...).Consistency(gocql.One).Iter()
}
