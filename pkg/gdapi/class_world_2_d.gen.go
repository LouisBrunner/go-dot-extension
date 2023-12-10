// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type World2D struct {
  obj gdc.ObjectPtr
}

func (me *World2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *World2D) BaseClass() string {
  return "World2D"
}



// Enums

func (me *World2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *World2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *World2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *World2D) GetCanvas() RID {
  classNameV := StringNameFromStr("World2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World2D) GetSpace() RID {
  classNameV := StringNameFromStr("World2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World2D) GetNavigationMap() RID {
  classNameV := StringNameFromStr("World2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World2D) GetDirectSpaceState() PhysicsDirectSpaceState2D {
  classNameV := StringNameFromStr("World2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direct_space_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2506717822) // FIXME: should cache?
  var ret PhysicsDirectSpaceState2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *World2D) GetPropCanvas() RID {
  panic("TODO: implement")
}

func (me *World2D) SetPropCanvas(value RID) {
  panic("TODO: implement")
}

func (me *World2D) GetPropSpace() RID {
  panic("TODO: implement")
}

func (me *World2D) SetPropSpace(value RID) {
  panic("TODO: implement")
}

func (me *World2D) GetPropNavigationMap() RID {
  panic("TODO: implement")
}

func (me *World2D) SetPropNavigationMap(value RID) {
  panic("TODO: implement")
}

func (me *World2D) GetPropDirectSpaceState() PhysicsDirectSpaceState2D {
  panic("TODO: implement")
}

func (me *World2D) SetPropDirectSpaceState(value PhysicsDirectSpaceState2D) {
  panic("TODO: implement")
}