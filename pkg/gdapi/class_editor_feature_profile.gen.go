// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  EditorFeatureProfileFeatureFeature3D EditorFeatureProfileFeature = 0
  EditorFeatureProfileFeatureFeatureScript EditorFeatureProfileFeature = 1
  EditorFeatureProfileFeatureFeatureAssetLib EditorFeatureProfileFeature = 2
  EditorFeatureProfileFeatureFeatureSceneTree EditorFeatureProfileFeature = 3
  EditorFeatureProfileFeatureFeatureNodeDock EditorFeatureProfileFeature = 4
  EditorFeatureProfileFeatureFeatureFilesystemDock EditorFeatureProfileFeature = 5
  EditorFeatureProfileFeatureFeatureImportDock EditorFeatureProfileFeature = 6
  EditorFeatureProfileFeatureFeatureHistoryDock EditorFeatureProfileFeature = 7
  EditorFeatureProfileFeatureFeatureMax EditorFeatureProfileFeature = 8
)

func  (me *EditorFeatureProfile) SetDisableClass(class_name StringName, disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) IsClassDisabled(class_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) SetDisableClassEditor(class_name StringName, disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) IsClassEditorDisabled(class_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) SetDisableClassProperty(class_name StringName, property StringName, disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) IsClassPropertyDisabled(class_name StringName, property StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) SetDisableFeature(feature EditorFeatureProfileFeature, disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) IsFeatureDisabled(feature EditorFeatureProfileFeature, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) GetFeatureName(feature EditorFeatureProfileFeature, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) SaveToFile(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFeatureProfile) LoadFromFile(path String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
