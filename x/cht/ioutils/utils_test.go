package ioutils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func GetTestData() ([]byte, []byte, []byte, error) {
	chtCode, err := ioutil.ReadFile("../keeper/testdata/hackatom.cht")
	if err != nil {
		return nil, nil, nil, err
	}

	gzipData, err := GzipIt(chtCode)
	if err != nil {
		return nil, nil, nil, err
	}

	someRandomStr := []byte("hello world")

	return chtCode, someRandomStr, gzipData, nil
}

func TestIsCht(t *testing.T) {
	chtCode, someRandomStr, gzipData, err := GetTestData()
	require.NoError(t, err)

	t.Log("should return false for some random string data")
	require.False(t, IsCht(someRandomStr))
	t.Log("should return false for gzip data")
	require.False(t, IsCht(gzipData))
	t.Log("should return true for exact cht")
	require.True(t, IsCht(chtCode))
}

func TestIsGzip(t *testing.T) {
	chtCode, someRandomStr, gzipData, err := GetTestData()
	require.NoError(t, err)

	require.False(t, IsGzip(chtCode))
	require.False(t, IsGzip(someRandomStr))
	require.True(t, IsGzip(gzipData))
}

func TestGzipIt(t *testing.T) {
	chtCode, someRandomStr, _, err := GetTestData()
	originalGzipData := []byte{
		31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 202, 72, 205, 201, 201, 87, 40, 207, 47, 202, 73, 1,
		4, 0, 0, 255, 255, 133, 17, 74, 13, 11, 0, 0, 0,
	}

	require.NoError(t, err)

	t.Log("gzip cht with no error")
	_, err = GzipIt(chtCode)
	require.NoError(t, err)

	t.Log("gzip of a string should return exact gzip data")
	strToGzip, err := GzipIt(someRandomStr)

	require.True(t, IsGzip(strToGzip))
	require.NoError(t, err)
	require.Equal(t, originalGzipData, strToGzip)
}
