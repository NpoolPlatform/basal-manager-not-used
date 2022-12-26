// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/basal-manager/pkg/db/ent/api"
	"github.com/NpoolPlatform/basal-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeApi = "Api"
)

// APIMutation represents an operation that mutates the Api nodes in the graph.
type APIMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *uint32
	addcreated_at *int32
	updated_at    *uint32
	addupdated_at *int32
	deleted_at    *uint32
	adddeleted_at *int32
	protocol      *string
	service_name  *string
	method        *string
	method_name   *string
	_path         *string
	exported      *bool
	path_prefix   *string
	domains       *[]string
	depracated    *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Api, error)
	predicates    []predicate.Api
}

var _ ent.Mutation = (*APIMutation)(nil)

// apiOption allows management of the mutation configuration using functional options.
type apiOption func(*APIMutation)

// newAPIMutation creates new mutation for the Api entity.
func newAPIMutation(c config, op Op, opts ...apiOption) *APIMutation {
	m := &APIMutation{
		config:        c,
		op:            op,
		typ:           TypeApi,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withApiID sets the ID field of the mutation.
func withApiID(id uuid.UUID) apiOption {
	return func(m *APIMutation) {
		var (
			err   error
			once  sync.Once
			value *Api
		)
		m.oldValue = func(ctx context.Context) (*Api, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Api.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withApi sets the old Api of the mutation.
func withApi(node *Api) apiOption {
	return func(m *APIMutation) {
		m.oldValue = func(context.Context) (*Api, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m APIMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m APIMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Api entities.
func (m *APIMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *APIMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *APIMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Api.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *APIMutation) SetCreatedAt(u uint32) {
	m.created_at = &u
	m.addcreated_at = nil
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *APIMutation) CreatedAt() (r uint32, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldCreatedAt(ctx context.Context) (v uint32, err error) {
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

// AddCreatedAt adds u to the "created_at" field.
func (m *APIMutation) AddCreatedAt(u int32) {
	if m.addcreated_at != nil {
		*m.addcreated_at += u
	} else {
		m.addcreated_at = &u
	}
}

// AddedCreatedAt returns the value that was added to the "created_at" field in this mutation.
func (m *APIMutation) AddedCreatedAt() (r int32, exists bool) {
	v := m.addcreated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *APIMutation) ResetCreatedAt() {
	m.created_at = nil
	m.addcreated_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *APIMutation) SetUpdatedAt(u uint32) {
	m.updated_at = &u
	m.addupdated_at = nil
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *APIMutation) UpdatedAt() (r uint32, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldUpdatedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// AddUpdatedAt adds u to the "updated_at" field.
func (m *APIMutation) AddUpdatedAt(u int32) {
	if m.addupdated_at != nil {
		*m.addupdated_at += u
	} else {
		m.addupdated_at = &u
	}
}

// AddedUpdatedAt returns the value that was added to the "updated_at" field in this mutation.
func (m *APIMutation) AddedUpdatedAt() (r int32, exists bool) {
	v := m.addupdated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *APIMutation) ResetUpdatedAt() {
	m.updated_at = nil
	m.addupdated_at = nil
}

// SetDeletedAt sets the "deleted_at" field.
func (m *APIMutation) SetDeletedAt(u uint32) {
	m.deleted_at = &u
	m.adddeleted_at = nil
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *APIMutation) DeletedAt() (r uint32, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldDeletedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// AddDeletedAt adds u to the "deleted_at" field.
func (m *APIMutation) AddDeletedAt(u int32) {
	if m.adddeleted_at != nil {
		*m.adddeleted_at += u
	} else {
		m.adddeleted_at = &u
	}
}

// AddedDeletedAt returns the value that was added to the "deleted_at" field in this mutation.
func (m *APIMutation) AddedDeletedAt() (r int32, exists bool) {
	v := m.adddeleted_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *APIMutation) ResetDeletedAt() {
	m.deleted_at = nil
	m.adddeleted_at = nil
}

// SetProtocol sets the "protocol" field.
func (m *APIMutation) SetProtocol(s string) {
	m.protocol = &s
}

// Protocol returns the value of the "protocol" field in the mutation.
func (m *APIMutation) Protocol() (r string, exists bool) {
	v := m.protocol
	if v == nil {
		return
	}
	return *v, true
}

// OldProtocol returns the old "protocol" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldProtocol(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProtocol is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProtocol requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProtocol: %w", err)
	}
	return oldValue.Protocol, nil
}

// ClearProtocol clears the value of the "protocol" field.
func (m *APIMutation) ClearProtocol() {
	m.protocol = nil
	m.clearedFields[api.FieldProtocol] = struct{}{}
}

// ProtocolCleared returns if the "protocol" field was cleared in this mutation.
func (m *APIMutation) ProtocolCleared() bool {
	_, ok := m.clearedFields[api.FieldProtocol]
	return ok
}

// ResetProtocol resets all changes to the "protocol" field.
func (m *APIMutation) ResetProtocol() {
	m.protocol = nil
	delete(m.clearedFields, api.FieldProtocol)
}

// SetServiceName sets the "service_name" field.
func (m *APIMutation) SetServiceName(s string) {
	m.service_name = &s
}

// ServiceName returns the value of the "service_name" field in the mutation.
func (m *APIMutation) ServiceName() (r string, exists bool) {
	v := m.service_name
	if v == nil {
		return
	}
	return *v, true
}

// OldServiceName returns the old "service_name" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldServiceName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldServiceName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldServiceName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldServiceName: %w", err)
	}
	return oldValue.ServiceName, nil
}

// ClearServiceName clears the value of the "service_name" field.
func (m *APIMutation) ClearServiceName() {
	m.service_name = nil
	m.clearedFields[api.FieldServiceName] = struct{}{}
}

// ServiceNameCleared returns if the "service_name" field was cleared in this mutation.
func (m *APIMutation) ServiceNameCleared() bool {
	_, ok := m.clearedFields[api.FieldServiceName]
	return ok
}

// ResetServiceName resets all changes to the "service_name" field.
func (m *APIMutation) ResetServiceName() {
	m.service_name = nil
	delete(m.clearedFields, api.FieldServiceName)
}

// SetMethod sets the "method" field.
func (m *APIMutation) SetMethod(s string) {
	m.method = &s
}

// Method returns the value of the "method" field in the mutation.
func (m *APIMutation) Method() (r string, exists bool) {
	v := m.method
	if v == nil {
		return
	}
	return *v, true
}

// OldMethod returns the old "method" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldMethod(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMethod is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMethod requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMethod: %w", err)
	}
	return oldValue.Method, nil
}

// ClearMethod clears the value of the "method" field.
func (m *APIMutation) ClearMethod() {
	m.method = nil
	m.clearedFields[api.FieldMethod] = struct{}{}
}

// MethodCleared returns if the "method" field was cleared in this mutation.
func (m *APIMutation) MethodCleared() bool {
	_, ok := m.clearedFields[api.FieldMethod]
	return ok
}

// ResetMethod resets all changes to the "method" field.
func (m *APIMutation) ResetMethod() {
	m.method = nil
	delete(m.clearedFields, api.FieldMethod)
}

// SetMethodName sets the "method_name" field.
func (m *APIMutation) SetMethodName(s string) {
	m.method_name = &s
}

// MethodName returns the value of the "method_name" field in the mutation.
func (m *APIMutation) MethodName() (r string, exists bool) {
	v := m.method_name
	if v == nil {
		return
	}
	return *v, true
}

// OldMethodName returns the old "method_name" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldMethodName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMethodName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMethodName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMethodName: %w", err)
	}
	return oldValue.MethodName, nil
}

// ClearMethodName clears the value of the "method_name" field.
func (m *APIMutation) ClearMethodName() {
	m.method_name = nil
	m.clearedFields[api.FieldMethodName] = struct{}{}
}

// MethodNameCleared returns if the "method_name" field was cleared in this mutation.
func (m *APIMutation) MethodNameCleared() bool {
	_, ok := m.clearedFields[api.FieldMethodName]
	return ok
}

// ResetMethodName resets all changes to the "method_name" field.
func (m *APIMutation) ResetMethodName() {
	m.method_name = nil
	delete(m.clearedFields, api.FieldMethodName)
}

// SetPath sets the "path" field.
func (m *APIMutation) SetPath(s string) {
	m._path = &s
}

// Path returns the value of the "path" field in the mutation.
func (m *APIMutation) Path() (r string, exists bool) {
	v := m._path
	if v == nil {
		return
	}
	return *v, true
}

// OldPath returns the old "path" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldPath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPath: %w", err)
	}
	return oldValue.Path, nil
}

// ClearPath clears the value of the "path" field.
func (m *APIMutation) ClearPath() {
	m._path = nil
	m.clearedFields[api.FieldPath] = struct{}{}
}

// PathCleared returns if the "path" field was cleared in this mutation.
func (m *APIMutation) PathCleared() bool {
	_, ok := m.clearedFields[api.FieldPath]
	return ok
}

// ResetPath resets all changes to the "path" field.
func (m *APIMutation) ResetPath() {
	m._path = nil
	delete(m.clearedFields, api.FieldPath)
}

// SetExported sets the "exported" field.
func (m *APIMutation) SetExported(b bool) {
	m.exported = &b
}

// Exported returns the value of the "exported" field in the mutation.
func (m *APIMutation) Exported() (r bool, exists bool) {
	v := m.exported
	if v == nil {
		return
	}
	return *v, true
}

// OldExported returns the old "exported" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldExported(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExported is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExported requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExported: %w", err)
	}
	return oldValue.Exported, nil
}

// ClearExported clears the value of the "exported" field.
func (m *APIMutation) ClearExported() {
	m.exported = nil
	m.clearedFields[api.FieldExported] = struct{}{}
}

// ExportedCleared returns if the "exported" field was cleared in this mutation.
func (m *APIMutation) ExportedCleared() bool {
	_, ok := m.clearedFields[api.FieldExported]
	return ok
}

// ResetExported resets all changes to the "exported" field.
func (m *APIMutation) ResetExported() {
	m.exported = nil
	delete(m.clearedFields, api.FieldExported)
}

// SetPathPrefix sets the "path_prefix" field.
func (m *APIMutation) SetPathPrefix(s string) {
	m.path_prefix = &s
}

// PathPrefix returns the value of the "path_prefix" field in the mutation.
func (m *APIMutation) PathPrefix() (r string, exists bool) {
	v := m.path_prefix
	if v == nil {
		return
	}
	return *v, true
}

// OldPathPrefix returns the old "path_prefix" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldPathPrefix(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPathPrefix is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPathPrefix requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPathPrefix: %w", err)
	}
	return oldValue.PathPrefix, nil
}

// ClearPathPrefix clears the value of the "path_prefix" field.
func (m *APIMutation) ClearPathPrefix() {
	m.path_prefix = nil
	m.clearedFields[api.FieldPathPrefix] = struct{}{}
}

// PathPrefixCleared returns if the "path_prefix" field was cleared in this mutation.
func (m *APIMutation) PathPrefixCleared() bool {
	_, ok := m.clearedFields[api.FieldPathPrefix]
	return ok
}

// ResetPathPrefix resets all changes to the "path_prefix" field.
func (m *APIMutation) ResetPathPrefix() {
	m.path_prefix = nil
	delete(m.clearedFields, api.FieldPathPrefix)
}

// SetDomains sets the "domains" field.
func (m *APIMutation) SetDomains(s []string) {
	m.domains = &s
}

// Domains returns the value of the "domains" field in the mutation.
func (m *APIMutation) Domains() (r []string, exists bool) {
	v := m.domains
	if v == nil {
		return
	}
	return *v, true
}

// OldDomains returns the old "domains" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldDomains(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDomains is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDomains requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDomains: %w", err)
	}
	return oldValue.Domains, nil
}

// ClearDomains clears the value of the "domains" field.
func (m *APIMutation) ClearDomains() {
	m.domains = nil
	m.clearedFields[api.FieldDomains] = struct{}{}
}

// DomainsCleared returns if the "domains" field was cleared in this mutation.
func (m *APIMutation) DomainsCleared() bool {
	_, ok := m.clearedFields[api.FieldDomains]
	return ok
}

// ResetDomains resets all changes to the "domains" field.
func (m *APIMutation) ResetDomains() {
	m.domains = nil
	delete(m.clearedFields, api.FieldDomains)
}

// SetDepracated sets the "depracated" field.
func (m *APIMutation) SetDepracated(b bool) {
	m.depracated = &b
}

// Depracated returns the value of the "depracated" field in the mutation.
func (m *APIMutation) Depracated() (r bool, exists bool) {
	v := m.depracated
	if v == nil {
		return
	}
	return *v, true
}

// OldDepracated returns the old "depracated" field's value of the Api entity.
// If the Api object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *APIMutation) OldDepracated(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDepracated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDepracated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDepracated: %w", err)
	}
	return oldValue.Depracated, nil
}

// ClearDepracated clears the value of the "depracated" field.
func (m *APIMutation) ClearDepracated() {
	m.depracated = nil
	m.clearedFields[api.FieldDepracated] = struct{}{}
}

// DepracatedCleared returns if the "depracated" field was cleared in this mutation.
func (m *APIMutation) DepracatedCleared() bool {
	_, ok := m.clearedFields[api.FieldDepracated]
	return ok
}

// ResetDepracated resets all changes to the "depracated" field.
func (m *APIMutation) ResetDepracated() {
	m.depracated = nil
	delete(m.clearedFields, api.FieldDepracated)
}

// Where appends a list predicates to the APIMutation builder.
func (m *APIMutation) Where(ps ...predicate.Api) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *APIMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Api).
func (m *APIMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *APIMutation) Fields() []string {
	fields := make([]string, 0, 12)
	if m.created_at != nil {
		fields = append(fields, api.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, api.FieldUpdatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, api.FieldDeletedAt)
	}
	if m.protocol != nil {
		fields = append(fields, api.FieldProtocol)
	}
	if m.service_name != nil {
		fields = append(fields, api.FieldServiceName)
	}
	if m.method != nil {
		fields = append(fields, api.FieldMethod)
	}
	if m.method_name != nil {
		fields = append(fields, api.FieldMethodName)
	}
	if m._path != nil {
		fields = append(fields, api.FieldPath)
	}
	if m.exported != nil {
		fields = append(fields, api.FieldExported)
	}
	if m.path_prefix != nil {
		fields = append(fields, api.FieldPathPrefix)
	}
	if m.domains != nil {
		fields = append(fields, api.FieldDomains)
	}
	if m.depracated != nil {
		fields = append(fields, api.FieldDepracated)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *APIMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case api.FieldCreatedAt:
		return m.CreatedAt()
	case api.FieldUpdatedAt:
		return m.UpdatedAt()
	case api.FieldDeletedAt:
		return m.DeletedAt()
	case api.FieldProtocol:
		return m.Protocol()
	case api.FieldServiceName:
		return m.ServiceName()
	case api.FieldMethod:
		return m.Method()
	case api.FieldMethodName:
		return m.MethodName()
	case api.FieldPath:
		return m.Path()
	case api.FieldExported:
		return m.Exported()
	case api.FieldPathPrefix:
		return m.PathPrefix()
	case api.FieldDomains:
		return m.Domains()
	case api.FieldDepracated:
		return m.Depracated()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *APIMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case api.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case api.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case api.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	case api.FieldProtocol:
		return m.OldProtocol(ctx)
	case api.FieldServiceName:
		return m.OldServiceName(ctx)
	case api.FieldMethod:
		return m.OldMethod(ctx)
	case api.FieldMethodName:
		return m.OldMethodName(ctx)
	case api.FieldPath:
		return m.OldPath(ctx)
	case api.FieldExported:
		return m.OldExported(ctx)
	case api.FieldPathPrefix:
		return m.OldPathPrefix(ctx)
	case api.FieldDomains:
		return m.OldDomains(ctx)
	case api.FieldDepracated:
		return m.OldDepracated(ctx)
	}
	return nil, fmt.Errorf("unknown Api field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *APIMutation) SetField(name string, value ent.Value) error {
	switch name {
	case api.FieldCreatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case api.FieldUpdatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case api.FieldDeletedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	case api.FieldProtocol:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProtocol(v)
		return nil
	case api.FieldServiceName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetServiceName(v)
		return nil
	case api.FieldMethod:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMethod(v)
		return nil
	case api.FieldMethodName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMethodName(v)
		return nil
	case api.FieldPath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPath(v)
		return nil
	case api.FieldExported:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExported(v)
		return nil
	case api.FieldPathPrefix:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPathPrefix(v)
		return nil
	case api.FieldDomains:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDomains(v)
		return nil
	case api.FieldDepracated:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDepracated(v)
		return nil
	}
	return fmt.Errorf("unknown Api field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *APIMutation) AddedFields() []string {
	var fields []string
	if m.addcreated_at != nil {
		fields = append(fields, api.FieldCreatedAt)
	}
	if m.addupdated_at != nil {
		fields = append(fields, api.FieldUpdatedAt)
	}
	if m.adddeleted_at != nil {
		fields = append(fields, api.FieldDeletedAt)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *APIMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case api.FieldCreatedAt:
		return m.AddedCreatedAt()
	case api.FieldUpdatedAt:
		return m.AddedUpdatedAt()
	case api.FieldDeletedAt:
		return m.AddedDeletedAt()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *APIMutation) AddField(name string, value ent.Value) error {
	switch name {
	case api.FieldCreatedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreatedAt(v)
		return nil
	case api.FieldUpdatedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpdatedAt(v)
		return nil
	case api.FieldDeletedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Api numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *APIMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(api.FieldProtocol) {
		fields = append(fields, api.FieldProtocol)
	}
	if m.FieldCleared(api.FieldServiceName) {
		fields = append(fields, api.FieldServiceName)
	}
	if m.FieldCleared(api.FieldMethod) {
		fields = append(fields, api.FieldMethod)
	}
	if m.FieldCleared(api.FieldMethodName) {
		fields = append(fields, api.FieldMethodName)
	}
	if m.FieldCleared(api.FieldPath) {
		fields = append(fields, api.FieldPath)
	}
	if m.FieldCleared(api.FieldExported) {
		fields = append(fields, api.FieldExported)
	}
	if m.FieldCleared(api.FieldPathPrefix) {
		fields = append(fields, api.FieldPathPrefix)
	}
	if m.FieldCleared(api.FieldDomains) {
		fields = append(fields, api.FieldDomains)
	}
	if m.FieldCleared(api.FieldDepracated) {
		fields = append(fields, api.FieldDepracated)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *APIMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *APIMutation) ClearField(name string) error {
	switch name {
	case api.FieldProtocol:
		m.ClearProtocol()
		return nil
	case api.FieldServiceName:
		m.ClearServiceName()
		return nil
	case api.FieldMethod:
		m.ClearMethod()
		return nil
	case api.FieldMethodName:
		m.ClearMethodName()
		return nil
	case api.FieldPath:
		m.ClearPath()
		return nil
	case api.FieldExported:
		m.ClearExported()
		return nil
	case api.FieldPathPrefix:
		m.ClearPathPrefix()
		return nil
	case api.FieldDomains:
		m.ClearDomains()
		return nil
	case api.FieldDepracated:
		m.ClearDepracated()
		return nil
	}
	return fmt.Errorf("unknown Api nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *APIMutation) ResetField(name string) error {
	switch name {
	case api.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case api.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case api.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	case api.FieldProtocol:
		m.ResetProtocol()
		return nil
	case api.FieldServiceName:
		m.ResetServiceName()
		return nil
	case api.FieldMethod:
		m.ResetMethod()
		return nil
	case api.FieldMethodName:
		m.ResetMethodName()
		return nil
	case api.FieldPath:
		m.ResetPath()
		return nil
	case api.FieldExported:
		m.ResetExported()
		return nil
	case api.FieldPathPrefix:
		m.ResetPathPrefix()
		return nil
	case api.FieldDomains:
		m.ResetDomains()
		return nil
	case api.FieldDepracated:
		m.ResetDepracated()
		return nil
	}
	return fmt.Errorf("unknown Api field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *APIMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *APIMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *APIMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *APIMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *APIMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *APIMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *APIMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Api unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *APIMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Api edge %s", name)
}
