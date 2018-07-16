package list

import (
	"testing"
	"os"
	"errors"
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


func TestListOpenFile(t *testing.T) {
	type tcase struct {
		file string
		data []byte
		err  error
	}

	fn := func(tc tcase, t *testing.T) {
		if tc.data != nil {
			fd, err := os.OpenFile(tc.file, os.O_CREATE|os.O_RDWR, 0600)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}

			_, err = fd.Write(tc.data)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
		}

		defer os.Remove(tc.file)

		heap := Allocator{
			Filename: tc.file,
		}

		_, err := heap.openFile()
		if err != nil {
			if !errEq(err, tc.err) {
				t.Errorf("unexpected error %v, expected %v", err, tc.err)
			}
			return
		}
	}

	testcases := map[string]tcase{
		"1": {
			file: "test.heap",
			data: nil,
			err:  nil,
		},
		"2": {
			file: "test.heap",
			data: append(Magic[:], []byte{0, 0, 0, 0}...),
		},
		"err magic": {
			file: "test.heap",
			data: make([]byte, 20),
			err:  errors.New("unrecognized magic numbers"),
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}

