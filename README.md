create .env file with DB_NAME DB_USER DB_PASSWORD

than start docker_engine, and run 
docker-compose -d --build
that will create db, but will not start the project I did it because it easier to use CLI in IDE

than you should start the project with 
go run main.go 
or through 
go build main.go
./main.exe

Project uses GORM to develop faster, join code made by rough sql.
we have cashier and manager in the logic layer(usually api layer) in it we should validate data, and use service layer to operate with business-logic
--slog-- as logger in service layer
in service layer we have business logic, and if needed should start transactions and call storage(database layer)
in database layer we work directly with db, redis or other stuff we have there CRUD for any data

