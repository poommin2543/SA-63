// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/poommin2543/app/ent/migrate"

	"github.com/poommin2543/app/ent/medicalequipment"
	"github.com/poommin2543/app/ent/medicaltype"
	"github.com/poommin2543/app/ent/physician"
	"github.com/poommin2543/app/ent/systemequipment"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// MedicalEquipment is the client for interacting with the MedicalEquipment builders.
	MedicalEquipment *MedicalEquipmentClient
	// MedicalType is the client for interacting with the MedicalType builders.
	MedicalType *MedicalTypeClient
	// Physician is the client for interacting with the Physician builders.
	Physician *PhysicianClient
	// Systemequipment is the client for interacting with the Systemequipment builders.
	Systemequipment *SystemequipmentClient
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
	c.MedicalEquipment = NewMedicalEquipmentClient(c.config)
	c.MedicalType = NewMedicalTypeClient(c.config)
	c.Physician = NewPhysicianClient(c.config)
	c.Systemequipment = NewSystemequipmentClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		MedicalEquipment: NewMedicalEquipmentClient(cfg),
		MedicalType:      NewMedicalTypeClient(cfg),
		Physician:        NewPhysicianClient(cfg),
		Systemequipment:  NewSystemequipmentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:           cfg,
		MedicalEquipment: NewMedicalEquipmentClient(cfg),
		MedicalType:      NewMedicalTypeClient(cfg),
		Physician:        NewPhysicianClient(cfg),
		Systemequipment:  NewSystemequipmentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		MedicalEquipment.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.MedicalEquipment.Use(hooks...)
	c.MedicalType.Use(hooks...)
	c.Physician.Use(hooks...)
	c.Systemequipment.Use(hooks...)
}

// MedicalEquipmentClient is a client for the MedicalEquipment schema.
type MedicalEquipmentClient struct {
	config
}

