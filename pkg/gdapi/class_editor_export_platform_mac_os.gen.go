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

type ptrsForEditorExportPlatformMacOSList struct {
}

var ptrsForEditorExportPlatformMacOS ptrsForEditorExportPlatformMacOSList

func initEditorExportPlatformMacOSPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorExportPlatformMacOS")
  defer className.Destroy()
}

type EditorExportPlatformMacOS struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformMacOS) BaseClass() string {
  return "EditorExportPlatformMacOS"
}

func NewEditorExportPlatformMacOS() *EditorExportPlatformMacOS {
  str := StringNameFromStr("EditorExportPlatformMacOS") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatformMacOS{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorExportPlatformMacOS) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformMacOS) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformMacOS) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
