// Code generated by ent, DO NOT EDIT.
// @generated

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dev-hato/misskey-abuse-user-report-notifier/ent/predicate"
	"github.com/dev-hato/misskey-abuse-user-report-notifier/ent/userreport"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUserReport = "UserReport"
)

// UserReportMutation represents an operation that mutates the UserReport nodes in the graph.
type UserReportMutation struct {
	config
	op            Op
	typ           string
	id            *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*UserReport, error)
	predicates    []predicate.UserReport
}

var _ ent.Mutation = (*UserReportMutation)(nil)

// userreportOption allows management of the mutation configuration using functional options.
type userreportOption func(*UserReportMutation)

// newUserReportMutation creates new mutation for the UserReport entity.
func newUserReportMutation(c config, op Op, opts ...userreportOption) *UserReportMutation {
	m := &UserReportMutation{
		config:        c,
		op:            op,
		typ:           TypeUserReport,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserReportID sets the ID field of the mutation.
func withUserReportID(id string) userreportOption {
	return func(m *UserReportMutation) {
		var (
			err   error
			once  sync.Once
			value *UserReport
		)
		m.oldValue = func(ctx context.Context) (*UserReport, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserReport.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserReport sets the old UserReport of the mutation.
func withUserReport(node *UserReport) userreportOption {
	return func(m *UserReportMutation) {
		m.oldValue = func(context.Context) (*UserReport, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserReportMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserReportMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of UserReport entities.
func (m *UserReportMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserReportMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserReportMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().UserReport.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *UserReportMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserReportMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the UserReport entity.
// If the UserReport object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserReportMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserReportMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the UserReportMutation builder.
func (m *UserReportMutation) Where(ps ...predicate.UserReport) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserReportMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserReportMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.UserReport, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserReportMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserReportMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (UserReport).
func (m *UserReportMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserReportMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.created_at != nil {
		fields = append(fields, userreport.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserReportMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case userreport.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserReportMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case userreport.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown UserReport field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserReportMutation) SetField(name string, value ent.Value) error {
	switch name {
	case userreport.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown UserReport field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserReportMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserReportMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserReportMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown UserReport numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserReportMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserReportMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserReportMutation) ClearField(name string) error {
	return fmt.Errorf("unknown UserReport nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserReportMutation) ResetField(name string) error {
	switch name {
	case userreport.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown UserReport field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserReportMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserReportMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserReportMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserReportMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserReportMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserReportMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserReportMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown UserReport unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserReportMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown UserReport edge %s", name)
}
