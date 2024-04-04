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

type WorldEnvironment struct {
  Node
}

func (me *WorldEnvironment) BaseClass() string {
  return "WorldEnvironment"
}

func NewWorldEnvironment() *WorldEnvironment {
  str := StringNameFromStr("WorldEnvironment") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WorldEnvironment{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WorldEnvironment) GetEnvironment() Environment {
  classNameV := StringNameFromStr("WorldEnvironment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3082064660) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEnvironment()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WorldEnvironment) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  classNameV := StringNameFromStr("WorldEnvironment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2817810567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{camera_attributes.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WorldEnvironment) GetCameraAttributes() CameraAttributes {
  classNameV := StringNameFromStr("WorldEnvironment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3921283215) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCameraAttributes()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
