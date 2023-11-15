// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dev-hato/misskey-abuse-user-report-notifier/ent/schema"
	"github.com/dev-hato/misskey-abuse-user-report-notifier/ent/userreport"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userreportFields := schema.UserReport{}.Fields()
	_ = userreportFields
	// userreportDescCreatedAt is the schema descriptor for created_at field.
	userreportDescCreatedAt := userreportFields[1].Descriptor()
	// userreport.DefaultCreatedAt holds the default value on creation for the created_at field.
	userreport.DefaultCreatedAt = userreportDescCreatedAt.Default.(func() time.Time)
}