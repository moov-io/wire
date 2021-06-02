// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators__validateOptionFName(t *testing.T) {
	v := &validator{}

	require.NoError(t, v.validateOptionFName("1/SMITH JOHN"))
	require.Error(t, v.validateOptionFName("1/"))
	require.Error(t, v.validateOptionFName("1"))
	require.Error(t, v.validateOptionFName(""))
	require.Error(t, v.validateOptionFName(" /"))
}

func TestValidators_isCentury(t *testing.T) {
	v := &validator{}
	require.NoError(t, v.isCentury("20"))
	require.Error(t, v.isCentury(""))
	require.Error(t, v.isCentury("2"))
	require.Error(t, v.isCentury("19"))
	require.Error(t, v.isCentury("2000"))
	require.Error(t, v.isCentury("AB"))
}

func TestValidators_isYear(t *testing.T) {
	v := &validator{}
	require.NoError(t, v.isYear("00"))
	require.NoError(t, v.isYear("35"))
	require.Error(t, v.isYear(""))
	require.Error(t, v.isYear("2"))
	require.Error(t, v.isYear("1998"))
	require.Error(t, v.isYear("AB"))
}

func TestValidators_isMonth(t *testing.T) {
	v := &validator{}
	require.NoError(t, v.isMonth("07"))
	require.Error(t, v.isMonth("7"))
	require.Error(t, v.isMonth(""))
	require.Error(t, v.isMonth("13"))
	require.Error(t, v.isMonth("AB"))
}

func TestValidators_isDay(t *testing.T) {
	v := &validator{}
	require.NoError(t, v.isDay("02", "29"))
	require.NoError(t, v.isDay("04", "30"))
	require.NoError(t, v.isDay("01", "31"))
	require.Error(t, v.isDay("", ""))
	require.Error(t, v.isDay("", "02"))
	require.Error(t, v.isDay("02", ""))
	require.Error(t, v.isDay("02", "30"))
	require.Error(t, v.isDay("04", "31"))
	require.Error(t, v.isDay("01", "32"))
	require.Error(t, v.isDay("1", "01"))
	require.Error(t, v.isDay("01", "1"))
	require.Error(t, v.isDay("AB", "AB"))
}

func TestValidators_validateDate(t *testing.T) {
	v := &validator{}
	require.NoError(t, v.validateDate("20080131"))
	require.Error(t, v.validateDate(""))
	require.Error(t, v.validateDate("19980131")) // invalid century (19)
	require.Error(t, v.validateDate("20AB0131")) // invalid year (AB)
	require.Error(t, v.validateDate("20089931")) // invalid month (99)
	require.Error(t, v.validateDate("20080199")) // invalid day (99)
}
