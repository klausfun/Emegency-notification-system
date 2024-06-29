package handler

import (
	"EmegencyNotificationSystem/profile_service/pkg/service"
	"EmegencyNotificationSystem/profile_service/schema_graphql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitGraphQL() http.Handler {
	schema := h.createSchema()

	return handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
}

func (h *Handler) createSchema() graphql.Schema {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"mwphing": &graphql.Field{
				Type:        schema_graphql.UserType,
				Description: "emw awmlpw nekwmkp",
				Args: graphql.FieldConfigArgument{
					"wmkemp": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(schema_graphql.SignUpInput),
					},
				},
				//Resolve: h.signUp,
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"signUp": &graphql.Field{
				Type:        schema_graphql.UserType,
				Description: "Sign up a new user",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(schema_graphql.SignUpInput),
					},
				},
				Resolve: h.signUp,
			},
			"signIn": &graphql.Field{
				Type:        schema_graphql.SignInResponse,
				Description: "Sign in an existing user",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(schema_graphql.SignInInput),
					},
				},
				Resolve: h.signIn,
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	return schema
}
