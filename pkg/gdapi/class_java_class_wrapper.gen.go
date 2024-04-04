// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type JavaClassWrapper struct {
  Object
}

func (me *JavaClassWrapper) BaseClass() string {
  return "JavaClassWrapper"
}

func NewJavaClassWrapper() *JavaClassWrapper {
  str := StringNameFromStr("JavaClassWrapper") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &JavaClassWrapper{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *JavaClassWrapper) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaClassWrapper) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaClassWrapper) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *JavaClassWrapper) Wrap(name String, ) JavaClass {
  classNameV := StringNameFromStr("JavaClassWrapper")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wrap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1124367868) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewJavaClass()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
