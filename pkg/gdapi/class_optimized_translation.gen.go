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

type OptimizedTranslation struct {
  Translation
}

func (me *OptimizedTranslation) BaseClass() string {
  return "OptimizedTranslation"
}

func NewOptimizedTranslation() *OptimizedTranslation {
  str := StringNameFromStr("OptimizedTranslation") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OptimizedTranslation{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
