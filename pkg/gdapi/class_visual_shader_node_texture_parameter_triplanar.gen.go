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

type ptrsForVisualShaderNodeTextureParameterTriplanarList struct {
}

var ptrsForVisualShaderNodeTextureParameterTriplanar ptrsForVisualShaderNodeTextureParameterTriplanarList

func initVisualShaderNodeTextureParameterTriplanarPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeTextureParameterTriplanar")
  defer className.Destroy()
}

type VisualShaderNodeTextureParameterTriplanar struct {
  VisualShaderNodeTextureParameter
}

func (me *VisualShaderNodeTextureParameterTriplanar) BaseClass() string {
  return "VisualShaderNodeTextureParameterTriplanar"
}

func NewVisualShaderNodeTextureParameterTriplanar() *VisualShaderNodeTextureParameterTriplanar {
  str := StringNameFromStr("VisualShaderNodeTextureParameterTriplanar") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeTextureParameterTriplanar{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeTextureParameterTriplanar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTextureParameterTriplanar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureParameterTriplanar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
