// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package expose

import (
	"context"
	_ "unsafe" // for go:linkname

	"xlr8d.io/common/grant"
	"xlr8d.io/common/macaroon"
	"xlr8d.io/common/rpc"
	"xlr8d.io/common/rpc/rpcpool"
	"xlr8d.io/uplink"
)

// ConfigSetConnectionPool exposes Config.setConnectionPool.
//
//go:linkname ConfigSetConnectionPool xlr8d.io/uplink.config_setConnectionPool
func ConfigSetConnectionPool(*uplink.Config, *rpcpool.Pool)

// ConfigSetSatelliteConnectionPool exposes Config.setSatelliteConnectionPool.
//
//go:linkname ConfigSetSatelliteConnectionPool xlr8d.io/uplink.config_setSatelliteConnectionPool
func ConfigSetSatelliteConnectionPool(*uplink.Config, *rpcpool.Pool)

// ConfigGetDialer exposes Config.getDialer.
//
//go:linkname ConfigGetDialer xlr8d.io/uplink.config_getDialer
//nolint:revive
func ConfigGetDialer(uplink.Config, context.Context) (rpc.Dialer, error)

// ConfigSetMaximumBufferSize exposes Config.setMaximumBufferSize.
//
//go:linkname ConfigSetMaximumBufferSize xlr8d.io/uplink.config_setMaximumBufferSize
func ConfigSetMaximumBufferSize(*uplink.Config, int)

// ConfigDisableObjectKeyEncryption exposes Config.disableObjectKeyEncryption.
//
//go:linkname ConfigDisableObjectKeyEncryption xlr8d.io/uplink.config_disableObjectKeyEncryption
func ConfigDisableObjectKeyEncryption(config *uplink.Config)

// AccessGetAPIKey exposes Access.getAPIKey.
//
//go:linkname AccessGetAPIKey xlr8d.io/uplink.access_getAPIKey
func AccessGetAPIKey(*uplink.Access) *macaroon.APIKey

// AccessGetEncAccess exposes Access.getEncAccess.
//
//go:linkname AccessGetEncAccess xlr8d.io/uplink.access_getEncAccess
func AccessGetEncAccess(*uplink.Access) *grant.EncryptionAccess

// AccessToInternal exposes uplink.access_toInternal.
//
//go:linkname AccessToInternal xlr8d.io/uplink.access_toInternal
func AccessToInternal(*uplink.Access) *grant.Access

// AccessFromInternal exposes uplink.access_fromInternal.
//
//go:linkname AccessFromInternal xlr8d.io/uplink.access_fromInternal
func AccessFromInternal(*grant.Access) *uplink.Access

// ConfigRequestAccessWithPassphraseAndConcurrency exposes Config.requestAccessWithPassphraseAndConcurrency.
//
//nolint:revive
//go:linkname ConfigRequestAccessWithPassphraseAndConcurrency xlr8d.io/uplink.config_requestAccessWithPassphraseAndConcurrency
func ConfigRequestAccessWithPassphraseAndConcurrency(config uplink.Config, ctx context.Context, satelliteAddress, apiKey, passphrase string, concurrency uint8) (*uplink.Access, error)

// ConfigDisableBackgroundQoS exposes Config.disableBackgroundQoS.
//
//go:linkname ConfigDisableBackgroundQoS xlr8d.io/uplink.config_disableBackgroundQoS
func ConfigDisableBackgroundQoS(config *uplink.Config, disabled bool)
