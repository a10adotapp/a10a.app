package schema

import (
	"fmt"

	"entgo.io/ent"
	"github.com/a10adotapp/a10a.app/lib/xslices"
)

func WrapFields(fields ...ent.Field) []ent.Field {
	return xslices.Map(fields, func(field ent.Field) ent.Field {
		field.Descriptor().Tag = fmt.Sprintf("json:\"%s\"", field.Descriptor().Name)

		return field
	})
}
