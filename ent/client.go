// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"bookstore/ent/migrate"

	"bookstore/ent/book"
	"bookstore/ent/category"
	"bookstore/ent/feature"
	"bookstore/ent/role"
	"bookstore/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Book is the client for interacting with the Book builders.
	Book *BookClient
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Feature is the client for interacting with the Feature builders.
	Feature *FeatureClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Book = NewBookClient(c.config)
	c.Category = NewCategoryClient(c.config)
	c.Feature = NewFeatureClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.User = NewUserClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Book:     NewBookClient(cfg),
		Category: NewCategoryClient(cfg),
		Feature:  NewFeatureClient(cfg),
		Role:     NewRoleClient(cfg),
		User:     NewUserClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		Book:     NewBookClient(cfg),
		Category: NewCategoryClient(cfg),
		Feature:  NewFeatureClient(cfg),
		Role:     NewRoleClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Book.
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
	c.Book.Use(hooks...)
	c.Category.Use(hooks...)
	c.Feature.Use(hooks...)
	c.Role.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Book.Intercept(interceptors...)
	c.Category.Intercept(interceptors...)
	c.Feature.Intercept(interceptors...)
	c.Role.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BookMutation:
		return c.Book.mutate(ctx, m)
	case *CategoryMutation:
		return c.Category.mutate(ctx, m)
	case *FeatureMutation:
		return c.Feature.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BookClient is a client for the Book schema.
type BookClient struct {
	config
}

// NewBookClient returns a client for the Book from the given config.
func NewBookClient(c config) *BookClient {
	return &BookClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `book.Hooks(f(g(h())))`.
func (c *BookClient) Use(hooks ...Hook) {
	c.hooks.Book = append(c.hooks.Book, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `book.Intercept(f(g(h())))`.
func (c *BookClient) Intercept(interceptors ...Interceptor) {
	c.inters.Book = append(c.inters.Book, interceptors...)
}

// Create returns a builder for creating a Book entity.
func (c *BookClient) Create() *BookCreate {
	mutation := newBookMutation(c.config, OpCreate)
	return &BookCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Book entities.
func (c *BookClient) CreateBulk(builders ...*BookCreate) *BookCreateBulk {
	return &BookCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Book.
func (c *BookClient) Update() *BookUpdate {
	mutation := newBookMutation(c.config, OpUpdate)
	return &BookUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookClient) UpdateOne(b *Book) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBook(b))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookClient) UpdateOneID(id int) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBookID(id))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Book.
func (c *BookClient) Delete() *BookDelete {
	mutation := newBookMutation(c.config, OpDelete)
	return &BookDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BookClient) DeleteOne(b *Book) *BookDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BookClient) DeleteOneID(id int) *BookDeleteOne {
	builder := c.Delete().Where(book.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookDeleteOne{builder}
}

// Query returns a query builder for Book.
func (c *BookClient) Query() *BookQuery {
	return &BookQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBook},
		inters: c.Interceptors(),
	}
}

