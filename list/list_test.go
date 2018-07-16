package list_test

import (
	"testing"
	"os"
	"errors"
	"io/ioutil"
	"fmt"
	"reflect"

	"github.com/ear7h/ras/list"
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

func TestListResize(t *testing.T) {
	const (
		verbAlloc  = iota
		verbFree
		verbResize
	)

	type manip struct {
		verb int
		val  uint
	}

	type tcase struct {
		manips []manip
		resize uint
		err    error
	}

	fn := func(tc tcase, t *testing.T) {
		heap := list.Allocator{Filename: "test.heap"}
		defer os.Remove("test.heap")

		var err error

		for _, v := range tc.manips {

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

		err = heap.Resize(tc.resize)
		if err != nil {
			if !errEq(err, tc.err) {
				t.Errorf("unexpected error %v, expected %v", err, tc.err)
			}
			return
		}

		stat, err := os.Stat("test.heap")
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
		"1": {
			manips: []manip{},
			resize: 10,
			err:    nil,
		},
		"2": {
			manips: []manip{
				{verbResize, 18},
				{verbAlloc, 2},
			},
			resize: 16,
			err:    nil,
		},
		"err allocated": {
			manips: []manip{
				{verbResize, 18},
				{verbAlloc, 2},
			},
			resize: 12,
			err:    errors.New("shrink over allocared area 15"),
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}

func TestAllocAndFree(t *testing.T) {
	const (
		verbAlloc  = iota
		verbFree
		verbResize
	)

	type manip struct {
		verb int
		val  uint
	}

	type tcase struct {
		manips []manip
		res    []byte
		err    error
	}

	fn := func(tc tcase, t *testing.T) {
		heap := list.Allocator{Filename: "test.heap"}
		defer os.Remove("test.heap")

		var err error
		var i int
		var v manip

		for i, v = range tc.manips {

			switch v.verb {
			case verbAlloc:
				_, err = heap.Alloc(v.val)
			case verbFree:
				err = heap.Free(v.val)
			case verbResize:
				err = heap.Resize(v.val)
			}

			if err != nil {
				break
			}
		}
		res := dumpFile("test.heap")

		if !errEq(err, tc.err) {
			fmt.Println(res)
			t.Errorf("unexpected error in manip %d %v, expected %v", i, err, tc.err)
			return
		}

		if !reflect.DeepEqual(res, tc.res) {
			t.Errorf("incorrect result in file \n%v\nexpected\n%v", res, tc.res)
			return
		}
	}

	testcases := map[string]tcase{
		"simple alloc": {
			manips: []manip{
				{verbResize, 16},
				{verbAlloc, 2},
			},
			res: append(list.Magic[:],
				14, 0, 0, 0, // head ptr
				2, 0, 0, 0, // header.len
				0, 0, 0, 0, // header.next
				0, 0), // val
			err: nil,
		},
		"simple free": {
			manips: []manip{
				{verbResize, 16},
				{verbAlloc, 2},
				{verbFree, 14},
			},
			res: append(list.Magic[:],
			0, 0, 0, 0, // head ptr
			2, 0, 0, 0, // header.len
			0, 0, 0, 0, // header.next
			0, 0), // val
			err: nil,
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
