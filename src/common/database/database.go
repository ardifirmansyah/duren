package database

import (
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ardifirmansyah/duren/src/common/config"
)

var (
	DBConfig struct {
		Master  ConnConfig
		Slave   ConnConfig
		Replica ConnConfig
	}
)

type (
	DB struct {
		psqlConnection *sqlx.DB

		ConnConfig

		Mock MockDatabase
	}

	DBConnection struct {
		Master  *DB
		Slave   *DB
		Replica *DB
	}

	ConnConfig struct {
		Conn          string
		RetryInterval int
		MaxConn       int
		MaxIdle       int
		Type          string
	}

	DBService interface {
		Preparex(query string) *sqlx.Stmt
	}
)

func GetDatabase() *DBConnection {
	// read from config

	config.MustReadModuleConfig(&DBConfig, []string{
		"./files/etc/app-config/database",
		"/etc/app-config/database",
	}, "database")
	log.Println(DBConfig)

	master := &DB{
		ConnConfig: DBConfig.Master,
	}
	master.ConnectAndMonitor()

	slave := &DB{
		ConnConfig: DBConfig.Slave,
	}
	slave.ConnectAndMonitor()

	replica := &DB{
		ConnConfig: DBConfig.Replica,
	}
	replica.ConnectAndMonitor()

	return &DBConnection{
		Master:  master,
		Slave:   slave,
		Replica: replica,
	}
}

//Connect to database from config and Ping the connection
func (d *DB) Connect() error {
	db, err := sqlx.Connect(d.Type, d.Conn)
	if err != nil {
		log.Println("[Error]: DB open connection error", err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Println("[Error]: DB ping connection error", err.Error())
		return err
	}

	if d.MaxConn > 0 {
		db.SetMaxOpenConns(d.MaxConn)
	}

	if d.MaxIdle > 0 {
		db.SetMaxIdleConns(d.MaxIdle)
	}

	d.psqlConnection = db

	return nil
}

func (d *DB) ConnectAndMonitor() {
	err := d.Connect()
	if err != nil {
		log.Printf("Not connected to database %s, trying", d.Conn)
	} else {
		log.Printf("Success connecting to database %s", d.Conn)
	}

	ticker := time.NewTicker(time.Duration(d.RetryInterval) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if d.psqlConnection == nil {
					d.Connect()
				} else {
					err := d.psqlConnection.Ping()
					if err != nil {
						log.Printf("[Error]: DB reconnect error %s", err.Error())
					}
				}
			}
		}
	}()
}

func (db *DB) Preparex(query string) *sqlx.Stmt {
	stmt, err := db.psqlConnection.Preparex(query)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return stmt
}

//Connect to database from config and Ping the connection
func (d *DB) Status() error {
	if d.psqlConnection == nil {
		return errors.New("not connected")
	}

	if err := d.psqlConnection.Ping(); err != nil {
		return err
	}

	return errors.New("database connected")
}
