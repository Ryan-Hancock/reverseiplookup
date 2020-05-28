# ReverseIPLookup

ReverseIPLookup is Go package suite including React frontend and library resolver.

## Installation

Server running the command below. This should use go mods to install the required packages.

```bash
go build . && ./reverseiplookup 
```

Frontend using normal create-react-app methods.
```bash
cd frontend && yarn install && yarn start
```

Docker can installed and ran by using `docker-compose`.
```bash
docker-compose up -d
```

the `docker-compose.yaml` file hosts the envs (ports, db). A side note is the use of network which will have to be changed depending on how you connect to your `db`. 

## Usage
Hitting the endpoints will validate and insert/update DNS hostnames.

```json
GET /ip/:address
response: {
  "message": [{domain: abc, validFrom: date}]
}  

GET /domain/:address
response: {
  "message": [{domain: abc, validFrom: date}]
}
```
### Resolver
To use the package you have to satisfy the resolver interface.
```golang
type ResolverStorage interface {
	GetByIP(IP string) ([]DomainRecord, error)
	GetOldest(limit int) ([]DomainRecord, error)
	Insert(DomainRecord) (ID int64, err error)
	InsertOrUpdate(DomainRecord) error
	Update(ID int64, record DomainRecord) error
}
```

### Cron
A running cron job every hour grabs the oldest domains and validates them. 

## Possible Improvements
1. Tests for the resolver package
2. Frontend UI and error messages