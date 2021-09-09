package schema

import (
	"github.com/go-pg/pg/v10"
	"github.com/graphql-go/graphql"
)

// GetClientSchema creates GQL schema for clients entity
func GetClientSchema(db *pg.DB) *graphql.Schema {
	var (
		clientType   = getClientType()
		queryType    = getQueryType(db, clientType)
		mutationType = getMutationType(db, clientType)
	)

	clientSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	return &clientSchema
}
