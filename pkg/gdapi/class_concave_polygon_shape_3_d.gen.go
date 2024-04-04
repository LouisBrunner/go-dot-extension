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

type ConcavePolygonShape3D struct {
  Shape3D
}

func (me *ConcavePolygonShape3D) BaseClass() string {
  return "ConcavePolygonShape3D"
}

func NewConcavePolygonShape3D() *ConcavePolygonShape3D {
  str := StringNameFromStr("ConcavePolygonShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ConcavePolygonShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ConcavePolygonShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConcavePolygonShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConcavePolygonShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ConcavePolygonShape3D) SetFaces(faces PackedVector3Array, )  {
  classNameV := StringNameFromStr("ConcavePolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{faces.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConcavePolygonShape3D) GetFaces() PackedVector3Array {
  classNameV := StringNameFromStr("ConcavePolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ConcavePolygonShape3D) SetBackfaceCollisionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("ConcavePolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_backface_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConcavePolygonShape3D) IsBackfaceCollisionEnabled() bool {
  classNameV := StringNameFromStr("ConcavePolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_backface_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
