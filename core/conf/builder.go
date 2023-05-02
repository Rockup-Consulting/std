package conf

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"text/tabwriter"
)

const (
	defaultMaskVal = "*****"
)

type Builder struct {
	prefix           string
	m                Map
	overrides        []Override
	disallowDefaults bool
	unsafeBuild      bool // builds the config without panicing if a required value has not been set
	maskStr          string
}

/*
Note: disallowDefaults, in practise, just removes the default value. If we get the the end and that
value hasn't been set to anything other than nil (eg. overwritten), we panic...unless safeBuild has
been set to true. In that case we just write a warning string.
*/

// NewBuilder returns a builder.
//
// Regarding overrides: The overrides are applied in the order that they are passed in as. Meaning
// that the last passed in override takes preference over everything else.
func NewBuilder(m Map) *Builder {
	return &Builder{"", m, make([]Override, 0), false, false, defaultMaskVal}
}

// AddOverride appends a conf override to the conf Builder
func (b *Builder) AddOverride(o Override) {
	b.overrides = append(b.overrides, o)
}

// SetPrefix sets the builder Prefix. This is useful for avoiding configuration variable clashes.
// For example, if you have APP_CLI_WEB_HOST vs APP_CLI_DEV_WEB_HOST, you can separate working
// against localhost vs production
func (b *Builder) SetPrefix(prefix string) {
	b.prefix = prefix
}

// UnsafeBuild sets the safeBuild flag to true, this will allow the conf build to succeed even if a
// required override has not been provided
func (b *Builder) UnsafeBuild() {
	b.unsafeBuild = true
}

// DissalowDefaults sets the dissalowDefaults boolean flag on the conf.Builder object to true.
// When dissalowDefaults is true, all fields _for which the required flag is enabled_ have to
// receive a variable from an override. If no override is provided, conf will panic
func (b *Builder) DissalowDefaults() {
	b.disallowDefaults = true
}

// SetMaskString sets the string used to mask sensitive values when printing Stringified conf
// or conf Usage
func (b *Builder) SetMaskString(mask string) {
	b.maskStr = mask
}

// Build constructs and returns a Conf struct.
//
// The steps taken to get to the final outcome can be viewed by calling the String method on the
// conf struct
func (b *Builder) Build() (*Conf, error) {
	if b.m == nil {
		return nil, ErrNilMap
	}

	var usageSb strings.Builder
	usageW := tabwriter.NewWriter(&usageSb, 14, 4, 4, ' ', 0)

	var stringSb strings.Builder
	stringW := tabwriter.NewWriter(&stringSb, 14, 4, 4, ' ', 0)

	var warnSb strings.Builder

	addTitles(usageW, stringW, b)

	if err := b.m.findDuplicateFieldKeys(); err != nil {
		return nil, err
	}

	for k, f := range b.m {
		if err := f.Valid(k); err != nil {
			panic(err)
		}

		// reset required fields to nil
		if f.Required && b.disallowDefaults {
			f.Value = nil
			b.m[k] = f
		}

		// Usage String
		val := f.Value

		if f.Sensitive {
			val = b.maskStr
		}
		if f.Value == nil {
			val = "-"
		}

		var usage string
		var confString string

		urlStruct, isURL := val.(url.URL)
		confString = fmt.Sprintf("%s\t", k)
		usage = fmt.Sprintf("%s\t%s\t%s", k, f.Type, f.Description)

		usage = fmt.Sprintf("%s\t%t", usage, f.Required)

		if isURL {
			usage = fmt.Sprintf("%s\t%s", usage, urlStruct.String())
			confString = fmt.Sprint(confString, urlStruct.String())
		} else {
			usage = fmt.Sprintf("%s\t%s", usage, val)
			confString = fmt.Sprint(confString, val)
		}

		for _, o := range b.overrides {
			overrideKey := o.Key(b.prefix, f)
			usage = fmt.Sprintf("%s\t%s", usage, overrideKey)

			var err error
			var val any

			switch f.Type {
			case TYPE_STRING:
				val, err = o.GetSTRING(overrideKey, f, overrideKey)
			case TYPE_INT:
				val, err = o.GetINT(overrideKey, f, overrideKey)
			case TYPE_BOOL:
				val, err = o.GetBOOL(overrideKey, f, overrideKey)
			case TYPE_DURATION:
				val, err = o.GetDURATION(overrideKey, f, overrideKey)
			case TYPE_URL:
				val, err = o.GetURL(overrideKey, f, overrideKey)
			default:
				panic(fmt.Sprintf("unimplemented conf.Type: %q", f.Type))
			}

			if err != nil {
				if errors.Is(err, errOverrideNotProvided) {
					confString = fmt.Sprintf("%s\t-", confString)
					continue
				}
				panic(err.Error())
			}

			if f.Sensitive {
				confString = fmt.Sprintf("%s\t%s", confString, b.maskStr)
			} else {
				urlStruct, isURL := val.(url.URL)
				confString = fmt.Sprintf("%s\t", confString)

				if isURL {
					confString = fmt.Sprint(confString, urlStruct.String())
				} else {
					confString = fmt.Sprint(confString, val)
				}

			}

			newField := f
			newField.Value = val
			b.m[k] = newField
		}

		finalFieldVal, ok := b.m[k]
		if !ok {
			panic("impossible case: value should be on map")
		}

		val = finalFieldVal.Value
		if finalFieldVal.Value == nil {
			if b.unsafeBuild {
				_, _ = fmt.Fprintf(&warnSb, "- no value has been set for config field '%s', this will panic if used in production\n", k)
				val = "-"
			} else {
				return nil, fmt.Errorf("default value has been dissalowed for conf field: %q with no overrides provided", k)
			}
		}

		if f.Sensitive && finalFieldVal.Value != nil {
			val = b.maskStr
		}

		confString = fmt.Sprintf("%s\t", confString)

		urlStruct, isURL = val.(url.URL)
		if isURL {
			confString = fmt.Sprint(confString, urlStruct.String())
		} else {
			confString = fmt.Sprint(confString, val)
		}

		_, _ = fmt.Fprintln(usageW, usage)
		_, _ = fmt.Fprintln(stringW, confString)
	}

	_ = usageW.Flush()
	_ = stringW.Flush()

	for _, o := range b.overrides {
		err := o.Cleanup()
		if err != nil {
			panic(err.Error())
		}
	}

	return &Conf{
		cnf:        b.m,
		prefix:     b.prefix,
		usage:      usageSb.String(),
		cString:    stringSb.String(),
		warnString: warnSb.String(),
	}, nil
}

// addTitles appends titles to the string builders based on the overrides
func addTitles(usageW *tabwriter.Writer, stringW *tabwriter.Writer, c *Builder) {
	_, _ = fmt.Fprintf(usageW, "KEY\tTYPE\tDESCRIPTION\tREQUIRED\tDEFAULT")
	// usageEmptyLine := "\t\t\t\t"
	_, _ = fmt.Fprintf(stringW, "KEY\tDEFAULT")

	for _, o := range c.overrides {
		_, _ = fmt.Fprint(usageW, "\t", o.Name())
		// usageEmptyLine = fmt.Sprint(usageEmptyLine, "\t")
		_, _ = fmt.Fprintf(stringW, "\t%s", o.Name())
	}
	_, _ = fmt.Fprint(usageW, "\n")
	// usageEmptyLine = fmt.Sprint(usageEmptyLine, "\n")
	_, _ = fmt.Fprintf(stringW, "\tRESULT")
	// _, _ = fmt.Fprint(usageW, usageEmptyLine)
	_, _ = fmt.Fprintln(stringW)
}
