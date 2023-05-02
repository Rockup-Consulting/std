// Package auth exposes primitives useful for building authentication and authorization features for
// applications. More specifically, auth exposes functionality for converting a map of secrets to
// a pool of cryptox.AesService(s) that can be used safely across threads.
//
// The important thing for this package is to not overreach its boundaries. Package auth has to
// remain unopinionated about the specifics, and rather provide useful tooling.
package auth
