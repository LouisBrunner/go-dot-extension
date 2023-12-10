// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceFormatImporterSaver struct {
  obj gdc.ObjectPtr
}

func (me *ResourceFormatImporterSaver) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceFormatImporterSaver) BaseClass() string {
  return "ResourceFormatImporterSaver"
}



// Enums

func (me *ResourceFormatImporterSaver) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceFormatImporterSaver) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceFormatImporterSaver) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties