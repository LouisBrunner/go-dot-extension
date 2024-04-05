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

type ptrsForEditorScenePostImportPluginList struct {
  fnXGetInternalImportOptions gdc.MethodBindPtr
  fnXGetInternalOptionVisibility gdc.MethodBindPtr
  fnXGetInternalOptionUpdateViewRequired gdc.MethodBindPtr
  fnXInternalProcess gdc.MethodBindPtr
  fnXGetImportOptions gdc.MethodBindPtr
  fnXGetOptionVisibility gdc.MethodBindPtr
  fnXPreProcess gdc.MethodBindPtr
  fnXPostProcess gdc.MethodBindPtr
  fnGetOptionValue gdc.MethodBindPtr
  fnAddImportOption gdc.MethodBindPtr
  fnAddImportOptionAdvanced gdc.MethodBindPtr
}

var ptrsForEditorScenePostImportPlugin ptrsForEditorScenePostImportPluginList

func initEditorScenePostImportPluginPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorScenePostImportPlugin")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_option_value")
    defer methodName.Destroy()
    ptrsForEditorScenePostImportPlugin.fnGetOptionValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
  }
  {
    methodName := StringNameFromStr("add_import_option")
    defer methodName.Destroy()
    ptrsForEditorScenePostImportPlugin.fnAddImportOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 402577236))
  }
  {
    methodName := StringNameFromStr("add_import_option_advanced")
    defer methodName.Destroy()
    ptrsForEditorScenePostImportPlugin.fnAddImportOptionAdvanced = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3674075649))
  }
}

type EditorScenePostImportPlugin struct {
  RefCounted
}

func (me *EditorScenePostImportPlugin) BaseClass() string {
  return "EditorScenePostImportPlugin"
}

func NewEditorScenePostImportPlugin() *EditorScenePostImportPlugin {
  str := StringNameFromStr("EditorScenePostImportPlugin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorScenePostImportPlugin{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type EditorScenePostImportPluginInternalImportCategory int
const (
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategoryNode EditorScenePostImportPluginInternalImportCategory = 0
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategoryMesh3DNode EditorScenePostImportPluginInternalImportCategory = 1
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategoryMesh EditorScenePostImportPluginInternalImportCategory = 2
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategoryMaterial EditorScenePostImportPluginInternalImportCategory = 3
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategoryAnimation EditorScenePostImportPluginInternalImportCategory = 4
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategoryAnimationNode EditorScenePostImportPluginInternalImportCategory = 5
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategorySkeleton3DNode EditorScenePostImportPluginInternalImportCategory = 6
  EditorScenePostImportPluginInternalImportCategoryInternalImportCategoryMax EditorScenePostImportPluginInternalImportCategory = 7
)

func (me *EditorScenePostImportPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorScenePostImportPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorScenePostImportPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorScenePostImportPlugin) GetOptionValue(name StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScenePostImportPlugin.fnGetOptionValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorScenePostImportPlugin) AddImportOption(name String, value Variant, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScenePostImportPlugin.fnAddImportOption), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorScenePostImportPlugin) AddImportOptionAdvanced(type_ VariantType, name String, default_value Variant, hint PropertyHint, hint_string String, usage_flags int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , name.AsCTypePtr(), default_value.AsCTypePtr(), gdc.ConstTypePtr(&hint) , hint_string.AsCTypePtr(), gdc.ConstTypePtr(&usage_flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScenePostImportPlugin.fnAddImportOptionAdvanced), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
