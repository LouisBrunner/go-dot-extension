// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
