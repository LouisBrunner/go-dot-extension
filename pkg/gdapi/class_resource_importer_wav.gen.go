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

type ResourceImporterWAV struct {
  ResourceImporter
}

func (me *ResourceImporterWAV) BaseClass() string {
  return "ResourceImporterWAV"
}

func NewResourceImporterWAV() *ResourceImporterWAV {
  str := StringNameFromStr("ResourceImporterWAV") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterWAV{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterWAV) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterWAV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterWAV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
