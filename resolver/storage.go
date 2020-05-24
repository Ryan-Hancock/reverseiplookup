package resolver

import "time"

type DomainRecord struct {
	ID    int64
	Name  string
	IP    string
	Valid time.Time
}

type ResolverStorage interface {
	GetByDomain(domain string) (DomainRecord, error)
	GetByIP(IP string) ([]DomainRecord, error)
	GetByID(ID int64) (DomainRecord, error)
	BulkInsert([]DomainRecord) (bool, error)
	Insert(DomainRecord) (ID int64, err error)
}
