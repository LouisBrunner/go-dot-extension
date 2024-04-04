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

type RemoteTransform3D struct {
  Node3D
}

func (me *RemoteTransform3D) BaseClass() string {
  return "RemoteTransform3D"
}

func NewRemoteTransform3D() *RemoteTransform3D {
  str := StringNameFromStr("RemoteTransform3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RemoteTransform3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RemoteTransform3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RemoteTransform3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RemoteTransform3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RemoteTransform3D) SetRemoteNode(path NodePath, )  {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_remote_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetRemoteNode() NodePath {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_remote_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RemoteTransform3D) ForceUpdateCache()  {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) SetUseGlobalCoordinates(use_global_coordinates bool, )  {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_global_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_global_coordinates) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUseGlobalCoordinates() bool {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_global_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RemoteTransform3D) SetUpdatePosition(update_remote_position bool, )  {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_remote_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUpdatePosition() bool {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RemoteTransform3D) SetUpdateRotation(update_remote_rotation bool, )  {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_remote_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUpdateRotation() bool {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RemoteTransform3D) SetUpdateScale(update_remote_scale bool, )  {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_remote_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUpdateScale() bool {
  classNameV := StringNameFromStr("RemoteTransform3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
