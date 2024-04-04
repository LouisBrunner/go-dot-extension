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

type ResourceImporterImage struct {
  ResourceImporter
}

func (me *ResourceImporterImage) BaseClass() string {
  return "ResourceImporterImage"
}

func NewResourceImporterImage() *ResourceImporterImage {
  str := StringNameFromStr("ResourceImporterImage") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterImage{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterImage) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterImage) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterImage) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
