// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

//go:generate go run gen.go

// Package cursed copies the code from the oss module necessary to unpack
// and repack internal satellite segment ids.
package cursed

import (
	"github.com/spacemonkeygo/monkit/v3"

	"xlr8d.io/common/oss"
)

var mon = monkit.Package()

// PieceID is an alias to oss.PieceID for use in generated protobuf code.
type PieceID = oss.PieceID