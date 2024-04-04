package gdapi

import (
	"fmt"
	"reflect"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

func findExpected(obj any, expected reflect.Type) (interface{}, error) {
	val := reflect.ValueOf(obj).Elem()

	if expected == nil || val.Type() == expected {
		return val.Interface(), nil
	}

	for i := 0; i < val.NumField(); i += 1 {
		field := val.Field(i)
		if field.Type() == expected {
			return field.Interface(), nil
		}
		if field.Kind() == reflect.Struct {
			res, err := findExpected(field.Addr().Interface(), expected)
			if err == nil {
				return res, nil
			}
		}
	}

	return nil, fmt.Errorf("could not find expected type %s in %s", expected, val.Type())
}

func NativeFromVariantT[T any](va Variant) (*T, error) {
	resRaw, err := NativeFromVariant(va, reflect.TypeFor[T]())
	if err != nil {
		return nil, err
	}
	res, cast := resRaw.(T)
	if !cast {
		return nil, fmt.Errorf("could not cast %v to %s", resRaw, reflect.TypeFor[T]())
	}
	return &res, nil
}

func NativeFromVariantPtr(vaRaw gdc.ConstVariantPtr, expected reflect.Type) (any, error) {
	va := NewVariantWithC(vaRaw)
	return NativeFromVariant(*va, expected)
}

func NativeFromVariant(va Variant, expected reflect.Type) (interface{}, error) {
	switch va.Type() {
	case gdc.VariantTypeNil:
		if expected != nil && expected.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("cannot convert nil to %s", expected.Kind())
		}
		return reflect.Zero(expected), nil
	case gdc.VariantTypeBool:
		if expected != nil && expected.Kind() != reflect.Bool {
			return nil, fmt.Errorf("cannot convert bool to %s", expected.Kind())
		}
		b, err := va.AsBool()
		if err != nil {
			return nil, err
		}
		return bool(b.Get()), nil
	case gdc.VariantTypeInt:
		i, err := va.AsInt()
		if err != nil {
			return nil, err
		}
		kind := reflect.Int64
		if expected != nil {
			kind = expected.Kind()
		}
		switch kind {
		case reflect.Int:
			return int(i.Get()), nil
		case reflect.Int8:
			return int8(i.Get()), nil
		case reflect.Int16:
			return int16(i.Get()), nil
		case reflect.Int32:
			return int32(i.Get()), nil
		case reflect.Int64:
			return int64(i.Get()), nil
		case reflect.Uint:
			return uint(i.Get()), nil
		case reflect.Uint8:
			return uint8(i.Get()), nil
		case reflect.Uint16:
			return uint16(i.Get()), nil
		case reflect.Uint32:
			return uint32(i.Get()), nil
		case reflect.Uint64:
			return uint64(i.Get()), nil
		}
		return nil, fmt.Errorf("cannot convert int to %s", kind)
	case gdc.VariantTypeFloat:
		f, err := va.AsFloat()
		if err != nil {
			return nil, err
		}
		kind := reflect.Float64
		if expected != nil {
			kind = expected.Kind()
		}
		switch kind {
		case reflect.Float32:
			return float32(f.Get()), nil
		case reflect.Float64:
			return float64(f.Get()), nil
		}
		return nil, fmt.Errorf("cannot convert float to %s", kind)
	case gdc.VariantTypeString:
		if expected != nil && expected.Kind() != reflect.String {
			return nil, fmt.Errorf("cannot convert string to %s", expected.Kind())
		}
		s, err := va.AsString()
		if err != nil {
			return nil, err
		}
		return string(s.String()), nil
	case gdc.VariantTypeObject:
		if expected != nil && expected.Kind() != reflect.Struct {
			return nil, fmt.Errorf("cannot convert object to %s", expected.Kind())
		}
		obj, err := va.AsObject()
		if err != nil {
			return nil, err
		}
		return findExpected(obj, expected)
	default:
		if expected != nil && expected.Kind() != reflect.Struct && expected.Kind() != reflect.Interface {
			return nil, fmt.Errorf("cannot convert bclass %v to %s", va.Type(), expected.Kind())
		}
		bclass, err := va.AsBClass()
		if err != nil {
			return nil, err
		}
		if expected != nil && expected.Kind() == reflect.Interface {
			return bclass, nil
		}
		return reflect.ValueOf(bclass).Elem().Interface(), nil
	}
}
