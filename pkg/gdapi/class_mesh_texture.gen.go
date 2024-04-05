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

type ptrsForMeshTextureList struct {
  fnSetMesh gdc.MethodBindPtr
  fnGetMesh gdc.MethodBindPtr
  fnSetImageSize gdc.MethodBindPtr
  fnGetImageSize gdc.MethodBindPtr
  fnSetBaseTexture gdc.MethodBindPtr
  fnGetBaseTexture gdc.MethodBindPtr
}

var ptrsForMeshTexture ptrsForMeshTextureList

func initMeshTexturePtrs(iface gdc.Interface) {

  className := StringNameFromStr("MeshTexture")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_mesh")
    defer methodName.Destroy()
    ptrsForMeshTexture.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
  }
  {
    methodName := StringNameFromStr("get_mesh")
    defer methodName.Destroy()
    ptrsForMeshTexture.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1808005922))
  }
  {
    methodName := StringNameFromStr("set_image_size")
    defer methodName.Destroy()
    ptrsForMeshTexture.fnSetImageSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_image_size")
    defer methodName.Destroy()
    ptrsForMeshTexture.fnGetImageSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_base_texture")
    defer methodName.Destroy()
    ptrsForMeshTexture.fnSetBaseTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_base_texture")
    defer methodName.Destroy()
    ptrsForMeshTexture.fnGetBaseTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
}

type MeshTexture struct {
  Texture2D
}

func (me *MeshTexture) BaseClass() string {
  return "MeshTexture"
}

func NewMeshTexture() *MeshTexture {
  str := StringNameFromStr("MeshTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MeshTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MeshTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MeshTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MeshTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MeshTexture) SetMesh(mesh Mesh, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshTexture.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshTexture) GetMesh() Mesh {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshTexture.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshTexture) SetImageSize(size Vector2, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshTexture.fnSetImageSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshTexture) GetImageSize() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshTexture.fnGetImageSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshTexture) SetBaseTexture(texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshTexture.fnSetBaseTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshTexture) GetBaseTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshTexture.fnGetBaseTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
