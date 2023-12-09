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



// Enums

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

func (me *Camera2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Camera2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Camera2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Camera2D) SetOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetOffset()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetAnchorMode(anchor_mode Camera2DAnchorMode, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetAnchorMode()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetIgnoreRotation(ignore bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsIgnoringRotation()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetProcessCallback(mode Camera2DCamera2DProcessCallback, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetProcessCallback()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) MakeCurrent()  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsCurrent()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetLimit(margin Side, limit int, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetLimit(margin Side, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetLimitSmoothingEnabled(limit_smoothing_enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsLimitSmoothingEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetDragVerticalEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsDragVerticalEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetDragHorizontalEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsDragHorizontalEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetDragVerticalOffset(offset float32, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetDragVerticalOffset()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetDragHorizontalOffset(offset float32, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetDragHorizontalOffset()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetDragMargin(margin Side, drag_margin float32, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetDragMargin(margin Side, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetScreenCenterPosition()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetZoom(zoom Vector2, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetZoom()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetCustomViewport(viewport Node, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetCustomViewport()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetPositionSmoothingSpeed(position_smoothing_speed float32, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetPositionSmoothingSpeed()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetPositionSmoothingEnabled(position_smoothing_speed bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsPositionSmoothingEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetRotationSmoothingEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsRotationSmoothingEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetRotationSmoothingSpeed(speed float32, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) GetRotationSmoothingSpeed()  {
  panic("TODO: implement")
}

func  (me *Camera2D) ForceUpdateScroll()  {
  panic("TODO: implement")
}

func  (me *Camera2D) ResetSmoothing()  {
  panic("TODO: implement")
}

func  (me *Camera2D) Align()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetScreenDrawingEnabled(screen_drawing_enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsScreenDrawingEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetLimitDrawingEnabled(limit_drawing_enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsLimitDrawingEnabled()  {
  panic("TODO: implement")
}

func  (me *Camera2D) SetMarginDrawingEnabled(margin_drawing_enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera2D) IsMarginDrawingEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
