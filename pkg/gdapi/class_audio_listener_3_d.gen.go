// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioListener3D struct {
  obj gdc.ObjectPtr
}

func (me *AudioListener3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioListener3D) BaseClass() string {
  return "AudioListener3D"
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
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioListener3D) ClearCurrent()  {
  classNameV := StringNameFromStr("AudioListener3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioListener3D) IsCurrent() bool {
  classNameV := StringNameFromStr("AudioListener3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioListener3D) GetListenerTransform() Transform3D {
  classNameV := StringNameFromStr("AudioListener3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_listener_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
