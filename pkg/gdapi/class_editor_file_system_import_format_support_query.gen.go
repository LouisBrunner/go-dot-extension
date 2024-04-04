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

type EditorFileSystemImportFormatSupportQuery struct {
  RefCounted
}

func (me *EditorFileSystemImportFormatSupportQuery) BaseClass() string {
  return "EditorFileSystemImportFormatSupportQuery"
}

func NewEditorFileSystemImportFormatSupportQuery() *EditorFileSystemImportFormatSupportQuery {
  str := StringNameFromStr("EditorFileSystemImportFormatSupportQuery") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorFileSystemImportFormatSupportQuery{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
