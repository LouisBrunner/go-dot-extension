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
  classNameV := StringNameFromStr("AudioListener3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioListener3D) ClearCurrent()  {
  classNameV := StringNameFromStr("AudioListener3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioListener3D) IsCurrent() bool {
  classNameV := StringNameFromStr("AudioListener3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioListener3D) GetListenerTransform() Transform3D {
  classNameV := StringNameFromStr("AudioListener3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_listener_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
