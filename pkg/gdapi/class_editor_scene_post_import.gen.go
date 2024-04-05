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

type ptrsForEditorScenePostImportList struct {
  fnXPostImport gdc.MethodBindPtr
  fnGetSourceFile gdc.MethodBindPtr
}

var ptrsForEditorScenePostImport ptrsForEditorScenePostImportList

func initEditorScenePostImportPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorScenePostImport")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_source_file")
    defer methodName.Destroy()
    ptrsForEditorScenePostImport.fnGetSourceFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScenePostImport.fnGetSourceFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
