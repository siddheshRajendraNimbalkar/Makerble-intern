.PHONY: 

postgres:
	docker run --name Makerble -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it Makerble createdb --username=root --owner=root  Makerble-backend

dropdb:
	docker exec -it Makerble dropdb Makerble-backend

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/Makerble-backend?sslmode=disable" -verbose up


migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/Makerble-backend?sslmode=disable" -verbose down

sqlc:
	sqlc generate
