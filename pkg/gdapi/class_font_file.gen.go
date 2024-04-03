// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type FontFile struct {
  Font
}

func (me *FontFile) BaseClass() string {
  return "FontFile"
}

func NewFontFile() *FontFile {
  str := StringNameFromStr("FontFile") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &FontFile{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *FontFile) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FontFile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FontFile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *FontFile) LoadBitmapFont(path String, ) Error {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_bitmap_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FontFile) LoadDynamicFont(path String, ) Error {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_dynamic_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FontFile) SetData(data PackedByteArray, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2971499966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetData() PackedByteArray {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetFontName(name String, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetFontStyleName(name String, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_style_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetFontStyle(style TextServerFontStyle, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_style")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 918070724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&style), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetFontWeight(weight int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&weight), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetFontStretch(stretch int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetAntialiasing(antialiasing TextServerFontAntialiasing, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1669900) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiasing), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetAntialiasing() TextServerFontAntialiasing {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4262718649) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TextServerFontAntialiasing

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FontFile) SetGenerateMipmaps(generate_mipmaps bool, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&generate_mipmaps), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetGenerateMipmaps() bool {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetMultichannelSignedDistanceField(msdf bool, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multichannel_signed_distance_field")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) IsMultichannelSignedDistanceField() bool {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_multichannel_signed_distance_field")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetMsdfPixelRange(msdf_pixel_range int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_pixel_range), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetMsdfPixelRange() int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetMsdfSize(msdf_size int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msdf_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetMsdfSize() int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msdf_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetFixedSize(fixed_size int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fixed_size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetFixedSize() int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetFixedSizeScaleMode(fixed_size_scale_mode TextServerFixedSizeScaleMode, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_size_scale_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1660989956) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fixed_size_scale_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetFixedSizeScaleMode() TextServerFixedSizeScaleMode {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_size_scale_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 753873478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TextServerFixedSizeScaleMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FontFile) SetAllowSystemFallback(allow_system_fallback bool, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_system_fallback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_system_fallback), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) IsAllowSystemFallback() bool {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_allow_system_fallback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetForceAutohinter(force_autohinter bool, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_force_autohinter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_autohinter), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) IsForceAutohinter() bool {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_force_autohinter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetHinting(hinting TextServerHinting, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hinting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1827459492) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hinting), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetHinting() TextServerHinting {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hinting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3683214614) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TextServerHinting

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FontFile) SetSubpixelPositioning(subpixel_positioning TextServerSubpixelPositioning, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subpixel_positioning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4225742182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subpixel_positioning), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetSubpixelPositioning() TextServerSubpixelPositioning {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subpixel_positioning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1069238588) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TextServerSubpixelPositioning

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FontFile) SetOversampling(oversampling float64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversampling), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetOversampling() float64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) GetCacheCount() int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cache_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) ClearCache()  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) RemoveCache(cache_index int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetSizeCacheList(cache_index int64, ) []Vector2i {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_cache_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), }
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Vector2i](ret)
}

