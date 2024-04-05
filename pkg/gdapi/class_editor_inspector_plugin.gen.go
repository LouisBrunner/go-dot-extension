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

type ptrsForEditorInspectorPluginList struct {
  fnXCanHandle gdc.MethodBindPtr
  fnXParseBegin gdc.MethodBindPtr
  fnXParseCategory gdc.MethodBindPtr
  fnXParseGroup gdc.MethodBindPtr
  fnXParseProperty gdc.MethodBindPtr
  fnXParseEnd gdc.MethodBindPtr
  fnAddCustomControl gdc.MethodBindPtr
  fnAddPropertyEditor gdc.MethodBindPtr
  fnAddPropertyEditorForMultipleProperties gdc.MethodBindPtr
}

var ptrsForEditorInspectorPlugin ptrsForEditorInspectorPluginList

func initEditorInspectorPluginPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorInspectorPlugin")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_custom_control")
    defer methodName.Destroy()
    ptrsForEditorInspectorPlugin.fnAddCustomControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
  }
  {
    methodName := StringNameFromStr("add_property_editor")
    defer methodName.Destroy()
    ptrsForEditorInspectorPlugin.fnAddPropertyEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3406284123))
  }
  {
    methodName := StringNameFromStr("add_property_editor_for_multiple_properties")
    defer methodName.Destroy()
    ptrsForEditorInspectorPlugin.fnAddPropertyEditorForMultipleProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 788598683))
  }
}

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
  cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInspectorPlugin.fnAddCustomControl), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInspectorPlugin) AddPropertyEditor(property String, editor Control, add_to_end bool, )  {
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), editor.AsCTypePtr(), gdc.ConstTypePtr(&add_to_end) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInspectorPlugin.fnAddPropertyEditor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInspectorPlugin) AddPropertyEditorForMultipleProperties(label String, properties PackedStringArray, editor Control, )  {
  cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), properties.AsCTypePtr(), editor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInspectorPlugin.fnAddPropertyEditorForMultipleProperties), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
