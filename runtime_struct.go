package ras

import (
	"reflect"
	"errors"
)

func (r Runtime) refStruct(v interface{}) (reflect.Value, error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return reflect.Value{}, errors.New("v must be a pointer to a struct")
	}
	return val, nil
}

func (r Runtime) NewStruct(v interface{}) error {
	val, err := r.refStruct(v)
	if err != nil {
		return err
	}

	var size uint
	for i := 0; i<val.NumField(); i++ {
		vari, ok := val.Field(i).Interface().(Variable)
		if ok {
			size += vari.SizeOf()
		}
	}

	addr, err := r.Alloc(size)
	if err != nil {
		return err
	}

	return r.Struct(addr, v)
}

func (r Runtime) Struct(addr uint, v interface{}) error {
	val, err := r.refStruct(v)
	if err != nil {
		return err
	}

	for i := 0; i < val.NumField(); i++ {

		field := val.Field(i)

		vari, ok := field.Interface().(Variable)
		if !ok {
			continue
		}

		vari.init(addr, r)
		addr += vari.SizeOf()
	}

	return nil
}
