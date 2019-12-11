// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package license

// IdentityProvider specifies an API for generating unique identifiers.
type IdentityProvider interface {
	// ID generates the unique identifier.
	ID() (string, error)
}
