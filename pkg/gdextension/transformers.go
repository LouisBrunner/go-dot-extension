package gdextension

import (
	"fmt"
	"reflect"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
	"github.com/iancoleman/strcase"
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

	name, err := nameForClass(expected, reflect.New(expected).Elem().Interface())
	if err != nil {
		return nil, err
	}
	regClass, isReg := me.registered[name]
	if !isReg {
		return nil, fmt.Errorf("no class registered for %s", name)
	}
	objRef := gdapi.ObjectFromPtr(obj)
	instanceID := objRef.GetInstanceId()

	inst, err := regClass.lookupInstance(uint64(instanceID))
	if err != nil {
		return nil, err
	}
	return inst.instance, nil
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
		case Class:
			valObj := *(*gdc.ObjectPtr)(val)
			obj, err := me.objectFromPtr(expected, gdc.ObjectPtr(valObj))
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

func assignTypeFromReflect(pRet gdc.TypePtr, val reflect.Value) error {
	switch val.Kind() {
	case reflect.Bool:
		v := val.Bool()
		vInt := 0
		if v {
			vInt = 1
		}
		*(*uint8)(pRet) = uint8(vInt)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := val.Int()
		*(*int64)(pRet) = v
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := val.Uint()
		*(*uint64)(pRet) = v
		return nil
	case reflect.Float32, reflect.Float64:
		v := val.Float()
		*(*float64)(pRet) = v
		return nil
	case reflect.String:
		v := val.String()
		str := gdapi.StringFromStr(v)
		str.ToTypePtr(pRet)
		return nil
	case reflect.Struct:
		addrV := reflect.New(val.Type())
		addrV.Elem().Set(val)
		v := addrV.Interface()
		bclazz, isBclass := v.(gdapi.BClass)
		obj, isObj := v.(Class)
		switch {
		case isBclass:
			bclazz.ToTypePtr(pRet)
			return nil
		case isObj:
			*(*gdc.ObjectPtr)(pRet) = (gdc.ObjectPtr)(obj.AsTypePtr()) // FIXME: silly cast
			return nil
		}
	case reflect.Interface:
		v := val.Interface()
		if bclazz, isBclass := v.(gdapi.BClass); isBclass {
			bclazz.ToTypePtr(pRet)
			return nil
		} else if obj, isObj := v.(Class); isObj {
			*(*gdc.ObjectPtr)(pRet) = (gdc.ObjectPtr)(obj.AsTypePtr()) // FIXME: silly cast
			return nil
		}
	case reflect.Pointer:
		return assignTypeFromReflect(pRet, val.Elem())
	}
	return fmt.Errorf("unsupported type: %s (%T)", val.Kind(), val.Interface())
}

func variantFromReflect(val reflect.Value) (*gdapi.Variant, error) {
	return variantFromReflectWithFallback(val, nil)
}

func variantFromReflectWithFallback(val reflect.Value, onUnsupported func(val reflect.Value) (*gdapi.Variant, error)) (*gdapi.Variant, error) {
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
			va, err := variantFromReflectWithFallback(val.Index(i), onUnsupported)
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
		return variantFromReflectWithFallback(val.Elem(), onUnsupported)
	}

	if onUnsupported != nil {
		return onUnsupported(val)
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

func typeToVariantNClass(val reflect.Type, enums map[string]string) (gdc.VariantType, bool, gdc.StringNamePtr, gdapi.PropertyUsageFlags, error) {
	va, isBclass, err := typeToVariant(val)
	if err != nil {
		return gdc.VariantTypeNil, false, nil, 0, err
	}
	if va != gdc.VariantTypeObject {
		usage := gdapi.PropertyUsageDefault
		var className gdapi.StringName
		if _, ok := enums[val.Name()]; va == gdc.VariantTypeInt && ok {
			className = *gdapi.StringNameFromStr(strcase.ToCamel(val.Name()))
			usage |= gdapi.PropertyUsageClassIsEnum
		} else {
			className = *gdapi.NewStringName()
		}
		return va, isBclass, className.AsPtr(), usage, nil
	}
	instRaw := reflect.New(val).Interface()
	inst, ok := instRaw.(Class)
	if !ok {
		return gdc.VariantTypeNil, false, nil, 0, fmt.Errorf("unsupported object type %s (%T)", val, inst)
	}
	name, err := nameForClass(val, instRaw)
	if err != nil {
		return gdc.VariantTypeNil, false, nil, 0, err
	}
	return va, true, gdapi.StringNameFromStr(name).AsPtr(), gdapi.PropertyUsageDefault, nil
}
