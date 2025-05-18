# Link Shortener

Make configurations in `.env` file

### Test Values
1. Launch app via
```shell
go run ./cmd/main.go
```
This will run automigration on DB and create necessary tables and entities

2. Run this to add `Test Values` to DB
```shell
go run ./cmd/seeder/seed.go
```