// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorInspector struct {
  obj gdc.ObjectPtr
}

func (me *EditorInspector) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorInspector) BaseClass() string {
  return "EditorInspector"
}



// Enums

func (me *EditorInspector) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorInspector) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorInspector) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorInspector) GetSelectedPath() String {
  classNameV := StringNameFromStr("EditorInspector")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API