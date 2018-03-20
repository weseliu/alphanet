package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"log"
)

type Connector struct {
	conn *sql.DB
	dsn string
	models map[string] reflect.Type
}

var defaultConnector *Connector = nil
func Instance() *Connector {
	if defaultConnector == nil{
		defaultConnector = &Connector{
		}
	}
	return defaultConnector
}

func (Self *Connector)Open(dsn string)  {
	if Self.conn == nil{
		var err error = nil
		Self.conn, err = sql.Open("mysql", dsn)
		if err != nil{
			panic(err)
		}
		Self.dsn = dsn
	}
}

func (Self *Connector)Close()  {
	if Self.conn != nil{
		Self.conn.Close()
	}
}

func (Self *Connector)RegisterModel(model interface{})  {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr{
		t = t.Elem()
	}
	tableName := t.Name()
	Self.models[tableName] = t
}

func (Self *Connector)Query(table interface{}, query string, args ...interface{}) interface{} {
	rows, err := Self.conn.Query(query, args...)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return nil
	}

	s := reflect.ValueOf(table).Elem()
	length := s.NumField()
	scanner := make([]interface{}, length)
	for i := 0; i < length; i++ {
		scanner[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		err = rows.Scan(scanner...)
		if err != nil {
			panic(err)
		}
		return table
	}
	return nil
}

func (Self *Connector)QueryAll(table interface{}, query string, args ...interface{}) []interface{} {
	rows, err := Self.conn.Query(query, args...)
	defer rows.Close()

	if err != nil {
		return nil
	}

	result := make([]interface{}, 0)
	s := reflect.ValueOf(table).Elem()
	length := s.NumField()
	scanner := make([]interface{}, length)
	for i := 0; i < length; i++ {
		scanner[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		err = rows.Scan(scanner...)
		if err != nil {
			panic(err)
		}
		result = append(result, s.Interface())
	}
	return result
}


