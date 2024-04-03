// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorInspectorPlugin struct {
  RefCounted
}

func (me *EditorInspectorPlugin) BaseClass() string {
  return "EditorInspectorPlugin"
}

func NewEditorInspectorPlugin() *EditorInspectorPlugin {
  str := StringNameFromStr("EditorInspectorPlugin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorInspectorPlugin{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorInspectorPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorInspectorPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorInspectorPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorInspectorPlugin) AddCustomControl(control Control, )  {
  classNameV := StringNameFromStr("EditorInspectorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_custom_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(control.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInspectorPlugin) AddPropertyEditor(property String, editor Control, add_to_end bool, )  {
  classNameV := StringNameFromStr("EditorInspectorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_property_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3406284123) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(editor.AsCTypePtr()), gdc.ConstTypePtr(&add_to_end), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInspectorPlugin) AddPropertyEditorForMultipleProperties(label String, properties PackedStringArray, editor Control, )  {
  classNameV := StringNameFromStr("EditorInspectorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_property_editor_for_multiple_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 788598683) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(properties.AsCTypePtr()), gdc.ConstTypePtr(editor.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
