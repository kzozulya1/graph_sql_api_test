### Модуль для выполнения CRUD операций с таблицей БД (Postgres) для GraphQL интерфейса

1) Запустить выполнение скрипта для создания схемы таблицы (файл `scripts/sql/clients.sql`)
2) Установить переменную окружения DB_CONN в креденшлы БД, напр. `postgres://postgres:postgres@127.0.0.1:5432/test_db?sslmode=disable`
3) Запустить модуль в папке cmd/graph_sql_api_test_serviced:
   `go run . `

4) В Postman создать POST запрос на `http://localhost:2121/client`


#### 4.1) Листинг всех записей (c возможностью фильтрации по имени по частичному вхождению поискового запроса):

##### QUERY:
```go
query GetClients ($client_name: String, $first: Int!, $offset: Int!)
{
  totalCount
  clients (client_name: $client_name, first: $first, offset: $offset){
     id
     client_name
     ur_adr
  }
}
```


##### GRAPHQL VARIABLES:
```go
{
    "client_name": "",
    "first": 5,
    "offset": 1
}
```


#### 4.2) Получение записи по её ID:

##### QUERY:
```go
query GetClientByID ($id: Int!){
  client (id: $id) {
      id
      client_name
      ur_adr
    }
  }
```

##### GRAPHQL VARIABLES:
```go
{
    "id": 155
}
```

#### 4.2.1) Псевдонимы записей - получение набора записей:

##### QUERY:
```go
{
  client151: client(id: 151) {
    client_name
  }
  client154: client(id: 154) {
    client_name
  }
}
```

#### 4.2.2) Фрагменты записей - получение набора записей:

##### QUERY:
```go
{
  client151: client(id: 151) {
    ...fields
  }
  client154: client(id: 154) {
    ...fields
  }
}

fragment fields on Client {
  client_name
  ur_adr
  
}
```

#### 4.2.3) Директивы:

##### QUERY:
```go
query GetClient($id: Int!, $withUrAdr: Boolean!)
{
  client(id: $id){
     id
     client_name
     ur_adr @include (if: $withUrAdr)
  }
}
```

##### GRAPHQL VARIABLES:
```go
{
    "id": 154,
    "withUrAdr": true
}
```

#### 4.2.4) Инлайн-фрагменты (по условию типа объекта):

##### QUERY:
```go
{
  client(id: 153){
     __typename
     client_name
     ... on Client{
         ur_adr
     }
  }
}

```

#### 4.2.5) Мета данные (тип объекта):

##### QUERY:
```go
{
  client(id: 153){
     __typename
     client_name
  }
}

```


###### Refs.:

> https://pkg.go.dev/github.com/graphql-go/graphql
