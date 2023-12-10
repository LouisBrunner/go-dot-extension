// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WorldEnvironment struct {
  obj gdc.ObjectPtr
}

func (me *WorldEnvironment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorldEnvironment) BaseClass() string {
  return "WorldEnvironment"
}



// Enums

func (me *WorldEnvironment) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WorldEnvironment) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorldEnvironment) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WorldEnvironment) SetEnvironment(env Environment, )  {
  classNameV := StringNameFromStr("WorldEnvironment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4143518816) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(env.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WorldEnvironment) GetEnvironment() Environment {
  classNameV := StringNameFromStr("WorldEnvironment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3082064660) // FIXME: should cache?
  var ret Environment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorldEnvironment) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  classNameV := StringNameFromStr("WorldEnvironment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2817810567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(camera_attributes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WorldEnvironment) GetCameraAttributes() CameraAttributes {
  classNameV := StringNameFromStr("WorldEnvironment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3921283215) // FIXME: should cache?
  var ret CameraAttributes
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *WorldEnvironment) GetPropEnvironment() Environment {
  panic("TODO: implement")
}

func (me *WorldEnvironment) SetPropEnvironment(value Environment) {
  panic("TODO: implement")
}

func (me *WorldEnvironment) GetPropCameraAttributes() any {
  panic("TODO: implement")
}

func (me *WorldEnvironment) SetPropCameraAttributes(value any) {
  panic("TODO: implement")
}