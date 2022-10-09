package dbc

import (
	"time"
)

const (
	MinTimeout = 3 * time.Second
	MaxTimeout = 1800 * time.Second
)

type Option struct {
	timeout time.Duration
	debug   bool
	ignores map[ClientType]struct{}
}

func (o Option) Debug() bool {
	return o.debug
}

func (o Option) Timeout() time.Duration {
	if o.timeout < MinTimeout {
		return MinTimeout
	}
	if o.timeout > MaxTimeout {
		return MaxTimeout
	}
	return o.timeout
}

func (o Option) Ignores() map[ClientType]struct{} {
	return o.ignores
}

type Opts func(opt *Option)

func parseOption(opts []Opts) *Option {
	o := &Option{
		ignores: make(map[ClientType]struct{}),
		timeout: MinTimeout,
		debug:   false,
	}
	for i := range opts {
		opts[i](o)
	}
	return o
}

func Debug(b bool) Opts {
	return func(opt *Option) {
		opt.debug = b
	}
}

func TimeoutOpt(t time.Duration) Opts {
	return func(opt *Option) {
		opt.timeout = t
	}
}

func IgnoresOpt(ps ...ClientType) Opts {
	return func(opt *Option) {
		for i := range ps {
			opt.ignores[ps[i]] = struct{}{}
		}
	}
}
