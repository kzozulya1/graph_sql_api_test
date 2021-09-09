package schema

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/graphql-go/graphql"
)

// getMutationType constructs mutation for create/update/delete routines
func getMutationType(db *pg.DB, modelType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createClient": &graphql.Field{
				Type: modelType,
				Args: graphql.FieldConfigArgument{
					"client_name": &graphql.ArgumentConfig{
						Description: "Client name",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"ur_adr": &graphql.ArgumentConfig{
						Description: "Client juridical address",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					client, err := createClient(db, p.Args["client_name"].(string), p.Args["ur_adr"].(string))
					if err != nil {
						return nil, fmt.Errorf("client create err: %s", err.Error())
					}
					return client, nil
				},
			},

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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					client, err := updateClient(db, p.Args["id"].(int), p.Args["client_name"].(string), p.Args["ur_adr"].(string))
					if err != nil {
						return nil, fmt.Errorf("client %d update err: %s", p.Args["id"].(int), err.Error())
					}
					return client, nil
				},
			},

			"deleteClient": &graphql.Field{
				Type: modelType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Client ID",
						Type:        graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					client, err := deleteClient(db, p.Args["id"].(int))
					if err != nil {
						return nil, fmt.Errorf("client %d delete err: %s", p.Args["id"].(int), err.Error())
					}
					return client, nil
				},
			},
		},
	})
}
