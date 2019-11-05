package tests

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/pkg/errors"
	"database/sql"
	time "time"
)

/**
 * Declare User
 */
	

var baseSchemaUser = ceous.NewBaseSchema(
	"users",
	"",
	ceous.NewSchemaField("id", ceous.FieldPK, ceous.FieldAutoIncrement),
	ceous.NewSchemaField("name"),
	ceous.NewSchemaField("password"),
	ceous.NewSchemaField("role"),
	ceous.NewSchemaField("created_at"),
	ceous.NewSchemaField("updated_at"),
)

// GetID returns the primary key for model `User`.
func (model *User) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *User) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &model.ID, nil
	case "name":
		return &model.Name, nil
	case "password":
		return &model.Password, nil
	case "role":
		return &model.Role, nil
	case "created_at":
		return &model.CreatedAt, nil
	case "updated_at":
		return &model.UpdatedAt, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *User) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return model.ID, nil
	case "name":
		return model.Name, nil
	case "password":
		return model.Password, nil
	case "role":
		return model.Role, nil
	case "created_at":
		return model.CreatedAt, nil
	case "updated_at":
		return model.UpdatedAt, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

type schemaUser struct {
	*ceous.BaseSchema
	ID ceous.SchemaField
	Name ceous.SchemaField
	Password ceous.SchemaField
	Role ceous.SchemaField
	CreatedAt ceous.SchemaField
	UpdatedAt ceous.SchemaField
}

/**
 * Declare Group
 */
	

var baseSchemaGroup = ceous.NewBaseSchema(
	"groups",
	"",
	ceous.NewSchemaField("id", ceous.FieldPK, ceous.FieldAutoIncrement),
	ceous.NewSchemaField("name"),
)

// GetID returns the primary key for model `Group`.
func (model *Group) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *Group) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &model.ID, nil
	case "name":
		return &model.Name, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *Group) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return model.ID, nil
	case "name":
		return model.Name, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

type schemaGroup struct {
	*ceous.BaseSchema
	ID ceous.SchemaField
	Name ceous.SchemaField
}

/**
 * Declare UserGroup
 */
	

var baseSchemaUserGroup = ceous.NewBaseSchema(
	"user_groups",
	"",
	ceous.NewSchemaField("user_id", ceous.FieldPK),
	ceous.NewSchemaField("group_id", ceous.FieldPK),
	ceous.NewSchemaField("admin"),
)

