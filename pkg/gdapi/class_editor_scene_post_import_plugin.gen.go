// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorScenePostImportPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorScenePostImportPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorScenePostImportPlugin) BaseClass() string {
  return "EditorScenePostImportPlugin"
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
  classNameV := StringNameFromStr("EditorScenePostImportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_option_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorScenePostImportPlugin) AddImportOption(name String, value Variant, )  {
  classNameV := StringNameFromStr("EditorScenePostImportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_import_option")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402577236) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorScenePostImportPlugin) AddImportOptionAdvanced(type_ VariantType, name String, default_value Variant, hint PropertyHint, hint_string String, usage_flags int, )  {
  classNameV := StringNameFromStr("EditorScenePostImportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_import_option_advanced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3774155785) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(default_value.AsCTypePtr()), gdc.ConstTypePtr(&hint), gdc.ConstTypePtr(hint_string.AsCTypePtr()), gdc.ConstTypePtr(&usage_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties