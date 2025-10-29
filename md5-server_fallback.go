//go:build !amd64 || appengine || !gc || noasm
// +build !amd64 appengine !gc noasm

// SPDX-FileCopyrightText: 2025 openstor contributors
// SPDX-FileCopyrightText: 2015-2025 MinIO, Inc.
// SPDX-License-Identifier: Apache-2.0-or-later

package md5simd

// NewServer - Create new object for parallel processing handling
func NewServer() *fallbackServer {
	return &fallbackServer{}
}

func NewServerWithOptions(opts ServerOptions) *fallbackServer {
	return &fallbackServer{}
}
