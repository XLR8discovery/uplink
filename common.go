// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package splitter

import (
	"xlr8d.io/common/encryption"
	"xlr8d.io/common/oss"
	"xlr8d.io/uplink/private/metaclient"
)

// TODO: move it to separate package?
func encryptETag(etag []byte, cipherSuite oss.CipherSuite, contentKey *oss.Key) ([]byte, error) {
	etagKey, err := encryption.DeriveKey(contentKey, "oss-etag-v1")
	if err != nil {
		return nil, err
	}

	encryptedETag, err := encryption.Encrypt(etag, cipherSuite, etagKey, &oss.Nonce{})
	if err != nil {
		return nil, err
	}

	return encryptedETag, nil
}

func nonceForPosition(position metaclient.SegmentPosition) (oss.Nonce, error) {
	var nonce oss.Nonce
	inc := (int64(position.PartNumber) << 32) | (int64(position.Index) + 1)
	_, err := encryption.Increment(&nonce, inc)
	return nonce, err
}