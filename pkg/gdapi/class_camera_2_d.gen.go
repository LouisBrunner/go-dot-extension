// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera2D struct {
  obj gdc.ObjectPtr
}

func (me *Camera2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Camera2D) BaseClass() string {
  return "Camera2D"
}

type Camera2DAnchorMode int
const (
  Camera2DAnchorModeAnchorModeFixedTopLeft Camera2DAnchorMode = 0
  Camera2DAnchorModeAnchorModeDragCenter Camera2DAnchorMode = 1
)

type Camera2DCamera2DProcessCallback int
const (
  Camera2DCamera2DProcessCallbackCamera2DProcessPhysics Camera2DCamera2DProcessCallback = 0
  Camera2DCamera2DProcessCallbackCamera2DProcessIdle Camera2DCamera2DProcessCallback = 1
)

func  (me *Camera2D) SetOffset(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetOffset() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetAnchorMode(anchor_mode Camera2DAnchorMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetAnchorMode() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetIgnoreRotation(ignore bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsIgnoringRotation() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetProcessCallback(mode Camera2DCamera2DProcessCallback, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetProcessCallback() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) MakeCurrent() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsCurrent() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetLimit(margin Side, limit int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetLimit(margin Side, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetLimitSmoothingEnabled(limit_smoothing_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsLimitSmoothingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetDragVerticalEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsDragVerticalEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetDragHorizontalEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsDragHorizontalEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetDragVerticalOffset(offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetDragVerticalOffset() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetDragHorizontalOffset(offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetDragHorizontalOffset() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetDragMargin(margin Side, drag_margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetDragMargin(margin Side, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetTargetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetScreenCenterPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetZoom(zoom Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetZoom() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetCustomViewport(viewport Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetCustomViewport() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetPositionSmoothingSpeed(position_smoothing_speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetPositionSmoothingSpeed() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetPositionSmoothingEnabled(position_smoothing_speed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsPositionSmoothingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetRotationSmoothingEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsRotationSmoothingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetRotationSmoothingSpeed(speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) GetRotationSmoothingSpeed() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) ForceUpdateScroll() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) ResetSmoothing() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) Align() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetScreenDrawingEnabled(screen_drawing_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsScreenDrawingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetLimitDrawingEnabled(limit_drawing_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsLimitDrawingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) SetMarginDrawingEnabled(margin_drawing_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Camera2D) IsMarginDrawingEnabled() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
