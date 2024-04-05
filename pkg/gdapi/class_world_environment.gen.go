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

type ptrsForWorldEnvironmentList struct {
  fnSetEnvironment gdc.MethodBindPtr
  fnGetEnvironment gdc.MethodBindPtr
  fnSetCameraAttributes gdc.MethodBindPtr
  fnGetCameraAttributes gdc.MethodBindPtr
}

var ptrsForWorldEnvironment ptrsForWorldEnvironmentList

func initWorldEnvironmentPtrs(iface gdc.Interface) {

  className := StringNameFromStr("WorldEnvironment")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_environment")
    defer methodName.Destroy()
    ptrsForWorldEnvironment.fnSetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4143518816))
  }
  {
    methodName := StringNameFromStr("get_environment")
    defer methodName.Destroy()
    ptrsForWorldEnvironment.fnGetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3082064660))
  }
  {
    methodName := StringNameFromStr("set_camera_attributes")
    defer methodName.Destroy()
    ptrsForWorldEnvironment.fnSetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2817810567))
  }
  {
    methodName := StringNameFromStr("get_camera_attributes")
    defer methodName.Destroy()
    ptrsForWorldEnvironment.fnGetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3921283215))
  }
}

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
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldEnvironment.fnSetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WorldEnvironment) GetEnvironment() Environment {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEnvironment()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldEnvironment.fnGetEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WorldEnvironment) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  cargs := []gdc.ConstTypePtr{camera_attributes.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldEnvironment.fnSetCameraAttributes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WorldEnvironment) GetCameraAttributes() CameraAttributes {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCameraAttributes()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldEnvironment.fnGetCameraAttributes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
