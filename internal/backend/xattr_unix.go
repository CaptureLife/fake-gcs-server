// Copyright 2017 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows
// +build !windows

package backend

import (
	"github.com/pkg/xattr"
	"strings"
	"os"
)

const xattrKey = "user.metadata"

func writeXattr(path string, encoded []byte) error {
	err := xattr.Set(path, xattrKey, encoded)
	if err != nil {
		return os.WriteFile(path + ".xattr", encoded, 0666)
	}
	return err
}

func readXattr(path string) ([]byte, error) {
	val, err := xattr.Get(path, xattrKey)
	if err != nil {
		return os.ReadFile(path + ".xattr")
	}
	return val, err
}

func isXattrFile(path string) bool {
	return strings.HasSuffix(path, ".xattr")
}

func removeXattrFile(path string) error {
	return os.Remove(path + ".xattr")
}
