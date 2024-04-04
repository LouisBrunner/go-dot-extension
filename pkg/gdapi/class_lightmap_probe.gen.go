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

type LightmapProbe struct {
  Node3D
}

func (me *LightmapProbe) BaseClass() string {
  return "LightmapProbe"
}

func NewLightmapProbe() *LightmapProbe {
  str := StringNameFromStr("LightmapProbe") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &LightmapProbe{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *LightmapProbe) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LightmapProbe) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightmapProbe) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
