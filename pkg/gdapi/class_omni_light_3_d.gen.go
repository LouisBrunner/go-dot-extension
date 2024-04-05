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

type ptrsForOmniLight3DList struct {
  fnSetShadowMode gdc.MethodBindPtr
  fnGetShadowMode gdc.MethodBindPtr
}

var ptrsForOmniLight3D ptrsForOmniLight3DList

func initOmniLight3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("OmniLight3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_shadow_mode")
    defer methodName.Destroy()
    ptrsForOmniLight3D.fnSetShadowMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121862228))
  }
  {
    methodName := StringNameFromStr("get_shadow_mode")
    defer methodName.Destroy()
    ptrsForOmniLight3D.fnGetShadowMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4181586331))
  }
}

type OmniLight3D struct {
  Light3D
}

func (me *OmniLight3D) BaseClass() string {
  return "OmniLight3D"
}

func NewOmniLight3D() *OmniLight3D {
  str := StringNameFromStr("OmniLight3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OmniLight3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type OmniLight3DShadowMode int
const (
  OmniLight3DShadowModeShadowDualParaboloid OmniLight3DShadowMode = 0
  OmniLight3DShadowModeShadowCube OmniLight3DShadowMode = 1
)

func (me *OmniLight3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OmniLight3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OmniLight3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OmniLight3D) SetShadowMode(mode OmniLight3DShadowMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOmniLight3D.fnSetShadowMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OmniLight3D) GetShadowMode() OmniLight3DShadowMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret OmniLight3DShadowMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOmniLight3D.fnGetShadowMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
