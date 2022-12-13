// Code generated by ent, DO NOT EDIT.

package kernel

import (
	"context"
	"fmt"
	"log"

	"github.com/tikafog/of/dbc/kernel/migrate"

	"github.com/tikafog/of/dbc/kernel/instruct"
	"github.com/tikafog/of/dbc/kernel/resource"
	"github.com/tikafog/of/dbc/kernel/version"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Instruct is the client for interacting with the Instruct builders.
	Instruct *InstructClient
	// Resource is the client for interacting with the Resource builders.
	Resource *ResourceClient
	// Version is the client for interacting with the Version builders.
	Version *VersionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Instruct = NewInstructClient(c.config)
	c.Resource = NewResourceClient(c.config)
	c.Version = NewVersionClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("kernel: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("kernel: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Instruct: NewInstructClient(cfg),
		Resource: NewResourceClient(cfg),
		Version:  NewVersionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Instruct: NewInstructClient(cfg),
		Resource: NewResourceClient(cfg),
		Version:  NewVersionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Instruct.
//		DB().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Instruct.Use(hooks...)
	c.Resource.Use(hooks...)
	c.Version.Use(hooks...)
}

// InstructClient is a client for the Instruct schema.
type InstructClient struct {
	config
}

// NewInstructClient returns a client for the Instruct from the given config.
func NewInstructClient(c config) *InstructClient {
	return &InstructClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instruct.Hooks(f(g(h())))`.
func (c *InstructClient) Use(hooks ...Hook) {
	c.hooks.Instruct = append(c.hooks.Instruct, hooks...)
}

// Create returns a builder for creating a Instruct entity.
func (c *InstructClient) Create() *InstructCreate {
	mutation := newInstructMutation(c.config, OpCreate)
	return &InstructCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Instruct entities.
func (c *InstructClient) CreateBulk(builders ...*InstructCreate) *InstructCreateBulk {
	return &InstructCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Instruct.
func (c *InstructClient) Update() *InstructUpdate {
	mutation := newInstructMutation(c.config, OpUpdate)
	return &InstructUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstructClient) UpdateOne(i *Instruct) *InstructUpdateOne {
	mutation := newInstructMutation(c.config, OpUpdateOne, withInstruct(i))
	return &InstructUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstructClient) UpdateOneID(id int) *InstructUpdateOne {
	mutation := newInstructMutation(c.config, OpUpdateOne, withInstructID(id))
	return &InstructUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Instruct.
func (c *InstructClient) Delete() *InstructDelete {
	mutation := newInstructMutation(c.config, OpDelete)
	return &InstructDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *InstructClient) DeleteOne(i *Instruct) *InstructDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *InstructClient) DeleteOneID(id int) *InstructDeleteOne {
	builder := c.Delete().Where(instruct.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstructDeleteOne{builder}
}

// Query returns a query builder for Instruct.
func (c *InstructClient) Query() *InstructQuery {
	return &InstructQuery{
		config: c.config,
	}
}

// Get returns a Instruct entity by its id.
func (c *InstructClient) Get(ctx context.Context, id int) (*Instruct, error) {
	return c.Query().Where(instruct.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstructClient) GetX(ctx context.Context, id int) *Instruct {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *InstructClient) Hooks() []Hook {
	return c.hooks.Instruct
}

// ResourceClient is a client for the Resource schema.
type ResourceClient struct {
	config
}

// NewResourceClient returns a client for the Resource from the given config.
func NewResourceClient(c config) *ResourceClient {
	return &ResourceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `resource.Hooks(f(g(h())))`.
func (c *ResourceClient) Use(hooks ...Hook) {
	c.hooks.Resource = append(c.hooks.Resource, hooks...)
}

// Create returns a builder for creating a Resource entity.
func (c *ResourceClient) Create() *ResourceCreate {
	mutation := newResourceMutation(c.config, OpCreate)
	return &ResourceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Resource entities.
func (c *ResourceClient) CreateBulk(builders ...*ResourceCreate) *ResourceCreateBulk {
	return &ResourceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Resource.
func (c *ResourceClient) Update() *ResourceUpdate {
	mutation := newResourceMutation(c.config, OpUpdate)
	return &ResourceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ResourceClient) UpdateOne(r *Resource) *ResourceUpdateOne {
	mutation := newResourceMutation(c.config, OpUpdateOne, withResource(r))
	return &ResourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ResourceClient) UpdateOneID(id int) *ResourceUpdateOne {
	mutation := newResourceMutation(c.config, OpUpdateOne, withResourceID(id))
	return &ResourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Resource.
func (c *ResourceClient) Delete() *ResourceDelete {
	mutation := newResourceMutation(c.config, OpDelete)
	return &ResourceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ResourceClient) DeleteOne(r *Resource) *ResourceDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ResourceClient) DeleteOneID(id int) *ResourceDeleteOne {
	builder := c.Delete().Where(resource.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ResourceDeleteOne{builder}
}

// Query returns a query builder for Resource.
func (c *ResourceClient) Query() *ResourceQuery {
	return &ResourceQuery{
		config: c.config,
	}
}

// Get returns a Resource entity by its id.
func (c *ResourceClient) Get(ctx context.Context, id int) (*Resource, error) {
	return c.Query().Where(resource.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ResourceClient) GetX(ctx context.Context, id int) *Resource {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ResourceClient) Hooks() []Hook {
	return c.hooks.Resource
}

// VersionClient is a client for the Version schema.
type VersionClient struct {
	config
}

// NewVersionClient returns a client for the Version from the given config.
func NewVersionClient(c config) *VersionClient {
	return &VersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `version.Hooks(f(g(h())))`.
func (c *VersionClient) Use(hooks ...Hook) {
	c.hooks.Version = append(c.hooks.Version, hooks...)
}

// Create returns a builder for creating a Version entity.
func (c *VersionClient) Create() *VersionCreate {
	mutation := newVersionMutation(c.config, OpCreate)
	return &VersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Version entities.
func (c *VersionClient) CreateBulk(builders ...*VersionCreate) *VersionCreateBulk {
	return &VersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Version.
func (c *VersionClient) Update() *VersionUpdate {
	mutation := newVersionMutation(c.config, OpUpdate)
	return &VersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VersionClient) UpdateOne(v *Version) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersion(v))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VersionClient) UpdateOneID(id int) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersionID(id))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Version.
func (c *VersionClient) Delete() *VersionDelete {
	mutation := newVersionMutation(c.config, OpDelete)
	return &VersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VersionClient) DeleteOne(v *Version) *VersionDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *VersionClient) DeleteOneID(id int) *VersionDeleteOne {
	builder := c.Delete().Where(version.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VersionDeleteOne{builder}
}

// Query returns a query builder for Version.
func (c *VersionClient) Query() *VersionQuery {
	return &VersionQuery{
		config: c.config,
	}
}

// Get returns a Version entity by its id.
func (c *VersionClient) Get(ctx context.Context, id int) (*Version, error) {
	return c.Query().Where(version.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VersionClient) GetX(ctx context.Context, id int) *Version {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VersionClient) Hooks() []Hook {
	return c.hooks.Version
}
