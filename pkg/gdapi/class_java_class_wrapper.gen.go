// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type JavaClassWrapper struct {
  obj gdc.ObjectPtr
}

func (me *JavaClassWrapper) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JavaClassWrapper) BaseClass() string {
  return "JavaClassWrapper"
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
  var ret JavaClass
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
