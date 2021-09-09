### Модуль для выполнения CRUD операций с таблицей БД (Postgres) для GraphQL интерфейса

1) Запустить выполнение скрипта для создания схемы таблицы (файл `scripts/sql/clients.sql`)
2) Установить переменную окружения DB_CONN в креденшлы БД, напр. `postgres://postgres:postgres@127.0.0.1:5432/test_db?sslmode=disable`
3) Запустить модуль в папке cmd/graph_sql_api_test_serviced:
   `go run . `

4) В Postman создать POST запрос на `http://localhost:2121/client`


#### 4.1) Листинг всех записей (c возможностью фильтрации по имени по частичному вхождению поискового запроса):

##### QUERY:
```go
query GetClients ($client_name: String, $first: Int, $offset: Int)
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
Все поля - необязательные
```go
{
    "client_name": "ООО",
    "first": 5, 
    "offset": 0
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
id - обязательное поле
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

#### 4.3) Создание новой записи:

##### QUERY:
```go
mutation CreateClientObj($clientName: String!, $urAdr: String! ) {
  createClient(client_name: $clientName, ur_adr: $urAdr) {
    id,
    client_name,
    ur_adr
  }
}
```

##### GRAPHQL VARIABLES:
```go
{
    "clientName": "Ksenia R.",
    "urAdr": "Moscow, Kievskaya, 23"
}
```

#### 4.4) Обновление существующей записи:

##### QUERY:
```go
mutation UpdateClientObj($id: Int!, $clientName: String, $urAdr: String) {
  updateClient(id: $id, client_name: $clientName, ur_adr: $urAdr) {
    id,
    client_name,
    ur_adr
  }
}
```

##### GRAPHQL VARIABLES:
```go
{
    "id": 1,
    "clientName": "Petr Ivanov",
    "urAdr": ""
}
```


#### 4.4.1) Обновление 2х существующих записей:

##### QUERY:
```go
mutation Update2Clients($id: Int!, $clientName: String, $urAdr: String) {
  firstUpd: updateClient(id: $id, client_name: $clientName, ur_adr: $urAdr) {
    ... ClientFields
  }

  secUpd: updateClient(id: 2, client_name: "Second client", ur_adr: "Praha, Zlata 45") {
    ... ClientFields
  }
}

fragment ClientFields on Client {
    id,
    client_name,
    ur_adr
}

```

##### GRAPHQL VARIABLES:

```go
{
    "id": 1,
    "clientName": "Kosta Z.",
    "urAdr": "Piterburg"
}
```



#### 4.5) Удаление существующей записи:

##### QUERY:
```go
mutation DeleteClientObj($id: Int!) {
  deleteClient(id: $id) {
    id,
    client_name,
    ur_adr
  }
}
```

##### GRAPHQL VARIABLES:
```go
{
    "id": 4
}
```

#### 4.5.1) Удаление нескольких существующих записей:

##### QUERY:
```go
mutation DeleteFewClientsObj {
  rem1: deleteClient(id: 1) {
    id
  }
  rem2:  deleteClient(id: 2) {
    id
  }
  rem3: deleteClient(id: 3) {
    id
  }
}
```


> Лайвхак - чтобы в graphQL бекэнде сделать какой то аргумент необязательным (nullable), нужно указать `DefaultValue: ""`, напр.:
```go
"updateClient": &graphql.Field{
				Type: modelType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Client ID",
						Type:        graphql.NewNonNull(graphql.Int),
					},
					"client_name": &graphql.ArgumentConfig{
						Description:  "Client name",
						Type:         graphql.String,
						DefaultValue: "",
					},
					"ur_adr": &graphql.ArgumentConfig{
						Description:  "Client juridical address",
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				...
			}
```



###### Refs.:

> https://pkg.go.dev/github.com/graphql-go/graphql
