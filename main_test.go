package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAntiBF(t *testing.T) {
	cfg = Config{
		PosArgs: struct {
			Path string `description:"Path to FNF mod" default:"."`
		}{
			Path: "./testdata/AntiBF",
		},
		Verbose:   false,
		Color:     false,
		ShowLines: false,
	}

	r, err := checkDir(cfg.PosArgs.Path)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, 4, r.NumInstances)
}

func TestNotAntiBF(t *testing.T) {
	cfg = Config{
		PosArgs: struct {
			Path string `description:"Path to FNF mod" default:"."`
		}{
			Path: "./testdata/NotAntiBF",
		},
		Verbose:   false,
		Color:     false,
		ShowLines: false,
	}

	r, err := checkDir(cfg.PosArgs.Path)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, 0, r.NumInstances)
}
