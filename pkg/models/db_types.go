package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DBModel interface{}

type Id string

func (Id) GormDataType() string {
	return "varchar(64) NOT NULL"
}

func MarshalValue(v interface{}) (driver.Value, error) {
	if v == nil {
		return nil, nil
	}
	bs, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return string(bs), nil
}

func UnmarshalValue(src interface{}, dst interface{}) error {
	if src == nil {
		return nil
	}

	bs, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("invalid type %T, value: %v", src, src)
	}
	return json.Unmarshal(bs, dst)
}

type StrSlice []string

func (v StrSlice) Value() (driver.Value, error) {
	return MarshalValue(v)
}

func (v *StrSlice) Scan(value interface{}) error {
	return UnmarshalValue(value, v)
}

type JSON json.RawMessage

func (v JSON) Value() (driver.Value, error) {
	return []byte(v), nil
}

func (v *JSON) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bs, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type %T, value: %v", value, value)
	}
	*v = bs
	return nil
}
