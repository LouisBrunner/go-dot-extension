// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SystemFont struct {
  Font
}

func (me *SystemFont) BaseClass() string {
  return "SystemFont"
}



// Enums

func (me *SystemFont) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SystemFont) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SystemFont) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SystemFont) SetAntialiasing(antialiasing TextServerFontAntialiasing, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1669900) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiasing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetAntialiasing() TextServerFontAntialiasing {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4262718649) // FIXME: should cache?
  var ret TextServerFontAntialiasing
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetGenerateMipmaps(generate_mipmaps bool, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&generate_mipmaps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetGenerateMipmaps() bool {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetAllowSystemFallback(allow_system_fallback bool, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_system_fallback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_system_fallback), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) IsAllowSystemFallback() bool {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_allow_system_fallback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetForceAutohinter(force_autohinter bool, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_force_autohinter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_autohinter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) IsForceAutohinter() bool {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_force_autohinter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetHinting(hinting TextServerHinting, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hinting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1827459492) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hinting), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetHinting() TextServerHinting {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hinting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3683214614) // FIXME: should cache?
  var ret TextServerHinting
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetSubpixelPositioning(subpixel_positioning TextServerSubpixelPositioning, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subpixel_positioning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4225742182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subpixel_positioning), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetSubpixelPositioning() TextServerSubpixelPositioning {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subpixel_positioning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1069238588) // FIXME: should cache?
  var ret TextServerSubpixelPositioning
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetMultichannelSignedDistanceField(msdf bool, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multichannel_signed_distance_field")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) IsMultichannelSignedDistanceField() bool {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_multichannel_signed_distance_field")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetMsdfPixelRange(msdf_pixel_range int, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_pixel_range), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetMsdfPixelRange() int {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetMsdfSize(msdf_size int, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msdf_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetMsdfSize() int {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msdf_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetOversampling(oversampling float32, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversampling), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetOversampling() float32 {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) GetFontNames() PackedStringArray {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetFontNames(names PackedStringArray, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(names.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) GetFontItalic() bool {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_italic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SystemFont) SetFontItalic(italic bool, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_italic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&italic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) SetFontWeight(weight int, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&weight), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SystemFont) SetFontStretch(stretch int, )  {
  classNameV := StringNameFromStr("SystemFont")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
