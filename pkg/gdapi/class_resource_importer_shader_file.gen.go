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

type ResourceImporterShaderFile struct {
  ResourceImporter
}

func (me *ResourceImporterShaderFile) BaseClass() string {
  return "ResourceImporterShaderFile"
}

func NewResourceImporterShaderFile() *ResourceImporterShaderFile {
  str := StringNameFromStr("ResourceImporterShaderFile") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterShaderFile{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterShaderFile) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterShaderFile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterShaderFile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
