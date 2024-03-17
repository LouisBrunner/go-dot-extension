// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationLibrary struct {
  Resource
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

// Signals

type AnimationLibraryAnimationAddedSignalFn func(name StringName, )

func (me *AnimationLibrary) ConnectAnimationAdded(subs SignalSubscribers, fn AnimationLibraryAnimationAddedSignalFn) {
  sig := StringNameFromStr("animation_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationAdded(subs SignalSubscribers, fn AnimationLibraryAnimationAddedSignalFn) {
  sig := StringNameFromStr("animation_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationLibraryAnimationRemovedSignalFn func(name StringName, )

func (me *AnimationLibrary) ConnectAnimationRemoved(subs SignalSubscribers, fn AnimationLibraryAnimationRemovedSignalFn) {
  sig := StringNameFromStr("animation_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationRemoved(subs SignalSubscribers, fn AnimationLibraryAnimationRemovedSignalFn) {
  sig := StringNameFromStr("animation_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationLibraryAnimationRenamedSignalFn func(name StringName, to_name StringName, )

func (me *AnimationLibrary) ConnectAnimationRenamed(subs SignalSubscribers, fn AnimationLibraryAnimationRenamedSignalFn) {
  sig := StringNameFromStr("animation_renamed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationRenamed(subs SignalSubscribers, fn AnimationLibraryAnimationRenamedSignalFn) {
  sig := StringNameFromStr("animation_renamed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationLibraryAnimationChangedSignalFn func(name StringName, )

func (me *AnimationLibrary) ConnectAnimationChanged(subs SignalSubscribers, fn AnimationLibraryAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationChanged(subs SignalSubscribers, fn AnimationLibraryAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
