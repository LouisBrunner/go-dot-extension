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

type ResourceImporterDynamicFont struct {
  ResourceImporter
}

func (me *ResourceImporterDynamicFont) BaseClass() string {
  return "ResourceImporterDynamicFont"
}

func NewResourceImporterDynamicFont() *ResourceImporterDynamicFont {
  str := StringNameFromStr("ResourceImporterDynamicFont") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterDynamicFont{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterDynamicFont) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterDynamicFont) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterDynamicFont) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
