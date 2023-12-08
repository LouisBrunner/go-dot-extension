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

func  (me *XRInterfaceExtension) XGetName() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetCapabilities() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XIsInitialized() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XInitialize() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XUninitialize() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetSystemInfo() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XSupportsPlayAreaMode(mode XRInterfacePlayAreaMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetPlayAreaMode() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XSetPlayAreaMode(mode XRInterfacePlayAreaMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetPlayArea() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetRenderTargetSize() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetViewCount() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetCameraTransform() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetTransformForView(view int, cam_transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetProjectionForView(view int, aspect float32, z_near float32, z_far float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetVrsTexture() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XProcess() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XPreRender() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XPreDrawViewport(render_target RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XPostDrawViewport(render_target RID, screen_rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XEndFrame() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetSuggestedTrackerNames() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetSuggestedPoseNames(tracker_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetTrackingStatus() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XTriggerHapticPulse(action_name String, tracker_name StringName, frequency float32, amplitude float32, duration_sec float32, delay_sec float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetAnchorDetectionIsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XSetAnchorDetectionIsEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetCameraFeedId() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetColorTexture() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetDepthTexture() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) XGetVelocityTexture() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) GetColorTexture() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) GetDepthTexture() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) GetVelocityTexture() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) AddBlit(render_target RID, src_rect Rect2, dst_rect Rect2i, use_layer bool, layer int, apply_lens_distortion bool, eye_center Vector2, k1 float32, k2 float32, upscale float32, aspect_ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterfaceExtension) GetRenderTargetTexture(render_target RID, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
