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

type ptrsForSpriteBase3DList struct {
  fnSetCentered gdc.MethodBindPtr
  fnIsCentered gdc.MethodBindPtr
  fnSetOffset gdc.MethodBindPtr
  fnGetOffset gdc.MethodBindPtr
  fnSetFlipH gdc.MethodBindPtr
  fnIsFlippedH gdc.MethodBindPtr
  fnSetFlipV gdc.MethodBindPtr
  fnIsFlippedV gdc.MethodBindPtr
  fnSetModulate gdc.MethodBindPtr
  fnGetModulate gdc.MethodBindPtr
  fnSetRenderPriority gdc.MethodBindPtr
  fnGetRenderPriority gdc.MethodBindPtr
  fnSetPixelSize gdc.MethodBindPtr
  fnGetPixelSize gdc.MethodBindPtr
  fnSetAxis gdc.MethodBindPtr
  fnGetAxis gdc.MethodBindPtr
  fnSetDrawFlag gdc.MethodBindPtr
  fnGetDrawFlag gdc.MethodBindPtr
  fnSetAlphaCutMode gdc.MethodBindPtr
  fnGetAlphaCutMode gdc.MethodBindPtr
  fnSetAlphaScissorThreshold gdc.MethodBindPtr
  fnGetAlphaScissorThreshold gdc.MethodBindPtr
  fnSetAlphaHashScale gdc.MethodBindPtr
  fnGetAlphaHashScale gdc.MethodBindPtr
  fnSetAlphaAntialiasing gdc.MethodBindPtr
  fnGetAlphaAntialiasing gdc.MethodBindPtr
  fnSetAlphaAntialiasingEdge gdc.MethodBindPtr
  fnGetAlphaAntialiasingEdge gdc.MethodBindPtr
  fnSetBillboardMode gdc.MethodBindPtr
  fnGetBillboardMode gdc.MethodBindPtr
  fnSetTextureFilter gdc.MethodBindPtr
  fnGetTextureFilter gdc.MethodBindPtr
  fnGetItemRect gdc.MethodBindPtr
  fnGenerateTriangleMesh gdc.MethodBindPtr
}

var ptrsForSpriteBase3D ptrsForSpriteBase3DList

func initSpriteBase3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SpriteBase3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_centered")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_centered")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnIsCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_offset")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_offset")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_flip_h")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetFlipH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flipped_h")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnIsFlippedH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_flip_v")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetFlipV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flipped_v")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnIsFlippedV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_modulate")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_modulate")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_render_priority")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_render_priority")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_pixel_size")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetPixelSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_pixel_size")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetPixelSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_axis")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1144690656))
  }
  {
    methodName := StringNameFromStr("get_axis")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3050976882))
  }
  {
    methodName := StringNameFromStr("set_draw_flag")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetDrawFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1135633219))
  }
  {
    methodName := StringNameFromStr("get_draw_flag")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetDrawFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1733036628))
  }
  {
    methodName := StringNameFromStr("set_alpha_cut_mode")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetAlphaCutMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 227561226))
  }
  {
    methodName := StringNameFromStr("get_alpha_cut_mode")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetAlphaCutMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 336003791))
  }
  {
    methodName := StringNameFromStr("set_alpha_scissor_threshold")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetAlphaScissorThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_alpha_scissor_threshold")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetAlphaScissorThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_alpha_hash_scale")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetAlphaHashScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_alpha_hash_scale")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetAlphaHashScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_alpha_antialiasing")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetAlphaAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3212649852))
  }
  {
    methodName := StringNameFromStr("get_alpha_antialiasing")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetAlphaAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2889939400))
  }
  {
    methodName := StringNameFromStr("set_alpha_antialiasing_edge")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetAlphaAntialiasingEdge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_alpha_antialiasing_edge")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetAlphaAntialiasingEdge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_billboard_mode")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetBillboardMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4202036497))
  }
  {
    methodName := StringNameFromStr("get_billboard_mode")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetBillboardMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1283840139))
  }
  {
    methodName := StringNameFromStr("set_texture_filter")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnSetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 22904437))
  }
  {
    methodName := StringNameFromStr("get_texture_filter")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289213076))
  }
  {
    methodName := StringNameFromStr("get_item_rect")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGetItemRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
  }
  {
    methodName := StringNameFromStr("generate_triangle_mesh")
    defer methodName.Destroy()
    ptrsForSpriteBase3D.fnGenerateTriangleMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3476533166))
  }
}

