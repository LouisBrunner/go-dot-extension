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

type Theme struct {
  Resource
}

func (me *Theme) BaseClass() string {
  return "Theme"
}

func NewTheme() *Theme {
  str := StringNameFromStr("Theme") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Theme{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ThemeDataType int
const (
  ThemeDataTypeDataTypeColor ThemeDataType = 0
  ThemeDataTypeDataTypeConstant ThemeDataType = 1
  ThemeDataTypeDataTypeFont ThemeDataType = 2
  ThemeDataTypeDataTypeFontSize ThemeDataType = 3
  ThemeDataTypeDataTypeIcon ThemeDataType = 4
  ThemeDataTypeDataTypeStylebox ThemeDataType = 5
  ThemeDataTypeDataTypeMax ThemeDataType = 6
)

func (me *Theme) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Theme) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Theme) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Theme) SetIcon(name StringName, theme_type StringName, texture Texture2D, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2188371082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetIcon(name StringName, theme_type StringName, ) Texture2D {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 934555193) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) HasIcon(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) RenameIcon(old_name StringName, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642128662) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) ClearIcon(name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetIconList(theme_type String, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291131558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetIconTypeList() PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) SetStylebox(name StringName, theme_type StringName, texture StyleBox, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2075907568) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetStylebox(name StringName, theme_type StringName, ) StyleBox {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3405608165) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStyleBox()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) HasStylebox(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) RenameStylebox(old_name StringName, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642128662) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) ClearStylebox(name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetStyleboxList(theme_type String, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stylebox_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291131558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetStyleboxTypeList() PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stylebox_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) SetFont(name StringName, theme_type StringName, font Font, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 177292320) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), font.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetFont(name StringName, theme_type StringName, ) Font {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3445063586) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFont()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) HasFont(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) RenameFont(old_name StringName, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642128662) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) ClearFont(name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetFontList(theme_type String, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291131558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetFontTypeList() PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) SetFontSize(name StringName, theme_type StringName, font_size int64, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 281601298) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), gdc.ConstTypePtr(&font_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetFontSize(name StringName, theme_type StringName, ) int64 {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2419549490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) HasFontSize(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) RenameFontSize(old_name StringName, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642128662) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) ClearFontSize(name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetFontSizeList(theme_type String, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_size_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291131558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetFontSizeTypeList() PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_size_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) SetColor(name StringName, theme_type StringName, color Color, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4111215154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetColor(name StringName, theme_type StringName, ) Color {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2015923404) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) HasColor(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) RenameColor(old_name StringName, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642128662) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) ClearColor(name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetColorList(theme_type String, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291131558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetColorTypeList() PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) SetConstant(name StringName, theme_type StringName, constant int64, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 281601298) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), gdc.ConstTypePtr(&constant) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetConstant(name StringName, theme_type StringName, ) int64 {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2419549490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) HasConstant(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) RenameConstant(old_name StringName, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642128662) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) ClearConstant(name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetConstantList(theme_type String, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291131558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetConstantTypeList() PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) SetDefaultBaseScale(base_scale float64, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&base_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetDefaultBaseScale() float64 {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) HasDefaultBaseScale() bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_default_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) SetDefaultFont(font Font, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262170328) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetDefaultFont() Font {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229501585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFont()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) HasDefaultFont() bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_default_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) SetDefaultFontSize(font_size int64, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetDefaultFontSize() int64 {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) HasDefaultFontSize() bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_default_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) SetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, value Variant, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_theme_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2492983623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type) , name.AsCTypePtr(), theme_type.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, ) Variant {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2191024021) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type) , name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&data_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) HasThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1739311056) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type) , name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&data_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) RenameThemeItem(data_type ThemeDataType, old_name StringName, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_theme_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3900867553) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type) , old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) ClearThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_theme_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2965505587) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type) , name.AsCTypePtr(), theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetThemeItemList(data_type ThemeDataType, theme_type String, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_item_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3726716710) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type) , theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()
  pinner.Pin(&data_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetThemeItemTypeList(data_type ThemeDataType, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_item_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1316004935) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()
  pinner.Pin(&data_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) SetTypeVariation(theme_type StringName, base_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_type_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), base_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) IsTypeVariation(theme_type StringName, base_type StringName, ) bool {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_type_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), base_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Theme) ClearTypeVariation(theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_type_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetTypeVariationBase(theme_type StringName, ) StringName {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_type_variation_base")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) GetTypeVariationList(base_type StringName, ) PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_type_variation_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1761182771) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{base_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) AddType(theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) RemoveType(theme_type StringName, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) GetTypeList() PackedStringArray {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_type_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Theme) MergeWith(other Theme, )  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("merge_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2326690814) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{other.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Theme) Clear()  {
  classNameV := StringNameFromStr("Theme")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
