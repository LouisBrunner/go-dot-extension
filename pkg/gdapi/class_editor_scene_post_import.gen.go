// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorScenePostImport struct {
  obj gdc.ObjectPtr
}

func (me *EditorScenePostImport) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorScenePostImport) BaseClass() string {
  return "EditorScenePostImport"
}



// Enums

func (me *EditorScenePostImport) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorScenePostImport) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorScenePostImport) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorScenePostImport) GetSourceFile() String {
  classNameV := StringNameFromStr("EditorScenePostImport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
