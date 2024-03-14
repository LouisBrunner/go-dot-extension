// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GDScript struct {
  Script
}

func (me *GDScript) BaseClass() string {
  return "GDScript"
}



// Enums

func (me *GDScript) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GDScript) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDScript) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GDScript) New(varargs ...Variant) Variant {
  classNameV := StringNameFromStr("GDScript")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("new")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1545262638) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstVariantPtr{}
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), gdc.UninitializedVariantPtr(&ret), err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

  return ret
}

// Signals
