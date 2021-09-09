package schema

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/graphql-go/graphql"
)

// getQueryType construct main Query type for schema
func getQueryType(db *pg.DB, modelType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"client": &graphql.Field{
				Type: modelType,
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
				Type: graphql.NewList(modelType),
				Args: graphql.FieldConfigArgument{
					//Data filter
					"client_name": &graphql.ArgumentConfig{
						Description:  "Client name filter",
						Type:         graphql.String,
						DefaultValue: "",
					},
					//Pagination filter
					"first": &graphql.ArgumentConfig{ //is a limit replacement
						Description:  "Pagination limit filter",
						Type:         graphql.Int,
						DefaultValue: 10,
					},
					"offset": &graphql.ArgumentConfig{
						Description:  "Pagination offset filter",
						Type:         graphql.Int,
						DefaultValue: 0,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					clients, err := getClients(db, p.Args["client_name"].(string), p.Args["first"].(int), p.Args["offset"].(int))
					if err != nil {
						return nil, fmt.Errorf("clients fetch err: %s", err.Error())
					}
					return clients, nil
				},
			},

			"totalCount": &graphql.Field{
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
}
