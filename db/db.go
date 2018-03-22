package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strings"
)

type Connector struct {
	conn *sql.DB
	dsn  string
}

var defaultConnector *Connector = nil

func Instance() *Connector {
	if defaultConnector == nil {
		defaultConnector = &Connector{}
	}
	return defaultConnector
}

func (Self *Connector) Open(dsn string) {
	if Self.conn == nil {
		var err error = nil
		Self.conn, err = sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		Self.dsn = dsn
	}
}

func (Self *Connector) Close() {
	if Self.conn != nil {
		Self.conn.Close()
	}
}

func (Self *Connector) Query(model interface{}, query string, args ...interface{}) interface{} {
	value := reflect.ValueOf(model)
	if value.Kind() != reflect.Ptr {
		log.Println("need ptr model interface !")
		return nil
	} else {
		value = value.Elem()
	}

	rows, err := Self.conn.Query(query, args...)

	if err != nil {
		log.Println(err)
		return nil
	}

	defer rows.Close()

	s := value
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
		return model
	}
	return nil
}

func (Self *Connector) QueryAll(model interface{}, query string, args ...interface{}) []interface{} {
	value := reflect.ValueOf(model)
	if value.Kind() != reflect.Ptr {
		log.Println("need ptr model interface !")
		return nil
	} else {
		value = value.Elem()
	}

	rows, err := Self.conn.Query(query, args...)
	defer rows.Close()

	if err != nil {
		return nil
	}

	result := make([]interface{}, 0)
	s := reflect.ValueOf(model).Elem()
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

func (Self *Connector) GetTableName(model interface{}) (tableName string) {
	typ := reflect.TypeOf(model)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	tableName = typ.Name()
	for i := 0; i < typ.NumField(); i++ {
		var tag = typ.Field(i).Tag.Get("table")
		if len(tag) > 0 {
			tableName = tag
			break
		}
	}
	return
}

func (Self *Connector) Insert(model interface{}) (insetId int64, err error) {
	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	typ := value.Type()
	length := value.NumField()
	data := make([]interface{}, length)
	names := make([]string, length)
	flags := make([]string, length)

	for i := 0; i < length; i++ {
		data[i] = value.Field(i).Addr().Interface()
		names[i] = typ.Field(i).Name
		flags[i] = "?"
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", Self.GetTableName(model), strings.Join(names, ","), strings.Join(flags, ","))

	stmt, err := Self.conn.Prepare(query)
	if err != nil {
		return
	}

	res, err := stmt.Exec(data...)
	if err != nil {
		return
	}

	insetId, err = res.LastInsertId()
	return
}
