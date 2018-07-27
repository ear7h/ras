
package ras_test

import (
	"testing"
	"github.com/ear7h/ras"
	"github.com/ear7h/ras/list"
	"os"
	"fmt"
	"strings"
	"path/filepath"
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
func TestUint(t *testing.T) {
	type tcase struct {
		vals []uint
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
			variable, err := rt.NewUint()
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
			variable := rt.Uint(addrs[i])
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
			vals: []uint{0},
		},
		"2": {
			vals: []uint{1},
		},
		"3": {
			vals: []uint{1, 4},
		},
		"4": {
			vals: []uint{1, MaxUint},
		},
		"5": {
			vals: []uint{0, 5, 10, 4, MaxUint},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestUint8(t *testing.T) {
	type tcase struct {
		vals []uint8
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
			variable, err := rt.NewUint8()
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
			variable := rt.Uint8(addrs[i])
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
			vals: []uint8{0},
		},
		"2": {
			vals: []uint8{1},
		},
		"3": {
			vals: []uint8{1, 4},
		},
		"4": {
			vals: []uint8{1, MaxUint8},
		},
		"5": {
			vals: []uint8{0, 5, 10, 4, MaxUint8},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestUint16(t *testing.T) {
	type tcase struct {
		vals []uint16
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
			variable, err := rt.NewUint16()
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
			variable := rt.Uint16(addrs[i])
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
			vals: []uint16{0},
		},
		"2": {
			vals: []uint16{1},
		},
		"3": {
			vals: []uint16{1, 4},
		},
		"4": {
			vals: []uint16{1, MaxUint16},
		},
		"5": {
			vals: []uint16{0, 5, 10, 4, MaxUint16},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestUint32(t *testing.T) {
	type tcase struct {
		vals []uint32
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
			variable, err := rt.NewUint32()
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
			variable := rt.Uint32(addrs[i])
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
			vals: []uint32{0},
		},
		"2": {
			vals: []uint32{1},
		},
		"3": {
			vals: []uint32{1, 4},
		},
		"4": {
			vals: []uint32{1, MaxUint32},
		},
		"5": {
			vals: []uint32{0, 5, 10, 4, MaxUint32},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestUint64(t *testing.T) {
	type tcase struct {
		vals []uint64
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
			variable, err := rt.NewUint64()
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
			variable := rt.Uint64(addrs[i])
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
			vals: []uint64{0},
		},
		"2": {
			vals: []uint64{1},
		},
		"3": {
			vals: []uint64{1, 4},
		},
		"4": {
			vals: []uint64{1, MaxUint64},
		},
		"5": {
			vals: []uint64{0, 5, 10, 4, MaxUint64},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestInt(t *testing.T) {
	type tcase struct {
		vals []int
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
			variable, err := rt.NewInt()
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
			variable := rt.Int(addrs[i])
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
			vals: []int{0},
		},
		"2": {
			vals: []int{1},
		},
		"3": {
			vals: []int{1, 4},
		},
		"4": {
			vals: []int{1, MaxInt},
		},
		"5": {
			vals: []int{0, 5, 10, 4, MaxInt},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestInt8(t *testing.T) {
	type tcase struct {
		vals []int8
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
			variable, err := rt.NewInt8()
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
			variable := rt.Int8(addrs[i])
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
			vals: []int8{0},
		},
		"2": {
			vals: []int8{1},
		},
		"3": {
			vals: []int8{1, 4},
		},
		"4": {
			vals: []int8{1, MaxInt8},
		},
		"5": {
			vals: []int8{0, 5, 10, 4, MaxInt8},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestInt16(t *testing.T) {
	type tcase struct {
		vals []int16
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
			variable, err := rt.NewInt16()
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
			variable := rt.Int16(addrs[i])
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
			vals: []int16{0},
		},
		"2": {
			vals: []int16{1},
		},
		"3": {
			vals: []int16{1, 4},
		},
		"4": {
			vals: []int16{1, MaxInt16},
		},
		"5": {
			vals: []int16{0, 5, 10, 4, MaxInt16},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestInt32(t *testing.T) {
	type tcase struct {
		vals []int32
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
			variable, err := rt.NewInt32()
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
			variable := rt.Int32(addrs[i])
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
			vals: []int32{0},
		},
		"2": {
			vals: []int32{1},
		},
		"3": {
			vals: []int32{1, 4},
		},
		"4": {
			vals: []int32{1, MaxInt32},
		},
		"5": {
			vals: []int32{0, 5, 10, 4, MaxInt32},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
func TestInt64(t *testing.T) {
	type tcase struct {
		vals []int64
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
			variable, err := rt.NewInt64()
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
			variable := rt.Int64(addrs[i])
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
			vals: []int64{0},
		},
		"2": {
			vals: []int64{1},
		},
		"3": {
			vals: []int64{1, 4},
		},
		"4": {
			vals: []int64{1, MaxInt64},
		},
		"5": {
			vals: []int64{0, 5, 10, 4, MaxInt64},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
//go:proto clear
