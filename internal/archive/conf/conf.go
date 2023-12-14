// Package conf is a highly opinionated package, it is sortof inspired by spf13/viper. Differences are
// mostly around panic vs fallback. Viper tends to do a fallback to default values when no value is
// found. Eg. if we create a config value `WEB_HOST: string`, and that value cannot be found for any
// overrides, viper returns an empty string. conf, on the other hand, panics. This falls in with our
// mantra of failing fast and loud rather than silently and slowly. If there is a config value
// missing, this is cause for a runtime panic at startup.
//
// conf is also slightly inspired by ardanlabs/conf but with more tooling and less opiniontated on
// where the config comes from. Anyone can implement their own config override.
//
// The main theme for this package is: panic first ask questions later. Since this is a crucial
// _startup package_, don't expect any error handling. Usage of this package is considered "UNSAFE".
// It will panic and crash your app, that's exactly what we want. Removing error handling also
// simplifies the code usage, and enforces proper usage of the package.
package conf

import (
	"fmt"
	"net/url"

	"time"
)

type Key string

type Map map[string]Field

type Conf struct {
	cnf     Map
	prefix  string
	usage   string
	cString string
	// warning string prints all the cases where we don't want to crash the confbuilder.Build, but
	// which will probably crash your app if you don't fix them. For example, forgetting to set a
	// required field, meaning that the outcome is a nil value.
	warnString string
}

func (c Conf) GetField(field string) Field {
	return unsafeGetMapVal(c.cnf, field)
}

func (c Conf) GetString(field string) string {
	f := unsafeGetMapVal(c.cnf, field)
	x, ok := f.Value.(string)
	if !ok {
		panic(expectedVal(field, TYPE_STRING, f.Value))
	}

	return x
}

func (c Conf) GetInt(field string) int {
	f := unsafeGetMapVal(c.cnf, field)
	x, ok := f.Value.(int)
	if !ok {
		panic(expectedVal(field, TYPE_INT, f.Value))
	}

	return x
}

func (c Conf) GetBool(field string) bool {
	f := unsafeGetMapVal(c.cnf, field)
	x, ok := f.Value.(bool)
	if !ok {
		panic(expectedVal(field, TYPE_BOOL, f.Value))
	}

	return x
}

func (c Conf) GetDuration(field string) time.Duration {
	f := unsafeGetMapVal(c.cnf, field)
	x, ok := f.Value.(time.Duration)
	if !ok {
		panic(expectedVal(field, TYPE_DURATION, f.Value))
	}

	return x
}

func (c Conf) GetURL(field string) url.URL {
	f := unsafeGetMapVal(c.cnf, field)
	x, ok := f.Value.(url.URL)
	if !ok {
		panic(expectedVal(field, TYPE_URL, f.Value))
	}

	return x
}

func (c Conf) String() string {
	// 	Key					|	DEFAULT			|	JSON 			|	ENV 	 		| 	RESULT
	//
	// 	HOST	 			| 	0.0.0.0:4000	|	0.0.0.0:80		|	0.0.0.0:3000	|	0.0.0.0:3000
	// 	BUILD_TIME			|	-				|	1662620745		|	-				|	1662620745
	// 	GRPC_SERVICE_HOST	|	0.0.0.0:4000	|	-	 			|	0.0.0.0:9000	| 	0.0.0.0:9000

	return c.cString
}

func (c Conf) Usage() string {
	// Key					|	TYPE 	 	|	DEFAULT	 		| 	JSON		        	|	ENV
	//
	// API_HOST				|	String		|	0.0.0.0:4000	|	--prefix-hello-world	|	PREFIX_API_HOST
	// BUILD_TIME			|	Int	 		|	-		 		|	--prefix-hello-world-2	|	PREFIX_BUILD_TIME
	// WEB_SHUTDOWN_TIMEOUT	|	Duration 	|	5sec	 		|	--prefix-hello-world-2	|	PREFIX_WEB_SHUTDOWN_TIMEOUT
	//
	// This table doesn't currently show the DESCRIPTION column, but you get the point

	return c.usage
}

func (c Conf) Warn() string {
	return c.warnString
}

// MaybePrintWarn prints the Warning string if it is non empty
func (c Conf) MaybePrintWarn() {
	if c.Warn() != "" {
		fmt.Println()
		fmt.Println("WARNING ")
		fmt.Print(c.Warn())
	}
}

func (c Conf) GetPrefix() string {
	return c.prefix
}

func unsafeGetMapVal[T comparable, V any](m map[T]V, t T) V {
	v, ok := m[t]
	if !ok {
		panic("value missing from map")
	}

	return v
}
