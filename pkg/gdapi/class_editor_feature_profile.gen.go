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



// Enums

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

func (me *EditorFeatureProfile) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorFeatureProfile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorFeatureProfile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorFeatureProfile) SetDisableClass(class_name StringName, disable bool, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) IsClassDisabled(class_name StringName, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) SetDisableClassEditor(class_name StringName, disable bool, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) IsClassEditorDisabled(class_name StringName, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) SetDisableClassProperty(class_name StringName, property StringName, disable bool, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) IsClassPropertyDisabled(class_name StringName, property StringName, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) SetDisableFeature(feature EditorFeatureProfileFeature, disable bool, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) IsFeatureDisabled(feature EditorFeatureProfileFeature, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) GetFeatureName(feature EditorFeatureProfileFeature, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) SaveToFile(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorFeatureProfile) LoadFromFile(path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