func  (me *FontFile) ClearSizeCache(cache_index int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_size_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) RemoveSizeCache(cache_index int64, size Vector2i, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_size_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311374912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetVariationCoordinates(cache_index int64, variation_coordinates Dictionary, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_variation_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 64545446) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(variation_coordinates.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetVariationCoordinates(cache_index int64, ) Dictionary {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_variation_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), }
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetEmbolden(cache_index int64, strength float64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_embolden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&strength), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetEmbolden(cache_index int64, ) float64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_embolden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetTransform(cache_index int64, transform Transform2D, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 30160968) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(transform.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetTransform(cache_index int64, ) Transform2D {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3836996910) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), }
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetExtraSpacing(cache_index int64, spacing TextServerSpacingType, value int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_extra_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 62942285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&spacing), gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetExtraSpacing(cache_index int64, spacing TextServerSpacingType, ) int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_extra_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1924257185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&spacing), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetFaceIndex(cache_index int64, face_index int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_face_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&face_index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetFaceIndex(cache_index int64, ) int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_face_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetCacheAscent(cache_index int64, size int64, ascent float64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cache_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&ascent), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetCacheAscent(cache_index int64, size int64, ) float64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cache_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetCacheDescent(cache_index int64, size int64, descent float64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cache_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&descent), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetCacheDescent(cache_index int64, size int64, ) float64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cache_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetCacheUnderlinePosition(cache_index int64, size int64, underline_position float64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cache_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_position), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetCacheUnderlinePosition(cache_index int64, size int64, ) float64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cache_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetCacheUnderlineThickness(cache_index int64, size int64, underline_thickness float64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cache_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_thickness), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetCacheUnderlineThickness(cache_index int64, size int64, ) float64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cache_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) SetCacheScale(cache_index int64, size int64, scale float64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cache_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&scale), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetCacheScale(cache_index int64, size int64, ) float64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cache_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) GetTextureCount(cache_index int64, size Vector2i, ) int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1987661582) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) ClearTextures(cache_index int64, size Vector2i, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311374912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) RemoveTexture(cache_index int64, size Vector2i, texture_index int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2328951467) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetTextureImage(cache_index int64, size Vector2i, texture_index int64, image Image, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4157974066) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), gdc.ConstTypePtr(image.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetTextureImage(cache_index int64, size Vector2i, texture_index int64, ) Image {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3878418953) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), }
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetTextureOffsets(cache_index int64, size Vector2i, texture_index int64, offset PackedInt32Array, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_offsets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2849993437) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), gdc.ConstTypePtr(offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetTextureOffsets(cache_index int64, size Vector2i, texture_index int64, ) PackedInt32Array {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_offsets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3703444828) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), }
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) GetGlyphList(cache_index int64, size Vector2i, ) PackedInt32Array {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 681709689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), }
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) ClearGlyphs(cache_index int64, size Vector2i, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_glyphs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311374912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) RemoveGlyph(cache_index int64, size Vector2i, glyph int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_glyph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2328951467) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetGlyphAdvance(cache_index int64, size int64, glyph int64, advance Vector2, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 947991729) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(advance.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetGlyphAdvance(cache_index int64, size int64, glyph int64, ) Vector2 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1601573536) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph), }
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetGlyphOffset(cache_index int64, size Vector2i, glyph int64, offset Vector2, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 921719850) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetGlyphOffset(cache_index int64, size Vector2i, glyph int64, ) Vector2 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3205412300) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetGlyphSize(cache_index int64, size Vector2i, glyph int64, gl_size Vector2, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 921719850) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(gl_size.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetGlyphSize(cache_index int64, size Vector2i, glyph int64, ) Vector2 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3205412300) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetGlyphUvRect(cache_index int64, size Vector2i, glyph int64, uv_rect Rect2, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_uv_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3821620992) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(uv_rect.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetGlyphUvRect(cache_index int64, size Vector2i, glyph int64, ) Rect2 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_uv_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927917900) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetGlyphTextureIdx(cache_index int64, size Vector2i, glyph int64, texture_idx int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_texture_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 355564111) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(&texture_idx), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetGlyphTextureIdx(cache_index int64, size Vector2i, glyph int64, ) int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_texture_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1629411054) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) GetKerningList(cache_index int64, size int64, ) []Vector2i {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_kerning_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2345056839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), }
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Vector2i](ret)
}

func  (me *FontFile) ClearKerningMap(cache_index int64, size int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_kerning_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) RemoveKerning(cache_index int64, size int64, glyph_pair Vector2i, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_kerning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3930204747) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(glyph_pair.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetKerning(cache_index int64, size int64, glyph_pair Vector2i, kerning Vector2, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_kerning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3182200918) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(glyph_pair.AsCTypePtr()), gdc.ConstTypePtr(kerning.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetKerning(cache_index int64, size int64, glyph_pair Vector2i, ) Vector2 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_kerning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611912865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(glyph_pair.AsCTypePtr()), }
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) RenderRange(cache_index int64, size Vector2i, start int64, end int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("render_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 355564111) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) RenderGlyph(cache_index int64, size Vector2i, index int64, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("render_glyph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2328951467) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) SetLanguageSupportOverride(language String, supported bool, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), gdc.ConstTypePtr(&supported), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetLanguageSupportOverride(language String, ) bool {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) RemoveLanguageSupportOverride(language String, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_language_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetLanguageSupportOverrides() PackedStringArray {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language_support_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetScriptSupportOverride(script String, supported bool, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_script_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(script.AsCTypePtr()), gdc.ConstTypePtr(&supported), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetScriptSupportOverride(script String, ) bool {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(script.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) RemoveScriptSupportOverride(script String, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_script_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(script.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetScriptSupportOverrides() PackedStringArray {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script_support_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) SetOpentypeFeatureOverrides(overrides Dictionary, )  {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_opentype_feature_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(overrides.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontFile) GetOpentypeFeatureOverrides() Dictionary {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_opentype_feature_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontFile) GetGlyphIndex(size int64, char int64, variation_selector int64, ) int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 864943070) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&variation_selector), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontFile) GetCharFromGlyphIndex(size int64, glyph_index int64, ) int64 {
  classNameV := StringNameFromStr("FontFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_char_from_glyph_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph_index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
