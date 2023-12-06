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
	goToCGo(v string) string
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
	return withGoPointers(fmt.Sprintf("C.%s", me.ctype), me.pointers)
}

func (me *baseType) getCType() string {
	return withCPointers(me.ctype, me.isConst, me.pointers)
}

func (me *baseType) goToCGo(v string) string {
	if me.gotype == "string" {
		return fmt.Sprintf("C.CString(%s)", v) // TODO: leaks memory
	}
	if me.ctype == "" {
		return v
	}
	return withGoCast(fmt.Sprintf("C.%s", me.ctype), me.pointers, v)
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

func (me *enumType) goToCGo(v string) string {
	return withGoCast(fmt.Sprintf("C.%s", me.name), me.pointers, v)
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

func (me *structType) goToCGo(v string) string {
	if me.pointers == 0 {
		return fmt.Sprintf("*(*C.%s)(unsafe.Pointer(&%s))", me.name, v)
	}
	return withGoCast(fmt.Sprintf("C.%s", me.name), me.pointers, fmt.Sprintf("unsafe.Pointer(%s)", v))
}

type funcType struct {
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

func (me *funcType) goToCGo(v string) string {
	return withGoCast(fmt.Sprintf("C.%s", me.ctype), me.pointers-1, v)
}

type funcAnonType struct {
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

func (me *funcAnonType) goToCGo(v string) string {
	return fmt.Sprintf("(*[0]byte)(%s)", v)
}

func (me *funcAnonType) addAnonTypes(w io.Writer, structName, memberName string) error {
	me.name = fmt.Sprintf("%s%sFn", structName, memberName)
	return writeTemplate(w, "types_struct_anon_fn.go", map[string]interface{}{
		"Name":       me.name,
		"StructName": structName,
		"MemberName": memberName,
	})
}
