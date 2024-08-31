package main

import (
	"testing"
)

func TestSocketKey(t *testing.T) {
	var str string

	str = generateWebSocketAccept("dGhlIHNhbXBsZSBub25jZQ==")
	if str != "s3pPLMBiTxaQ9kYGzzhZRbK+xOo=" {
		t.Errorf("generateWebSocketAccept(\"dGhlIHNhbXBsZSBub25jZQ==\") = %s; want \"s3pPLMBiTxaQ9kYGzzhZRbK+xOo=\"", str)
	}

}
