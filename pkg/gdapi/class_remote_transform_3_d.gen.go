// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RemoteTransform3D struct {
  obj gdc.ObjectPtr
}

func (me *RemoteTransform3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RemoteTransform3D) BaseClass() string {
  return "RemoteTransform3D"
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
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) GetRemoteNode()  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) ForceUpdateCache()  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) SetUseGlobalCoordinates(use_global_coordinates bool, )  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) GetUseGlobalCoordinates()  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) SetUpdatePosition(update_remote_position bool, )  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) GetUpdatePosition()  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) SetUpdateRotation(update_remote_rotation bool, )  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) GetUpdateRotation()  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) SetUpdateScale(update_remote_scale bool, )  {
  panic("TODO: implement")
}

func  (me *RemoteTransform3D) GetUpdateScale()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
