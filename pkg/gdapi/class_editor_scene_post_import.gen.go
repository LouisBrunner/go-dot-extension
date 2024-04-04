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

type EditorScenePostImport struct {
  RefCounted
}

func (me *EditorScenePostImport) BaseClass() string {
  return "EditorScenePostImport"
}

func NewEditorScenePostImport() *EditorScenePostImport {
  str := StringNameFromStr("EditorScenePostImport") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorScenePostImport{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
