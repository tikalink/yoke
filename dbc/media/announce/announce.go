// Code generated by entc, DO NOT EDIT.

package announce

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the announce type in the database.
	Label = "announce"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedUnix holds the string denoting the created_unix field in the database.
	FieldCreatedUnix = "created_unix"
	// FieldUpdatedUnix holds the string denoting the updated_unix field in the database.
	FieldUpdatedUnix = "updated_unix"
	// FieldDeletedUnix holds the string denoting the deleted_unix field in the database.
	FieldDeletedUnix = "deleted_unix"
	// FieldAnnounceNo holds the string denoting the announce_no field in the database.
	FieldAnnounceNo = "announce_no"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldSign holds the string denoting the sign field in the database.
	FieldSign = "sign"
	// Table holds the table name of the announce in the database.
	Table = "announces"
)

// Columns holds all SQL columns for announce fields.
var Columns = []string{
	FieldID,
	FieldCreatedUnix,
	FieldUpdatedUnix,
	FieldDeletedUnix,
	FieldAnnounceNo,
	FieldTitle,
	FieldKind,
	FieldContent,
	FieldLink,
	FieldSign,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedUnix holds the default value on creation for the "created_unix" field.
	DefaultCreatedUnix int64
	// DefaultUpdatedUnix holds the default value on creation for the "updated_unix" field.
	DefaultUpdatedUnix int64
	// DefaultAnnounceNo holds the default value on creation for the "announce_no" field.
	DefaultAnnounceNo string
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultLink holds the default value on creation for the "link" field.
	DefaultLink string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Kind defines the type for the "kind" enum field.
type Kind string

// KindNotice is the default value of the Kind enum.
const DefaultKind = KindNotice

// Kind values.
const (
	KindNotice       Kind = "notice"
	KindEvent        Kind = "event"
	KindAnnouncement Kind = "announcement"
	KindUpdate       Kind = "update"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindNotice, KindEvent, KindAnnouncement, KindUpdate:
		return nil
	default:
		return fmt.Errorf("announce: invalid enum value for kind field: %q", k)
	}
}