// NewMedicalEquipmentClient returns a client for the MedicalEquipment from the given config.
func NewMedicalEquipmentClient(c config) *MedicalEquipmentClient {
	return &MedicalEquipmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `medicalequipment.Hooks(f(g(h())))`.
func (c *MedicalEquipmentClient) Use(hooks ...Hook) {
	c.hooks.MedicalEquipment = append(c.hooks.MedicalEquipment, hooks...)
}

// Create returns a create builder for MedicalEquipment.
func (c *MedicalEquipmentClient) Create() *MedicalEquipmentCreate {
	mutation := newMedicalEquipmentMutation(c.config, OpCreate)
	return &MedicalEquipmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for MedicalEquipment.
func (c *MedicalEquipmentClient) Update() *MedicalEquipmentUpdate {
	mutation := newMedicalEquipmentMutation(c.config, OpUpdate)
	return &MedicalEquipmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MedicalEquipmentClient) UpdateOne(me *MedicalEquipment) *MedicalEquipmentUpdateOne {
	mutation := newMedicalEquipmentMutation(c.config, OpUpdateOne, withMedicalEquipment(me))
	return &MedicalEquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MedicalEquipmentClient) UpdateOneID(id int) *MedicalEquipmentUpdateOne {
	mutation := newMedicalEquipmentMutation(c.config, OpUpdateOne, withMedicalEquipmentID(id))
	return &MedicalEquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MedicalEquipment.
func (c *MedicalEquipmentClient) Delete() *MedicalEquipmentDelete {
	mutation := newMedicalEquipmentMutation(c.config, OpDelete)
	return &MedicalEquipmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MedicalEquipmentClient) DeleteOne(me *MedicalEquipment) *MedicalEquipmentDeleteOne {
	return c.DeleteOneID(me.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MedicalEquipmentClient) DeleteOneID(id int) *MedicalEquipmentDeleteOne {
	builder := c.Delete().Where(medicalequipment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MedicalEquipmentDeleteOne{builder}
}

// Create returns a query builder for MedicalEquipment.
func (c *MedicalEquipmentClient) Query() *MedicalEquipmentQuery {
	return &MedicalEquipmentQuery{config: c.config}
}

// Get returns a MedicalEquipment entity by its id.
func (c *MedicalEquipmentClient) Get(ctx context.Context, id int) (*MedicalEquipment, error) {
	return c.Query().Where(medicalequipment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MedicalEquipmentClient) GetX(ctx context.Context, id int) *MedicalEquipment {
	me, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return me
}

// QuerySystemequipment queries the systemequipment edge of a MedicalEquipment.
func (c *MedicalEquipmentClient) QuerySystemequipment(me *MedicalEquipment) *SystemequipmentQuery {
	query := &SystemequipmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := me.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalequipment.Table, medicalequipment.FieldID, id),
			sqlgraph.To(systemequipment.Table, systemequipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, medicalequipment.SystemequipmentTable, medicalequipment.SystemequipmentColumn),
		)
		fromV = sqlgraph.Neighbors(me.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MedicalEquipmentClient) Hooks() []Hook {
	return c.hooks.MedicalEquipment
}

// MedicalTypeClient is a client for the MedicalType schema.
type MedicalTypeClient struct {
	config
}

// NewMedicalTypeClient returns a client for the MedicalType from the given config.
func NewMedicalTypeClient(c config) *MedicalTypeClient {
	return &MedicalTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `medicaltype.Hooks(f(g(h())))`.
func (c *MedicalTypeClient) Use(hooks ...Hook) {
	c.hooks.MedicalType = append(c.hooks.MedicalType, hooks...)
}

// Create returns a create builder for MedicalType.
func (c *MedicalTypeClient) Create() *MedicalTypeCreate {
	mutation := newMedicalTypeMutation(c.config, OpCreate)
	return &MedicalTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for MedicalType.
func (c *MedicalTypeClient) Update() *MedicalTypeUpdate {
	mutation := newMedicalTypeMutation(c.config, OpUpdate)
	return &MedicalTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MedicalTypeClient) UpdateOne(mt *MedicalType) *MedicalTypeUpdateOne {
	mutation := newMedicalTypeMutation(c.config, OpUpdateOne, withMedicalType(mt))
	return &MedicalTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MedicalTypeClient) UpdateOneID(id int) *MedicalTypeUpdateOne {
	mutation := newMedicalTypeMutation(c.config, OpUpdateOne, withMedicalTypeID(id))
	return &MedicalTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MedicalType.
func (c *MedicalTypeClient) Delete() *MedicalTypeDelete {
	mutation := newMedicalTypeMutation(c.config, OpDelete)
	return &MedicalTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MedicalTypeClient) DeleteOne(mt *MedicalType) *MedicalTypeDeleteOne {
	return c.DeleteOneID(mt.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MedicalTypeClient) DeleteOneID(id int) *MedicalTypeDeleteOne {
	builder := c.Delete().Where(medicaltype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MedicalTypeDeleteOne{builder}
}

// Create returns a query builder for MedicalType.
func (c *MedicalTypeClient) Query() *MedicalTypeQuery {
	return &MedicalTypeQuery{config: c.config}
}

// Get returns a MedicalType entity by its id.
func (c *MedicalTypeClient) Get(ctx context.Context, id int) (*MedicalType, error) {
	return c.Query().Where(medicaltype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MedicalTypeClient) GetX(ctx context.Context, id int) *MedicalType {
	mt, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return mt
}

// QuerySystemequipment queries the systemequipment edge of a MedicalType.
func (c *MedicalTypeClient) QuerySystemequipment(mt *MedicalType) *SystemequipmentQuery {
	query := &SystemequipmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(medicaltype.Table, medicaltype.FieldID, id),
			sqlgraph.To(systemequipment.Table, systemequipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, medicaltype.SystemequipmentTable, medicaltype.SystemequipmentColumn),
		)
		fromV = sqlgraph.Neighbors(mt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MedicalTypeClient) Hooks() []Hook {
	return c.hooks.MedicalType
}

// PhysicianClient is a client for the Physician schema.
type PhysicianClient struct {
	config
}

// NewPhysicianClient returns a client for the Physician from the given config.
func NewPhysicianClient(c config) *PhysicianClient {
	return &PhysicianClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `physician.Hooks(f(g(h())))`.
func (c *PhysicianClient) Use(hooks ...Hook) {
	c.hooks.Physician = append(c.hooks.Physician, hooks...)
}

// Create returns a create builder for Physician.
func (c *PhysicianClient) Create() *PhysicianCreate {
	mutation := newPhysicianMutation(c.config, OpCreate)
	return &PhysicianCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Physician.
func (c *PhysicianClient) Update() *PhysicianUpdate {
	mutation := newPhysicianMutation(c.config, OpUpdate)
	return &PhysicianUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PhysicianClient) UpdateOne(ph *Physician) *PhysicianUpdateOne {
	mutation := newPhysicianMutation(c.config, OpUpdateOne, withPhysician(ph))
	return &PhysicianUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PhysicianClient) UpdateOneID(id int) *PhysicianUpdateOne {
	mutation := newPhysicianMutation(c.config, OpUpdateOne, withPhysicianID(id))
	return &PhysicianUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Physician.
func (c *PhysicianClient) Delete() *PhysicianDelete {
	mutation := newPhysicianMutation(c.config, OpDelete)
	return &PhysicianDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PhysicianClient) DeleteOne(ph *Physician) *PhysicianDeleteOne {
	return c.DeleteOneID(ph.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PhysicianClient) DeleteOneID(id int) *PhysicianDeleteOne {
	builder := c.Delete().Where(physician.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PhysicianDeleteOne{builder}
}

// Create returns a query builder for Physician.
func (c *PhysicianClient) Query() *PhysicianQuery {
	return &PhysicianQuery{config: c.config}
}

// Get returns a Physician entity by its id.
func (c *PhysicianClient) Get(ctx context.Context, id int) (*Physician, error) {
	return c.Query().Where(physician.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PhysicianClient) GetX(ctx context.Context, id int) *Physician {
	ph, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ph
}

// QuerySystemequipment queries the systemequipment edge of a Physician.
func (c *PhysicianClient) QuerySystemequipment(ph *Physician) *SystemequipmentQuery {
	query := &SystemequipmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ph.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(physician.Table, physician.FieldID, id),
			sqlgraph.To(systemequipment.Table, systemequipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, physician.SystemequipmentTable, physician.SystemequipmentColumn),
		)
		fromV = sqlgraph.Neighbors(ph.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PhysicianClient) Hooks() []Hook {
	return c.hooks.Physician
}

// SystemequipmentClient is a client for the Systemequipment schema.
type SystemequipmentClient struct {
	config
}

// NewSystemequipmentClient returns a client for the Systemequipment from the given config.
func NewSystemequipmentClient(c config) *SystemequipmentClient {
	return &SystemequipmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `systemequipment.Hooks(f(g(h())))`.
func (c *SystemequipmentClient) Use(hooks ...Hook) {
	c.hooks.Systemequipment = append(c.hooks.Systemequipment, hooks...)
}

// Create returns a create builder for Systemequipment.
func (c *SystemequipmentClient) Create() *SystemequipmentCreate {
	mutation := newSystemequipmentMutation(c.config, OpCreate)
	return &SystemequipmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Systemequipment.
func (c *SystemequipmentClient) Update() *SystemequipmentUpdate {
	mutation := newSystemequipmentMutation(c.config, OpUpdate)
	return &SystemequipmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SystemequipmentClient) UpdateOne(s *Systemequipment) *SystemequipmentUpdateOne {
	mutation := newSystemequipmentMutation(c.config, OpUpdateOne, withSystemequipment(s))
	return &SystemequipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SystemequipmentClient) UpdateOneID(id int) *SystemequipmentUpdateOne {
	mutation := newSystemequipmentMutation(c.config, OpUpdateOne, withSystemequipmentID(id))
	return &SystemequipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Systemequipment.
func (c *SystemequipmentClient) Delete() *SystemequipmentDelete {
	mutation := newSystemequipmentMutation(c.config, OpDelete)
	return &SystemequipmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SystemequipmentClient) DeleteOne(s *Systemequipment) *SystemequipmentDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SystemequipmentClient) DeleteOneID(id int) *SystemequipmentDeleteOne {
	builder := c.Delete().Where(systemequipment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SystemequipmentDeleteOne{builder}
}

// Create returns a query builder for Systemequipment.
func (c *SystemequipmentClient) Query() *SystemequipmentQuery {
	return &SystemequipmentQuery{config: c.config}
}

// Get returns a Systemequipment entity by its id.
func (c *SystemequipmentClient) Get(ctx context.Context, id int) (*Systemequipment, error) {
	return c.Query().Where(systemequipment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SystemequipmentClient) GetX(ctx context.Context, id int) *Systemequipment {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryPhysician queries the physician edge of a Systemequipment.
func (c *SystemequipmentClient) QueryPhysician(s *Systemequipment) *PhysicianQuery {
	query := &PhysicianQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(systemequipment.Table, systemequipment.FieldID, id),
			sqlgraph.To(physician.Table, physician.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, systemequipment.PhysicianTable, systemequipment.PhysicianColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMedicaltype queries the medicaltype edge of a Systemequipment.
func (c *SystemequipmentClient) QueryMedicaltype(s *Systemequipment) *MedicalTypeQuery {
	query := &MedicalTypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(systemequipment.Table, systemequipment.FieldID, id),
			sqlgraph.To(medicaltype.Table, medicaltype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, systemequipment.MedicaltypeTable, systemequipment.MedicaltypeColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMedicalequipment queries the medicalequipment edge of a Systemequipment.
func (c *SystemequipmentClient) QueryMedicalequipment(s *Systemequipment) *MedicalEquipmentQuery {
	query := &MedicalEquipmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(systemequipment.Table, systemequipment.FieldID, id),
			sqlgraph.To(medicalequipment.Table, medicalequipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, systemequipment.MedicalequipmentTable, systemequipment.MedicalequipmentColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SystemequipmentClient) Hooks() []Hook {
	return c.hooks.Systemequipment
}