// GetID returns the primary key for model `UserGroup`.
func (model *UserGroup) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *UserGroup) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return &model.ID.UserID, nil
	case "group_id":
		return &model.ID.GroupID, nil
	case "admin":
		return &model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *UserGroup) Value(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return model.ID.UserID, nil
	case "group_id":
		return model.ID.GroupID, nil
	case "admin":
		return model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

type schemaUserGroup struct {
	*ceous.BaseSchema
	ID schemaUserGroupPK
	Admin ceous.SchemaField
}



// ColumnAddress returns the pointer address of a field given its column name.
func (model *UserGroupPK) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return &model.UserID, nil
	case "group_id":
		return &model.GroupID, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *UserGroupPK) Value(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return model.UserID, nil
	case "group_id":
		return model.GroupID, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

type schemaUserGroupPK struct {
	UserID  ceous.SchemaField
	GroupID  ceous.SchemaField
}

var UserGroupPKSchema = schemaUserGroupPK {
	UserID: ceous.NewSchemaField("user_id"),
	GroupID: ceous.NewSchemaField("group_id"),
}

type Connection interface {
	ceous.Connection
	// UserQuery creates a new query related with the connection set.
	UserQuery(options ...ceous.CeousOption) *userQuery
	// UserStore creates a new store related with the connection set.
	UserStore(options ...ceous.CeousOption) *userStore
	// GroupQuery creates a new query related with the connection set.
	GroupQuery(options ...ceous.CeousOption) *groupQuery
	// GroupStore creates a new store related with the connection set.
	GroupStore(options ...ceous.CeousOption) *groupStore
	// UserGroupQuery creates a new query related with the connection set.
	UserGroupQuery(options ...ceous.CeousOption) *userGroupQuery
	// UserGroupStore creates a new store related with the connection set.
	UserGroupStore(options ...ceous.CeousOption) *userGroupStore
}
type DefaultConnection  struct {
	*ceous.BaseConnection
}

	
// UserQuery creates a new query related with the connection Default set.
func (c *DefaultConnection) UserQuery(options ...ceous.CeousOption) *userQuery {
	return NewUserQuery(append(options, ceous.WithConn(c))...)
}

// UserStore creates a new store related with the connection Default set.
func (c *DefaultConnection) UserStore(options ...ceous.CeousOption) *userStore {
	return NewUserStore(append(options, ceous.WithConn(c))...)
}
	
// GroupQuery creates a new query related with the connection Default set.
func (c *DefaultConnection) GroupQuery(options ...ceous.CeousOption) *groupQuery {
	return NewGroupQuery(append(options, ceous.WithConn(c))...)
}

// GroupStore creates a new store related with the connection Default set.
func (c *DefaultConnection) GroupStore(options ...ceous.CeousOption) *groupStore {
	return NewGroupStore(append(options, ceous.WithConn(c))...)
}
	
// UserGroupQuery creates a new query related with the connection Default set.
func (c *DefaultConnection) UserGroupQuery(options ...ceous.CeousOption) *userGroupQuery {
	return NewUserGroupQuery(append(options, ceous.WithConn(c))...)
}

// UserGroupStore creates a new store related with the connection Default set.
func (c *DefaultConnection) UserGroupStore(options ...ceous.CeousOption) *userGroupStore {
	return NewUserGroupStore(append(options, ceous.WithConn(c))...)
}
	

var (
	// Default is a database connection reference.
	Default Connection
)
// InitDefault initializes the connection `Default:`.
func InitDefault(db *sql.DB) {
	Default = &DefaultConnection{
		BaseConnection: ceous.NewConnection(db),
	}
}

type schema struct {
	User *schemaUser
	Group *schemaGroup
	UserGroup *schemaUserGroup
}

// Schema represents the schema of the package "tests".
var Schema = schema{
	User: &schemaUser {
	BaseSchema: baseSchemaUser,

	
		ID: baseSchemaUser.ColumnsArr[0],
	
		Name: baseSchemaUser.ColumnsArr[1],
	
		Password: baseSchemaUser.ColumnsArr[2],
	
		Role: baseSchemaUser.ColumnsArr[3],
	
		CreatedAt: baseSchemaUser.ColumnsArr[4],
	
		UpdatedAt: baseSchemaUser.ColumnsArr[5],
	
	},
	Group: &schemaGroup {
	BaseSchema: baseSchemaGroup,

	
		ID: baseSchemaGroup.ColumnsArr[0],
	
		Name: baseSchemaGroup.ColumnsArr[1],
	
	},
	UserGroup: &schemaUserGroup {
	BaseSchema: baseSchemaUserGroup,

	
		ID: UserGroupPKSchema,
	
		Admin: baseSchemaUserGroup.ColumnsArr[2],
	
	},
}
	

type userResultSet struct {
	*ceous.RecordResultSet
}

func NewUserResultSet(rs ceous.ResultSet, err error) *userResultSet {
	return &userResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}
	

// userQuery is the query for the model `User`.
type userQuery struct {
	*ceous.BaseQuery
	ID	ceous.SchemaField
	Name	ceous.SchemaField
	Password	ceous.SchemaField
	Role	ceous.SchemaField
	CreatedAt	ceous.SchemaField
	UpdatedAt	ceous.SchemaField
}

// NewUserQuery creates a new query for model `User`.
func NewUserQuery(options ...ceous.CeousOption) *userQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = Schema.User.BaseSchema
	}
	return &userQuery{
		BaseQuery: bq,
	}
}
// ByID add a filter by `ID`.
func (q *userQuery) ByID(value int) *userQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.User.ID, value))
	
	return q
}
// ByName add a filter by `Name`.
func (q *userQuery) ByName(value string) *userQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.User.Name, value))
	
	return q
}
// ByPassword add a filter by `Password`.
func (q *userQuery) ByPassword(value string) *userQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.User.Password, value))
	
	return q
}
// ByRole add a filter by `Role`.
func (q *userQuery) ByRole(value string) *userQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.User.Role, value))
	
	return q
}
// ByCreatedAt add a filter by `CreatedAt`.
func (q *userQuery) ByCreatedAt(value time.Time) *userQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.User.CreatedAt, value))
	
	return q
}
// ByUpdatedAt add a filter by `UpdatedAt`.
func (q *userQuery) ByUpdatedAt(value time.Time) *userQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.User.UpdatedAt, value))
	
	return q
}

// Select defines what fields should be selected from the database.
func (q *userQuery) Select(fields ...ceous.SchemaField) *userQuery {
	q.BaseQuery.Select(fields...)
	return q
}

// ExcludeFields defines what fields should not be selected from the database.
func (q *userQuery) ExcludeFields(fields ...ceous.SchemaField) *userQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

// Where defines the conditions for 
func (q *userQuery) Where(pred interface{}, args ...interface{}) *userQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *userQuery) Limit(limit uint64) *userQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *userQuery) Offset(offset uint64) *userQuery {
	q.BaseQuery.Offset(offset)
	return q
}

