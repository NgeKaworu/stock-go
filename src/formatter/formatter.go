package formatter

import (
	"time"

	"github.com/NgeKaworu/maplization"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Formatter formatter funcs
var Formatter = map[string]maplization.Formatter{
	"now": func(i interface{}) (interface{}, error) {
		return time.Now(), nil
	},
	"nowLocal": func(i interface{}) (interface{}, error) {
		return time.Now().Local(), nil
	},
	"local": func(i interface{}) (interface{}, error) {
		t, err := time.Parse(time.RFC3339, i.(string))
		return t.Local(), err
	},
	"oid": func(i interface{}) (interface{}, error) {
		return primitive.ObjectIDFromHex(i.(string))
	},
}
