package main

import (
	"encoding/json"

	"github.com/ztrue/tracerr"
)

type Report struct {
	Time         int64      `json:"time"`
	Path         string     `json:"path"`
	Instances    []Instance `json:"instances"`
	NumInstances int        `json:"numInstances"`
}

type Instance struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

func (r *Report) Encode() ([]byte, error) {
	ret, err := json.Marshal(r)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return ret, nil
}
