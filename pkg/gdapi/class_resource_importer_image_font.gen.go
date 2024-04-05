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

type ptrsForResourceImporterImageFontList struct {
}

var ptrsForResourceImporterImageFont ptrsForResourceImporterImageFontList

func initResourceImporterImageFontPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ResourceImporterImageFont")
  defer className.Destroy()
}

type ResourceImporterImageFont struct {
  ResourceImporter
}

func (me *ResourceImporterImageFont) BaseClass() string {
  return "ResourceImporterImageFont"
}

func NewResourceImporterImageFont() *ResourceImporterImageFont {
  str := StringNameFromStr("ResourceImporterImageFont") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterImageFont{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterImageFont) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterImageFont) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterImageFont) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
