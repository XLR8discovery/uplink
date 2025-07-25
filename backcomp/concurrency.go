// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

// Package backcomp contains utilities for handling backwards incompatible changes.
package backcomp

import (
	"context"

	"xlr8d.io/uplink"
	"xlr8d.io/uplink/internal/expose"
)

// RequestAccessWithPassphraseAndConcurrency requests satellite for a new access grant using a passhprase and specific concurrency for the Argon2 key derivation.
func RequestAccessWithPassphraseAndConcurrency(ctx context.Context, config uplink.Config, satelliteNodeURL, apiKey, passphrase string, concurrency uint8) (_ *uplink.Access, err error) {
	return expose.ConfigRequestAccessWithPassphraseAndConcurrency(config, ctx, satelliteNodeURL, apiKey, passphrase, concurrency)
}
