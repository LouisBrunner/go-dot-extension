// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeSample3D struct {
  VisualShaderNode
}

func (me *VisualShaderNodeSample3D) BaseClass() string {
  return "VisualShaderNodeSample3D"
}

func NewVisualShaderNodeSample3D() *VisualShaderNodeSample3D {
  str := StringNameFromStr("VisualShaderNodeSample3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeSample3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeSample3DSource int
const (
  VisualShaderNodeSample3DSourceSourceTexture VisualShaderNodeSample3DSource = 0
  VisualShaderNodeSample3DSourceSourcePort VisualShaderNodeSample3DSource = 1
  VisualShaderNodeSample3DSourceSourceMax VisualShaderNodeSample3DSource = 2
)

func (me *VisualShaderNodeSample3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSample3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSample3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeSample3D) SetSource(value VisualShaderNodeSample3DSource, )  {
  classNameV := StringNameFromStr("VisualShaderNodeSample3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3315130991) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeSample3D) GetSource() VisualShaderNodeSample3DSource {
  classNameV := StringNameFromStr("VisualShaderNodeSample3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1079494121) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret VisualShaderNodeSample3DSource

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
