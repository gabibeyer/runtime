// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package rootless

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var uidMapPathStore = uidMapPath

func createTestUIDMapFile(input string) error {
	f, err := os.Create(uidMapPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(input)
	if err != nil {
		return err
	}

	return nil
}

// TestIstRootlessRootlessUID1000 tests that isRootless is set to
// true when a host UID is not 0
func TestIsRootlessRootlessUID1000(t *testing.T) {
	assert := assert.New(t)

	// by default isRootless should be set to false initially
	assert.False(isRootless)

	tmpDir, err := ioutil.TempDir("", "")
	assert.NoError(err)

	uidMapPath = filepath.Join(tmpDir, "testUIDMapFile")
	defer func() {
		uidMapPath = uidMapPathStore
		os.RemoveAll(tmpDir)
		isRootless = false
	}()

	mapping := "\t0\t1000\t5555"
	err = createTestUIDMapFile(mapping)
	assert.NoError(err)

	// make call to IsRootless, this should also call
	// SetRootless
	assert.True(IsRootless())
}

// TestIsRootlessRootUID0 tests that isRootless is not set when
// the host UID is 0
func TestIsRootlessRootUID0(t *testing.T) {
	assert := assert.New(t)

	// by default isRootless should be set to false initially
	assert.False(isRootless)

	tmpDir, err := ioutil.TempDir("", "")
	assert.NoError(err)

	uidMapPath = filepath.Join(tmpDir, "testUIDMapFile")
	defer func() {
		uidMapPath = uidMapPathStore
		os.RemoveAll(uidMapPath)
		isRootless = false
	}()

	mapping := "\t0\t0\t5555"
	err = createTestUIDMapFile(mapping)
	assert.NoError(err)

	// make call to IsRootless, this should also call
	// SetRootless
	assert.False(IsRootless())
}
