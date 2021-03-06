// Copyright (2017) Sandia Corporation.
// Under the terms of Contract DE-AC04-94AL85000 with Sandia Corporation,
// the U.S. Government retains certain rights in this software.

// A fifo for bytes.

package main

import (
	"sync"
)

type byteFifo struct {
	data  []byte
	limit int // max size
	sync.Mutex
}

func NewByteFifo(limit int) *byteFifo {
	b := make([]byte, 0)
	return &byteFifo{data: b, limit: limit}
}

func (bf *byteFifo) Get() []byte {
	bf.Lock()
	defer bf.Unlock()

	return bf.data
}

func (bf *byteFifo) Write(p []byte) (n int, err error) {
	bf.Lock()
	defer bf.Unlock()

	bf.data = append(bf.data, p...)
	if len(bf.data) > bf.limit {
		bf.data = bf.data[len(bf.data)-bf.limit-1:]
	}
	n = len(p)
	return
}
