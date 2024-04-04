// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type InstancePlaceholder struct {
  Node
}

func (me *InstancePlaceholder) BaseClass() string {
  return "InstancePlaceholder"
}

func NewInstancePlaceholder() *InstancePlaceholder {
  str := StringNameFromStr("InstancePlaceholder") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InstancePlaceholder{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&with_order) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&with_order)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *InstancePlaceholder) CreateInstance(replace bool, custom_scene PackedScene, ) Node {
  classNameV := StringNameFromStr("InstancePlaceholder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3794612210) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&replace) , custom_scene.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()
  pinner.Pin(&replace)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *InstancePlaceholder) GetInstancePath() String {
  classNameV := StringNameFromStr("InstancePlaceholder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
