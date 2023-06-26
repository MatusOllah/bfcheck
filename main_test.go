package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAntiBF(t *testing.T) {
	cfg = Config{
		Verbose:   false,
		Path:      "./testdata/AntiBF",
		Color:     false,
		ShowLines: false,
	}

	r, err := checkDir(cfg.Path)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, 4, r.NumInstances)
}

func TestNotAntiBF(t *testing.T) {
	cfg = Config{
		Verbose:   false,
		Path:      "./testdata/NotAntiBF",
		Color:     false,
		ShowLines: false,
	}

	r, err := checkDir(cfg.Path)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, 0, r.NumInstances)
}
