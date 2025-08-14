// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package etag_test

import (
	"bytes"
	"crypto/sha256"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"xlr8d.io/common/memory"
	"xlr8d.io/common/testrand"
	"xlr8d.io/uplink/private/etag"
)

func TestHashReader(t *testing.T) {
	inputData := testrand.Bytes(1 * memory.KiB)
	expectedETag := sha256.Sum256(inputData)

	reader := etag.NewHashReader(bytes.NewReader(inputData), sha256.New())
	readData, err := io.ReadAll(reader)
	require.NoError(t, err)
	require.Equal(t, inputData, readData)
	require.Equal(t, expectedETag[:], reader.CurrentETag())
}