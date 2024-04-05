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

type ptrsForAudioListener3DList struct {
  fnMakeCurrent gdc.MethodBindPtr
  fnClearCurrent gdc.MethodBindPtr
  fnIsCurrent gdc.MethodBindPtr
  fnGetListenerTransform gdc.MethodBindPtr
}

var ptrsForAudioListener3D ptrsForAudioListener3DList

func initAudioListener3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioListener3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("make_current")
    defer methodName.Destroy()
    ptrsForAudioListener3D.fnMakeCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("clear_current")
    defer methodName.Destroy()
    ptrsForAudioListener3D.fnClearCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("is_current")
    defer methodName.Destroy()
    ptrsForAudioListener3D.fnIsCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_listener_transform")
    defer methodName.Destroy()
    ptrsForAudioListener3D.fnGetListenerTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
  }
}

type AudioListener3D struct {
  Node3D
}

func (me *AudioListener3D) BaseClass() string {
  return "AudioListener3D"
}

func NewAudioListener3D() *AudioListener3D {
  str := StringNameFromStr("AudioListener3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioListener3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioListener3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioListener3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioListener3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioListener3D) MakeCurrent()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioListener3D.fnMakeCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioListener3D) ClearCurrent()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioListener3D.fnClearCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioListener3D) IsCurrent() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioListener3D.fnIsCurrent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioListener3D) GetListenerTransform() Transform3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioListener3D.fnGetListenerTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
