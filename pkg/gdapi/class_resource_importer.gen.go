// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceImporter struct {
  obj gdc.ObjectPtr
}

func (me *ResourceImporter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceImporter) BaseClass() string {
  return "ResourceImporter"
}

type ResourceImporterImportOrder int
const (
  ResourceImporterImportOrderImportOrderDefault ResourceImporterImportOrder = 0
  ResourceImporterImportOrderImportOrderScene ResourceImporterImportOrder = 100
)

// TODO: properties

// TODO: signals
