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
	"github.com/tikafog/of/dbc/media/page"
	"github.com/tikafog/of/dbc/media/toplist"
)

// PageCreate is the builder for creating a Page entity.
type PageCreate struct {
	config
	mutation *PageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedUnix sets the "created_unix" field.
func (pc *PageCreate) SetCreatedUnix(i int64) *PageCreate {
	pc.mutation.SetCreatedUnix(i)
	return pc
}

// SetNillableCreatedUnix sets the "created_unix" field if the given value is not nil.
func (pc *PageCreate) SetNillableCreatedUnix(i *int64) *PageCreate {
	if i != nil {
		pc.SetCreatedUnix(*i)
	}
	return pc
}

// SetUpdatedUnix sets the "updated_unix" field.
func (pc *PageCreate) SetUpdatedUnix(i int64) *PageCreate {
	pc.mutation.SetUpdatedUnix(i)
	return pc
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (pc *PageCreate) SetNillableUpdatedUnix(i *int64) *PageCreate {
	if i != nil {
		pc.SetUpdatedUnix(*i)
	}
	return pc
}

// SetDeletedUnix sets the "deleted_unix" field.
func (pc *PageCreate) SetDeletedUnix(i int64) *PageCreate {
	pc.mutation.SetDeletedUnix(i)
	return pc
}

// SetNillableDeletedUnix sets the "deleted_unix" field if the given value is not nil.
func (pc *PageCreate) SetNillableDeletedUnix(i *int64) *PageCreate {
	if i != nil {
		pc.SetDeletedUnix(*i)
	}
	return pc
}

// SetParentID sets the "parent_id" field.
func (pc *PageCreate) SetParentID(u uuid.UUID) *PageCreate {
	pc.mutation.SetParentID(u)
	return pc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pc *PageCreate) SetNillableParentID(u *uuid.UUID) *PageCreate {
	if u != nil {
		pc.SetParentID(*u)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PageCreate) SetTitle(s string) *PageCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pc *PageCreate) SetNillableTitle(s *string) *PageCreate {
	if s != nil {
		pc.SetTitle(*s)
	}
	return pc
}

// SetFeaturedIndex sets the "featured_index" field.
func (pc *PageCreate) SetFeaturedIndex(i int8) *PageCreate {
	pc.mutation.SetFeaturedIndex(i)
	return pc
}

// SetNillableFeaturedIndex sets the "featured_index" field if the given value is not nil.
func (pc *PageCreate) SetNillableFeaturedIndex(i *int8) *PageCreate {
	if i != nil {
		pc.SetFeaturedIndex(*i)
	}
	return pc
}

// SetFeaturedContent sets the "featured_content" field.
func (pc *PageCreate) SetFeaturedContent(s string) *PageCreate {
	pc.mutation.SetFeaturedContent(s)
	return pc
}

// SetNillableFeaturedContent sets the "featured_content" field if the given value is not nil.
func (pc *PageCreate) SetNillableFeaturedContent(s *string) *PageCreate {
	if s != nil {
		pc.SetFeaturedContent(*s)
	}
	return pc
}

// SetRecommend sets the "recommend" field.
func (pc *PageCreate) SetRecommend(i int64) *PageCreate {
	pc.mutation.SetRecommend(i)
	return pc
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (pc *PageCreate) SetNillableRecommend(i *int64) *PageCreate {
	if i != nil {
		pc.SetRecommend(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PageCreate) SetID(u uuid.UUID) *PageCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PageCreate) SetNillableID(u *uuid.UUID) *PageCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetParent sets the "parent" edge to the Page entity.
func (pc *PageCreate) SetParent(p *Page) *PageCreate {
	return pc.SetParentID(p.ID)
}

// AddSubpageIDs adds the "subpages" edge to the Page entity by IDs.
func (pc *PageCreate) AddSubpageIDs(ids ...uuid.UUID) *PageCreate {
	pc.mutation.AddSubpageIDs(ids...)
	return pc
}

// AddSubpages adds the "subpages" edges to the Page entity.
func (pc *PageCreate) AddSubpages(p ...*Page) *PageCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddSubpageIDs(ids...)
}

// AddTopListIDs adds the "top_lists" edge to the TopList entity by IDs.
func (pc *PageCreate) AddTopListIDs(ids ...uuid.UUID) *PageCreate {
	pc.mutation.AddTopListIDs(ids...)
	return pc
}

// AddTopLists adds the "top_lists" edges to the TopList entity.
func (pc *PageCreate) AddTopLists(t ...*TopList) *PageCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTopListIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (pc *PageCreate) Mutation() *PageMutation {
	return pc.mutation
}

// Save creates the Page in the database.
func (pc *PageCreate) Save(ctx context.Context) (*Page, error) {
	var (
		err  error
		node *Page
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("media: uninitialized hook (forgotten import media/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PageCreate) SaveX(ctx context.Context) *Page {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PageCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PageCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PageCreate) defaults() {
	if _, ok := pc.mutation.CreatedUnix(); !ok {
		v := page.DefaultCreatedUnix
		pc.mutation.SetCreatedUnix(v)
	}
	if _, ok := pc.mutation.UpdatedUnix(); !ok {
		v := page.DefaultUpdatedUnix
		pc.mutation.SetUpdatedUnix(v)
	}
	if _, ok := pc.mutation.Title(); !ok {
		v := page.DefaultTitle
		pc.mutation.SetTitle(v)
	}
	if _, ok := pc.mutation.FeaturedIndex(); !ok {
		v := page.DefaultFeaturedIndex
		pc.mutation.SetFeaturedIndex(v)
	}
	if _, ok := pc.mutation.FeaturedContent(); !ok {
		v := page.DefaultFeaturedContent
		pc.mutation.SetFeaturedContent(v)
	}
	if _, ok := pc.mutation.Recommend(); !ok {
		v := page.DefaultRecommend
		pc.mutation.SetRecommend(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := page.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PageCreate) check() error {
	if _, ok := pc.mutation.CreatedUnix(); !ok {
		return &ValidationError{Name: "created_unix", err: errors.New(`media: missing required field "Page.created_unix"`)}
	}
	if _, ok := pc.mutation.UpdatedUnix(); !ok {
		return &ValidationError{Name: "updated_unix", err: errors.New(`media: missing required field "Page.updated_unix"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`media: missing required field "Page.title"`)}
	}
	if _, ok := pc.mutation.FeaturedIndex(); !ok {
		return &ValidationError{Name: "featured_index", err: errors.New(`media: missing required field "Page.featured_index"`)}
	}
	if _, ok := pc.mutation.FeaturedContent(); !ok {
		return &ValidationError{Name: "featured_content", err: errors.New(`media: missing required field "Page.featured_content"`)}
	}
	if _, ok := pc.mutation.Recommend(); !ok {
		return &ValidationError{Name: "recommend", err: errors.New(`media: missing required field "Page.recommend"`)}
	}
	return nil
}

func (pc *PageCreate) sqlSave(ctx context.Context) (*Page, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *PageCreate) createSpec() (*Page, *sqlgraph.CreateSpec) {
	var (
		_node = &Page{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: page.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: page.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: page.FieldCreatedUnix,
		})
		_node.CreatedUnix = value
	}
	if value, ok := pc.mutation.UpdatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: page.FieldUpdatedUnix,
		})
		_node.UpdatedUnix = value
	}
	if value, ok := pc.mutation.DeletedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: page.FieldDeletedUnix,
		})
		_node.DeletedUnix = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.FeaturedIndex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: page.FieldFeaturedIndex,
		})
		_node.FeaturedIndex = value
	}
	if value, ok := pc.mutation.FeaturedContent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldFeaturedContent,
		})
		_node.FeaturedContent = value
	}
	if value, ok := pc.mutation.Recommend(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: page.FieldRecommend,
		})
		_node.Recommend = value
	}
	if nodes := pc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.ParentTable,
			Columns: []string{page.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: page.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SubpagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.SubpagesTable,
			Columns: []string{page.SubpagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: page.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TopListsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.TopListsTable,
			Columns: []string{page.TopListsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: toplist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Page.Create().
//		SetCreatedUnix(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PageUpsert) {
//			SetCreatedUnix(v+v).
//		}).
//		Exec(ctx)
//
func (pc *PageCreate) OnConflict(opts ...sql.ConflictOption) *PageUpsertOne {
	pc.conflict = opts
	return &PageUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Page.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pc *PageCreate) OnConflictColumns(columns ...string) *PageUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PageUpsertOne{
		create: pc,
	}
}

type (
	// PageUpsertOne is the builder for "upsert"-ing
	//  one Page node.
	PageUpsertOne struct {
		create *PageCreate
	}

	// PageUpsert is the "OnConflict" setter.
	PageUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedUnix sets the "created_unix" field.
func (u *PageUpsert) SetCreatedUnix(v int64) *PageUpsert {
	u.Set(page.FieldCreatedUnix, v)
	return u
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *PageUpsert) UpdateCreatedUnix() *PageUpsert {
	u.SetExcluded(page.FieldCreatedUnix)
	return u
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *PageUpsert) AddCreatedUnix(v int64) *PageUpsert {
	u.Add(page.FieldCreatedUnix, v)
	return u
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *PageUpsert) SetUpdatedUnix(v int64) *PageUpsert {
	u.Set(page.FieldUpdatedUnix, v)
	return u
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *PageUpsert) UpdateUpdatedUnix() *PageUpsert {
	u.SetExcluded(page.FieldUpdatedUnix)
	return u
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *PageUpsert) AddUpdatedUnix(v int64) *PageUpsert {
	u.Add(page.FieldUpdatedUnix, v)
	return u
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *PageUpsert) SetDeletedUnix(v int64) *PageUpsert {
	u.Set(page.FieldDeletedUnix, v)
	return u
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *PageUpsert) UpdateDeletedUnix() *PageUpsert {
	u.SetExcluded(page.FieldDeletedUnix)
	return u
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *PageUpsert) AddDeletedUnix(v int64) *PageUpsert {
	u.Add(page.FieldDeletedUnix, v)
	return u
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *PageUpsert) ClearDeletedUnix() *PageUpsert {
	u.SetNull(page.FieldDeletedUnix)
	return u
}

// SetParentID sets the "parent_id" field.
func (u *PageUpsert) SetParentID(v uuid.UUID) *PageUpsert {
	u.Set(page.FieldParentID, v)
	return u
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *PageUpsert) UpdateParentID() *PageUpsert {
	u.SetExcluded(page.FieldParentID)
	return u
}

// ClearParentID clears the value of the "parent_id" field.
func (u *PageUpsert) ClearParentID() *PageUpsert {
	u.SetNull(page.FieldParentID)
	return u
}

// SetTitle sets the "title" field.
func (u *PageUpsert) SetTitle(v string) *PageUpsert {
	u.Set(page.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PageUpsert) UpdateTitle() *PageUpsert {
	u.SetExcluded(page.FieldTitle)
	return u
}

// SetFeaturedIndex sets the "featured_index" field.
func (u *PageUpsert) SetFeaturedIndex(v int8) *PageUpsert {
	u.Set(page.FieldFeaturedIndex, v)
	return u
}

// UpdateFeaturedIndex sets the "featured_index" field to the value that was provided on create.
func (u *PageUpsert) UpdateFeaturedIndex() *PageUpsert {
	u.SetExcluded(page.FieldFeaturedIndex)
	return u
}

// AddFeaturedIndex adds v to the "featured_index" field.
func (u *PageUpsert) AddFeaturedIndex(v int8) *PageUpsert {
	u.Add(page.FieldFeaturedIndex, v)
	return u
}

// SetFeaturedContent sets the "featured_content" field.
func (u *PageUpsert) SetFeaturedContent(v string) *PageUpsert {
	u.Set(page.FieldFeaturedContent, v)
	return u
}

// UpdateFeaturedContent sets the "featured_content" field to the value that was provided on create.
func (u *PageUpsert) UpdateFeaturedContent() *PageUpsert {
	u.SetExcluded(page.FieldFeaturedContent)
	return u
}

// SetRecommend sets the "recommend" field.
func (u *PageUpsert) SetRecommend(v int64) *PageUpsert {
	u.Set(page.FieldRecommend, v)
	return u
}

// UpdateRecommend sets the "recommend" field to the value that was provided on create.
func (u *PageUpsert) UpdateRecommend() *PageUpsert {
	u.SetExcluded(page.FieldRecommend)
	return u
}

// AddRecommend adds v to the "recommend" field.
func (u *PageUpsert) AddRecommend(v int64) *PageUpsert {
	u.Add(page.FieldRecommend, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Page.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(page.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PageUpsertOne) UpdateNewValues() *PageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(page.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Page.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PageUpsertOne) Ignore() *PageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PageUpsertOne) DoNothing() *PageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PageCreate.OnConflict
// documentation for more info.
func (u *PageUpsertOne) Update(set func(*PageUpsert)) *PageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedUnix sets the "created_unix" field.
func (u *PageUpsertOne) SetCreatedUnix(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetCreatedUnix(v)
	})
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *PageUpsertOne) AddCreatedUnix(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.AddCreatedUnix(v)
	})
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateCreatedUnix() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateCreatedUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *PageUpsertOne) SetUpdatedUnix(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *PageUpsertOne) AddUpdatedUnix(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateUpdatedUnix() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *PageUpsertOne) SetDeletedUnix(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetDeletedUnix(v)
	})
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *PageUpsertOne) AddDeletedUnix(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.AddDeletedUnix(v)
	})
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateDeletedUnix() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateDeletedUnix()
	})
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *PageUpsertOne) ClearDeletedUnix() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.ClearDeletedUnix()
	})
}

// SetParentID sets the "parent_id" field.
func (u *PageUpsertOne) SetParentID(v uuid.UUID) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateParentID() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *PageUpsertOne) ClearParentID() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.ClearParentID()
	})
}

// SetTitle sets the "title" field.
func (u *PageUpsertOne) SetTitle(v string) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateTitle() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateTitle()
	})
}

// SetFeaturedIndex sets the "featured_index" field.
func (u *PageUpsertOne) SetFeaturedIndex(v int8) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetFeaturedIndex(v)
	})
}

// AddFeaturedIndex adds v to the "featured_index" field.
func (u *PageUpsertOne) AddFeaturedIndex(v int8) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.AddFeaturedIndex(v)
	})
}

// UpdateFeaturedIndex sets the "featured_index" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateFeaturedIndex() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateFeaturedIndex()
	})
}

// SetFeaturedContent sets the "featured_content" field.
func (u *PageUpsertOne) SetFeaturedContent(v string) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetFeaturedContent(v)
	})
}

// UpdateFeaturedContent sets the "featured_content" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateFeaturedContent() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateFeaturedContent()
	})
}

// SetRecommend sets the "recommend" field.
func (u *PageUpsertOne) SetRecommend(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.SetRecommend(v)
	})
}

// AddRecommend adds v to the "recommend" field.
func (u *PageUpsertOne) AddRecommend(v int64) *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.AddRecommend(v)
	})
}

// UpdateRecommend sets the "recommend" field to the value that was provided on create.
func (u *PageUpsertOne) UpdateRecommend() *PageUpsertOne {
	return u.Update(func(s *PageUpsert) {
		s.UpdateRecommend()
	})
}

// Exec executes the query.
func (u *PageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("media: missing options for PageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PageUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("media: PageUpsertOne.ID is not supported by MySQL driver. Use PageUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PageUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PageCreateBulk is the builder for creating many Page entities in bulk.
type PageCreateBulk struct {
	config
	builders []*PageCreate
	conflict []sql.ConflictOption
}

// Save creates the Page entities in the database.
func (pcb *PageCreateBulk) Save(ctx context.Context) ([]*Page, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Page, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PageMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PageCreateBulk) SaveX(ctx context.Context) []*Page {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PageCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PageCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Page.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PageUpsert) {
//			SetCreatedUnix(v+v).
//		}).
//		Exec(ctx)
//
func (pcb *PageCreateBulk) OnConflict(opts ...sql.ConflictOption) *PageUpsertBulk {
	pcb.conflict = opts
	return &PageUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Page.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pcb *PageCreateBulk) OnConflictColumns(columns ...string) *PageUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PageUpsertBulk{
		create: pcb,
	}
}

// PageUpsertBulk is the builder for "upsert"-ing
// a bulk of Page nodes.
type PageUpsertBulk struct {
	create *PageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Page.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(page.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PageUpsertBulk) UpdateNewValues() *PageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(page.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Page.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PageUpsertBulk) Ignore() *PageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PageUpsertBulk) DoNothing() *PageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PageCreateBulk.OnConflict
// documentation for more info.
func (u *PageUpsertBulk) Update(set func(*PageUpsert)) *PageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedUnix sets the "created_unix" field.
func (u *PageUpsertBulk) SetCreatedUnix(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetCreatedUnix(v)
	})
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *PageUpsertBulk) AddCreatedUnix(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.AddCreatedUnix(v)
	})
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateCreatedUnix() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateCreatedUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *PageUpsertBulk) SetUpdatedUnix(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *PageUpsertBulk) AddUpdatedUnix(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateUpdatedUnix() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *PageUpsertBulk) SetDeletedUnix(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetDeletedUnix(v)
	})
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *PageUpsertBulk) AddDeletedUnix(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.AddDeletedUnix(v)
	})
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateDeletedUnix() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateDeletedUnix()
	})
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *PageUpsertBulk) ClearDeletedUnix() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.ClearDeletedUnix()
	})
}

// SetParentID sets the "parent_id" field.
func (u *PageUpsertBulk) SetParentID(v uuid.UUID) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateParentID() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *PageUpsertBulk) ClearParentID() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.ClearParentID()
	})
}

// SetTitle sets the "title" field.
func (u *PageUpsertBulk) SetTitle(v string) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateTitle() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateTitle()
	})
}

// SetFeaturedIndex sets the "featured_index" field.
func (u *PageUpsertBulk) SetFeaturedIndex(v int8) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetFeaturedIndex(v)
	})
}

// AddFeaturedIndex adds v to the "featured_index" field.
func (u *PageUpsertBulk) AddFeaturedIndex(v int8) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.AddFeaturedIndex(v)
	})
}

// UpdateFeaturedIndex sets the "featured_index" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateFeaturedIndex() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateFeaturedIndex()
	})
}

// SetFeaturedContent sets the "featured_content" field.
func (u *PageUpsertBulk) SetFeaturedContent(v string) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetFeaturedContent(v)
	})
}

// UpdateFeaturedContent sets the "featured_content" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateFeaturedContent() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateFeaturedContent()
	})
}

// SetRecommend sets the "recommend" field.
func (u *PageUpsertBulk) SetRecommend(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.SetRecommend(v)
	})
}

// AddRecommend adds v to the "recommend" field.
func (u *PageUpsertBulk) AddRecommend(v int64) *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.AddRecommend(v)
	})
}

// UpdateRecommend sets the "recommend" field to the value that was provided on create.
func (u *PageUpsertBulk) UpdateRecommend() *PageUpsertBulk {
	return u.Update(func(s *PageUpsert) {
		s.UpdateRecommend()
	})
}

// Exec executes the query.
func (u *PageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("media: OnConflict was set for builder %d. Set it on the PageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("media: missing options for PageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
