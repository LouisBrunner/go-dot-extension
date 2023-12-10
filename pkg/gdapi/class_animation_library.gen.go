// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationLibrary struct {
  obj gdc.ObjectPtr
}

func (me *AnimationLibrary) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationLibrary) BaseClass() string {
  return "AnimationLibrary"
}



// Enums

func (me *AnimationLibrary) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationLibrary) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationLibrary) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationLibrary) AddAnimation(name StringName, animation Animation, ) Error {
  classNameV := StringNameFromStr("AnimationLibrary")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1811855551) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(animation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationLibrary) RemoveAnimation(name StringName, )  {
  classNameV := StringNameFromStr("AnimationLibrary")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationLibrary) RenameAnimation(name StringName, newname StringName, )  {
  classNameV := StringNameFromStr("AnimationLibrary")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(newname.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationLibrary) HasAnimation(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationLibrary")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationLibrary) GetAnimation(name StringName, ) Animation {
  classNameV := StringNameFromStr("AnimationLibrary")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2933122410) // FIXME: should cache?
  var ret Animation
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationLibrary) GetAnimationList() StringName {
  classNameV := StringNameFromStr("AnimationLibrary")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API