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

type ptrsForVisualShaderNodeCubemapList struct {
  fnSetSource gdc.MethodBindPtr
  fnGetSource gdc.MethodBindPtr
  fnSetCubeMap gdc.MethodBindPtr
  fnGetCubeMap gdc.MethodBindPtr
  fnSetTextureType gdc.MethodBindPtr
  fnGetTextureType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeCubemap ptrsForVisualShaderNodeCubemapList

func initVisualShaderNodeCubemapPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeCubemap")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_source")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCubemap.fnSetSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1625400621))
  }
  {
    methodName := StringNameFromStr("get_source")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCubemap.fnGetSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2222048781))
  }
  {
    methodName := StringNameFromStr("set_cube_map")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCubemap.fnSetCubeMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2219800736))
  }
  {
    methodName := StringNameFromStr("get_cube_map")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCubemap.fnGetCubeMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1772111058))
  }
  {
    methodName := StringNameFromStr("set_texture_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCubemap.fnSetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1899718876))
  }
  {
    methodName := StringNameFromStr("get_texture_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCubemap.fnGetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3356498888))
  }
}

type VisualShaderNodeCubemap struct {
  VisualShaderNode
}

func (me *VisualShaderNodeCubemap) BaseClass() string {
  return "VisualShaderNodeCubemap"
}

func NewVisualShaderNodeCubemap() *VisualShaderNodeCubemap {
  str := StringNameFromStr("VisualShaderNodeCubemap") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeCubemap{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeCubemapSource int
const (
  VisualShaderNodeCubemapSourceSourceTexture VisualShaderNodeCubemapSource = 0
  VisualShaderNodeCubemapSourceSourcePort VisualShaderNodeCubemapSource = 1
  VisualShaderNodeCubemapSourceSourceMax VisualShaderNodeCubemapSource = 2
)

type VisualShaderNodeCubemapTextureType int
const (
  VisualShaderNodeCubemapTextureTypeTypeData VisualShaderNodeCubemapTextureType = 0
  VisualShaderNodeCubemapTextureTypeTypeColor VisualShaderNodeCubemapTextureType = 1
  VisualShaderNodeCubemapTextureTypeTypeNormalMap VisualShaderNodeCubemapTextureType = 2
  VisualShaderNodeCubemapTextureTypeTypeMax VisualShaderNodeCubemapTextureType = 3
)

func (me *VisualShaderNodeCubemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCubemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCubemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeCubemap) SetSource(value VisualShaderNodeCubemapSource, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCubemap.fnSetSource), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCubemap) GetSource() VisualShaderNodeCubemapSource {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeCubemapSource

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCubemap.fnGetSource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShaderNodeCubemap) SetCubeMap(value Cubemap, )  {
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCubemap.fnSetCubeMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCubemap) GetCubeMap() Cubemap {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCubemap()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCubemap.fnGetCubeMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNodeCubemap) SetTextureType(value VisualShaderNodeCubemapTextureType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCubemap.fnSetTextureType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCubemap) GetTextureType() VisualShaderNodeCubemapTextureType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeCubemapTextureType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCubemap.fnGetTextureType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
