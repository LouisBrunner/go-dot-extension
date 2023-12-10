// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OptimizedTranslation struct {
  obj gdc.ObjectPtr
}

func (me *OptimizedTranslation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OptimizedTranslation) BaseClass() string {
  return "OptimizedTranslation"
}



// Enums

func (me *OptimizedTranslation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OptimizedTranslation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OptimizedTranslation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OptimizedTranslation) Generate(from Translation, )  {
  classNameV := StringNameFromStr("OptimizedTranslation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1466479800) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties