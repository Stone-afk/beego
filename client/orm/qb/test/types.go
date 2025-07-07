package test

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/ecodeclub/ekit"
)

type SimpleStruct struct {
	Id      uint64 `orm:"primary_key,column=int_c"`
	Bool    bool
	BoolPtr *bool

	Int    int
	IntPtr *int

	Int8    int8
	Int8Ptr *int8

	Int16    int16
	Int16Ptr *int16

	Int32    int32
	Int32Ptr *int32

	Int64    int64
	Int64Ptr *int64

	Uint    uint
	UintPtr *uint

	Uint8    uint8
	Uint8Ptr *uint8

	Uint16    uint16
	Uint16Ptr *uint16

	Uint32    uint32
	Uint32Ptr *uint32

	Uint64    uint64
	Uint64Ptr *uint64

	Float32    float32
	Float32Ptr *float32

	Float64    float64
	Float64Ptr *float64

	String string

	NullString  sql.NullString
	NullInt64   sql.NullInt64
	NullBool    sql.NullBool
	NullFloat64 sql.NullFloat64
}

// JsonColumn is a custom JSON type field
// The Val field must be a struct pointer
type JsonColumn struct {
	Val   User
	Valid bool
}

type User struct {
	Name string
}

func (j *JsonColumn) Scan(src any) error {
	if src == nil {
		return nil
	}
	var bs []byte
	switch val := src.(type) {
	case string:
		bs = []byte(val)
	case []byte:
		bs = val
	case *[]byte:
		if val == nil {
			return nil
		}
		bs = *val
	default:
		return fmt.Errorf("illegal type %+v", src)
	}
	if len(bs) == 0 {
		return nil
	}
	err := json.Unmarshal(bs, &j.Val)
	if err != nil {
		return err
	}
	j.Valid = true
	return nil
}

func (j *JsonColumn) Value() (driver.Value, error) {
	if j == nil || !j.Valid {
		return nil, nil
	}
	bs, err := json.Marshal(j.Val)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func NewSimpleStruct(id uint64) *SimpleStruct {
	return &SimpleStruct{
		Id:         id,
		Bool:       true,
		BoolPtr:    ekit.ToPtr[bool](false),
		Int:        12,
		IntPtr:     ekit.ToPtr[int](13),
		Int8:       8,
		Int8Ptr:    ekit.ToPtr[int8](-8),
		Int16:      16,
		Int16Ptr:   ekit.ToPtr[int16](-16),
		Int32:      32,
		Int32Ptr:   ekit.ToPtr[int32](-32),
		Int64:      64,
		Int64Ptr:   ekit.ToPtr[int64](-64),
		Uint:       14,
		UintPtr:    ekit.ToPtr[uint](15),
		Uint8:      8,
		Uint8Ptr:   ekit.ToPtr[uint8](18),
		Uint16:     16,
		Uint16Ptr:  ekit.ToPtr[uint16](116),
		Uint32:     32,
		Uint32Ptr:  ekit.ToPtr[uint32](132),
		Uint64:     64,
		Uint64Ptr:  ekit.ToPtr[uint64](164),
		Float32:    3.2,
		Float32Ptr: ekit.ToPtr[float32](-3.2),
		Float64:    6.4,
		Float64Ptr: ekit.ToPtr[float64](-6.4),
		String:     "world",

		NullString:  sql.NullString{String: "null string", Valid: true},
		NullInt64:   sql.NullInt64{Int64: 64, Valid: true},
		NullBool:    sql.NullBool{Bool: true, Valid: true},
		NullFloat64: sql.NullFloat64{Float64: 6.4, Valid: true},
	}
}
