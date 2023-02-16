LOCAL_DB_NAME="wb"
LOCAL_DB_DSN="user=delinack dbname=wb sslmode=disable"
CMD     	= cmd/main.go

all:
		go run $(CMD)

clean:
		rm -rf main

db:
		psql -c "drop database if exists $(LOCAL_DB_NAME)"
		psql -c "create database $(LOCAL_DB_NAME)"
		goose -allow-missing -dir schema postgres $(LOCAL_DB_DSN) up
		make schema

schema:
		pg_dump -d $(LOCAL_DB_NAME) --schema-only --no-owner --no-privileges --no-tablespaces --no-security-labels -s > schema.sql

.PHONY:	all clean db schema