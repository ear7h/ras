// +build proto

package ras_test

import (
	"testing"
	"os"
	"fmt"
	"strings"
	"path/filepath"

	"github.com/ear7h/ras"
	"github.com/ear7h/ras/list"
)

const (
	MaxUint8  = uint8(0xff)
	MaxUint16 = uint16(0xffff)
	MaxUint32 = uint32(0xffffffff)
	MaxUint64 = uint64(0xffffffffffffffff)
	MaxUint = uint(MaxUint64)

	MaxInt8  = int8(0xff / 2)
	MaxInt16 = int16(0xffff / 2)
	MaxInt32 = int32(0xffffffff / 2)
	MaxInt64 = int64(0xffffffffffffffff / 2)
	MaxInt = int(MaxInt64)
)

func filename(s string) string {
	return strings.Replace(s, string(filepath.Separator), "_", -1)
}

//go:proto V=Uints,Ints K=V.lower
func TestV(t *testing.T) {
	type tcase struct {
		vals []K
	}

	fn := func(tc tcase, t *testing.T) {
		t.Parallel()

		fname := filename(t.Name()) + ".ras"
		defer func() {
			err := os.Remove(fname)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
		}()

		rt := ras.Runtime{
			Allocator: &list.Allocator{
				Filename: fname,
			},
		}

		err := rt.Resize(1024)
		if err != nil {
			t.Fatalf("unexpeceted error %v", err)
		}

		addrs := make([]uint, len(tc.vals))
		for i, v := range tc.vals {
			variable, err := rt.NewV()
			if err != nil {
				t.Fatalf("unexpecet error %v", err)
			}

			err = variable.Set(v)
			if err != nil {
				t.Fatalf("unexpecet error %v", err)
			}

			addrs[i] = variable.Addr()
		}

		fmt.Println("addrs: ", addrs)
		for i, v := range tc.vals {
			variable := rt.V(addrs[i])
			val, err := variable.Get()
			if err != nil {
				t.Fatalf("unexpecet error %v", err)
			}

			if val != v {
				t.Fatalf("unexpected value %v, expected %v", val, v)
			}
		}
	}

	testcases := map[string]tcase {
		"1": {
			vals: []K{0},
		},
		"2": {
			vals: []K{1},
		},
		"3": {
			vals: []K{1, 4},
		},
		"4": {
			vals: []K{1, MaxV},
		},
		"5": {
			vals: []K{0, 5, 10, 4, MaxV},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
//go:proto clear