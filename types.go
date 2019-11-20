package ceous

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/oklog/ulid"
)

// ULID is an ID type provided by kallax that is a lexically sortable UUID.
// The internal representation is an ULID (https://github.com/oklog/ulid).
// It already implements sql.Scanner and driver.Valuer, so it's perfectly
// safe for database usage.
//
// This ULID implementation was copied from the go-kallax library (https://github.com/src-d/go-kallax/blob/c3e1e4d85f44dd6ed4f0b65c7fed60c0c576ba85/model.go).
type ULID uuid.UUID

// NewULID returns a new ULID, which is a lexically sortable UUID.
func NewULID() ULID {
	return ULID(ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader))
}

// NewULIDFromText creates a new ULID from its string representation. Will
// return an error if the text is not a valid ULID.
func NewULIDFromText(text string) (ULID, error) {
	var id ULID
	err := id.UnmarshalText([]byte(text))
	return id, err
}

// Scan implements the Scanner interface.
func (id *ULID) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		if len(src) != 16 {
			return id.UnmarshalText(src)
		}

		var ulid ulid.ULID
		if err := ulid.UnmarshalBinary(src); err != nil {
			return err
		}
		*id = ULID(ulid)
		return nil
	case string:
		return id.Scan([]byte(src))
	default:
		return fmt.Errorf("kallax: cannot scan %T into ULID", src)
	}
}

var (
	urnPrefix  = []byte("urn:uuid:")
	byteGroups = []int{8, 4, 4, 4, 12}
)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Following formats are supported:
// "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
// "{6ba7b810-9dad-11d1-80b4-00c04fd430c8}",
// "urn:uuid:6ba7b810-9dad-11d1-80b4-00c04fd430c8"
// Implements the exact same code as the UUID UnmarshalText removing the
// version check.
func (u *ULID) UnmarshalText(text []byte) (err error) {
	if len(text) < 32 {
		err = fmt.Errorf("uuid: UUID string too short: %s", text)
		return
	}

	t := text[:]
	braced := false

	if bytes.Equal(t[:9], urnPrefix) {
		t = t[9:]
	} else if t[0] == '{' {
		braced = true
		t = t[1:]
	}

	b := u[:]

	for i, byteGroup := range byteGroups {
		if i > 0 && t[0] == '-' {
			t = t[1:]
		} else if i > 0 && t[0] != '-' {
			err = ErrULIDInvalidStringFormat
			return
		}

		if len(t) < byteGroup {
			err = fmt.Errorf("kallax: ulid string too short: %s", text)
			return
		}

		if i == 4 && len(t) > byteGroup &&
			((braced && t[byteGroup] != '}') || len(t[byteGroup:]) > 1 || !braced) {
			err = fmt.Errorf("kallax: ulid string too long: %s", t)
			return
		}

		_, err = hex.Decode(b[:byteGroup/2], t[:byteGroup])

		if err != nil {
			return
		}

		t = t[byteGroup:]
		b = b[byteGroup/2:]
	}

	return
}

func (id ULID) MarshalText() ([]byte, error) {
	return []byte(id.String()), nil
}

// Value implements the Valuer interface.
func (id ULID) Value() (driver.Value, error) {
	return uuid.UUID(id).Value()
}

// IsEmpty returns whether the ID is empty or not. An empty ID means it has not
// been set yet.
func (id ULID) IsEmpty() bool {
	return uuid.UUID(id) == uuid.Nil
}

// String returns the string representation of the ID.
func (id ULID) String() string {
	return uuid.UUID(id).String()
}

/*
// Equals reports whether the ID and the given one are equals.
func (id ULID) Equals(other Identifier) bool {
	v, ok := other.(*ULID)
	if !ok {
		return false
	}

	return uuid.UUID(id) == uuid.UUID(*v)
}
*/

// Raw returns the underlying raw value.
func (id ULID) Raw() interface{} {
	return id
}
