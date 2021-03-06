CSV TO DB
=====

# Quick Start

### Create Mysql container

` docker-compose.yml up -d `

### Build the application

` go build `

### Run the application when mysql's container is ready

`./csv-to-db  -db-username=root -db-password=password -db-address=localhost -db-name=dbtest -db-table=data --buffer-sql=1000 --file static/data.csv`

### View results

`docker-compose exec mysql mysql -u root -p -e 'SELECT COUNT(*) FROM data' dbtest `


### Launch tests

` go test -v ./... `

### Parameters :

```
  -buffer-sql int
    	Values to insert in DB (default 1000)
  -db-address string
    	Database url (default "127.0.0.1")
  -db-name string
    	Database name (default "test")
  -db-password string
    	Database password
  -db-table string
    	Table to create
  -db-username string
    	Database user (default "root")
  -file string
    	CSV File
```
