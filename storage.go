package main

import (
	"log"
	"reverseiplookup/resolver"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS records (
	id INT(11) NOT NULL AUTO_INCREMENT,
	domain VARCHAR(255) NOT NULL,
	ip VARCHAR(16) NOT NULL,
	valid DATETIME NOT NULL,
	PRIMARY KEY (id) USING BTREE
)
`

type Storage struct {
	Record resolver.DomainRecord
	DB     *sqlx.DB
}

// SetupDB brings back the storage object
func SetupDB() Storage {
	db, err := sqlx.Connect("mysql", "root:dev@(127.0.0.1:3306)/domains?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	return Storage{DB: db}
}

// GetByDomain ...
func (s *Storage) GetByDomain(domain string) (resolver.DomainRecord, error) {
	return s.Record, nil
}

// GetByIP ...
func (s *Storage) GetByIP(IP string) ([]resolver.DomainRecord, error) {
	return nil, nil
}

// GetByID ...
func (s *Storage) GetByID(ID int64) (resolver.DomainRecord, error) {
	return s.Record, nil
}

// BulkInsert ...
func (s *Storage) BulkInsert([]resolver.DomainRecord) (bool, error) {
	return true, nil
}

// Insert ...
func (s *Storage) Insert(dr resolver.DomainRecord) (ID int64, err error) {
	r, err := s.DB.NamedExec(`INSERT INTO records (domain,ip,valid) VALUES (:domain,:ip,:valid)`,
		map[string]interface{}{
			"domain": dr.Name,
			"ip":     dr.IP,
			"valid":  time.Now(),
		})
	if err != nil {
		return 0, err
	}
	id, _ := r.LastInsertId()

	return id, nil
}
