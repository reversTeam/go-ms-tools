package scylla

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gocql/gocql"
)

const (
	hosts       = "127.0.0.1:9042"
	keyspace    = "glb"
	consistency = gocql.LocalOne
)

type ScyllaDBManager struct {
	session *gocql.Session
}

func NewScyllaDBManager() (*ScyllaDBManager, error) {
	cluster := gocql.NewCluster(hosts)
	cluster.Keyspace = keyspace
	cluster.Consistency = consistency

	log.Printf("NEW SCYLLA")

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

func (manager *ScyllaDBManager) GetOne(out interface{}, query string, values ...interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("out must be a pointer to a struct")
	}

	v = v.Elem()
	num := v.NumField()
	fields := make([]interface{}, 0, num)

	for i := 0; i < num; i++ {
		field := v.Field(i)
		if field.CanAddr() && field.CanInterface() {
			fields = append(fields, field.Addr().Interface())
		}
	}

	if len(fields) == 0 {
		return errors.New("no exported fields in the struct")
	}

	if err := manager.session.Query(query, values...).Consistency(gocql.One).Scan(fields...); err != nil {
		return err
	}

	return nil
}

func (manager *ScyllaDBManager) Get(out interface{}, query string, values ...interface{}) error {
	slicePtrVal := reflect.ValueOf(out)
	if slicePtrVal.Kind() != reflect.Ptr || slicePtrVal.Elem().Kind() != reflect.Slice {
		return errors.New("out doit être un pointeur vers une slice")
	}

	sliceVal := slicePtrVal.Elem()
	elemType := sliceVal.Type().Elem()

	iter := manager.session.Query(query, values...).Iter()
	columnInfos := iter.Columns()
	if columnInfos == nil {
		return errors.New("erreur lors de la récupération des infos de colonnes")
	}

	for {
		newRowPtr := reflect.New(elemType)
		scanResult := make([]interface{}, len(columnInfos))

		for i, columnInfo := range columnInfos {
			fieldIndex, found := findFieldIndex(elemType, columnInfo.Name)
			if !found {
				return fmt.Errorf("colonne %s non trouvée dans le type %s", columnInfo.Name, elemType.Name())
			}
			field := newRowPtr.Elem().Field(fieldIndex)
			scanResult[i] = field.Addr().Interface()
		}

		if !iter.Scan(scanResult...) {
			break
		}

		sliceVal.Set(reflect.Append(sliceVal, newRowPtr.Elem()))
	}

	if err := iter.Close(); err != nil {
		return err
	}

	return nil
}

func findFieldIndex(t reflect.Type, name string) (int, bool) {
	nameLower := strings.ToLower(name)
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		cqlTag := t.Field(i).Tag.Get("cql")
		if strings.ToLower(fieldName) == nameLower || strings.ToLower(cqlTag) == nameLower {
			return i, true
		}
	}
	return 0, false
}
