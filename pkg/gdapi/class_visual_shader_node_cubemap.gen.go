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
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1625400621) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCubemap) GetSource() VisualShaderNodeCubemapSource {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2222048781) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeCubemapSource

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShaderNodeCubemap) SetCubeMap(value Cubemap, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cube_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2219800736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCubemap) GetCubeMap() Cubemap {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cube_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1772111058) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCubemap()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNodeCubemap) SetTextureType(value VisualShaderNodeCubemapTextureType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1899718876) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCubemap) GetTextureType() VisualShaderNodeCubemapTextureType {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3356498888) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeCubemapTextureType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
