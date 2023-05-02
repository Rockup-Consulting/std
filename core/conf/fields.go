package conf

import (
	"fmt"
	"net/url"
	"time"
)

type ConfType int

func (c ConfType) String() string {
	switch c {
	case 1:
		return "<string>"
	case 2:
		return "<int>"
	case 3:
		return "<bool>"
	case 4:
		return "<time.Duration>"
	case 5:
		return "<url>"
	default:
		panic(fmt.Sprintf("Field type %d does not exist", c))
	}
}

const (
	TYPE_STRING ConfType = iota + 1
	TYPE_INT
	TYPE_BOOL

	// TYPE_DURATION is a time.Duration. The parsing function is time.ParseDuration, see package
	// time for more information
	TYPE_DURATION
	TYPE_URL
)

// FieldKeys are used to create override values
//
// We use an array of fixed size to make FieldKeys directly comparables
type FieldKeys [4]string

type Field struct {
	Type        ConfType
	Value       any
	Keys        FieldKeys
	Sensitive   bool
	Description string
	Help        string
	Required    bool
}

func (f Field) checkDefault(k string) error {
	// If the default value is nil we assume that the user wants to set it at a later stage.
	// The error is effectively deffered until errConfValueNotSet is returned.
	if f.Value == nil {
		return nil
	}

	switch f.Type {
	case TYPE_STRING:
		_, ok := f.Value.(string)
		if !ok {
			return errIncorrectTypeForDefaultValue(k, f)
		}

	case TYPE_INT:
		_, ok := f.Value.(int)
		if !ok {
			return errIncorrectTypeForDefaultValue(k, f)
		}
	case TYPE_BOOL:
		_, ok := f.Value.(bool)
		if !ok {
			return errIncorrectTypeForDefaultValue(k, f)
		}
	case TYPE_DURATION:
		_, ok := f.Value.(time.Duration)
		if !ok {
			return errIncorrectTypeForDefaultValue(k, f)
		}
	case TYPE_URL:
		_, ok := f.Value.(url.URL)
		if !ok {
			return errIncorrectTypeForDefaultValue(k, f)
		}
	default:
		panic(fmt.Sprintf("unimplemented for conf.Type %s", f.Type))
	}

	return nil
}

func (l FieldKeys) Slice() []string {
	var out []string

	for _, s := range l {
		if s == "" {
			break
		}
		out = append(out, s)
	}

	return out
}

func (f Field) Valid(key string) error {
	// type must be set
	if f.Type == 0 {
		return errTypeNotSet(key)
	}

	// assert that the Field has valid FieldKeys

	err := f.Keys.Valid(key)
	if err != nil {
		return err
	}

	// if default value is set, check that it has a valid type
	err = f.checkDefault(key)
	if err != nil {
		return err
	}

	return nil
}

func (l FieldKeys) Valid(key string) error {
	// assert that the entry has atleast one non empty string at the start
	if l[0] == "" {
		return errEmptyKey(key)
	}

	// FieldKeys can't have a broken chain, eg. ["ONE", "", "TWO", ""] is not allowed
	for i, s := range l {
		if s == "" {
			j := i
			for j < len(l) {
				ss := l[j]
				if ss != "" {
					return errFieldKeyBrokenChain(key)
				}

				j++
			}
			break
		}
	}

	return nil
}
