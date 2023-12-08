// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScenePostImportPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorScenePostImportPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorScenePostImportPlugin) BaseClass() string {
  return "EditorScenePostImportPlugin"
}

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

func  (me *EditorScenePostImportPlugin) XGetInternalImportOptions(category int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) XGetInternalOptionVisibility(category int, for_animation bool, option String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) XGetInternalOptionUpdateViewRequired(category int, option String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) XInternalProcess(category int, base_node Node, node Node, resource Resource, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) XGetImportOptions(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) XGetOptionVisibility(path String, for_animation bool, option String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) XPreProcess(scene Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) XPostProcess(scene Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) GetOptionValue(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) AddImportOption(name String, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorScenePostImportPlugin) AddImportOptionAdvanced(type_ VariantType, name String, default_value Variant, hint PropertyHint, hint_string String, usage_flags int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
