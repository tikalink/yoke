// Code generated by entc, DO NOT EDIT.

package enttest

import (
	"context"

	"github.com/tikafog/of/dbc/media"
	// required by schema hooks.
	_ "github.com/tikafog/of/dbc/media/runtime"

	"entgo.io/ent/dialect/sql/schema"
)

type (
	// TestingT is the interface that is shared between
	// testing.T and testing.B and used by enttest.
	TestingT interface {
		FailNow()
		Error(...interface{})
	}

	// Option configures client creation.
	Option func(*options)

	options struct {
		opts        []media.Option
		migrateOpts []schema.MigrateOption
	}
)

// WithOptions forwards options to client creation.
func WithOptions(opts ...media.Option) Option {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

// WithMigrateOptions forwards options to auto migration.
func WithMigrateOptions(opts ...schema.MigrateOption) Option {
	return func(o *options) {
		o.migrateOpts = append(o.migrateOpts, opts...)
	}
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Open calls media.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, opts ...Option) *media.Client {
	o := newOptions(opts)
	c, err := media.Open(driverName, dataSourceName, o.opts...)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := c.Schema.Create(context.Background(), o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
	return c
}

// NewClient calls media.NewClient and auto-run migration.
func NewClient(t TestingT, opts ...Option) *media.Client {
	o := newOptions(opts)
	c := media.NewClient(o.opts...)
	if err := c.Schema.Create(context.Background(), o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
	return c
}
