package conf

import (
	"errors"
	"fmt"
)

var (
	ErrNilMap = errors.New("config cannot be initialized with a nil conf.Map")
)

type ErrTypeNotSet struct {
	key string
}

func (e ErrTypeNotSet) Error() string {
	return fmt.Sprintf("config value %q type not set", e.key)
}

var errTypeNotSet = func(key string) ErrTypeNotSet {
	return ErrTypeNotSet{key}
}

type ErrEmptyKey struct {
	key string
}

func (e ErrEmptyKey) Error() string {
	return fmt.Sprintf("config value %q cannot have an empty key", e.key)
}

var errEmptyKey = func(key string) ErrEmptyKey {
	return ErrEmptyKey{key}
}

type ErrFieldKeyBrokenChain struct {
	key string
}

func (e ErrFieldKeyBrokenChain) Error() string {
	return fmt.Sprintf("config value %q should not have a broken chain", e.key)
}

var errFieldKeyBrokenChain = func(key string) ErrFieldKeyBrokenChain {
	return ErrFieldKeyBrokenChain{key}
}

type ErrIncorrectTypeForDefaultValue struct {
	key string
	f   Field
}

func (e ErrIncorrectTypeForDefaultValue) Error() string {
	return fmt.Sprintf("default value has incorrect type for config field %q. \n\tValue - %s\n\tExpected Type - %s\n\tGot Type - <%T>", e.key, e.f.Value, e.f.Type, e.f.Value)
}

var errIncorrectTypeForDefaultValue = func(key string, f Field) ErrIncorrectTypeForDefaultValue {
	return ErrIncorrectTypeForDefaultValue{key, f}
}

type ErrTypeParse struct {
	key string
	f   Field
}

func (e ErrTypeParse) Error() string {
	return fmt.Sprintf("parse conf: expected field %q to have type %q, but got type %T", e.key, e.f.Type, e.f.Value)
}

// If you see this again and it hasn't been used...delete it
// var errTypeParse = func(key string, f Field) ErrTypeParse {
// 	return ErrTypeParse{key, f}
// }

type ErrDuplicateFieldKeys struct {
	keys []string
}

func (e ErrDuplicateFieldKeys) Error() string {
	return fmt.Sprintf("found duplicate Field Keys on Conf Fields %s, this is a mistake and must be fixed", e.keys)
}

var errDuplicateFieldKeys = func(keys []string) ErrDuplicateFieldKeys {
	return ErrDuplicateFieldKeys{keys}
}

var expectedVal = func(field string, t ConfType, val any) string {
	return fmt.Sprintf("tried to get conf field '%s' as %s, but actual value is %T", field, t, val)
}
