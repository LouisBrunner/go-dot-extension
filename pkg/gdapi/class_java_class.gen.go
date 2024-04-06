// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForJavaClassList struct {
}

var ptrsForJavaClass ptrsForJavaClassList

func initJavaClassPtrs(iface gdc.Interface) {

	className := StringNameFromStr("JavaClass")
	defer className.Destroy()
}

type JavaClass struct {
	RefCounted
}

func (me *JavaClass) BaseClass() string {
	return "JavaClass"
}

func NewJavaClass() *JavaClass {
	str := StringNameFromStr("JavaClass") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &JavaClass{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *JavaClass) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *JavaClass) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *JavaClass) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
