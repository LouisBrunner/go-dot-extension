// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInspectorPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorInspectorPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorInspectorPlugin) BaseClass() string {
  return "EditorInspectorPlugin"
}



// Enums

func (me *EditorInspectorPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorInspectorPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorInspectorPlugin) XCanHandle(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) XParseBegin(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) XParseCategory(object Object, category String, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) XParseGroup(object Object, group String, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) XParseProperty(object Object, type_ VariantType, name String, hint_type PropertyHint, hint_string String, usage_flags PropertyUsageFlags, wide bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) XParseEnd(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) AddCustomControl(control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) AddPropertyEditor(property String, editor Control, add_to_end bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInspectorPlugin) AddPropertyEditorForMultipleProperties(label String, properties PackedStringArray, editor Control, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
