// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiMeshInstance2D struct {
  Node2D
}

func (me *MultiMeshInstance2D) BaseClass() string {
  return "MultiMeshInstance2D"
}



// Enums

func (me *MultiMeshInstance2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiMeshInstance2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiMeshInstance2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MultiMeshInstance2D) SetMultimesh(multimesh MultiMesh, )  {
  classNameV := StringNameFromStr("MultiMeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multimesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2246127404) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(multimesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMeshInstance2D) GetMultimesh() MultiMesh {
  classNameV := StringNameFromStr("MultiMeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multimesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1385450523) // FIXME: should cache?
  var ret MultiMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMeshInstance2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("MultiMeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMeshInstance2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("MultiMeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiMeshInstance2DTextureChangedSignalFn func()

func (me *MultiMeshInstance2D) ConnectTextureChanged(subs SignalSubscribers, fn MultiMeshInstance2DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiMeshInstance2D) DisconnectTextureChanged(subs SignalSubscribers, fn MultiMeshInstance2DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
