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

type ptrsForCSGPrimitive3DList struct {
  fnSetFlipFaces gdc.MethodBindPtr
  fnGetFlipFaces gdc.MethodBindPtr
}

var ptrsForCSGPrimitive3D ptrsForCSGPrimitive3DList

func initCSGPrimitive3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CSGPrimitive3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_flip_faces")
    defer methodName.Destroy()
    ptrsForCSGPrimitive3D.fnSetFlipFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_flip_faces")
    defer methodName.Destroy()
    ptrsForCSGPrimitive3D.fnGetFlipFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
}

type CSGPrimitive3D struct {
  CSGShape3D
}

func (me *CSGPrimitive3D) BaseClass() string {
  return "CSGPrimitive3D"
}

func NewCSGPrimitive3D() *CSGPrimitive3D {
  str := StringNameFromStr("CSGPrimitive3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CSGPrimitive3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CSGPrimitive3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGPrimitive3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGPrimitive3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGPrimitive3D) SetFlipFaces(flip_faces bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_faces) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPrimitive3D.fnSetFlipFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPrimitive3D) GetFlipFaces() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPrimitive3D.fnGetFlipFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
