// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForRemoteTransform3DList struct {
  fnSetRemoteNode gdc.MethodBindPtr
  fnGetRemoteNode gdc.MethodBindPtr
  fnForceUpdateCache gdc.MethodBindPtr
  fnSetUseGlobalCoordinates gdc.MethodBindPtr
  fnGetUseGlobalCoordinates gdc.MethodBindPtr
  fnSetUpdatePosition gdc.MethodBindPtr
  fnGetUpdatePosition gdc.MethodBindPtr
  fnSetUpdateRotation gdc.MethodBindPtr
  fnGetUpdateRotation gdc.MethodBindPtr
  fnSetUpdateScale gdc.MethodBindPtr
  fnGetUpdateScale gdc.MethodBindPtr
}

var ptrsForRemoteTransform3D ptrsForRemoteTransform3DList

func initRemoteTransform3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RemoteTransform3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_remote_node")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnSetRemoteNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_remote_node")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnGetRemoteNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("force_update_cache")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnForceUpdateCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_use_global_coordinates")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnSetUseGlobalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_use_global_coordinates")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnGetUseGlobalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_update_position")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnSetUpdatePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_update_position")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnGetUpdatePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_update_rotation")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnSetUpdateRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_update_rotation")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnGetUpdateRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_update_scale")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnSetUpdateScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_update_scale")
    defer methodName.Destroy()
    ptrsForRemoteTransform3D.fnGetUpdateScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnSetRemoteNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetRemoteNode() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnGetRemoteNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RemoteTransform3D) ForceUpdateCache()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnForceUpdateCache), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) SetUseGlobalCoordinates(use_global_coordinates bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_global_coordinates) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnSetUseGlobalCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUseGlobalCoordinates() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnGetUseGlobalCoordinates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RemoteTransform3D) SetUpdatePosition(update_remote_position bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_remote_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnSetUpdatePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUpdatePosition() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnGetUpdatePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RemoteTransform3D) SetUpdateRotation(update_remote_rotation bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_remote_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnSetUpdateRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUpdateRotation() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnGetUpdateRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RemoteTransform3D) SetUpdateScale(update_remote_scale bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_remote_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnSetUpdateScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RemoteTransform3D) GetUpdateScale() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRemoteTransform3D.fnGetUpdateScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