// One results only one record matching the query.
func (q *userQuery) One() (m User, err error) {
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return
	}

	rs := NewUserResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&m)
		if err != nil {
			return
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&m)
			if err != nil {
				return User{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}

	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return User{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}

	return
}

// All return all records that match the query.
func (q *userQuery) All() ([]*User, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserResultSet(query, nil)
	defer rs.Close()

	ms := make([]*User, 0)
	for rs.Next() {
		m := &User{}
		err = rs.ToModel(m)
		if err != nil {
			return nil, err
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(m)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
		ms = append(ms, m)
	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
	}

	return ms, nil
}

func (q *userQuery) OrderBy(fields ...interface{}) *userQuery {
	q.BaseQuery.OrderBy(fields...)
	return q
}
	

// userStore is the query for the model `User`.
type userStore struct {
	*ceous.BaseStore
}

// NewUserStore creates a new query for model `User`.
func NewUserStore(options ...ceous.CeousOption) *userStore {
	return &userStore{
		BaseStore: ceous.NewStore(baseSchemaUser, options...),
	}
}

func (store *userStore) Insert(record *User, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}

func (store *userStore) Update(record *User, fields ...ceous.SchemaField) (int64, error) {
	return store.BaseStore.Update(record, fields...)
}
	

type groupResultSet struct {
	*ceous.RecordResultSet
}

func NewGroupResultSet(rs ceous.ResultSet, err error) *groupResultSet {
	return &groupResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}
	

// groupQuery is the query for the model `Group`.
type groupQuery struct {
	*ceous.BaseQuery
	ID	ceous.SchemaField
	Name	ceous.SchemaField
}

// NewGroupQuery creates a new query for model `Group`.
func NewGroupQuery(options ...ceous.CeousOption) *groupQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = Schema.Group.BaseSchema
	}
	return &groupQuery{
		BaseQuery: bq,
	}
}
// ByID add a filter by `ID`.
func (q *groupQuery) ByID(value int) *groupQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.Group.ID, value))
	
	return q
}
// ByName add a filter by `Name`.
func (q *groupQuery) ByName(value string) *groupQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.Group.Name, value))
	
	return q
}

// Select defines what fields should be selected from the database.
func (q *groupQuery) Select(fields ...ceous.SchemaField) *groupQuery {
	q.BaseQuery.Select(fields...)
	return q
}

// ExcludeFields defines what fields should not be selected from the database.
func (q *groupQuery) ExcludeFields(fields ...ceous.SchemaField) *groupQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

// Where defines the conditions for 
func (q *groupQuery) Where(pred interface{}, args ...interface{}) *groupQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *groupQuery) Limit(limit uint64) *groupQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *groupQuery) Offset(offset uint64) *groupQuery {
	q.BaseQuery.Offset(offset)
	return q
}

// One results only one record matching the query.
func (q *groupQuery) One() (m Group, err error) {
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return
	}

	rs := NewGroupResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&m)
		if err != nil {
			return
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&m)
			if err != nil {
				return Group{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}

	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return Group{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}

	return
}

// All return all records that match the query.
func (q *groupQuery) All() ([]*Group, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewGroupResultSet(query, nil)
	defer rs.Close()

	ms := make([]*Group, 0)
	for rs.Next() {
		m := &Group{}
		err = rs.ToModel(m)
		if err != nil {
			return nil, err
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(m)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
		ms = append(ms, m)
	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
	}

	return ms, nil
}

func (q *groupQuery) OrderBy(fields ...interface{}) *groupQuery {
	q.BaseQuery.OrderBy(fields...)
	return q
}
	

// groupStore is the query for the model `Group`.
type groupStore struct {
	*ceous.BaseStore
}

// NewGroupStore creates a new query for model `Group`.
func NewGroupStore(options ...ceous.CeousOption) *groupStore {
	return &groupStore{
		BaseStore: ceous.NewStore(baseSchemaGroup, options...),
	}
}

func (store *groupStore) Insert(record *Group, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}

func (store *groupStore) Update(record *Group, fields ...ceous.SchemaField) (int64, error) {
	return store.BaseStore.Update(record, fields...)
}
	

type userGroupResultSet struct {
	*ceous.RecordResultSet
}

func NewUserGroupResultSet(rs ceous.ResultSet, err error) *userGroupResultSet {
	return &userGroupResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}
	

// userGroupQuery is the query for the model `UserGroup`.
type userGroupQuery struct {
	*ceous.BaseQuery
	ID	ceous.SchemaField
	Admin	ceous.SchemaField
}

// NewUserGroupQuery creates a new query for model `UserGroup`.
func NewUserGroupQuery(options ...ceous.CeousOption) *userGroupQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = Schema.UserGroup.BaseSchema
	}
	return &userGroupQuery{
		BaseQuery: bq,
	}
}
// ByID add a filter by `ID`.
func (q *userGroupQuery) ByID(value UserGroupPK) *userGroupQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.UserGroup.ID.UserID, value.UserID))
	
	q.BaseQuery.Where(ceous.Eq(Schema.UserGroup.ID.GroupID, value.GroupID))
	
	return q
}
// ByAdmin add a filter by `Admin`.
func (q *userGroupQuery) ByAdmin(value bool) *userGroupQuery {
	
	q.BaseQuery.Where(ceous.Eq(Schema.UserGroup.Admin, value))
	
	return q
}

