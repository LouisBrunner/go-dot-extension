// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasItem struct {
  obj gdc.ObjectPtr
}

func (me *CanvasItem) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasItem) BaseClass() string {
  return "CanvasItem"
}

const (
  CanvasItemNOTIFICATION_TRANSFORM_CHANGED = 2000
  CanvasItemNOTIFICATION_LOCAL_TRANSFORM_CHANGED = 35
  CanvasItemNOTIFICATION_DRAW = 30
  CanvasItemNOTIFICATION_VISIBILITY_CHANGED = 31
  CanvasItemNOTIFICATION_ENTER_CANVAS = 32
  CanvasItemNOTIFICATION_EXIT_CANVAS = 33
  CanvasItemNOTIFICATION_WORLD_2D_CHANGED = 36
)

type CanvasItemTextureFilter int
const (
  CanvasItemTextureFilterParentNode CanvasItemTextureFilter = 0
  CanvasItemTextureFilterNearest CanvasItemTextureFilter = 1
  CanvasItemTextureFilterLinear CanvasItemTextureFilter = 2
  CanvasItemTextureFilterNearestWithMipmaps CanvasItemTextureFilter = 3
  CanvasItemTextureFilterLinearWithMipmaps CanvasItemTextureFilter = 4
  CanvasItemTextureFilterNearestWithMipmapsAnisotropic CanvasItemTextureFilter = 5
  CanvasItemTextureFilterLinearWithMipmapsAnisotropic CanvasItemTextureFilter = 6
  CanvasItemTextureFilterMax CanvasItemTextureFilter = 7
)

type CanvasItemTextureRepeat int
const (
  CanvasItemTextureRepeatParentNode CanvasItemTextureRepeat = 0
  CanvasItemTextureRepeatDisabled CanvasItemTextureRepeat = 1
  CanvasItemTextureRepeatEnabled CanvasItemTextureRepeat = 2
  CanvasItemTextureRepeatMirror CanvasItemTextureRepeat = 3
  CanvasItemTextureRepeatMax CanvasItemTextureRepeat = 4
)

type CanvasItemClipChildrenMode int
const (
  CanvasItemClipChildrenDisabled CanvasItemClipChildrenMode = 0
  CanvasItemClipChildrenOnly CanvasItemClipChildrenMode = 1
  CanvasItemClipChildrenAndDraw CanvasItemClipChildrenMode = 2
  CanvasItemClipChildrenMax CanvasItemClipChildrenMode = 3
)
