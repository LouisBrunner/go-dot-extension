// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&centered) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) IsCentered() bool {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteBase3D) SetFlipH(flip_h bool, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) IsFlippedH() bool {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetFlipV(flip_v bool, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) IsFlippedV() bool {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetModulate(modulate Color, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetModulate() Color {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteBase3D) SetRenderPriority(priority int64, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_render_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetRenderPriority() int64 {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetPixelSize(pixel_size float64, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pixel_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixel_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetPixelSize() float64 {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pixel_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAxis(axis Vector3Axis, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1144690656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAxis() Vector3Axis {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050976882) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Vector3Axis

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetDrawFlag(flag SpriteBase3DDrawFlags, enabled bool, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1135633219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetDrawFlag(flag SpriteBase3DDrawFlags, ) bool {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1733036628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&flag)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAlphaCutMode(mode SpriteBase3DAlphaCutMode, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_cut_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 227561226) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaCutMode() SpriteBase3DAlphaCutMode {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_cut_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 336003791) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SpriteBase3DAlphaCutMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetAlphaScissorThreshold(threshold float64, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_scissor_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaScissorThreshold() float64 {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_scissor_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAlphaHashScale(threshold float64, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_hash_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaHashScale() float64 {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_hash_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3212649852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alpha_aa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaAntialiasing() BaseMaterial3DAlphaAntiAliasing {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2889939400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DAlphaAntiAliasing

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetAlphaAntialiasingEdge(edge float64, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_antialiasing_edge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetAlphaAntialiasingEdge() float64 {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_antialiasing_edge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteBase3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_billboard_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4202036497) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetBillboardMode() BaseMaterial3DBillboardMode {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_billboard_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1283840139) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DBillboardMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, )  {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 22904437) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteBase3D) GetTextureFilter() BaseMaterial3DTextureFilter {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289213076) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTextureFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SpriteBase3D) GetItemRect() Rect2 {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteBase3D) GenerateTriangleMesh() TriangleMesh {
  classNameV := StringNameFromStr("SpriteBase3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_triangle_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3476533166) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTriangleMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
