package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	args []string
	err  error
	config
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, numTimes: 10},
		},
		{
			args:   []string{"abc"},
			err:    errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{printUsage: false, numTimes: 0},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if err != nil {
			if tc.err == nil {
				t.Errorf("Expected nil error, got %v\n", err)
			} else if err.Error() != tc.err.Error() {
				t.Fatalf("Expected error to be: %v, got %v\n", tc.err, err)
			}
		} else {
			if tc.err != nil {
				t.Errorf("Expected error to be: %v, got nil\n", tc.err)
			}
		}
		if c.printUsage != tc.config.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.config.printUsage, c.printUsage)
		}
		if c.numTimes != tc.config.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.config.numTimes, c.numTimes)
		}
	}
}
