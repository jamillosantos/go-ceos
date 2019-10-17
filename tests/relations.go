package tests

import "github.com/jamillosantos/go-ceous"

type (
	UserGroupModelUserRelation struct {
		keys    []interface{}
		records map[int][]*UserGroup
	}
)

func NewUserGroupModelUserRelation() *UserGroupModelUserRelation {
	return &UserGroupModelUserRelation{
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
	records, err := NewUserQuery(ceous.WithDB(DB)).Where(ceous.Eq(Schema.User.ID, relation.keys)).All()
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
