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

type ptrsForJavaScriptObjectList struct {
}

var ptrsForJavaScriptObject ptrsForJavaScriptObjectList

func initJavaScriptObjectPtrs(iface gdc.Interface) {

  className := StringNameFromStr("JavaScriptObject")
  defer className.Destroy()
}

type JavaScriptObject struct {
  RefCounted
}

func (me *JavaScriptObject) BaseClass() string {
  return "JavaScriptObject"
}

func NewJavaScriptObject() *JavaScriptObject {
  str := StringNameFromStr("JavaScriptObject") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &JavaScriptObject{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *JavaScriptObject) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaScriptObject) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaScriptObject) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
