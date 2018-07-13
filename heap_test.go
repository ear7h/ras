package ras_test

import (
	"testing"
	"os"
	"github.com/ear7h/ras"
	"errors"
	"io/ioutil"
	"fmt"
)

func errEq(e1, e2 error) bool {
	if (e1 == nil) != (e2 == nil) {
		return false
	}

	if e1 == nil {
		return true
	}

	if e1.Error() == e2.Error() {
		return true
	}

	return false
}

func dumpFile(name string) []byte {
	byt, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return byt
}

func TestNewHeap(t *testing.T) {
	type tcase struct {
		file string
		data []byte
		err error
	}

	fn := func(tc tcase, t *testing.T) {
		if tc.data != nil {
			fd, err := os.OpenFile(tc.file, os.O_CREATE | os.O_RDWR, 0600)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}

			_, err = fd.Write(tc.data)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
		}

		defer os.Remove(tc.file)

		heap := ras.Heap{
			Filename: tc.file,
		}

		err := heap.Load()
		if err != nil {
			if !errEq(err, tc.err) {
				t.Errorf("unexpected error %v, expected %v", err, tc.err)
			}
			return
		}
	}

	testcases := map[string]tcase{
		"1" : {
			file: "test.heap",
			data: nil,
			err: nil,
		},
		"2" : {
			file: "test.heap",
			data: append(ras.Magic[:], []byte{0, 0, 0, 0}...),
		},
		"err magic": {
			file: "test.heap",
			data: []byte{0, 0, 0, 0},
			err: errors.New("unrecognized magic numbers"),
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}

func TestHeapResize (t *testing.T) {
	const (
		verbAlloc = iota
		verbFree
		verbResize
	)

	type manip struct {
		verb int
		val uint
	}

	type tcase struct {
		file string
		manips []manip
		resize uint
		err error
	}

	fn := func(tc tcase, t *testing.T) {
		heap := ras.Heap{Filename:tc.file}
		defer os.Remove(tc.file)

		for _, v := range tc.manips {
			var err error
			switch v.verb {
			case verbAlloc:
				_, err = heap.Alloc(v.val)
			case verbFree:
				err = heap.Free(v.val)
			case verbResize:
				err = heap.Resize(v.val)
			}

			if err != nil {
				t.Errorf("unexpected error (%d) %v", v.verb, err)

				fmt.Println(dumpFile(heap.Filename))
				return
			}
		}

		err := heap.Resize(tc.resize)
		if err != nil {
			if !errEq(err, tc.err) {
				t.Errorf("unexpected error %v, expected %v", err, tc.err)
			}
			return
		}

		stat, err := os.Stat(tc.file)
		if err != nil {
			t.Errorf("unexpected errror %v", err)
			return
		}

		if stat.Size() != int64(tc.resize) {
			t.Errorf("incorrect file size %d, expected %d", stat.Size(), tc.resize)
			return
		}
	}

	testcases := map[string]tcase{
		"1" : {
			file: "test.heap",
			manips: []manip{},
			resize: 10,
			err: nil,
		},
		"2" : {
			file: "test.heap",
			manips: []manip{
				{verbResize, 16},
				{verbAlloc, 2},
			},
			resize: 15,
			err: nil,
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
