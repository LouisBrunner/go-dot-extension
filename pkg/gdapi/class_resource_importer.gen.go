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

type ResourceImporter struct {
  RefCounted
}

func (me *ResourceImporter) BaseClass() string {
  return "ResourceImporter"
}

func NewResourceImporter() *ResourceImporter {
  str := StringNameFromStr("ResourceImporter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ResourceImporterImportOrder int
const (
  ResourceImporterImportOrderImportOrderDefault ResourceImporterImportOrder = 0
  ResourceImporterImportOrderImportOrderScene ResourceImporterImportOrder = 100
)

func (me *ResourceImporter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
