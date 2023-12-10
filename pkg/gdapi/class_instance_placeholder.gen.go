// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InstancePlaceholder struct {
  obj gdc.ObjectPtr
}

func (me *InstancePlaceholder) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InstancePlaceholder) BaseClass() string {
  return "InstancePlaceholder"
}



// Enums

func (me *InstancePlaceholder) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InstancePlaceholder) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InstancePlaceholder) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InstancePlaceholder) GetStoredValues(with_order bool, ) Dictionary {
  classNameV := StringNameFromStr("InstancePlaceholder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stored_values")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2230153369) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&with_order), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InstancePlaceholder) CreateInstance(replace bool, custom_scene PackedScene, ) Node {
  classNameV := StringNameFromStr("InstancePlaceholder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3794612210) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&replace), gdc.ConstTypePtr(custom_scene.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InstancePlaceholder) GetInstancePath() String {
  classNameV := StringNameFromStr("InstancePlaceholder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties