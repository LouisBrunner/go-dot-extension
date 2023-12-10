// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorFileSystemImportFormatSupportQuery struct {
  obj gdc.ObjectPtr
}

func (me *EditorFileSystemImportFormatSupportQuery) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorFileSystemImportFormatSupportQuery) BaseClass() string {
  return "EditorFileSystemImportFormatSupportQuery"
}



// Enums

func (me *EditorFileSystemImportFormatSupportQuery) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorFileSystemImportFormatSupportQuery) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorFileSystemImportFormatSupportQuery) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties