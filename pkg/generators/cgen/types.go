package main

import (
	"fmt"
	"io"
	"strings"

	"modernc.org/cc/v4"
)

type cType interface {
	getGoType() string
	getCType() string
	getCGoType() string
	goToCGo(v string) (def string, free string, use string)
	cgoToGo(v string) string
	validate() error
}

type memberType struct {
	name string
	typ  cType
}

func withGoPointers(typ string, pointers uint) string {
	return fmt.Sprintf("%s%s", strings.Repeat("*", int(pointers)), typ)
}

func withCPointers(typ string, isConst bool, pointers uint) string {
	cst := ""
	if isConst {
		cst = "const "
	}
	return fmt.Sprintf("%s%s%s", cst, typ, strings.Repeat("*", int(pointers)))
}

func withGoCast(typ string, pointers uint, v string) string {
	typ = withGoPointers(typ, pointers)
	if pointers > 0 {
		typ = fmt.Sprintf("(%s)", typ)
	}
	return fmt.Sprintf("%s(%s)", typ, v)
}

type baseType struct {
	isConst  bool
	ctype    string
	gotype   string
	pointers uint
}

func (me *baseType) getGoType() string {
	if !strings.HasPrefix(me.ctype, prefixTypes) {
		return withGoPointers(me.gotype, me.pointers)
	}
	return withGoPointers(improveTypename(me.ctype), me.pointers)
}

func (me *baseType) getCGoType() string {
	if me.gotype == "string" {
		return "C.CString"
	}
	if me.ctype == "void*" {
		return "unsafe.Pointer"
	}
	return withGoPointers(fmt.Sprintf("C.%s", me.ctype), me.pointers)
}

func (me *baseType) getCType() string {
	ptr := me.pointers
	if me.gotype == "string" {
		ptr += 1
	}
	return withCPointers(me.ctype, me.isConst, ptr)
}

func (me *baseType) goToCGo(v string) (def string, free string, use string) {
	if me.gotype == "string" {
		varName := fmt.Sprintf("c%s", strings.ReplaceAll(v, ".", ""))
		define := fmt.Sprintf("%s := C.CString(%s)", varName, v)
		return define, fmt.Sprintf("C.free(unsafe.Pointer(%s))", varName), varName
	}
	if me.ctype == "void*" {
		return "", "", v
	}
	if me.pointers == 0 {
		return "", "", withGoCast(fmt.Sprintf("C.%s", me.ctype), me.pointers, v)
	}
	return "", "", fmt.Sprintf("*(%s)(unsafe.Pointer(&%s))", withGoPointers(fmt.Sprintf("C.%s", me.ctype), me.pointers+1), v)
}

func (me *baseType) cgoToGo(v string) string {
	if me.gotype == "string" {
		return fmt.Sprintf("C.GoString(%s)", v)
	}
	if me.ctype == "" {
		return v
	}
	if me.pointers == 0 {
		if !strings.HasPrefix(me.ctype, prefixTypes) {
			return withGoCast(me.gotype, me.pointers, v)
		}
		return withGoCast(improveTypename(me.ctype), me.pointers, v)
	}

	typ := improveTypename(me.ctype)
	if !strings.HasPrefix(me.ctype, prefixTypes) {
		typ = me.gotype
	}
	return fmt.Sprintf("*(%s)(unsafe.Pointer(&%s))", withGoPointers(typ, me.pointers+1), v)
}

func (me *baseType) validate() error {
	if me.ctype == "" {
		return fmt.Errorf("missing ctype for base type (gotype=%q)", me.gotype)
	}
	if me.gotype == "" && me.ctype != "void" {
		return fmt.Errorf("missing gotype for base type (ctype=%q)", me.ctype)
	}
	return nil
}

type enumMemberType struct {
	name  string
	value cc.Value
}

type enumType struct {
	name     string
	gotype   string
	pointers uint
	members  []enumMemberType
}

func (me *enumType) getGoType() string {
	return withGoPointers(improveTypename(me.name), me.pointers)
}

func (me *enumType) getCGoType() string {
	return withGoPointers(fmt.Sprintf("C.%s", me.name), me.pointers)
}

func (me *enumType) getCType() string {
	return withCPointers(me.name, false, me.pointers)
}

func (me *enumType) goToCGo(v string) (string, string, string) {
	return "", "", withGoCast(fmt.Sprintf("C.%s", me.name), me.pointers, v)
}

func (me *enumType) cgoToGo(v string) string {
	return withGoCast(improveTypename(me.name), me.pointers, v)
}

func (me *enumType) validate() error {
	if me.name == "" {
		return fmt.Errorf("missing name for enum")
	}
	if me.gotype == "" {
		return fmt.Errorf("missing gotype for enum %q", me.name)
	}
	for i, member := range me.members {
		if member.name == "" {
			return fmt.Errorf("missing name for enum member %d of %q", i, me.name)
		}
	}
	return nil
}

type structType struct {
	name     string
	pointers uint
	members  []memberType
}

func (me *structType) getGoType() string {
	return withGoPointers(improveTypename(me.name), me.pointers)
}

func (me *structType) getCGoType() string {
	return withGoPointers(fmt.Sprintf("C.%s", me.name), me.pointers)
}

func (me *structType) getCType() string {
	return withCPointers(me.name, false, me.pointers)
}

func (me *structType) goToCGo(v string) (string, string, string) {
	if me.pointers == 0 {
		return "", "", fmt.Sprintf("*(*C.%s)(unsafe.Pointer(&%s))", me.name, v)
	}
	return "", "", withGoCast(fmt.Sprintf("C.%s", me.name), me.pointers, fmt.Sprintf("unsafe.Pointer(%s)", v))
}

func (me *structType) cgoToGo(v string) string {
	if me.pointers == 0 {
		return fmt.Sprintf("*(%s)(unsafe.Pointer(&%s))", me.name, v)
	}
	return withGoCast(improveTypename(me.name), me.pointers, fmt.Sprintf("unsafe.Pointer(%s)", v))
}

func (me *structType) validate() error {
	if me.name == "" {
		return fmt.Errorf("missing name for struct")
	}
	for i, member := range me.members {
		if member.name == "" {
			return fmt.Errorf("missing name for struct member %d of %q", i, me.name)
		}
		err := member.typ.validate()
		if err != nil {
			return fmt.Errorf("invalid member type for %q.%q: %w", me.name, member.name, err)
		}
	}
	return nil
}

type funcType struct {
	name     string
	pointers uint
	ctype    string
	retType  cType
	params   []memberType
}

func (me *funcType) getGoType() string {
	return withGoPointers(improveTypename(me.ctype), me.pointers-1)
}

func (me *funcType) getCGoType() string {
	return withGoPointers(fmt.Sprintf("C.%s", me.ctype), me.pointers-1)
}

func (me *funcType) getCType() string {
	return withCPointers(me.ctype, false, me.pointers-1)
}

func (me *funcType) goToCGo(v string) (string, string, string) {
	return "", "", withGoCast(fmt.Sprintf("C.%s", me.ctype), me.pointers-1, v)
}

func (me *funcType) cgoToGo(v string) string {
	return withGoCast(improveTypename(me.ctype), me.pointers-1, v)
}

func (me *funcType) validate() error {
	for i, param := range me.params {
		if param.name == "" {
			return fmt.Errorf("missing name for function %q parameter %d", me.name, i)
		}
		err := param.typ.validate()
		if err != nil {
			return fmt.Errorf("invalid function %q parameter type for %q: %w", me.name, param.name, err)
		}
	}
	if me.retType != nil {
		err := me.retType.validate()
		if err != nil {
			return fmt.Errorf("invalid return type for %q: %w", me.name, err)
		}
	}
	if me.ctype == "" {
		return fmt.Errorf("missing ctype for function %q", me.name)
	}
	return nil
}

type funcAnonType struct {
	funcType
	name string
}

func (me *funcAnonType) getGoType() string {
	if me.name == "" {
		return "unsafe.Pointer"
	}
	return me.name
}

func (me *funcAnonType) getCGoType() string {
	if me.name == "" {
		return "unsafe.Pointer"
	}
	return me.name
}

func (me *funcAnonType) getCType() string {
	return "void*"
}

func (me *funcAnonType) goToCGo(v string) (string, string, string) {
	if me.name == "" {
		return "", "", v
	}
	return "", "", fmt.Sprintf("(*[0]byte)(%s)", v)
}

func (me *funcAnonType) cgoToGo(v string) string {
	return v
}

func (me *funcAnonType) validate() error {
	return nil
}

func (me *funcAnonType) addAnonTypes(w io.Writer, structName, memberName string) error {
	me.name = fmt.Sprintf("%s%sFn", structName, memberName)
	return writeTemplate(w, "types_struct_anon_fn.go", map[string]interface{}{
		"Name":       me.name,
		"StructName": structName,
		"MemberName": memberName,
	})
}
