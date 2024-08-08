// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/marko-durasic/olympics/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/marko-durasic/olympics/ent/combinedsport"
	"github.com/marko-durasic/olympics/ent/individualsport"
	"github.com/marko-durasic/olympics/ent/teamsport"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CombinedSport is the client for interacting with the CombinedSport builders.
	CombinedSport *CombinedSportClient
	// IndividualSport is the client for interacting with the IndividualSport builders.
	IndividualSport *IndividualSportClient
	// TeamSport is the client for interacting with the TeamSport builders.
	TeamSport *TeamSportClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CombinedSport = NewCombinedSportClient(c.config)
	c.IndividualSport = NewIndividualSportClient(c.config)
	c.TeamSport = NewTeamSportClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		CombinedSport:   NewCombinedSportClient(cfg),
		IndividualSport: NewIndividualSportClient(cfg),
		TeamSport:       NewTeamSportClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:             ctx,
		config:          cfg,
		CombinedSport:   NewCombinedSportClient(cfg),
		IndividualSport: NewIndividualSportClient(cfg),
		TeamSport:       NewTeamSportClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CombinedSport.
//		Query().
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
	c.CombinedSport.Use(hooks...)
	c.IndividualSport.Use(hooks...)
	c.TeamSport.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.CombinedSport.Intercept(interceptors...)
	c.IndividualSport.Intercept(interceptors...)
	c.TeamSport.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CombinedSportMutation:
		return c.CombinedSport.mutate(ctx, m)
	case *IndividualSportMutation:
		return c.IndividualSport.mutate(ctx, m)
	case *TeamSportMutation:
		return c.TeamSport.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CombinedSportClient is a client for the CombinedSport schema.
type CombinedSportClient struct {
	config
}

// NewCombinedSportClient returns a client for the CombinedSport from the given config.
func NewCombinedSportClient(c config) *CombinedSportClient {
	return &CombinedSportClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `combinedsport.Hooks(f(g(h())))`.
func (c *CombinedSportClient) Use(hooks ...Hook) {
	c.hooks.CombinedSport = append(c.hooks.CombinedSport, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `combinedsport.Intercept(f(g(h())))`.
func (c *CombinedSportClient) Intercept(interceptors ...Interceptor) {
	c.inters.CombinedSport = append(c.inters.CombinedSport, interceptors...)
}

// Create returns a builder for creating a CombinedSport entity.
func (c *CombinedSportClient) Create() *CombinedSportCreate {
	mutation := newCombinedSportMutation(c.config, OpCreate)
	return &CombinedSportCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CombinedSport entities.
func (c *CombinedSportClient) CreateBulk(builders ...*CombinedSportCreate) *CombinedSportCreateBulk {
	return &CombinedSportCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CombinedSportClient) MapCreateBulk(slice any, setFunc func(*CombinedSportCreate, int)) *CombinedSportCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CombinedSportCreateBulk{err: fmt.Errorf("calling to CombinedSportClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CombinedSportCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CombinedSportCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CombinedSport.
func (c *CombinedSportClient) Update() *CombinedSportUpdate {
	mutation := newCombinedSportMutation(c.config, OpUpdate)
	return &CombinedSportUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CombinedSportClient) UpdateOne(cs *CombinedSport) *CombinedSportUpdateOne {
	mutation := newCombinedSportMutation(c.config, OpUpdateOne, withCombinedSport(cs))
	return &CombinedSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CombinedSportClient) UpdateOneID(id int) *CombinedSportUpdateOne {
	mutation := newCombinedSportMutation(c.config, OpUpdateOne, withCombinedSportID(id))
	return &CombinedSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CombinedSport.
func (c *CombinedSportClient) Delete() *CombinedSportDelete {
	mutation := newCombinedSportMutation(c.config, OpDelete)
	return &CombinedSportDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CombinedSportClient) DeleteOne(cs *CombinedSport) *CombinedSportDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CombinedSportClient) DeleteOneID(id int) *CombinedSportDeleteOne {
	builder := c.Delete().Where(combinedsport.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CombinedSportDeleteOne{builder}
}

// Query returns a query builder for CombinedSport.
func (c *CombinedSportClient) Query() *CombinedSportQuery {
	return &CombinedSportQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCombinedSport},
		inters: c.Interceptors(),
	}
}

// Get returns a CombinedSport entity by its id.
func (c *CombinedSportClient) Get(ctx context.Context, id int) (*CombinedSport, error) {
	return c.Query().Where(combinedsport.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CombinedSportClient) GetX(ctx context.Context, id int) *CombinedSport {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CombinedSportClient) Hooks() []Hook {
	return c.hooks.CombinedSport
}

// Interceptors returns the client interceptors.
func (c *CombinedSportClient) Interceptors() []Interceptor {
	return c.inters.CombinedSport
}

func (c *CombinedSportClient) mutate(ctx context.Context, m *CombinedSportMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CombinedSportCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CombinedSportUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CombinedSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CombinedSportDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CombinedSport mutation op: %q", m.Op())
	}
}

// IndividualSportClient is a client for the IndividualSport schema.
type IndividualSportClient struct {
	config
}

// NewIndividualSportClient returns a client for the IndividualSport from the given config.
func NewIndividualSportClient(c config) *IndividualSportClient {
	return &IndividualSportClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `individualsport.Hooks(f(g(h())))`.
func (c *IndividualSportClient) Use(hooks ...Hook) {
	c.hooks.IndividualSport = append(c.hooks.IndividualSport, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `individualsport.Intercept(f(g(h())))`.
func (c *IndividualSportClient) Intercept(interceptors ...Interceptor) {
	c.inters.IndividualSport = append(c.inters.IndividualSport, interceptors...)
}

// Create returns a builder for creating a IndividualSport entity.
func (c *IndividualSportClient) Create() *IndividualSportCreate {
	mutation := newIndividualSportMutation(c.config, OpCreate)
	return &IndividualSportCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of IndividualSport entities.
func (c *IndividualSportClient) CreateBulk(builders ...*IndividualSportCreate) *IndividualSportCreateBulk {
	return &IndividualSportCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *IndividualSportClient) MapCreateBulk(slice any, setFunc func(*IndividualSportCreate, int)) *IndividualSportCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &IndividualSportCreateBulk{err: fmt.Errorf("calling to IndividualSportClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*IndividualSportCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &IndividualSportCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for IndividualSport.
func (c *IndividualSportClient) Update() *IndividualSportUpdate {
	mutation := newIndividualSportMutation(c.config, OpUpdate)
	return &IndividualSportUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IndividualSportClient) UpdateOne(is *IndividualSport) *IndividualSportUpdateOne {
	mutation := newIndividualSportMutation(c.config, OpUpdateOne, withIndividualSport(is))
	return &IndividualSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IndividualSportClient) UpdateOneID(id int) *IndividualSportUpdateOne {
	mutation := newIndividualSportMutation(c.config, OpUpdateOne, withIndividualSportID(id))
	return &IndividualSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for IndividualSport.
func (c *IndividualSportClient) Delete() *IndividualSportDelete {
	mutation := newIndividualSportMutation(c.config, OpDelete)
	return &IndividualSportDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *IndividualSportClient) DeleteOne(is *IndividualSport) *IndividualSportDeleteOne {
	return c.DeleteOneID(is.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *IndividualSportClient) DeleteOneID(id int) *IndividualSportDeleteOne {
	builder := c.Delete().Where(individualsport.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IndividualSportDeleteOne{builder}
}

// Query returns a query builder for IndividualSport.
func (c *IndividualSportClient) Query() *IndividualSportQuery {
	return &IndividualSportQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeIndividualSport},
		inters: c.Interceptors(),
	}
}

// Get returns a IndividualSport entity by its id.
func (c *IndividualSportClient) Get(ctx context.Context, id int) (*IndividualSport, error) {
	return c.Query().Where(individualsport.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IndividualSportClient) GetX(ctx context.Context, id int) *IndividualSport {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *IndividualSportClient) Hooks() []Hook {
	return c.hooks.IndividualSport
}

// Interceptors returns the client interceptors.
func (c *IndividualSportClient) Interceptors() []Interceptor {
	return c.inters.IndividualSport
}

func (c *IndividualSportClient) mutate(ctx context.Context, m *IndividualSportMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&IndividualSportCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&IndividualSportUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&IndividualSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&IndividualSportDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown IndividualSport mutation op: %q", m.Op())
	}
}

// TeamSportClient is a client for the TeamSport schema.
type TeamSportClient struct {
	config
}

// NewTeamSportClient returns a client for the TeamSport from the given config.
func NewTeamSportClient(c config) *TeamSportClient {
	return &TeamSportClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `teamsport.Hooks(f(g(h())))`.
func (c *TeamSportClient) Use(hooks ...Hook) {
	c.hooks.TeamSport = append(c.hooks.TeamSport, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `teamsport.Intercept(f(g(h())))`.
func (c *TeamSportClient) Intercept(interceptors ...Interceptor) {
	c.inters.TeamSport = append(c.inters.TeamSport, interceptors...)
}

// Create returns a builder for creating a TeamSport entity.
func (c *TeamSportClient) Create() *TeamSportCreate {
	mutation := newTeamSportMutation(c.config, OpCreate)
	return &TeamSportCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TeamSport entities.
func (c *TeamSportClient) CreateBulk(builders ...*TeamSportCreate) *TeamSportCreateBulk {
	return &TeamSportCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TeamSportClient) MapCreateBulk(slice any, setFunc func(*TeamSportCreate, int)) *TeamSportCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TeamSportCreateBulk{err: fmt.Errorf("calling to TeamSportClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TeamSportCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TeamSportCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TeamSport.
func (c *TeamSportClient) Update() *TeamSportUpdate {
	mutation := newTeamSportMutation(c.config, OpUpdate)
	return &TeamSportUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeamSportClient) UpdateOne(ts *TeamSport) *TeamSportUpdateOne {
	mutation := newTeamSportMutation(c.config, OpUpdateOne, withTeamSport(ts))
	return &TeamSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeamSportClient) UpdateOneID(id int) *TeamSportUpdateOne {
	mutation := newTeamSportMutation(c.config, OpUpdateOne, withTeamSportID(id))
	return &TeamSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TeamSport.
func (c *TeamSportClient) Delete() *TeamSportDelete {
	mutation := newTeamSportMutation(c.config, OpDelete)
	return &TeamSportDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TeamSportClient) DeleteOne(ts *TeamSport) *TeamSportDeleteOne {
	return c.DeleteOneID(ts.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TeamSportClient) DeleteOneID(id int) *TeamSportDeleteOne {
	builder := c.Delete().Where(teamsport.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeamSportDeleteOne{builder}
}

// Query returns a query builder for TeamSport.
func (c *TeamSportClient) Query() *TeamSportQuery {
	return &TeamSportQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTeamSport},
		inters: c.Interceptors(),
	}
}

// Get returns a TeamSport entity by its id.
func (c *TeamSportClient) Get(ctx context.Context, id int) (*TeamSport, error) {
	return c.Query().Where(teamsport.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeamSportClient) GetX(ctx context.Context, id int) *TeamSport {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TeamSportClient) Hooks() []Hook {
	return c.hooks.TeamSport
}

// Interceptors returns the client interceptors.
func (c *TeamSportClient) Interceptors() []Interceptor {
	return c.inters.TeamSport
}

func (c *TeamSportClient) mutate(ctx context.Context, m *TeamSportMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TeamSportCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TeamSportUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TeamSportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TeamSportDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TeamSport mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		CombinedSport, IndividualSport, TeamSport []ent.Hook
	}
	inters struct {
		CombinedSport, IndividualSport, TeamSport []ent.Interceptor
	}
)
