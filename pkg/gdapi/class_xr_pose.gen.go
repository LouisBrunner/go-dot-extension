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

type ptrsForXRPoseList struct {
  fnSetHasTrackingData gdc.MethodBindPtr
  fnGetHasTrackingData gdc.MethodBindPtr
  fnSetName gdc.MethodBindPtr
  fnGetName gdc.MethodBindPtr
  fnSetTransform gdc.MethodBindPtr
  fnGetTransform gdc.MethodBindPtr
  fnGetAdjustedTransform gdc.MethodBindPtr
  fnSetLinearVelocity gdc.MethodBindPtr
  fnGetLinearVelocity gdc.MethodBindPtr
  fnSetAngularVelocity gdc.MethodBindPtr
  fnGetAngularVelocity gdc.MethodBindPtr
  fnSetTrackingConfidence gdc.MethodBindPtr
  fnGetTrackingConfidence gdc.MethodBindPtr
}

var ptrsForXRPose ptrsForXRPoseList

func initXRPosePtrs(iface gdc.Interface) {

  className := StringNameFromStr("XRPose")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_has_tracking_data")
    defer methodName.Destroy()
    ptrsForXRPose.fnSetHasTrackingData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_has_tracking_data")
    defer methodName.Destroy()
    ptrsForXRPose.fnGetHasTrackingData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_name")
    defer methodName.Destroy()
    ptrsForXRPose.fnSetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_name")
    defer methodName.Destroy()
    ptrsForXRPose.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_transform")
    defer methodName.Destroy()
    ptrsForXRPose.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
  }
  {
    methodName := StringNameFromStr("get_transform")
    defer methodName.Destroy()
    ptrsForXRPose.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
  }
  {
    methodName := StringNameFromStr("get_adjusted_transform")
    defer methodName.Destroy()
    ptrsForXRPose.fnGetAdjustedTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
  }
  {
    methodName := StringNameFromStr("set_linear_velocity")
    defer methodName.Destroy()
    ptrsForXRPose.fnSetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_linear_velocity")
    defer methodName.Destroy()
    ptrsForXRPose.fnGetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_angular_velocity")
    defer methodName.Destroy()
    ptrsForXRPose.fnSetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_angular_velocity")
    defer methodName.Destroy()
    ptrsForXRPose.fnGetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_tracking_confidence")
    defer methodName.Destroy()
    ptrsForXRPose.fnSetTrackingConfidence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4171656666))
  }
  {
    methodName := StringNameFromStr("get_tracking_confidence")
    defer methodName.Destroy()
    ptrsForXRPose.fnGetTrackingConfidence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2064923680))
  }
}

type XRPose struct {
  RefCounted
}

func (me *XRPose) BaseClass() string {
  return "XRPose"
}

func NewXRPose() *XRPose {
  str := StringNameFromStr("XRPose") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &XRPose{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type XRPoseTrackingConfidence int
const (
  XRPoseTrackingConfidenceXrTrackingConfidenceNone XRPoseTrackingConfidence = 0
  XRPoseTrackingConfidenceXrTrackingConfidenceLow XRPoseTrackingConfidence = 1
  XRPoseTrackingConfidenceXrTrackingConfidenceHigh XRPoseTrackingConfidence = 2
)

func (me *XRPose) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRPose) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRPose) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XRPose) SetHasTrackingData(has_tracking_data bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&has_tracking_data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnSetHasTrackingData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPose) GetHasTrackingData() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnGetHasTrackingData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRPose) SetName(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnSetName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPose) GetName() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPose) SetTransform(transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPose) GetTransform() Transform3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPose) GetAdjustedTransform() Transform3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnGetAdjustedTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPose) SetLinearVelocity(velocity Vector3, )  {
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnSetLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPose) GetLinearVelocity() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnGetLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPose) SetAngularVelocity(velocity Vector3, )  {
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnSetAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPose) GetAngularVelocity() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnGetAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPose) SetTrackingConfidence(tracking_confidence XRPoseTrackingConfidence, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tracking_confidence) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnSetTrackingConfidence), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPose) GetTrackingConfidence() XRPoseTrackingConfidence {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret XRPoseTrackingConfidence

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPose.fnGetTrackingConfidence), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
