# Build REST-API on lambda
## Setup process

-- Build
sam build

-- Local Test
sam local start-api

-- Deploy with Config
sam deploy --config-file samconfig.toml


## API Doc
Query TodoItem GET localhost:3000/api/todo
Create TodoItem POST localhost:3000/api/todo
Update TodoItem POST localhost:3000/api/todo/:Id
Delete TodoItem DELETE localhost:3000/api/todo/:Id

Post man Collection:
