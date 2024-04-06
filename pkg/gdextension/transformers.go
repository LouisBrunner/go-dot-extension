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
		switch reflect.New(expected).Interface().(type) {
		case *gdapi.Object:
			obj, err := me.objectFromPtr(expected, gdc.ObjectPtr(val))
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			return reflect.ValueOf(obj).Elem(), nil
		case *gdapi.Variant:
			return reflect.ValueOf(gdapi.NewVariantWith(gdc.VariantPtr(val))).Elem(), nil
		case gdapi.BClass:
			bclassType, _ := bclassForType(expected)
			bclass, err := gdapi.BClassFromPtr(bclassType, val)
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			return reflect.ValueOf(bclass).Elem(), nil
		}
	case reflect.Pointer:
		typ, err := me.reflectFromType(val, expected.Elem())
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		return typ.Addr(), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("unsupported type: %s", expected)
}

func typeFromReflect(val reflect.Value) (gdc.TypePtr, error) {
	panic("unimplemented") // TODO: finish
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

func variantFromReflect(val reflect.Value) (*gdapi.Variant, error) {
	switch val.Kind() {
	case reflect.Bool:
		return gdapi.NewBoolFromBool(val.Bool()).AsVariant(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return gdapi.NewIntFromInt(val.Int()).AsVariant(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return gdapi.NewIntFromInt(int64(val.Uint())).AsVariant(), nil
	case reflect.Float32, reflect.Float64:
		return gdapi.NewFloatFromFloat32(val.Float()).AsVariant(), nil
	case reflect.String:
		return gdapi.StringFromStr(val.String()).AsVariant(), nil
	case reflect.Array, reflect.Slice:
		arr := gdapi.NewArray()
		for i := 0; i < val.Len(); i++ {
			va, err := variantFromReflect(val.Index(i))
			if err != nil {
				return nil, fmt.Errorf("error marshalling array element %d: %w", i, err)
			}
			arr.PushBack(*va)
		}
		return arr.AsVariant(), nil
	case reflect.Struct:
		val = val.Addr()
		fallthrough
	case reflect.Interface:
		bclazz, isBclass := val.Interface().(gdapi.BClass)
		if isBclass {
			return bclazz.AsVariant(), nil
		}
	case reflect.Ptr:
		if val.IsNil() {
			return nil, nil
		}
		return variantFromReflect(val.Elem())
	}
	return nil, fmt.Errorf("unsupported type: %s (%+v)", val.Type().String(), val.Interface())
}

func typeToVariant(val reflect.Type) (gdc.VariantType, bool, error) {
	switch val.Kind() {
	case reflect.Bool:
		return gdc.VariantTypeBool, false, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return gdc.VariantTypeInt, false, nil
	case reflect.Float32, reflect.Float64:
		return gdc.VariantTypeFloat, false, nil
	case reflect.String:
		return gdc.VariantTypeString, false, nil
	case reflect.Array, reflect.Slice:
		return gdc.VariantTypeArray, false, nil
	case reflect.Struct, reflect.Interface:
		bclazz, isBclass := bclassForType(val)
		if isBclass {
			return bclazz, true, nil
		}
		return gdc.VariantTypeObject, true, nil
	case reflect.Pointer:
		return typeToVariant(val.Elem())
	}
	return gdc.VariantTypeNil, false, fmt.Errorf("unsupported type %s", val.Kind())
}

func typeToVariantNClass(val reflect.Type) (gdc.VariantType, bool, gdc.StringNamePtr, error) {
	va, isBclass, err := typeToVariant(val)
	if err != nil {
		return gdc.VariantTypeNil, false, nil, err
	}
	if va != gdc.VariantTypeObject {
		return va, isBclass, gdapi.NewStringName().AsPtr(), nil
	}
	instRaw := reflect.New(val).Interface()
	inst, ok := instRaw.(Class)
	if !ok {
		return gdc.VariantTypeNil, false, nil, fmt.Errorf("unsupported object type %s (%T)", val, inst)
	}
	name, err := nameForClass(val, instRaw)
	if err != nil {
		return gdc.VariantTypeNil, false, nil, err
	}
	return va, true, gdapi.StringNameFromStr(name).AsPtr(), nil
}
