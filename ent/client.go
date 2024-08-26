// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Milkado/api-challenge-jornada-milhas/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/destinies"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/destinypictures"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/testimonies"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/users"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Destinies is the client for interacting with the Destinies builders.
	Destinies *DestiniesClient
	// DestinyPictures is the client for interacting with the DestinyPictures builders.
	DestinyPictures *DestinyPicturesClient
	// Testimonies is the client for interacting with the Testimonies builders.
	Testimonies *TestimoniesClient
	// Users is the client for interacting with the Users builders.
	Users *UsersClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Destinies = NewDestiniesClient(c.config)
	c.DestinyPictures = NewDestinyPicturesClient(c.config)
	c.Testimonies = NewTestimoniesClient(c.config)
	c.Users = NewUsersClient(c.config)
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
		Destinies:       NewDestiniesClient(cfg),
		DestinyPictures: NewDestinyPicturesClient(cfg),
		Testimonies:     NewTestimoniesClient(cfg),
		Users:           NewUsersClient(cfg),
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
		Destinies:       NewDestiniesClient(cfg),
		DestinyPictures: NewDestinyPicturesClient(cfg),
		Testimonies:     NewTestimoniesClient(cfg),
		Users:           NewUsersClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Destinies.
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
	c.Destinies.Use(hooks...)
	c.DestinyPictures.Use(hooks...)
	c.Testimonies.Use(hooks...)
	c.Users.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Destinies.Intercept(interceptors...)
	c.DestinyPictures.Intercept(interceptors...)
	c.Testimonies.Intercept(interceptors...)
	c.Users.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DestiniesMutation:
		return c.Destinies.mutate(ctx, m)
	case *DestinyPicturesMutation:
		return c.DestinyPictures.mutate(ctx, m)
	case *TestimoniesMutation:
		return c.Testimonies.mutate(ctx, m)
	case *UsersMutation:
		return c.Users.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DestiniesClient is a client for the Destinies schema.
type DestiniesClient struct {
	config
}

