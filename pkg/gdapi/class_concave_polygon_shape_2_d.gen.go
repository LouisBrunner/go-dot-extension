// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ConcavePolygonShape2D struct {
  obj gdc.ObjectPtr
}

func (me *ConcavePolygonShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConcavePolygonShape2D) BaseClass() string {
  return "ConcavePolygonShape2D"
}



// Enums

func (me *ConcavePolygonShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConcavePolygonShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConcavePolygonShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ConcavePolygonShape2D) SetSegments(segments PackedVector2Array, )  {
  classNameV := StringNameFromStr("ConcavePolygonShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(segments.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConcavePolygonShape2D) GetSegments() PackedVector2Array {
  classNameV := StringNameFromStr("ConcavePolygonShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
