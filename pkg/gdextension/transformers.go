package gdextension

import (
	"fmt"
	"reflect"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

func bclassForType(typ reflect.Type) (gdc.VariantType, bool) {
	bclazz, isBclass := reflect.New(typ).Interface().(gdapi.BClass)
	if isBclass {
		return bclazz.Type(), true
	}
	return gdc.VariantTypeNil, false
}

func (me *extension) objectFromPtr(expected reflect.Type, obj gdc.ObjectPtr) (interface{}, error) {
	dobj, err := gdapi.DObjectFromPtr(expected.String(), obj)
	if err == nil {
		return dobj, nil
	}
	regClass, isReg := me.registered[expected.String()]
	if !isReg {
		return nil, err
	}
	reg := regClass.constructor()
	reg.SetBaseObject(obj)
	return reg, nil
}

func (me *extension) reflectFromType(val gdc.ConstTypePtr, expected reflect.Type) (reflect.Value, error) {
	switch expected.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(bool(*(*bool)(val))), nil
	case reflect.Int:
		return reflect.ValueOf(int(*(*int)(val))), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(*(*int8)(val))), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(*(*int16)(val))), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(*(*int32)(val))), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(*(*int64)(val))), nil
	case reflect.Uint:
		return reflect.ValueOf(uint(*(*uint)(val))), nil
	case reflect.Uint8:
		return reflect.ValueOf(uint8(*(*uint8)(val))), nil
	case reflect.Uint16:
		return reflect.ValueOf(uint16(*(*uint16)(val))), nil
	case reflect.Uint32:
		return reflect.ValueOf(uint32(*(*uint32)(val))), nil
	case reflect.Uint64:
		return reflect.ValueOf(uint64(*(*uint64)(val))), nil
	case reflect.Float32:
		return reflect.ValueOf(float32(*(*float32)(val))), nil
	case reflect.Float64:
		return reflect.ValueOf(float64(*(*float64)(val))), nil
	case reflect.String:
		str := gdapi.StringFromPtr(val)
		return reflect.ValueOf(string(str.String())), nil
	case reflect.Struct, reflect.Interface:
		switch reflect.ValueOf(val).Elem().Interface().(type) {
		case gdapi.Object:
			obj, err := me.objectFromPtr(expected, gdc.ObjectPtr(val))
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			return reflect.ValueOf(obj), nil
		case gdapi.Variant:
			return reflect.ValueOf(gdapi.NewVariantWith(gdc.VariantPtr(val))), nil
		case gdapi.BClass:
			bclassType, _ := bclassForType(expected)
			bclass, err := gdapi.BClassFromPtr(bclassType, val)
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			return reflect.ValueOf(bclass), nil
		}
	case reflect.Pointer:
		typ, err := me.reflectFromType(val, expected.Elem())
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		return typ.Addr(), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("unsupported type: %s", expected.Kind())
}

func typeFromReflect(val reflect.Value) (gdc.TypePtr, error) {
	// TODO: finish
	// TODO: this will leak
	switch val.Kind() {
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
	case reflect.Array, reflect.Slice:
	case reflect.Struct, reflect.Interface:
	case reflect.Pointer:
		return typeFromReflect(val.Elem())
	}
	return nil, fmt.Errorf("unsupported type: %s", val.Kind())
}

func nativeFromVariant(vaRaw gdc.ConstVariantPtr, expected reflect.Type) (any, error) {
	va := gdapi.NewVariantWithC(vaRaw)
	switch va.Type() {
	case gdc.VariantTypeNil:
		if expected.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("cannot convert nil to %s", expected.Kind())
		}
		return reflect.Zero(expected), nil
	case gdc.VariantTypeBool:
		if expected.Kind() != reflect.Bool {
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
		switch expected.Kind() {
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
		return nil, fmt.Errorf("cannot convert int to %s", expected.Kind())
	case gdc.VariantTypeFloat:
		f, err := va.AsFloat()
		if err != nil {
			return nil, err
		}
		switch expected.Kind() {
		case reflect.Float32:
			return float32(f.Get()), nil
		case reflect.Float64:
			return float64(f.Get()), nil
		}
		return nil, fmt.Errorf("cannot convert float to %s", expected.Kind())
	case gdc.VariantTypeString:
		if expected.Kind() != reflect.String {
			return nil, fmt.Errorf("cannot convert string to %s", expected.Kind())
		}
		s, err := va.AsString()
		if err != nil {
			return nil, err
		}
		return string(s.String()), nil
	case gdc.VariantTypeObject:
		if expected.Kind() != reflect.Struct {
			return nil, fmt.Errorf("cannot convert object to %s", expected.Kind())
		}
		obj, err := va.AsObject()
		if err != nil {
			return nil, err
		}
		return *obj, nil
	default:
		if expected.Kind() != reflect.Struct {
			return nil, fmt.Errorf("cannot convert bclass to %s", expected.Kind())
		}
		bclass, err := va.AsBClass()
		if err != nil {
			return nil, err
		}
		return reflect.ValueOf(bclass).Elem().Interface(), nil
	}
}

func variantFromReflect(val reflect.Value) (gdc.VariantPtr, error) {
	// TODO: this leaks
	switch val.Kind() {
	case reflect.Bool:
		return gdapi.NewBoolFromBool(val.Bool()).AsVariant().AsPtr(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return gdapi.NewIntFromInt(val.Int()).AsVariant().AsPtr(), nil
	case reflect.Float32, reflect.Float64:
		return gdapi.NewFloatFromFloat32(val.Float()).AsVariant().AsPtr(), nil
	case reflect.String:
		return gdapi.StringFromStr(val.String()).AsVariant().AsPtr(), nil
	case reflect.Array, reflect.Slice:
		arr := gdapi.NewArray()
		for i := 0; i < val.Len(); i++ {
			va, err := variantFromReflect(val.Index(i))
			if err != nil {
				return nil, err
			}
			arr.PushBack(*gdapi.NewVariantWith(gdc.VariantPtr(va)))
		}
		return arr.AsVariant().AsPtr(), nil
	case reflect.Struct:
		val = val.Addr()
		fallthrough
	case reflect.Interface:
		bclazz, isBclass := val.Interface().(gdapi.BClass)
		if isBclass {
			return bclazz.AsVariant().AsPtr(), nil
		}
	case reflect.Ptr:
		if val.IsNil() {
			return nil, nil
		}
		return variantFromReflect(val.Elem())
	}
	return nil, fmt.Errorf("unsupported type: %s (%+v)", val.Type().String(), val.Interface())
}

func typeToVariant(val reflect.Type) (gdc.VariantType, error) {
	switch val.Kind() {
	case reflect.Bool:
		return gdc.VariantTypeBool, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return gdc.VariantTypeInt, nil
	case reflect.Float32, reflect.Float64:
		return gdc.VariantTypeFloat, nil
	case reflect.String:
		return gdc.VariantTypeString, nil
	case reflect.Array, reflect.Slice:
		return gdc.VariantTypeArray, nil
	case reflect.Struct, reflect.Interface:
		bclazz, isBclass := bclassForType(val)
		if isBclass {
			return bclazz, nil
		}
		return gdc.VariantTypeObject, nil
	case reflect.Pointer:
		return typeToVariant(val.Elem())
	}
	return gdc.VariantTypeNil, fmt.Errorf("unsupported type %s", val.Kind())
}

func typeToVariantNClass(val reflect.Type) (gdc.VariantType, gdc.StringNamePtr, error) {
	va, err := typeToVariant(val)
	if err != nil {
		return gdc.VariantTypeNil, nil, err
	}
	if va != gdc.VariantTypeObject {
		return va, gdapi.NewStringName().AsPtr(), nil
	}
	inst, ok := reflect.New(val).Interface().(Class)
	if !ok {
		return gdc.VariantTypeNil, nil, fmt.Errorf("unsupported object type %s (%T)", val, inst)
	}
	return va, gdapi.StringNameFromStr(inst.BaseClass()).AsPtr(), nil
}
