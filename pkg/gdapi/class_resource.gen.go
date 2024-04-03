// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Resource struct {
  RefCounted
}

func (me *Resource) BaseClass() string {
  return "Resource"
}

func NewResource() *Resource {
  str := StringNameFromStr("Resource") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Resource{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Resource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Resource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Resource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Resource) SetPath(path String, )  {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Resource) TakeOverPath(path String, )  {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("take_over_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Resource) GetPath() String {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Resource) SetName(name String, )  {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Resource) GetName() String {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Resource) GetRid() RID {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Resource) SetLocalToScene(enable bool, )  {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_local_to_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Resource) IsLocalToScene() bool {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_local_to_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Resource) GetLocalScene() Node {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Resource) SetupLocalToScene()  {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("setup_local_to_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Resource) EmitChanged()  {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("emit_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Resource) Duplicate(subresources bool, ) Resource {
  classNameV := StringNameFromStr("Resource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("duplicate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 482882304) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subresources), }
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ResourceChangedSignalFn func()

func (me *Resource) ConnectChanged(subs SignalSubscribers, fn ResourceChangedSignalFn) {
  sig := StringNameFromStr("changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Resource) DisconnectChanged(subs SignalSubscribers, fn ResourceChangedSignalFn) {
  sig := StringNameFromStr("changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ResourceSetupLocalToSceneRequestedSignalFn func()

func (me *Resource) ConnectSetupLocalToSceneRequested(subs SignalSubscribers, fn ResourceSetupLocalToSceneRequestedSignalFn) {
  sig := StringNameFromStr("setup_local_to_scene_requested")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Resource) DisconnectSetupLocalToSceneRequested(subs SignalSubscribers, fn ResourceSetupLocalToSceneRequestedSignalFn) {
  sig := StringNameFromStr("setup_local_to_scene_requested")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
