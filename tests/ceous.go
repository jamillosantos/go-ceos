
package tests

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/pkg/errors"
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
func (model *User) GetID() interface{} {
		return model.ID
	}

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
	ID	ceous.SchemaField
	Name	ceous.SchemaField
	Password	ceous.SchemaField
	Role	ceous.SchemaField
	CreatedAt	ceous.SchemaField
	UpdatedAt	ceous.SchemaField
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
func (model *Group) GetID() interface{} {
		return model.ID
	}

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
	ID	ceous.SchemaField
	Name	ceous.SchemaField
}

/** 
 * Declare UserGroup
 */

var baseSchemaUserGroup = ceous.NewBaseSchema(
	"user_groups",
	"",
	ceous.NewSchemaField("admin"),
)
func (model *UserGroup) GetID() interface{} {
		return model.ID
	}

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
	UserID	ceous.SchemaField
	GroupID	ceous.SchemaField
	Admin	ceous.SchemaField
}



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

type schema struct {
	User *schemaUser
	Group *schemaGroup
	UserGroup *schemaUserGroup
	UserGroupPK *schemaUserGroupPK
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
		Admin: baseSchemaUserGroup.ColumnsArr[0],
	
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



type groupResultSet struct {
	*ceous.RecordResultSet
}

func NewGroupResultSet(rs ceous.ResultSet, err error) *groupResultSet {
	return &groupResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}



type userGroupResultSet struct {
	*ceous.RecordResultSet
}

func NewUserGroupResultSet(rs ceous.ResultSet, err error) *userGroupResultSet {
	return &userGroupResultSet{
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
func (q *userQuery) ByID(value interface{}) *userQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.User.ID, value))
	return q
}

// ByName add a filter by `Name`.
func (q *userQuery) ByName(value interface{}) *userQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.User.Name, value))
	return q
}

// ByPassword add a filter by `Password`.
func (q *userQuery) ByPassword(value interface{}) *userQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.User.Password, value))
	return q
}

// ByRole add a filter by `Role`.
func (q *userQuery) ByRole(value interface{}) *userQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.User.Role, value))
	return q
}

// ByCreatedAt add a filter by `CreatedAt`.
func (q *userQuery) ByCreatedAt(value interface{}) *userQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.User.CreatedAt, value))
	return q
}

// ByUpdatedAt add a filter by `UpdatedAt`.
func (q *userQuery) ByUpdatedAt(value interface{}) *userQuery {
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
	} else {
		err = ceous.ErrNotFound
	}
	return
}

// All return all records that match the query.
func (q *userQuery) All() ([]User, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserResultSet(query, nil)
	defer rs.Close()

	ms := make([]User, 0)
	for rs.Next() {
		var m User
		err = rs.ToModel(&m)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
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
func (q *groupQuery) ByID(value interface{}) *groupQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.Group.ID, value))
	return q
}

// ByName add a filter by `Name`.
func (q *groupQuery) ByName(value interface{}) *groupQuery {
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
	} else {
		err = ceous.ErrNotFound
	}
	return
}

// All return all records that match the query.
func (q *groupQuery) All() ([]Group, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewGroupResultSet(query, nil)
	defer rs.Close()

	ms := make([]Group, 0)
	for rs.Next() {
		var m Group
		err = rs.ToModel(&m)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
}



// userGroupQuery is the query for the model `UserGroup`.
type userGroupQuery struct {
	*ceous.BaseQuery
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

// ByAdmin add a filter by `Admin`.
func (q *userGroupQuery) ByAdmin(value interface{}) *userGroupQuery {
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
	} else {
		err = ceous.ErrNotFound
	}
	return
}

// All return all records that match the query.
func (q *userGroupQuery) All() ([]UserGroup, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	ms := make([]UserGroup, 0)
	for rs.Next() {
		var m UserGroup
		err = rs.ToModel(&m)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
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