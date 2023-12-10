// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type World3D struct {
  obj gdc.ObjectPtr
}

func (me *World3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *World3D) BaseClass() string {
  return "World3D"
}



// Enums

func (me *World3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *World3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *World3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *World3D) GetSpace() RID {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World3D) GetNavigationMap() RID {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World3D) GetScenario() RID {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scenario")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World3D) SetEnvironment(env Environment, )  {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4143518816) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(env.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *World3D) GetEnvironment() Environment {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3082064660) // FIXME: should cache?
  var ret Environment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World3D) SetFallbackEnvironment(env Environment, )  {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4143518816) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(env.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *World3D) GetFallbackEnvironment() Environment {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3082064660) // FIXME: should cache?
  var ret Environment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World3D) SetCameraAttributes(attributes CameraAttributes, )  {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2817810567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(attributes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *World3D) GetCameraAttributes() CameraAttributes {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3921283215) // FIXME: should cache?
  var ret CameraAttributes
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *World3D) GetDirectSpaceState() PhysicsDirectSpaceState3D {
  classNameV := StringNameFromStr("World3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direct_space_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2069328350) // FIXME: should cache?
  var ret PhysicsDirectSpaceState3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *World3D) GetPropEnvironment() Environment {
  panic("TODO: implement")
}

func (me *World3D) SetPropEnvironment(value Environment) {
  panic("TODO: implement")
}

func (me *World3D) GetPropFallbackEnvironment() Environment {
  panic("TODO: implement")
}

func (me *World3D) SetPropFallbackEnvironment(value Environment) {
  panic("TODO: implement")
}

func (me *World3D) GetPropCameraAttributes() any {
  panic("TODO: implement")
}

func (me *World3D) SetPropCameraAttributes(value any) {
  panic("TODO: implement")
}

func (me *World3D) GetPropSpace() RID {
  panic("TODO: implement")
}

func (me *World3D) SetPropSpace(value RID) {
  panic("TODO: implement")
}

func (me *World3D) GetPropNavigationMap() RID {
  panic("TODO: implement")
}

func (me *World3D) SetPropNavigationMap(value RID) {
  panic("TODO: implement")
}

func (me *World3D) GetPropScenario() RID {
  panic("TODO: implement")
}

func (me *World3D) SetPropScenario(value RID) {
  panic("TODO: implement")
}

func (me *World3D) GetPropDirectSpaceState() PhysicsDirectSpaceState3D {
  panic("TODO: implement")
}

func (me *World3D) SetPropDirectSpaceState(value PhysicsDirectSpaceState3D) {
  panic("TODO: implement")
}