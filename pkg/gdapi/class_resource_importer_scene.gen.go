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

type ResourceImporterScene struct {
  ResourceImporter
}

func (me *ResourceImporterScene) BaseClass() string {
  return "ResourceImporterScene"
}

func NewResourceImporterScene() *ResourceImporterScene {
  str := StringNameFromStr("ResourceImporterScene") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterScene{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterScene) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterScene) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterScene) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
