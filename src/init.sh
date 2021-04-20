cat src/sql/initDatabase.sql | sqlite3 src/sql/iotcameradata.db
go build src/main.go
./main