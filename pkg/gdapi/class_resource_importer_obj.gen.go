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

type ResourceImporterOBJ struct {
  ResourceImporter
}

func (me *ResourceImporterOBJ) BaseClass() string {
  return "ResourceImporterOBJ"
}

func NewResourceImporterOBJ() *ResourceImporterOBJ {
  str := StringNameFromStr("ResourceImporterOBJ") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterOBJ{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterOBJ) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterOBJ) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterOBJ) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
