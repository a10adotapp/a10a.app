package schema

import (
	"fmt"

	"entgo.io/ent"
	"github.com/a10adotapp/a10a.app/lib/slices"
)

func WrapFields(fields ...ent.Field) []ent.Field {
	return slices.Map(fields, make([]ent.Field, 0), func(field ent.Field) ent.Field {
		field.Descriptor().Tag = fmt.Sprintf("json:\"%s\"", field.Descriptor().Name)

		return field
	})
}
