// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"testing"
)

func TestValidators__validateOptionFName(t *testing.T) {
	v := &validator{}
	if err := v.validateOptionFName("1/SMITH JOHN"); err != nil {
		t.Error(err)
	}
	if err := v.validateOptionFName("1/"); err == nil {
		t.Error("expected error")
	}
	if err := v.validateOptionFName("1"); err == nil {
		t.Error("expected error")
	}
	if err := v.validateOptionFName(""); err == nil {
		t.Error("expected error")
	}
	if err := v.validateOptionFName(" /"); err == nil {
		t.Error("expected error")
	}
}
