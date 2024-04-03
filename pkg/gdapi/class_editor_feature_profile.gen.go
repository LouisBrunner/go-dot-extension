// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorFeatureProfile struct {
  RefCounted
}

func (me *EditorFeatureProfile) BaseClass() string {
  return "EditorFeatureProfile"
}

func NewEditorFeatureProfile() *EditorFeatureProfile {
  str := StringNameFromStr("EditorFeatureProfile") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorFeatureProfile{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2524380260) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class_name.AsCTypePtr()), gdc.ConstTypePtr(&disable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFeatureProfile) IsClassDisabled(class_name StringName, ) bool {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_class_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class_name.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFeatureProfile) SetDisableClassEditor(class_name StringName, disable bool, )  {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_class_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2524380260) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class_name.AsCTypePtr()), gdc.ConstTypePtr(&disable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFeatureProfile) IsClassEditorDisabled(class_name StringName, ) bool {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_class_editor_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class_name.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFeatureProfile) SetDisableClassProperty(class_name StringName, property StringName, disable bool, )  {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_class_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 865197084) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class_name.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(&disable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFeatureProfile) IsClassPropertyDisabled(class_name StringName, property StringName, ) bool {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_class_property_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class_name.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFeatureProfile) SetDisableFeature(feature EditorFeatureProfileFeature, disable bool, )  {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_feature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1884871044) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature), gdc.ConstTypePtr(&disable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFeatureProfile) IsFeatureDisabled(feature EditorFeatureProfileFeature, ) bool {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_feature_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2974403161) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFeatureProfile) GetFeatureName(feature EditorFeatureProfileFeature, ) String {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feature_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3401335809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFeatureProfile) SaveToFile(path String, ) Error {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_to_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *EditorFeatureProfile) LoadFromFile(path String, ) Error {
  classNameV := StringNameFromStr("EditorFeatureProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals
