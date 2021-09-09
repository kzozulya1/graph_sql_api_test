Модуль для выполнения CRUD операций с таблицей БД (Postgres) для GraphQL интерфейса:

1) Запустить выполнение скрипта для создания схемы таблицы (файл E:\Projects\graph_sql_api_test\scripts\sql\clients.sql)
2) Установить переменную окружения DB_CONN в креденшлы БД, напр. "postgres://postgres:postgres@127.0.0.1:5432/test_db?sslmode=disable"
3) Запустить модуль в папке cmd/graph_sql_api_test_serviced:
   **go run .**

4) В Postman создать POST запрос на **localhost:2121/client**


4.1) Листинг всех записей (c возможностью фильтрации по имени по частичному вхождению поискового запроса):

**QUERY:**

query GetClients ($client_name: String, $limit: Int!, $offset: Int!)
{
  clients_count #counts total records, for pager in frontend
   
  clients (client_name: $client_name, limit: $limit, offset: $offset){
     id
     client_name
     ur_adr
  }
}


**GRAPHQL VARIABLES:**

{
    "client_name": "", #here can be some filter example нест
    "limit": 15,
    "offset": 0
}


4.2) Получение записи по её ID:

**QUERY:**


query GetClientByID ($id: Int!){
  client (id: $id) {
      id
      client_name
      ur_adr
    }
  }

**GRAPHQL VARIABLES:**


{
    "id": 155
}


4.2.1) Псевдонимы записей - получение набора записей:

**QUERY:**


{
  client151: client(id: 151) {
    client_name
  }
  client154: client(id: 154) {
    client_name
  }
}

4.2.2) Фрагменты записей - получение набора записей:

**QUERY:**


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


4.2.3) Директивы:


**QUERY:**



query GetClient($id: Int!, $withUrAdr: Boolean!)
{
  client(id: $id){
     id
     client_name
     ur_adr @include (if: $withUrAdr)
  }
}


**GRAPHQL VARIABLES:**


{
    "id": 154,
    "withUrAdr": true
}




Refs.:
https://pkg.go.dev/github.com/graphql-go/graphql
