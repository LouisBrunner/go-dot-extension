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

type MeshInstance2D struct {
  Node2D
}

func (me *MeshInstance2D) BaseClass() string {
  return "MeshInstance2D"
}

func NewMeshInstance2D() *MeshInstance2D {
  str := StringNameFromStr("MeshInstance2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MeshInstance2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MeshInstance2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MeshInstance2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MeshInstance2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MeshInstance2D) SetMesh(mesh Mesh, )  {
  classNameV := StringNameFromStr("MeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194775623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance2D) GetMesh() Mesh {
  classNameV := StringNameFromStr("MeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1808005922) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("MeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("MeshInstance2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MeshInstance2DTextureChangedSignalFn func()

func (me *MeshInstance2D) ConnectTextureChanged(subs SignalSubscribers, fn MeshInstance2DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MeshInstance2D) DisconnectTextureChanged(subs SignalSubscribers, fn MeshInstance2DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
