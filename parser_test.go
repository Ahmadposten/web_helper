// Tests parser functions
package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestParseRecord(t *testing.T) {
	sadScenarioInputs := []string{
		"",
		"12.2.3.4",
	}

	happyScenarioInputs := []struct {
		log string
		ip  string
	}{
		{
			"1.2.3.4 -- whatever [whatever] <whatever>",
			"1.2.3.4",
		},
	}

	for _, scenarioIp := range sadScenarioInputs {
		_, err := parseRecord(scenarioIp)
		require.Error(t, err)
	}
	for _, scenario := range happyScenarioInputs {
		rec, err := parseRecord(scenario.log)
		require.NoError(t, err)
		assert.Equal(t, rec.Ip, net.ParseIP(scenario.ip))
	}
}
