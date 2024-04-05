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

type ptrsForEditorExportPlatformPCList struct {
}

var ptrsForEditorExportPlatformPC ptrsForEditorExportPlatformPCList

func initEditorExportPlatformPCPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorExportPlatformPC")
  defer className.Destroy()
}

type EditorExportPlatformPC struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformPC) BaseClass() string {
  return "EditorExportPlatformPC"
}

func NewEditorExportPlatformPC() *EditorExportPlatformPC {
  str := StringNameFromStr("EditorExportPlatformPC") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatformPC{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorExportPlatformPC) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformPC) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformPC) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
