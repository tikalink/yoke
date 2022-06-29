// Code generated by entc, DO NOT EDIT.

package media

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/announce"
)

// AnnounceCreate is the builder for creating a Announce entity.
type AnnounceCreate struct {
	config
	mutation *AnnounceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedUnix sets the "created_unix" field.
func (ac *AnnounceCreate) SetCreatedUnix(i int64) *AnnounceCreate {
	ac.mutation.SetCreatedUnix(i)
	return ac
}

// SetNillableCreatedUnix sets the "created_unix" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableCreatedUnix(i *int64) *AnnounceCreate {
	if i != nil {
		ac.SetCreatedUnix(*i)
	}
	return ac
}

// SetUpdatedUnix sets the "updated_unix" field.
func (ac *AnnounceCreate) SetUpdatedUnix(i int64) *AnnounceCreate {
	ac.mutation.SetUpdatedUnix(i)
	return ac
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableUpdatedUnix(i *int64) *AnnounceCreate {
	if i != nil {
		ac.SetUpdatedUnix(*i)
	}
	return ac
}

// SetDeletedUnix sets the "deleted_unix" field.
func (ac *AnnounceCreate) SetDeletedUnix(i int64) *AnnounceCreate {
	ac.mutation.SetDeletedUnix(i)
	return ac
}

// SetNillableDeletedUnix sets the "deleted_unix" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableDeletedUnix(i *int64) *AnnounceCreate {
	if i != nil {
		ac.SetDeletedUnix(*i)
	}
	return ac
}

// SetAnnounceNo sets the "announce_no" field.
func (ac *AnnounceCreate) SetAnnounceNo(s string) *AnnounceCreate {
	ac.mutation.SetAnnounceNo(s)
	return ac
}

// SetNillableAnnounceNo sets the "announce_no" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableAnnounceNo(s *string) *AnnounceCreate {
	if s != nil {
		ac.SetAnnounceNo(*s)
	}
	return ac
}

// SetTitle sets the "title" field.
func (ac *AnnounceCreate) SetTitle(s string) *AnnounceCreate {
	ac.mutation.SetTitle(s)
	return ac
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableTitle(s *string) *AnnounceCreate {
	if s != nil {
		ac.SetTitle(*s)
	}
	return ac
}

// SetKind sets the "kind" field.
func (ac *AnnounceCreate) SetKind(a announce.Kind) *AnnounceCreate {
	ac.mutation.SetKind(a)
	return ac
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableKind(a *announce.Kind) *AnnounceCreate {
	if a != nil {
		ac.SetKind(*a)
	}
	return ac
}

// SetContent sets the "content" field.
func (ac *AnnounceCreate) SetContent(s string) *AnnounceCreate {
	ac.mutation.SetContent(s)
	return ac
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableContent(s *string) *AnnounceCreate {
	if s != nil {
		ac.SetContent(*s)
	}
	return ac
}

// SetLink sets the "link" field.
func (ac *AnnounceCreate) SetLink(s string) *AnnounceCreate {
	ac.mutation.SetLink(s)
	return ac
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableLink(s *string) *AnnounceCreate {
	if s != nil {
		ac.SetLink(*s)
	}
	return ac
}

// SetSign sets the "sign" field.
func (ac *AnnounceCreate) SetSign(s string) *AnnounceCreate {
	ac.mutation.SetSign(s)
	return ac
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableSign(s *string) *AnnounceCreate {
	if s != nil {
		ac.SetSign(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AnnounceCreate) SetID(u uuid.UUID) *AnnounceCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AnnounceCreate) SetNillableID(u *uuid.UUID) *AnnounceCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// Mutation returns the AnnounceMutation object of the builder.
func (ac *AnnounceCreate) Mutation() *AnnounceMutation {
	return ac.mutation
}

// Save creates the Announce in the database.
func (ac *AnnounceCreate) Save(ctx context.Context) (*Announce, error) {
	var (
		err  error
		node *Announce
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnounceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("media: uninitialized hook (forgotten import media/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnnounceCreate) SaveX(ctx context.Context) *Announce {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnnounceCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnnounceCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AnnounceCreate) defaults() {
	if _, ok := ac.mutation.CreatedUnix(); !ok {
		v := announce.DefaultCreatedUnix
		ac.mutation.SetCreatedUnix(v)
	}
	if _, ok := ac.mutation.UpdatedUnix(); !ok {
		v := announce.DefaultUpdatedUnix
		ac.mutation.SetUpdatedUnix(v)
	}
	if _, ok := ac.mutation.AnnounceNo(); !ok {
		v := announce.DefaultAnnounceNo
		ac.mutation.SetAnnounceNo(v)
	}
	if _, ok := ac.mutation.Title(); !ok {
		v := announce.DefaultTitle
		ac.mutation.SetTitle(v)
	}
	if _, ok := ac.mutation.Kind(); !ok {
		v := announce.DefaultKind
		ac.mutation.SetKind(v)
	}
	if _, ok := ac.mutation.Content(); !ok {
		v := announce.DefaultContent
		ac.mutation.SetContent(v)
	}
	if _, ok := ac.mutation.Link(); !ok {
		v := announce.DefaultLink
		ac.mutation.SetLink(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := announce.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnnounceCreate) check() error {
	if _, ok := ac.mutation.CreatedUnix(); !ok {
		return &ValidationError{Name: "created_unix", err: errors.New(`media: missing required field "Announce.created_unix"`)}
	}
	if _, ok := ac.mutation.UpdatedUnix(); !ok {
		return &ValidationError{Name: "updated_unix", err: errors.New(`media: missing required field "Announce.updated_unix"`)}
	}
	if _, ok := ac.mutation.AnnounceNo(); !ok {
		return &ValidationError{Name: "announce_no", err: errors.New(`media: missing required field "Announce.announce_no"`)}
	}
	if _, ok := ac.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`media: missing required field "Announce.title"`)}
	}
	if _, ok := ac.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`media: missing required field "Announce.kind"`)}
	}
	if v, ok := ac.mutation.Kind(); ok {
		if err := announce.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`media: validator failed for field "Announce.kind": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`media: missing required field "Announce.content"`)}
	}
	if _, ok := ac.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`media: missing required field "Announce.link"`)}
	}
	return nil
}

func (ac *AnnounceCreate) sqlSave(ctx context.Context) (*Announce, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ac *AnnounceCreate) createSpec() (*Announce, *sqlgraph.CreateSpec) {
	var (
		_node = &Announce{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: announce.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: announce.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: announce.FieldCreatedUnix,
		})
		_node.CreatedUnix = value
	}
	if value, ok := ac.mutation.UpdatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: announce.FieldUpdatedUnix,
		})
		_node.UpdatedUnix = value
	}
	if value, ok := ac.mutation.DeletedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: announce.FieldDeletedUnix,
		})
		_node.DeletedUnix = value
	}
	if value, ok := ac.mutation.AnnounceNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announce.FieldAnnounceNo,
		})
		_node.AnnounceNo = value
	}
	if value, ok := ac.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announce.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ac.mutation.Kind(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: announce.FieldKind,
		})
		_node.Kind = value
	}
	if value, ok := ac.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announce.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := ac.mutation.Link(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announce.FieldLink,
		})
		_node.Link = value
	}
	if value, ok := ac.mutation.Sign(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announce.FieldSign,
		})
		_node.Sign = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Announce.Create().
//		SetCreatedUnix(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnnounceUpsert) {
//			SetCreatedUnix(v+v).
//		}).
//		Exec(ctx)
//
func (ac *AnnounceCreate) OnConflict(opts ...sql.ConflictOption) *AnnounceUpsertOne {
	ac.conflict = opts
	return &AnnounceUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Announce.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *AnnounceCreate) OnConflictColumns(columns ...string) *AnnounceUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AnnounceUpsertOne{
		create: ac,
	}
}

type (
	// AnnounceUpsertOne is the builder for "upsert"-ing
	//  one Announce node.
	AnnounceUpsertOne struct {
		create *AnnounceCreate
	}

	// AnnounceUpsert is the "OnConflict" setter.
	AnnounceUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedUnix sets the "created_unix" field.
func (u *AnnounceUpsert) SetCreatedUnix(v int64) *AnnounceUpsert {
	u.Set(announce.FieldCreatedUnix, v)
	return u
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateCreatedUnix() *AnnounceUpsert {
	u.SetExcluded(announce.FieldCreatedUnix)
	return u
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *AnnounceUpsert) AddCreatedUnix(v int64) *AnnounceUpsert {
	u.Add(announce.FieldCreatedUnix, v)
	return u
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *AnnounceUpsert) SetUpdatedUnix(v int64) *AnnounceUpsert {
	u.Set(announce.FieldUpdatedUnix, v)
	return u
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateUpdatedUnix() *AnnounceUpsert {
	u.SetExcluded(announce.FieldUpdatedUnix)
	return u
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *AnnounceUpsert) AddUpdatedUnix(v int64) *AnnounceUpsert {
	u.Add(announce.FieldUpdatedUnix, v)
	return u
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *AnnounceUpsert) SetDeletedUnix(v int64) *AnnounceUpsert {
	u.Set(announce.FieldDeletedUnix, v)
	return u
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateDeletedUnix() *AnnounceUpsert {
	u.SetExcluded(announce.FieldDeletedUnix)
	return u
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *AnnounceUpsert) AddDeletedUnix(v int64) *AnnounceUpsert {
	u.Add(announce.FieldDeletedUnix, v)
	return u
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *AnnounceUpsert) ClearDeletedUnix() *AnnounceUpsert {
	u.SetNull(announce.FieldDeletedUnix)
	return u
}

// SetAnnounceNo sets the "announce_no" field.
func (u *AnnounceUpsert) SetAnnounceNo(v string) *AnnounceUpsert {
	u.Set(announce.FieldAnnounceNo, v)
	return u
}

// UpdateAnnounceNo sets the "announce_no" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateAnnounceNo() *AnnounceUpsert {
	u.SetExcluded(announce.FieldAnnounceNo)
	return u
}

// SetTitle sets the "title" field.
func (u *AnnounceUpsert) SetTitle(v string) *AnnounceUpsert {
	u.Set(announce.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateTitle() *AnnounceUpsert {
	u.SetExcluded(announce.FieldTitle)
	return u
}

// SetKind sets the "kind" field.
func (u *AnnounceUpsert) SetKind(v announce.Kind) *AnnounceUpsert {
	u.Set(announce.FieldKind, v)
	return u
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateKind() *AnnounceUpsert {
	u.SetExcluded(announce.FieldKind)
	return u
}

// SetContent sets the "content" field.
func (u *AnnounceUpsert) SetContent(v string) *AnnounceUpsert {
	u.Set(announce.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateContent() *AnnounceUpsert {
	u.SetExcluded(announce.FieldContent)
	return u
}

// SetLink sets the "link" field.
func (u *AnnounceUpsert) SetLink(v string) *AnnounceUpsert {
	u.Set(announce.FieldLink, v)
	return u
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateLink() *AnnounceUpsert {
	u.SetExcluded(announce.FieldLink)
	return u
}

// SetSign sets the "sign" field.
func (u *AnnounceUpsert) SetSign(v string) *AnnounceUpsert {
	u.Set(announce.FieldSign, v)
	return u
}

// UpdateSign sets the "sign" field to the value that was provided on create.
func (u *AnnounceUpsert) UpdateSign() *AnnounceUpsert {
	u.SetExcluded(announce.FieldSign)
	return u
}

// ClearSign clears the value of the "sign" field.
func (u *AnnounceUpsert) ClearSign() *AnnounceUpsert {
	u.SetNull(announce.FieldSign)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Announce.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(announce.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AnnounceUpsertOne) UpdateNewValues() *AnnounceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(announce.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Announce.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AnnounceUpsertOne) Ignore() *AnnounceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnnounceUpsertOne) DoNothing() *AnnounceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnnounceCreate.OnConflict
// documentation for more info.
func (u *AnnounceUpsertOne) Update(set func(*AnnounceUpsert)) *AnnounceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnnounceUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedUnix sets the "created_unix" field.
func (u *AnnounceUpsertOne) SetCreatedUnix(v int64) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetCreatedUnix(v)
	})
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *AnnounceUpsertOne) AddCreatedUnix(v int64) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.AddCreatedUnix(v)
	})
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateCreatedUnix() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateCreatedUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *AnnounceUpsertOne) SetUpdatedUnix(v int64) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *AnnounceUpsertOne) AddUpdatedUnix(v int64) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateUpdatedUnix() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *AnnounceUpsertOne) SetDeletedUnix(v int64) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetDeletedUnix(v)
	})
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *AnnounceUpsertOne) AddDeletedUnix(v int64) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.AddDeletedUnix(v)
	})
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateDeletedUnix() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateDeletedUnix()
	})
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *AnnounceUpsertOne) ClearDeletedUnix() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.ClearDeletedUnix()
	})
}

// SetAnnounceNo sets the "announce_no" field.
func (u *AnnounceUpsertOne) SetAnnounceNo(v string) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetAnnounceNo(v)
	})
}

// UpdateAnnounceNo sets the "announce_no" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateAnnounceNo() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateAnnounceNo()
	})
}

// SetTitle sets the "title" field.
func (u *AnnounceUpsertOne) SetTitle(v string) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateTitle() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateTitle()
	})
}

// SetKind sets the "kind" field.
func (u *AnnounceUpsertOne) SetKind(v announce.Kind) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateKind() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateKind()
	})
}

// SetContent sets the "content" field.
func (u *AnnounceUpsertOne) SetContent(v string) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateContent() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateContent()
	})
}

// SetLink sets the "link" field.
func (u *AnnounceUpsertOne) SetLink(v string) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateLink() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateLink()
	})
}

// SetSign sets the "sign" field.
func (u *AnnounceUpsertOne) SetSign(v string) *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetSign(v)
	})
}

// UpdateSign sets the "sign" field to the value that was provided on create.
func (u *AnnounceUpsertOne) UpdateSign() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateSign()
	})
}

// ClearSign clears the value of the "sign" field.
func (u *AnnounceUpsertOne) ClearSign() *AnnounceUpsertOne {
	return u.Update(func(s *AnnounceUpsert) {
		s.ClearSign()
	})
}

// Exec executes the query.
func (u *AnnounceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("media: missing options for AnnounceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnnounceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AnnounceUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("media: AnnounceUpsertOne.ID is not supported by MySQL driver. Use AnnounceUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AnnounceUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AnnounceCreateBulk is the builder for creating many Announce entities in bulk.
type AnnounceCreateBulk struct {
	config
	builders []*AnnounceCreate
	conflict []sql.ConflictOption
}

// Save creates the Announce entities in the database.
func (acb *AnnounceCreateBulk) Save(ctx context.Context) ([]*Announce, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Announce, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnnounceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnnounceCreateBulk) SaveX(ctx context.Context) []*Announce {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnnounceCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnnounceCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Announce.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnnounceUpsert) {
//			SetCreatedUnix(v+v).
//		}).
//		Exec(ctx)
//
func (acb *AnnounceCreateBulk) OnConflict(opts ...sql.ConflictOption) *AnnounceUpsertBulk {
	acb.conflict = opts
	return &AnnounceUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Announce.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *AnnounceCreateBulk) OnConflictColumns(columns ...string) *AnnounceUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AnnounceUpsertBulk{
		create: acb,
	}
}

// AnnounceUpsertBulk is the builder for "upsert"-ing
// a bulk of Announce nodes.
type AnnounceUpsertBulk struct {
	create *AnnounceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Announce.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(announce.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AnnounceUpsertBulk) UpdateNewValues() *AnnounceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(announce.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Announce.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AnnounceUpsertBulk) Ignore() *AnnounceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnnounceUpsertBulk) DoNothing() *AnnounceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnnounceCreateBulk.OnConflict
// documentation for more info.
func (u *AnnounceUpsertBulk) Update(set func(*AnnounceUpsert)) *AnnounceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnnounceUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedUnix sets the "created_unix" field.
func (u *AnnounceUpsertBulk) SetCreatedUnix(v int64) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetCreatedUnix(v)
	})
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *AnnounceUpsertBulk) AddCreatedUnix(v int64) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.AddCreatedUnix(v)
	})
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateCreatedUnix() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateCreatedUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *AnnounceUpsertBulk) SetUpdatedUnix(v int64) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *AnnounceUpsertBulk) AddUpdatedUnix(v int64) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateUpdatedUnix() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *AnnounceUpsertBulk) SetDeletedUnix(v int64) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetDeletedUnix(v)
	})
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *AnnounceUpsertBulk) AddDeletedUnix(v int64) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.AddDeletedUnix(v)
	})
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateDeletedUnix() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateDeletedUnix()
	})
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *AnnounceUpsertBulk) ClearDeletedUnix() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.ClearDeletedUnix()
	})
}

// SetAnnounceNo sets the "announce_no" field.
func (u *AnnounceUpsertBulk) SetAnnounceNo(v string) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetAnnounceNo(v)
	})
}

// UpdateAnnounceNo sets the "announce_no" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateAnnounceNo() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateAnnounceNo()
	})
}

// SetTitle sets the "title" field.
func (u *AnnounceUpsertBulk) SetTitle(v string) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateTitle() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateTitle()
	})
}

// SetKind sets the "kind" field.
func (u *AnnounceUpsertBulk) SetKind(v announce.Kind) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateKind() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateKind()
	})
}

// SetContent sets the "content" field.
func (u *AnnounceUpsertBulk) SetContent(v string) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateContent() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateContent()
	})
}

// SetLink sets the "link" field.
func (u *AnnounceUpsertBulk) SetLink(v string) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateLink() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateLink()
	})
}

// SetSign sets the "sign" field.
func (u *AnnounceUpsertBulk) SetSign(v string) *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.SetSign(v)
	})
}

// UpdateSign sets the "sign" field to the value that was provided on create.
func (u *AnnounceUpsertBulk) UpdateSign() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.UpdateSign()
	})
}

// ClearSign clears the value of the "sign" field.
func (u *AnnounceUpsertBulk) ClearSign() *AnnounceUpsertBulk {
	return u.Update(func(s *AnnounceUpsert) {
		s.ClearSign()
	})
}

// Exec executes the query.
func (u *AnnounceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("media: OnConflict was set for builder %d. Set it on the AnnounceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("media: missing options for AnnounceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnnounceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
