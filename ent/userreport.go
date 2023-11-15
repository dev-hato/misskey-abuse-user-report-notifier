// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dev-hato/misskey-abuse-user-report-notifier/ent/userreport"
)

// UserReport is the model entity for the UserReport schema.
type UserReport struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserReport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userreport.FieldID:
			values[i] = new(sql.NullString)
		case userreport.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserReport fields.
func (ur *UserReport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userreport.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ur.ID = value.String
			}
		case userreport.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ur.CreatedAt = value.Time
			}
		default:
			ur.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserReport.
// This includes values selected through modifiers, order, etc.
func (ur *UserReport) Value(name string) (ent.Value, error) {
	return ur.selectValues.Get(name)
}

// Update returns a builder for updating this UserReport.
// Note that you need to call UserReport.Unwrap() before calling this method if this UserReport
// was returned from a transaction, and the transaction was committed or rolled back.
func (ur *UserReport) Update() *UserReportUpdateOne {
	return NewUserReportClient(ur.config).UpdateOne(ur)
}

// Unwrap unwraps the UserReport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ur *UserReport) Unwrap() *UserReport {
	_tx, ok := ur.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserReport is not a transactional entity")
	}
	ur.config.driver = _tx.drv
	return ur
}

// String implements the fmt.Stringer.
func (ur *UserReport) String() string {
	var builder strings.Builder
	builder.WriteString("UserReport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ur.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ur.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserReports is a parsable slice of UserReport.
type UserReports []*UserReport