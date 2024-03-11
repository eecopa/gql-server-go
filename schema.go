package main

import (
	"github.com/graphql-go/graphql"
)

type ChoosePackagePage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var chooosePackagePagetype = graphql.NewObject(graphql.ObjectConfig{
	Name: "ChoosePackagePage",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"choosePackagePage": &graphql.Field{
			Type:        chooosePackagePagetype,
			Description: "Choose Package Page",
			Args: graphql.FieldConfigArgument{
				"cid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				cid, isOK := params.Args["cid"].(string)

				if isOK {
					return ChoosePackagePage{Name: cid}, nil
				}

				return ChoosePackagePage{}, nil

			},
		},
	},
})

var AcquisitionSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
