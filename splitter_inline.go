// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package splitter

import (
	"bytes"
	"io"

	"xlr8d.io/common/oss"
	"xlr8d.io/uplink/private/metaclient"
)

type splitterInline struct {
	position   metaclient.SegmentPosition
	encryption metaclient.SegmentEncryption
	encParams  oss.EncryptionParameters
	contentKey *oss.Key

	encData   []byte
	plainSize int64
}

func (s *splitterInline) Begin() metaclient.BatchItem {
	return &metaclient.MakeInlineSegmentParams{
		StreamID:            nil, // set by the stream batcher
		Position:            s.position,
		Encryption:          s.encryption,
		EncryptedInlineData: s.encData,
		PlainSize:           s.plainSize,
		EncryptedETag:       nil, // set by the segment tracker
	}
}

func (s *splitterInline) Position() metaclient.SegmentPosition { return s.position }
func (s *splitterInline) Inline() bool                         { return true }
func (s *splitterInline) Reader() io.Reader                    { return bytes.NewReader(s.encData) }
func (s *splitterInline) DoneReading(err error)                {}

func (s *splitterInline) EncryptETag(eTag []byte) ([]byte, error) {
	return encryptETag(eTag, s.encParams.CipherSuite, s.contentKey)
}

func (s *splitterInline) Finalize() *SegmentInfo {
	return &SegmentInfo{
		Encryption:    s.encryption,
		PlainSize:     s.plainSize,
		EncryptedSize: int64(len(s.encData)),
	}
}