// NewDestiniesClient returns a client for the Destinies from the given config.
func NewDestiniesClient(c config) *DestiniesClient {
	return &DestiniesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `destinies.Hooks(f(g(h())))`.
func (c *DestiniesClient) Use(hooks ...Hook) {
	c.hooks.Destinies = append(c.hooks.Destinies, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `destinies.Intercept(f(g(h())))`.
func (c *DestiniesClient) Intercept(interceptors ...Interceptor) {
	c.inters.Destinies = append(c.inters.Destinies, interceptors...)
}

// Create returns a builder for creating a Destinies entity.
func (c *DestiniesClient) Create() *DestiniesCreate {
	mutation := newDestiniesMutation(c.config, OpCreate)
	return &DestiniesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Destinies entities.
func (c *DestiniesClient) CreateBulk(builders ...*DestiniesCreate) *DestiniesCreateBulk {
	return &DestiniesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DestiniesClient) MapCreateBulk(slice any, setFunc func(*DestiniesCreate, int)) *DestiniesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DestiniesCreateBulk{err: fmt.Errorf("calling to DestiniesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DestiniesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DestiniesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Destinies.
func (c *DestiniesClient) Update() *DestiniesUpdate {
	mutation := newDestiniesMutation(c.config, OpUpdate)
	return &DestiniesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DestiniesClient) UpdateOne(d *Destinies) *DestiniesUpdateOne {
	mutation := newDestiniesMutation(c.config, OpUpdateOne, withDestinies(d))
	return &DestiniesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DestiniesClient) UpdateOneID(id int) *DestiniesUpdateOne {
	mutation := newDestiniesMutation(c.config, OpUpdateOne, withDestiniesID(id))
	return &DestiniesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Destinies.
func (c *DestiniesClient) Delete() *DestiniesDelete {
	mutation := newDestiniesMutation(c.config, OpDelete)
	return &DestiniesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DestiniesClient) DeleteOne(d *Destinies) *DestiniesDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DestiniesClient) DeleteOneID(id int) *DestiniesDeleteOne {
	builder := c.Delete().Where(destinies.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DestiniesDeleteOne{builder}
}

// Query returns a query builder for Destinies.
func (c *DestiniesClient) Query() *DestiniesQuery {
	return &DestiniesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDestinies},
		inters: c.Interceptors(),
	}
}

// Get returns a Destinies entity by its id.
func (c *DestiniesClient) Get(ctx context.Context, id int) (*Destinies, error) {
	return c.Query().Where(destinies.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DestiniesClient) GetX(ctx context.Context, id int) *Destinies {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTestimonies queries the testimonies edge of a Destinies.
func (c *DestiniesClient) QueryTestimonies(d *Destinies) *TestimoniesQuery {
	query := (&TestimoniesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(destinies.Table, destinies.FieldID, id),
			sqlgraph.To(testimonies.Table, testimonies.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, destinies.TestimoniesTable, destinies.TestimoniesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDestinyPictures queries the destiny_pictures edge of a Destinies.
func (c *DestiniesClient) QueryDestinyPictures(d *Destinies) *DestinyPicturesQuery {
	query := (&DestinyPicturesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(destinies.Table, destinies.FieldID, id),
			sqlgraph.To(destinypictures.Table, destinypictures.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, destinies.DestinyPicturesTable, destinies.DestinyPicturesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DestiniesClient) Hooks() []Hook {
	hooks := c.hooks.Destinies
	return append(hooks[:len(hooks):len(hooks)], destinies.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *DestiniesClient) Interceptors() []Interceptor {
	return c.inters.Destinies
}

func (c *DestiniesClient) mutate(ctx context.Context, m *DestiniesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DestiniesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DestiniesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DestiniesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DestiniesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Destinies mutation op: %q", m.Op())
	}
}

// DestinyPicturesClient is a client for the DestinyPictures schema.
type DestinyPicturesClient struct {
	config
}

// NewDestinyPicturesClient returns a client for the DestinyPictures from the given config.
func NewDestinyPicturesClient(c config) *DestinyPicturesClient {
	return &DestinyPicturesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `destinypictures.Hooks(f(g(h())))`.
func (c *DestinyPicturesClient) Use(hooks ...Hook) {
	c.hooks.DestinyPictures = append(c.hooks.DestinyPictures, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `destinypictures.Intercept(f(g(h())))`.
func (c *DestinyPicturesClient) Intercept(interceptors ...Interceptor) {
	c.inters.DestinyPictures = append(c.inters.DestinyPictures, interceptors...)
}

// Create returns a builder for creating a DestinyPictures entity.
func (c *DestinyPicturesClient) Create() *DestinyPicturesCreate {
	mutation := newDestinyPicturesMutation(c.config, OpCreate)
	return &DestinyPicturesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DestinyPictures entities.
func (c *DestinyPicturesClient) CreateBulk(builders ...*DestinyPicturesCreate) *DestinyPicturesCreateBulk {
	return &DestinyPicturesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DestinyPicturesClient) MapCreateBulk(slice any, setFunc func(*DestinyPicturesCreate, int)) *DestinyPicturesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DestinyPicturesCreateBulk{err: fmt.Errorf("calling to DestinyPicturesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DestinyPicturesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DestinyPicturesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DestinyPictures.
func (c *DestinyPicturesClient) Update() *DestinyPicturesUpdate {
	mutation := newDestinyPicturesMutation(c.config, OpUpdate)
	return &DestinyPicturesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DestinyPicturesClient) UpdateOne(dp *DestinyPictures) *DestinyPicturesUpdateOne {
	mutation := newDestinyPicturesMutation(c.config, OpUpdateOne, withDestinyPictures(dp))
	return &DestinyPicturesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DestinyPicturesClient) UpdateOneID(id int) *DestinyPicturesUpdateOne {
	mutation := newDestinyPicturesMutation(c.config, OpUpdateOne, withDestinyPicturesID(id))
	return &DestinyPicturesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DestinyPictures.
func (c *DestinyPicturesClient) Delete() *DestinyPicturesDelete {
	mutation := newDestinyPicturesMutation(c.config, OpDelete)
	return &DestinyPicturesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DestinyPicturesClient) DeleteOne(dp *DestinyPictures) *DestinyPicturesDeleteOne {
	return c.DeleteOneID(dp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DestinyPicturesClient) DeleteOneID(id int) *DestinyPicturesDeleteOne {
	builder := c.Delete().Where(destinypictures.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DestinyPicturesDeleteOne{builder}
}

// Query returns a query builder for DestinyPictures.
func (c *DestinyPicturesClient) Query() *DestinyPicturesQuery {
	return &DestinyPicturesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDestinyPictures},
		inters: c.Interceptors(),
	}
}

// Get returns a DestinyPictures entity by its id.
func (c *DestinyPicturesClient) Get(ctx context.Context, id int) (*DestinyPictures, error) {
	return c.Query().Where(destinypictures.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DestinyPicturesClient) GetX(ctx context.Context, id int) *DestinyPictures {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDestinies queries the destinies edge of a DestinyPictures.
func (c *DestinyPicturesClient) QueryDestinies(dp *DestinyPictures) *DestiniesQuery {
	query := (&DestiniesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := dp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(destinypictures.Table, destinypictures.FieldID, id),
			sqlgraph.To(destinies.Table, destinies.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, destinypictures.DestiniesTable, destinypictures.DestiniesColumn),
		)
		fromV = sqlgraph.Neighbors(dp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DestinyPicturesClient) Hooks() []Hook {
	return c.hooks.DestinyPictures
}

// Interceptors returns the client interceptors.
func (c *DestinyPicturesClient) Interceptors() []Interceptor {
	return c.inters.DestinyPictures
}

func (c *DestinyPicturesClient) mutate(ctx context.Context, m *DestinyPicturesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DestinyPicturesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DestinyPicturesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DestinyPicturesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DestinyPicturesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown DestinyPictures mutation op: %q", m.Op())
	}
}

// TestimoniesClient is a client for the Testimonies schema.
type TestimoniesClient struct {
	config
}

// NewTestimoniesClient returns a client for the Testimonies from the given config.
func NewTestimoniesClient(c config) *TestimoniesClient {
	return &TestimoniesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `testimonies.Hooks(f(g(h())))`.
func (c *TestimoniesClient) Use(hooks ...Hook) {
	c.hooks.Testimonies = append(c.hooks.Testimonies, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `testimonies.Intercept(f(g(h())))`.
func (c *TestimoniesClient) Intercept(interceptors ...Interceptor) {
	c.inters.Testimonies = append(c.inters.Testimonies, interceptors...)
}

// Create returns a builder for creating a Testimonies entity.
func (c *TestimoniesClient) Create() *TestimoniesCreate {
	mutation := newTestimoniesMutation(c.config, OpCreate)
	return &TestimoniesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Testimonies entities.
func (c *TestimoniesClient) CreateBulk(builders ...*TestimoniesCreate) *TestimoniesCreateBulk {
	return &TestimoniesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TestimoniesClient) MapCreateBulk(slice any, setFunc func(*TestimoniesCreate, int)) *TestimoniesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TestimoniesCreateBulk{err: fmt.Errorf("calling to TestimoniesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TestimoniesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TestimoniesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Testimonies.
func (c *TestimoniesClient) Update() *TestimoniesUpdate {
	mutation := newTestimoniesMutation(c.config, OpUpdate)
	return &TestimoniesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TestimoniesClient) UpdateOne(t *Testimonies) *TestimoniesUpdateOne {
	mutation := newTestimoniesMutation(c.config, OpUpdateOne, withTestimonies(t))
	return &TestimoniesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TestimoniesClient) UpdateOneID(id int) *TestimoniesUpdateOne {
	mutation := newTestimoniesMutation(c.config, OpUpdateOne, withTestimoniesID(id))
	return &TestimoniesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Testimonies.
func (c *TestimoniesClient) Delete() *TestimoniesDelete {
	mutation := newTestimoniesMutation(c.config, OpDelete)
	return &TestimoniesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TestimoniesClient) DeleteOne(t *Testimonies) *TestimoniesDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TestimoniesClient) DeleteOneID(id int) *TestimoniesDeleteOne {
	builder := c.Delete().Where(testimonies.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TestimoniesDeleteOne{builder}
}

// Query returns a query builder for Testimonies.
func (c *TestimoniesClient) Query() *TestimoniesQuery {
	return &TestimoniesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTestimonies},
		inters: c.Interceptors(),
	}
}

// Get returns a Testimonies entity by its id.
func (c *TestimoniesClient) Get(ctx context.Context, id int) (*Testimonies, error) {
	return c.Query().Where(testimonies.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TestimoniesClient) GetX(ctx context.Context, id int) *Testimonies {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDestinies queries the destinies edge of a Testimonies.
func (c *TestimoniesClient) QueryDestinies(t *Testimonies) *DestiniesQuery {
	query := (&DestiniesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(testimonies.Table, testimonies.FieldID, id),
			sqlgraph.To(destinies.Table, destinies.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, testimonies.DestiniesTable, testimonies.DestiniesColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TestimoniesClient) Hooks() []Hook {
	hooks := c.hooks.Testimonies
	return append(hooks[:len(hooks):len(hooks)], testimonies.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *TestimoniesClient) Interceptors() []Interceptor {
	return c.inters.Testimonies
}

func (c *TestimoniesClient) mutate(ctx context.Context, m *TestimoniesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TestimoniesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TestimoniesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TestimoniesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TestimoniesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Testimonies mutation op: %q", m.Op())
	}
}

// UsersClient is a client for the Users schema.
type UsersClient struct {
	config
}

// NewUsersClient returns a client for the Users from the given config.
func NewUsersClient(c config) *UsersClient {
	return &UsersClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `users.Hooks(f(g(h())))`.
func (c *UsersClient) Use(hooks ...Hook) {
	c.hooks.Users = append(c.hooks.Users, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `users.Intercept(f(g(h())))`.
func (c *UsersClient) Intercept(interceptors ...Interceptor) {
	c.inters.Users = append(c.inters.Users, interceptors...)
}

// Create returns a builder for creating a Users entity.
func (c *UsersClient) Create() *UsersCreate {
	mutation := newUsersMutation(c.config, OpCreate)
	return &UsersCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Users entities.
func (c *UsersClient) CreateBulk(builders ...*UsersCreate) *UsersCreateBulk {
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UsersClient) MapCreateBulk(slice any, setFunc func(*UsersCreate, int)) *UsersCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UsersCreateBulk{err: fmt.Errorf("calling to UsersClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UsersCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Users.
func (c *UsersClient) Update() *UsersUpdate {
	mutation := newUsersMutation(c.config, OpUpdate)
	return &UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersClient) UpdateOne(u *Users) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsers(u))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersClient) UpdateOneID(id int) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsersID(id))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Users.
func (c *UsersClient) Delete() *UsersDelete {
	mutation := newUsersMutation(c.config, OpDelete)
	return &UsersDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersClient) DeleteOne(u *Users) *UsersDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UsersClient) DeleteOneID(id int) *UsersDeleteOne {
	builder := c.Delete().Where(users.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersDeleteOne{builder}
}

// Query returns a query builder for Users.
func (c *UsersClient) Query() *UsersQuery {
	return &UsersQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUsers},
		inters: c.Interceptors(),
	}
}

// Get returns a Users entity by its id.
func (c *UsersClient) Get(ctx context.Context, id int) (*Users, error) {
	return c.Query().Where(users.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersClient) GetX(ctx context.Context, id int) *Users {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UsersClient) Hooks() []Hook {
	return c.hooks.Users
}

// Interceptors returns the client interceptors.
func (c *UsersClient) Interceptors() []Interceptor {
	return c.inters.Users
}

func (c *UsersClient) mutate(ctx context.Context, m *UsersMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UsersCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UsersDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Users mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Destinies, DestinyPictures, Testimonies, Users []ent.Hook
	}
	inters struct {
		Destinies, DestinyPictures, Testimonies, Users []ent.Interceptor
	}
)
