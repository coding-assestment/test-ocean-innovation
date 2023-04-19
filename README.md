## [Guide](...)

### Installation

```sh
go mod tidy
```

### running service

```sh
go run main.go
```


### [Collection](../Soccer Service.postman_collection.json)  


### API 
#### Create A team 

```sh
curl --location 'localhost:1323/api/v1/teams' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Team 4"
}'
```
#### Get A team 

```sh
curl --location --request GET 'localhost:1323/api/v1/teams/3' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Joe Smith"
}'
```

#### Update A team 

```sh
curl --location --request PUT 'localhost:1323/api/v1/teams/3' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Team A0 testt"
}'
```

#### Delete A team 

```sh
curl --location --request DELETE 'localhost:1323/api/v1/teams/2' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Team A0 Rev"
}'
```

#### Get All Team

```sh
curl --location --request GET 'localhost:1323/api/v1/teams' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Joe Smith"
}'
```

#### Add Player of a Team 

```sh
curl --location 'localhost:1323/api/v1/teams/0/player' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Player 3"
}'
```

#### Get Player of a Team 

```sh
curl --location 'localhost:1323/api/v1/teams/0/player' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Player 3"
}'
```

