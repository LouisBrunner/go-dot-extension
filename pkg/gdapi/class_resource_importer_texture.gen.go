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

type ResourceImporterTexture struct {
  ResourceImporter
}

func (me *ResourceImporterTexture) BaseClass() string {
  return "ResourceImporterTexture"
}

func NewResourceImporterTexture() *ResourceImporterTexture {
  str := StringNameFromStr("ResourceImporterTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
