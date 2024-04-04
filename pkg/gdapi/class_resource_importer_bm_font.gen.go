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

type ResourceImporterBMFont struct {
  ResourceImporter
}

func (me *ResourceImporterBMFont) BaseClass() string {
  return "ResourceImporterBMFont"
}

func NewResourceImporterBMFont() *ResourceImporterBMFont {
  str := StringNameFromStr("ResourceImporterBMFont") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterBMFont{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterBMFont) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterBMFont) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterBMFont) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
