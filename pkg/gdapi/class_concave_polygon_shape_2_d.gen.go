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

type ptrsForConcavePolygonShape2DList struct {
  fnSetSegments gdc.MethodBindPtr
  fnGetSegments gdc.MethodBindPtr
}

var ptrsForConcavePolygonShape2D ptrsForConcavePolygonShape2DList

func initConcavePolygonShape2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ConcavePolygonShape2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_segments")
    defer methodName.Destroy()
    ptrsForConcavePolygonShape2D.fnSetSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
  }
  {
    methodName := StringNameFromStr("get_segments")
    defer methodName.Destroy()
    ptrsForConcavePolygonShape2D.fnGetSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
  }
}

type ConcavePolygonShape2D struct {
  Shape2D
}

func (me *ConcavePolygonShape2D) BaseClass() string {
  return "ConcavePolygonShape2D"
}

func NewConcavePolygonShape2D() *ConcavePolygonShape2D {
  str := StringNameFromStr("ConcavePolygonShape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ConcavePolygonShape2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{segments.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConcavePolygonShape2D.fnSetSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConcavePolygonShape2D) GetSegments() PackedVector2Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConcavePolygonShape2D.fnGetSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