// Get returns a Book entity by its id.
func (c *BookClient) Get(ctx context.Context, id int) (*Book, error) {
	return c.Query().Where(book.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookClient) GetX(ctx context.Context, id int) *Book {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Book.
func (c *BookClient) QueryUser(b *Book) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.UserTable, book.UserColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCategory queries the category edge of a Book.
func (c *BookClient) QueryCategory(b *Book) *CategoryQuery {
	query := (&CategoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, book.CategoryTable, book.CategoryPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookClient) Hooks() []Hook {
	return c.hooks.Book
}

// Interceptors returns the client interceptors.
func (c *BookClient) Interceptors() []Interceptor {
	return c.inters.Book
}

func (c *BookClient) mutate(ctx context.Context, m *BookMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BookCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BookUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BookDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Book mutation op: %q", m.Op())
	}
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `category.Intercept(f(g(h())))`.
func (c *CategoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Category = append(c.inters.Category, interceptors...)
}

// Create returns a builder for creating a Category entity.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id int) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CategoryClient) DeleteOneID(id int) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCategory},
		inters: c.Interceptors(),
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id int) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id int) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBook queries the book edge of a Category.
func (c *CategoryClient) QueryBook(ca *Category) *BookQuery {
	query := (&BookClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, category.BookTable, category.BookPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// Interceptors returns the client interceptors.
func (c *CategoryClient) Interceptors() []Interceptor {
	return c.inters.Category
}

func (c *CategoryClient) mutate(ctx context.Context, m *CategoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Category mutation op: %q", m.Op())
	}
}

// FeatureClient is a client for the Feature schema.
type FeatureClient struct {
	config
}

// NewFeatureClient returns a client for the Feature from the given config.
func NewFeatureClient(c config) *FeatureClient {
	return &FeatureClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `feature.Hooks(f(g(h())))`.
func (c *FeatureClient) Use(hooks ...Hook) {
	c.hooks.Feature = append(c.hooks.Feature, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `feature.Intercept(f(g(h())))`.
func (c *FeatureClient) Intercept(interceptors ...Interceptor) {
	c.inters.Feature = append(c.inters.Feature, interceptors...)
}

// Create returns a builder for creating a Feature entity.
func (c *FeatureClient) Create() *FeatureCreate {
	mutation := newFeatureMutation(c.config, OpCreate)
	return &FeatureCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Feature entities.
func (c *FeatureClient) CreateBulk(builders ...*FeatureCreate) *FeatureCreateBulk {
	return &FeatureCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Feature.
func (c *FeatureClient) Update() *FeatureUpdate {
	mutation := newFeatureMutation(c.config, OpUpdate)
	return &FeatureUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FeatureClient) UpdateOne(f *Feature) *FeatureUpdateOne {
	mutation := newFeatureMutation(c.config, OpUpdateOne, withFeature(f))
	return &FeatureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FeatureClient) UpdateOneID(id int) *FeatureUpdateOne {
	mutation := newFeatureMutation(c.config, OpUpdateOne, withFeatureID(id))
	return &FeatureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Feature.
func (c *FeatureClient) Delete() *FeatureDelete {
	mutation := newFeatureMutation(c.config, OpDelete)
	return &FeatureDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FeatureClient) DeleteOne(f *Feature) *FeatureDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FeatureClient) DeleteOneID(id int) *FeatureDeleteOne {
	builder := c.Delete().Where(feature.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FeatureDeleteOne{builder}
}

// Query returns a query builder for Feature.
func (c *FeatureClient) Query() *FeatureQuery {
	return &FeatureQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFeature},
		inters: c.Interceptors(),
	}
}

// Get returns a Feature entity by its id.
func (c *FeatureClient) Get(ctx context.Context, id int) (*Feature, error) {
	return c.Query().Where(feature.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FeatureClient) GetX(ctx context.Context, id int) *Feature {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRole queries the role edge of a Feature.
func (c *FeatureClient) QueryRole(f *Feature) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(feature.Table, feature.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, feature.RoleTable, feature.RolePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FeatureClient) Hooks() []Hook {
	return c.hooks.Feature
}

// Interceptors returns the client interceptors.
func (c *FeatureClient) Interceptors() []Interceptor {
	return c.inters.Feature
}

func (c *FeatureClient) mutate(ctx context.Context, m *FeatureMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FeatureCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FeatureUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FeatureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FeatureDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Feature mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Role.
func (c *RoleClient) QueryUser(r *Role) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.UserTable, role.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFeature queries the feature edge of a Role.
func (c *RoleClient) QueryFeature(r *Role) *FeatureQuery {
	query := (&FeatureClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(feature.Table, feature.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.FeatureTable, role.FeaturePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCreatedBy queries the createdBy edge of a User.
func (c *UserClient) QueryCreatedBy(u *User) *BookQuery {
	query := (&BookClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CreatedByTable, user.CreatedByColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRole queries the role edge of a User.
func (c *UserClient) QueryRole(u *User) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.RoleTable, user.RolePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Book, Category, Feature, Role, User []ent.Hook
	}
	inters struct {
		Book, Category, Feature, Role, User []ent.Interceptor
	}
)