type SpriteBase3D struct {
  GeometryInstance3D
}

func (me *SpriteBase3D) BaseClass() string {
  return "SpriteBase3D"
}

func NewSpriteBase3D() *SpriteBase3D {
  str := StringNameFromStr("SpriteBase3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SpriteBase3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type SpriteBase3DDrawFlags int
const (
  SpriteBase3DDrawFlagsFlagTransparent SpriteBase3DDrawFlags = 0
  SpriteBase3DDrawFlagsFlagShaded SpriteBase3DDrawFlags = 1
  SpriteBase3DDrawFlagsFlagDoubleSided SpriteBase3DDrawFlags = 2
  SpriteBase3DDrawFlagsFlagDisableDepthTest SpriteBase3DDrawFlags = 3
  SpriteBase3DDrawFlagsFlagFixedSize SpriteBase3DDrawFlags = 4
  SpriteBase3DDrawFlagsFlagMax SpriteBase3DDrawFlags = 5
)

type SpriteBase3DAlphaCutMode int
const (
  SpriteBase3DAlphaCutModeAlphaCutDisabled SpriteBase3DAlphaCutMode = 0
  SpriteBase3DAlphaCutModeAlphaCutDiscard SpriteBase3DAlphaCutMode = 1
  SpriteBase3DAlphaCutModeAlphaCutOpaquePrepass SpriteBase3DAlphaCutMode = 2
  SpriteBase3DAlphaCutModeAlphaCutHash SpriteBase3DAlphaCutMode = 3
)

func (me *SpriteBase3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SpriteBase3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpriteBase3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SpriteBase3D) SetCentered(centered bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&centered) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetCentered), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) IsCentered() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnIsCentered), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteBase3D) SetFlipH(flip_h bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetFlipH), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) IsFlippedH() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnIsFlippedH), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetFlipV(flip_v bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetFlipV), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) IsFlippedV() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnIsFlippedV), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetModulate(modulate Color, )  {
  cargs := []gdc.ConstTypePtr{modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetModulate() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteBase3D) SetRenderPriority(priority int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetRenderPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetRenderPriority() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetRenderPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetPixelSize(pixel_size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixel_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetPixelSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetPixelSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetPixelSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAxis(axis Vector3Axis, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetAxis), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAxis() Vector3Axis {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Vector3Axis

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetAxis), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetDrawFlag(flag SpriteBase3DDrawFlags, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetDrawFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetDrawFlag(flag SpriteBase3DDrawFlags, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&flag)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetDrawFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAlphaCutMode(mode SpriteBase3DAlphaCutMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetAlphaCutMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaCutMode() SpriteBase3DAlphaCutMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SpriteBase3DAlphaCutMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetAlphaCutMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetAlphaScissorThreshold(threshold float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetAlphaScissorThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaScissorThreshold() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetAlphaScissorThreshold), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAlphaHashScale(threshold float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetAlphaHashScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaHashScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetAlphaHashScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alpha_aa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetAlphaAntialiasing), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaAntialiasing() BaseMaterial3DAlphaAntiAliasing {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DAlphaAntiAliasing

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetAlphaAntialiasing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetAlphaAntialiasingEdge(edge float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetAlphaAntialiasingEdge), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaAntialiasingEdge() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetAlphaAntialiasingEdge), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetBillboardMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetBillboardMode() BaseMaterial3DBillboardMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DBillboardMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetBillboardMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnSetTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetTextureFilter() BaseMaterial3DTextureFilter {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTextureFilter

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetTextureFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) GetItemRect() Rect2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGetItemRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteBase3D) GenerateTriangleMesh() TriangleMesh {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTriangleMesh()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteBase3D.fnGenerateTriangleMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
