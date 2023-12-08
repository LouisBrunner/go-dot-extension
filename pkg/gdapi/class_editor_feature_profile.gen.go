// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorFeatureProfile struct {
  obj gdc.ObjectPtr
}

func (me *EditorFeatureProfile) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorFeatureProfile) BaseClass() string {
  return "EditorFeatureProfile"
}

type EditorFeatureProfileFeature int
const (
  EditorFeatureProfileFeature3D EditorFeatureProfileFeature = 0
  EditorFeatureProfileFeatureScript EditorFeatureProfileFeature = 1
  EditorFeatureProfileFeatureAssetLib EditorFeatureProfileFeature = 2
  EditorFeatureProfileFeatureSceneTree EditorFeatureProfileFeature = 3
  EditorFeatureProfileFeatureNodeDock EditorFeatureProfileFeature = 4
  EditorFeatureProfileFeatureFilesystemDock EditorFeatureProfileFeature = 5
  EditorFeatureProfileFeatureImportDock EditorFeatureProfileFeature = 6
  EditorFeatureProfileFeatureHistoryDock EditorFeatureProfileFeature = 7
  EditorFeatureProfileFeatureMax EditorFeatureProfileFeature = 8
)
