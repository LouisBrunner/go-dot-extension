package main

import (
	"fmt"

	"modernc.org/cc/v4"
)

const bitsPerWord = 32 << (^uint(0) >> 63)

// These are heavily derived from github.com/xlab/c-for-go/blob/master/translator/ast_walker.go

func transformType(name string, typ cc.Type, isConst bool) (cType, error) {
	typedef := name
	if typ.Typedef() != nil {
		var err error
		typedef, err = identifierOf(typ.Typedef().DirectDeclarator)
		if err != nil {
			return nil, err
		}
	}

	spec := &baseType{
		ctype:   typedef,
		isConst: isConst,
	}

	for pointerType, ok := typ.(*cc.PointerType); ok; pointerType, ok = typ.(*cc.PointerType) {
		next := pointerType.Elem()
		if next.Kind() == cc.Void {
			spec.gotype = fmt.Sprintf("C.%s", typedef)
			if typedef == "" {
				spec.gotype = "unsafe.Pointer"
				if pointerType.Typedef() != nil {
					typedef, err := identifierOf(pointerType.Typedef().DirectDeclarator)
					if err != nil {
						return nil, err
					}
					spec.ctype = typedef
				}
			}
			return spec, nil
		}
		typ = next
		spec.pointers++
	}

	switch typ.Kind() {
	case cc.Void:
	case cc.Char:
		if spec.pointers == 1 {
			spec.pointers = 0
			spec.gotype = "string"
		} else {
			spec.gotype = "byte"
		}
	case cc.SChar:
		spec.gotype = "int8"
	case cc.UChar:
		spec.gotype = "uint8"
	case cc.Short:
		spec.gotype = "int16"
	case cc.UShort:
		spec.gotype = "uint16"
	case cc.Int:
		spec.gotype = "int"
	case cc.UInt:
		spec.gotype = "uint"
	case cc.Long:
		if bitsPerWord == 32 {
			spec.gotype = "int32"
		} else {
			spec.gotype = "int64"
		}
	case cc.ULong:
		if bitsPerWord == 32 {
			spec.gotype = "uint32"
		} else {
			spec.gotype = "uint64"
		}
	case cc.LongLong:
		spec.gotype = "int64"
	case cc.ULongLong:
		spec.gotype = "uint64"
	case cc.Float, cc.Float32, cc.Float32x:
		spec.gotype = "float32"
	case cc.Double, cc.Float64, cc.Float64x:
		spec.gotype = "float64"
	case cc.Bool:
		spec.gotype = "bool"
	case cc.ComplexFloat:
		spec.gotype = "complex64"
	case cc.ComplexDouble:
		spec.gotype = "complex128"
	case cc.Enum:
		return transformEnum(typ, spec)
	case cc.Struct:
		return transformStruct(typ, spec)
	case cc.Function:
		return transformFunc(typ, spec, isConst)
	default:
		return nil, fmt.Errorf("unsupported type kind %d", typ.Kind())
	}

	return spec, nil
}

func transformEnum(typ cc.Type, base *baseType) (*enumType, error) {
	enumTyp, ok := typ.(*cc.EnumType)
	if !ok {
		return nil, fmt.Errorf("not actually an enum: %T", typ)
	}

	spec := &enumType{
		name:     typ.Typedef().Name(),
		gotype:   "uint32",
		pointers: base.pointers,
	}

	for _, en := range enumTyp.Enumerators() {
		spec.members = append(spec.members, enumMemberType{
			name:  en.Token.SrcStr(),
			value: en.Value(),
		})
	}

	return spec, nil
}

func transformStruct(typ cc.Type, base *baseType) (*structType, error) {
	structTyp, ok := typ.(*cc.StructType)
	if !ok {
		return nil, fmt.Errorf("not actually a struct: %T", typ)
	}

	tag := structTyp.Tag()
	name := tag.SrcStr()
	if typ.Typedef() != nil {
		name = typ.Typedef().Name()
	}

	spec := &structType{
		name:     name,
		pointers: base.pointers,
	}

	for i := 0; i < structTyp.NumFields(); i++ {
		m := structTyp.FieldByIndex(i)
		decl := m.Declarator()
		if decl == nil {
			return nil, fmt.Errorf("struct field %q has no declarator", m.Name())
		}
		subType, err := transformType("", m.Type(), decl.IsConst())
		if err != nil {
			return nil, err
		}
		spec.members = append(spec.members, memberType{
			name: memberName(i, m),
			typ:  subType,
		})
	}
	return spec, nil
}

func transformFunc(typ cc.Type, base *baseType, isConst bool) (cType, error) {
	funcTyp, ok := typ.(*cc.FunctionType)
	if !ok {
		return nil, fmt.Errorf("not actually a function: %T", typ)
	}

	if base.ctype == "" {
		return &funcAnonType{}, nil
	}

	spec := &funcType{
		ctype:    base.ctype,
		pointers: base.pointers,
	}

	ret := funcTyp.Result()
	if ret != nil && ret.Kind() != cc.Void {
		var name string
		var isReturnConst bool
		typedef := ret.Typedef()
		if typedef != nil {
			name = typedef.Name()
			isReturnConst = typedef.IsConst()
		}
		var err error
		spec.retType, err = transformType(name, ret, isConst || isReturnConst)
		if err != nil {
			return nil, err
		}
	}

	for i, p := range funcTyp.Parameters() {
		if p.Type().Kind() == cc.Void {
			continue
		}

		var isConst bool
		decl := p.Declarator
		if decl != nil {
			isConst = decl.IsConst()
		}
		subType, err := transformType("", p.Type(), isConst)
		if err != nil {
			return nil, err
		}
		spec.params = append(spec.params, memberType{
			name: paramName(i, p),
			typ:  subType,
		})
	}
	return spec, nil
}
