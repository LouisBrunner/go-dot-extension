// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ConcavePolygonShape3D struct {
  obj gdc.ObjectPtr
}

func (me *ConcavePolygonShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConcavePolygonShape3D) BaseClass() string {
  return "ConcavePolygonShape3D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(faces.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConcavePolygonShape3D) GetFaces() PackedVector3Array {
  classNameV := StringNameFromStr("ConcavePolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ConcavePolygonShape3D) SetBackfaceCollisionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("ConcavePolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_backface_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConcavePolygonShape3D) IsBackfaceCollisionEnabled() bool {
  classNameV := StringNameFromStr("ConcavePolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_backface_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ConcavePolygonShape3D) GetPropData() PackedVector3Array {
  panic("TODO: implement")
}

func (me *ConcavePolygonShape3D) SetPropData(value PackedVector3Array) {
  panic("TODO: implement")
}

func (me *ConcavePolygonShape3D) GetPropBackfaceCollision() bool {
  panic("TODO: implement")
}

func (me *ConcavePolygonShape3D) SetPropBackfaceCollision(value bool) {
  panic("TODO: implement")
}