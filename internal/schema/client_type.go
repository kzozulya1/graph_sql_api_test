package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/kzozulya1/graph_sql_api_test/internal/storage"
)

// getClietType constructs `client` object type that represents db table model
func getClientType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
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
}
