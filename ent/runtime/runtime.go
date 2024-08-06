// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/Milkado/api-challenge-jornada-milhas/ent/destinies"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/schema"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/testimonies"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	destiniesHooks := schema.Destinies{}.Hooks()
	destinies.Hooks[0] = destiniesHooks[0]
	destiniesFields := schema.Destinies{}.Fields()
	_ = destiniesFields
	// destiniesDescCreatedAt is the schema descriptor for created_at field.
	destiniesDescCreatedAt := destiniesFields[3].Descriptor()
	// destinies.DefaultCreatedAt holds the default value on creation for the created_at field.
	destinies.DefaultCreatedAt = destiniesDescCreatedAt.Default.(func() time.Time)
	// destiniesDescUpdatedAt is the schema descriptor for updated_at field.
	destiniesDescUpdatedAt := destiniesFields[4].Descriptor()
	// destinies.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	destinies.DefaultUpdatedAt = destiniesDescUpdatedAt.Default.(func() time.Time)
	testimoniesFields := schema.Testimonies{}.Fields()
	_ = testimoniesFields
	// testimoniesDescCreatedAt is the schema descriptor for created_at field.
	testimoniesDescCreatedAt := testimoniesFields[3].Descriptor()
	// testimonies.DefaultCreatedAt holds the default value on creation for the created_at field.
	testimonies.DefaultCreatedAt = testimoniesDescCreatedAt.Default.(func() time.Time)
	// testimoniesDescUpdatedAt is the schema descriptor for updated_at field.
	testimoniesDescUpdatedAt := testimoniesFields[4].Descriptor()
	// testimonies.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	testimonies.DefaultUpdatedAt = testimoniesDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPasswordToken is the schema descriptor for password_token field.
	userDescPasswordToken := userFields[4].Descriptor()
	// user.PasswordTokenValidator is a validator for the "password_token" field. It is called by the builders before save.
	user.PasswordTokenValidator = userDescPasswordToken.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
