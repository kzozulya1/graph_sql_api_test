package schema

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/graphql-go/graphql"

	"github.com/kzozulya1/graph_sql_api_test/internal/storage"
)

// GetClientSchema creates GQL schema for clients entity
func GetClientSchema(db *pg.DB) *graphql.Schema {
	var (
		clientType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "Client",
			Description: "Organisation client entity",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "The ID of the client.",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if client, ok := p.Source.(*storage.Client); ok {
							return client.ID, nil
						}
						return nil, nil
					},
				},
				"client_name": &graphql.Field{
					Type:        graphql.String,
					Description: "The name of the client.",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if client, ok := p.Source.(*storage.Client); ok {
							return client.ClientName, nil
						}

						return nil, nil
					},
				},
				"ur_adr": &graphql.Field{
					Type:        graphql.String,
					Description: "The juridical address of client",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if client, ok := p.Source.(*storage.Client); ok {
							return client.UrAdr, nil
						}
						return nil, nil
					},
				},
			},
		})

		queryType = graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"client": &graphql.Field{
					Type: clientType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Description: "Client ID filter",
							Type:        graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						client, err := getClient(db, p.Args["id"].(int))
						if err != nil {
							return nil, fmt.Errorf("client for id %d fetch err: %s", p.Args["id"], err.Error())
						}
						return client, nil
					},
				},

				"clients": &graphql.Field{
					Type: graphql.NewList(clientType),
					Args: graphql.FieldConfigArgument{
						//Data filter
						"client_name": &graphql.ArgumentConfig{
							Description: "Client name filter",
							Type:        graphql.String,
						},
						//Pagination filter
						"limit": &graphql.ArgumentConfig{
							Description: "Pagination limit filter",
							Type:        graphql.Int,
						},
						"offset": &graphql.ArgumentConfig{
							Description: "Pagination offset filter",
							Type:        graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						clients, err := getClients(db, p.Args["client_name"].(string), p.Args["limit"].(int), p.Args["offset"].(int))
						if err != nil {
							return nil, fmt.Errorf("clients fetch err: %s", err.Error())
						}
						return clients, nil
					},
				},

				"clients_count": &graphql.Field{
					Type: graphql.Int,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						cnt, err := getClientsCount(db)
						if err != nil {
							return nil, fmt.Errorf("clients counting err: %s", err.Error())
						}
						return cnt, nil
					},
				},
			},
		})
	)

	clientSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	return &clientSchema
}

//getClient return client by filter
func getClient(db *pg.DB, id int) (interface{}, error) {
	var client = &storage.Client{}
	err := db.Model(client).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return client, nil
}

//getClients return clients slice by filter
func getClients(db *pg.DB, clientName string, limit, offset int) (interface{}, error) {
	var (
		clients []*storage.Client
		model   = db.Model(&clients)
	)

	if clientName != "" {
		model.Where("client_name ILIKE '%" + clientName + "%'")
	}
	if limit != 0 {
		model.Limit(limit)
	}
	if offset != 0 {
		model.Offset(offset)
	}

	if err := model.Select(); err != nil {
		return nil, err
	}
	return clients, nil
}

//getClientsCount counts record number in table
func getClientsCount(db *pg.DB) (interface{}, error) {
	return db.Model(&storage.Client{}).Count()
}
