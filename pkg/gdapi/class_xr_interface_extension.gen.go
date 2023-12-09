// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRInterfaceExtension struct {
  obj gdc.ObjectPtr
}

func (me *XRInterfaceExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRInterfaceExtension) BaseClass() string {
  return "XRInterfaceExtension"
}



// Enums

func (me *XRInterfaceExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRInterfaceExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRInterfaceExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XRInterfaceExtension) XGetName()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetCapabilities()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XIsInitialized()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XInitialize()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XUninitialize()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetSystemInfo()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XSupportsPlayAreaMode(mode XRInterfacePlayAreaMode, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetPlayAreaMode()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XSetPlayAreaMode(mode XRInterfacePlayAreaMode, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetPlayArea()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetRenderTargetSize()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetViewCount()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetCameraTransform()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetTransformForView(view int, cam_transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetProjectionForView(view int, aspect float32, z_near float32, z_far float32, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetVrsTexture()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XProcess()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XPreRender()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XPreDrawViewport(render_target RID, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XPostDrawViewport(render_target RID, screen_rect Rect2, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XEndFrame()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetSuggestedTrackerNames()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetSuggestedPoseNames(tracker_name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetTrackingStatus()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XTriggerHapticPulse(action_name String, tracker_name StringName, frequency float32, amplitude float32, duration_sec float32, delay_sec float32, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetAnchorDetectionIsEnabled()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XSetAnchorDetectionIsEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetCameraFeedId()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetColorTexture()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetDepthTexture()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) XGetVelocityTexture()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) GetColorTexture()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) GetDepthTexture()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) GetVelocityTexture()  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) AddBlit(render_target RID, src_rect Rect2, dst_rect Rect2i, use_layer bool, layer int, apply_lens_distortion bool, eye_center Vector2, k1 float32, k2 float32, upscale float32, aspect_ratio float32, )  {
  panic("TODO: implement")
}

func  (me *XRInterfaceExtension) GetRenderTargetTexture(render_target RID, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
