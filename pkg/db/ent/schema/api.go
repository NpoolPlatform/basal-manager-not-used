package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/basal-manager/pkg/db/mixin"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
)

// Api holds the schema definition for the Api entity.
type Api struct {
	ent.Schema
}

func (Api) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Api.
func (Api) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			String("protocol").
			Optional().
			Default(npool.Protocol_DefaultProtocol.String()),
		field.
			String("service_name").
			Optional().
			Default(""),
		field.
			String("method").
			Optional().
			Default(npool.Method_DefaultMethod.String()),
		field.
			String("method_name").
			Optional().
			Default(""),
		field.
			String("path").
			Optional().
			Default(""),
		field.
			Bool("exported").
			Optional().
			Default(false),
		field.
			String("path_prefix").
			Optional().
			Default(""),
		field.
			JSON("domains", []string{}).
			Optional().
			Default([]string{}),
		field.
			Bool("depracated").
			Optional().
			Default(false),
	}
}

// Edges of the Api.
func (Api) Edges() []ent.Edge {
	return nil
}
