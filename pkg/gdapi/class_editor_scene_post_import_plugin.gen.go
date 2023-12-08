// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  EditorScenePostImportPluginInternalImportCategoryNode EditorScenePostImportPluginInternalImportCategory = 0
  EditorScenePostImportPluginInternalImportCategoryMesh3DNode EditorScenePostImportPluginInternalImportCategory = 1
  EditorScenePostImportPluginInternalImportCategoryMesh EditorScenePostImportPluginInternalImportCategory = 2
  EditorScenePostImportPluginInternalImportCategoryMaterial EditorScenePostImportPluginInternalImportCategory = 3
  EditorScenePostImportPluginInternalImportCategoryAnimation EditorScenePostImportPluginInternalImportCategory = 4
  EditorScenePostImportPluginInternalImportCategoryAnimationNode EditorScenePostImportPluginInternalImportCategory = 5
  EditorScenePostImportPluginInternalImportCategorySkeleton3DNode EditorScenePostImportPluginInternalImportCategory = 6
  EditorScenePostImportPluginInternalImportCategoryMax EditorScenePostImportPluginInternalImportCategory = 7
)
