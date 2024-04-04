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

type AudioListener2D struct {
  Node2D
}

func (me *AudioListener2D) BaseClass() string {
  return "AudioListener2D"
}

func NewAudioListener2D() *AudioListener2D {
  str := StringNameFromStr("AudioListener2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioListener2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioListener2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioListener2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioListener2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioListener2D) MakeCurrent()  {
  classNameV := StringNameFromStr("AudioListener2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioListener2D) ClearCurrent()  {
  classNameV := StringNameFromStr("AudioListener2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioListener2D) IsCurrent() bool {
  classNameV := StringNameFromStr("AudioListener2D")
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

// Signals
