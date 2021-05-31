# go_store
Go store with golang and react

## Environment
- go 1.16.4+
- yarn 1.22.10+
- node 14.17.0+
- postgreSQL 13.3+

## Set-up

- Create user for database: 
  * CREATE ROLE my_role WITH LOGIN PASSWORD '1234';

## Create the database:
- In order to create the database structure we have to create
the schema first, with the following command:

    * CREATE DATABASE go_store OWNER my_role

## .env
- This is an example how the .env file must look like
```
  PORT=8000
  DATABASE_POSTGRES_URI="postgres://my_role:1234@127.0.0.1:5432/go_store?sslmode=disable"
  DATABASE_DRIVER="postgres"
```


