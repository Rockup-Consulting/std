// package neox provides a set of extensions on top of the Neo4j Go Driver. It plays a similar role
// for Neo4j that [SQLX](https://jmoiron.github.io/sqlx/) plays for the database/sql package.
//
// Neo4j Driver docs can be found here: neo4j.com/docs/go-manual/current/
package neox

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
)

type Conf struct {
	HOST              string
	UserName          string
	Password          string
	ConnectionTimeout time.Duration
}

var (
	DevConf = Conf{
		HOST:              "bolt://localhost:7687",
		UserName:          "neo4j",
		Password:          "password",
		ConnectionTimeout: time.Second,
	}
)

func NewDriver(conf Conf) (neo4j.DriverWithContext, error) {
	auth := neo4j.BasicAuth(conf.UserName, conf.Password, "")

	driver, err := neo4j.NewDriverWithContext(
		conf.HOST,
		auth,
	)

	if err != nil {
		return nil, err
	}

	return driver, nil
}

// Returns T, isNil
func UnsafeGet[T neo4j.RecordValue](record *db.Record, field string) (T, bool) {
	t, isNil, err := neo4j.GetRecordValue[T](record, field)

	if err != nil {
		panic(fmt.Sprintf("expected '%s' field to be on neo4j response: %s", field, err))
	}

	return t, isNil
}

func UnsafeAssert[T neo4j.RecordValue](val any) T {
	t, ok := val.(T)
	if !ok {
		panic(fmt.Sprintf("unexpected type for value: %T", val))
	}

	return t
}

type WriteWork func(tx neo4j.ManagedTransaction) error

func ExecuteWrite(ctx context.Context, db neo4j.DriverWithContext, work WriteWork) error {
	if _, err := db.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}).ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return nil, work(tx)
	}); err != nil {
		return err
	}

	return nil
}

func ExecuteWriteT[T any](ctx context.Context, db neo4j.DriverWithContext, tx neo4j.ManagedTransactionWorkT[T]) (T, error) {
	return neo4j.ExecuteWrite(ctx, db.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}), tx)
}

func ExecuteReadT[T any](ctx context.Context, db neo4j.DriverWithContext, tx neo4j.ManagedTransactionWorkT[T]) (T, error) {
	return neo4j.ExecuteRead(ctx, db.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead}), tx)
}

// Single returns => neo4j.Record, notFound (bool), error
// Warning, if more than one record is found, this method will panic. You should provde a useful
// panic message when calling this method. The panic message is:
//
//	"found multiple entries: '[your message]'"
func Single(ctx context.Context, result neo4j.ResultWithContext, panicMsg string) (*neo4j.Record, bool, error) {
	if !result.Next(ctx) {
		err := result.Err()
		if err != nil {
			return nil, false, err
		}

		return nil, true, nil
	}

	r := result.Record()
	if result.Peek(ctx) {
		panic(fmt.Sprintf("found multiple entries: %q", panicMsg))
	}

	return r, false, nil
}

// ClearDB is a test utility that clears all data from the current instance
func ClearDB(ctx context.Context, t testing.TB, db neo4j.DriverWithContext) {
	t.Helper()

	ExecuteWrite(ctx, db, func(tx neo4j.ManagedTransaction) error {
		_, err := tx.Run(ctx, "MATCH (n) DETACH DELETE n", nil)

		if err != nil {
			t.Fatalf("clearing Neo4j: %s", err)
		}
		return nil
	})
}
