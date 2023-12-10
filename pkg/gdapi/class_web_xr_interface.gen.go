// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WebXRInterface struct {
  obj gdc.ObjectPtr
}

func (me *WebXRInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebXRInterface) BaseClass() string {
  return "WebXRInterface"
}



// Enums

type WebXRInterfaceTargetRayMode int
const (
  WebXRInterfaceTargetRayModeTargetRayModeUnknown WebXRInterfaceTargetRayMode = 0
  WebXRInterfaceTargetRayModeTargetRayModeGaze WebXRInterfaceTargetRayMode = 1
  WebXRInterfaceTargetRayModeTargetRayModeTrackedPointer WebXRInterfaceTargetRayMode = 2
  WebXRInterfaceTargetRayModeTargetRayModeScreen WebXRInterfaceTargetRayMode = 3
)

func (me *WebXRInterface) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebXRInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebXRInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WebXRInterface) IsSessionSupported(session_mode String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_session_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(session_mode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) SetSessionMode(session_mode String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_session_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(session_mode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetSessionMode() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_session_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetRequiredFeatures(required_features String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_required_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(required_features.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetRequiredFeatures() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_required_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetOptionalFeatures(optional_features String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_optional_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(optional_features.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetOptionalFeatures() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_optional_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetReferenceSpaceType() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reference_space_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetRequestedReferenceSpaceTypes(requested_reference_space_types String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_requested_reference_space_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(requested_reference_space_types.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetRequestedReferenceSpaceTypes() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_requested_reference_space_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) IsInputSourceActive(input_source_id int, ) bool {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_input_source_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetInputSourceTracker(input_source_id int, ) XRPositionalTracker {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_source_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 636011756) // FIXME: should cache?
  var ret XRPositionalTracker
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetInputSourceTargetRayMode(input_source_id int, ) WebXRInterfaceTargetRayMode {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_source_target_ray_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2852387453) // FIXME: should cache?
  var ret WebXRInterfaceTargetRayMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetVisibilityState() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetDisplayRefreshRate() float32 {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_refresh_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetDisplayRefreshRate(refresh_rate float32, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_refresh_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refresh_rate), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetAvailableDisplayRefreshRates() Array {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_available_display_refresh_rates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *WebXRInterface) GetPropSessionMode() String {
  panic("TODO: implement")
}

func (me *WebXRInterface) SetPropSessionMode(value String) {
  panic("TODO: implement")
}

func (me *WebXRInterface) GetPropRequiredFeatures() String {
  panic("TODO: implement")
}

func (me *WebXRInterface) SetPropRequiredFeatures(value String) {
  panic("TODO: implement")
}

func (me *WebXRInterface) GetPropOptionalFeatures() String {
  panic("TODO: implement")
}

func (me *WebXRInterface) SetPropOptionalFeatures(value String) {
  panic("TODO: implement")
}

func (me *WebXRInterface) GetPropRequestedReferenceSpaceTypes() String {
  panic("TODO: implement")
}

func (me *WebXRInterface) SetPropRequestedReferenceSpaceTypes(value String) {
  panic("TODO: implement")
}

func (me *WebXRInterface) GetPropReferenceSpaceType() String {
  panic("TODO: implement")
}

func (me *WebXRInterface) SetPropReferenceSpaceType(value String) {
  panic("TODO: implement")
}

func (me *WebXRInterface) GetPropVisibilityState() String {
  panic("TODO: implement")
}

func (me *WebXRInterface) SetPropVisibilityState(value String) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API