// Select defines what fields should be selected from the database.
func (q *userGroupQuery) Select(fields ...ceous.SchemaField) *userGroupQuery {
	q.BaseQuery.Select(fields...)
	return q
}

// ExcludeFields defines what fields should not be selected from the database.
func (q *userGroupQuery) ExcludeFields(fields ...ceous.SchemaField) *userGroupQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

// Where defines the conditions for 
func (q *userGroupQuery) Where(pred interface{}, args ...interface{}) *userGroupQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *userGroupQuery) Limit(limit uint64) *userGroupQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *userGroupQuery) Offset(offset uint64) *userGroupQuery {
	q.BaseQuery.Offset(offset)
	return q
}

// One results only one record matching the query.
func (q *userGroupQuery) One() (m UserGroup, err error) {
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&m)
		if err != nil {
			return
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&m)
			if err != nil {
				return UserGroup{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}

	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return UserGroup{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}

	return
}

// All return all records that match the query.
func (q *userGroupQuery) All() ([]*UserGroup, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	ms := make([]*UserGroup, 0)
	for rs.Next() {
		m := &UserGroup{}
		err = rs.ToModel(m)
		if err != nil {
			return nil, err
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(m)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
		ms = append(ms, m)
	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
	}

	return ms, nil
}

func (q *userGroupQuery) OrderBy(fields ...interface{}) *userGroupQuery {
	q.BaseQuery.OrderBy(fields...)
	return q
}

type UserGroupModelUserRelation struct {
	conn ceous.Connection
	keys []interface{}
	records map[int][]*UserGroup
}

func NewUserGroupModelUserRelation(conn ceous.Connection) *UserGroupModelUserRelation {
	return &UserGroupModelUserRelation{
		conn:    conn,
		keys:    make([]interface{}, 0),
		records: make(map[int][]*UserGroup),
	}
}

func (relation *UserGroupModelUserRelation) Aggregate(record ceous.Record) error {
	ugRecord, ok := record.(*UserGroup)
	if !ok {
		return ceous.ErrInvalidRecordType
	}
	if rs, ok := relation.records[ugRecord.ID.UserID]; ok {
		relation.records[ugRecord.ID.UserID] = append(rs, ugRecord)
		// No need to add the key here, since its will be already in the `keys`.
	} else {
		relation.records[ugRecord.ID.UserID] = append(rs, ugRecord)
		relation.keys = append(relation.keys, ugRecord.ID.UserID)
	}
	return nil
}

func (relation *UserGroupModelUserRelation) Realize() error {
	records, err := NewUserQuery(ceous.WithConn(relation.conn)).Where(ceous.Eq(Schema.User.ID, relation.keys)).All()
	if err != nil {
		return err // TODO(jota): Shall this be wrapped into a custom error?
	}
	for _, record := range records {
		masterRecords, ok := relation.records[record.ID]
		if !ok {
			return ceous.ErrInconsistentRelationResult
		}
		for _, r := range masterRecords {
			r.User = record
		}
	}
	return nil
}

func (q *userGroupQuery) WithUser() *userGroupQuery {
	q.BaseQuery.Relations = append(q.BaseQuery.Relations, NewUserGroupModelUserRelation(q.BaseQuery.Connection))
	return q
}
	

// userGroupStore is the query for the model `UserGroup`.
type userGroupStore struct {
	*ceous.BaseStore
}

// NewUserGroupStore creates a new query for model `UserGroup`.
func NewUserGroupStore(options ...ceous.CeousOption) *userGroupStore {
	return &userGroupStore{
		BaseStore: ceous.NewStore(baseSchemaUserGroup, options...),
	}
}

func (store *userGroupStore) Insert(record *UserGroup, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}

func (store *userGroupStore) Update(record *UserGroup, fields ...ceous.SchemaField) (int64, error) {
	return store.BaseStore.Update(record, fields...)
}