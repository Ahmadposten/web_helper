package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestToCidr(t *testing.T) {
	ip := "1.2.3.4"
	cidrizedIp := toCidr(ip)

	assert.Equal(t, "1.2.3.4/32", cidrizedIp)
}

func TestFilterIps(t *testing.T) {
	f, err := os.Open("test_files/sample.log")
	require.NoError(t, err)

	cases := []struct {
		ip    string
		count int
	}{
		{
			"78.29.246.2",
			7,
		},
		{
			"2.24.146.2",
			2,
		},
		{
			"2.24.146.0/24",
			3,
		},
		{
			"0.0.0.0/0",
			11,
		},
	}

	for _, testCase := range cases {
		r, w, _ := os.Pipe()
		filterIps(testCase.ip, f, w)
		f.Seek(0, 0)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		assert.Equal(t, testCase.count, len(strings.Split(string(out), "\n")))
	}
}
