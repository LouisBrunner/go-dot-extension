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

type EditorExportPlatform struct {
  RefCounted
}

func (me *EditorExportPlatform) BaseClass() string {
  return "EditorExportPlatform"
}

func NewEditorExportPlatform() *EditorExportPlatform {
  str := StringNameFromStr("EditorExportPlatform") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatform{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorExportPlatform) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatform) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatform) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorExportPlatform) GetOsName() String {
  classNameV := StringNameFromStr("EditorExportPlatform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_os_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
