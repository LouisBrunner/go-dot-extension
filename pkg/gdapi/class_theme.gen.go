// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Theme struct {
  obj gdc.ObjectPtr
}

func (me *Theme) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Theme) BaseClass() string {
  return "Theme"
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
  panic("TODO: implement")
}

func  (me *Theme) GetIcon(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) HasIcon(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RenameIcon(old_name StringName, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearIcon(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetIconList(theme_type String, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetIconTypeList()  {
  panic("TODO: implement")
}

func  (me *Theme) SetStylebox(name StringName, theme_type StringName, texture StyleBox, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetStylebox(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) HasStylebox(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RenameStylebox(old_name StringName, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearStylebox(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetStyleboxList(theme_type String, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetStyleboxTypeList()  {
  panic("TODO: implement")
}

func  (me *Theme) SetFont(name StringName, theme_type StringName, font Font, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetFont(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) HasFont(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RenameFont(old_name StringName, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearFont(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetFontList(theme_type String, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetFontTypeList()  {
  panic("TODO: implement")
}

func  (me *Theme) SetFontSize(name StringName, theme_type StringName, font_size int, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetFontSize(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) HasFontSize(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RenameFontSize(old_name StringName, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearFontSize(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetFontSizeList(theme_type String, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetFontSizeTypeList()  {
  panic("TODO: implement")
}

func  (me *Theme) SetColor(name StringName, theme_type StringName, color Color, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetColor(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) HasColor(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RenameColor(old_name StringName, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearColor(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetColorList(theme_type String, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetColorTypeList()  {
  panic("TODO: implement")
}

func  (me *Theme) SetConstant(name StringName, theme_type StringName, constant int, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetConstant(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) HasConstant(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RenameConstant(old_name StringName, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearConstant(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetConstantList(theme_type String, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetConstantTypeList()  {
  panic("TODO: implement")
}

func  (me *Theme) SetDefaultBaseScale(base_scale float32, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetDefaultBaseScale()  {
  panic("TODO: implement")
}

func  (me *Theme) HasDefaultBaseScale()  {
  panic("TODO: implement")
}

func  (me *Theme) SetDefaultFont(font Font, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetDefaultFont()  {
  panic("TODO: implement")
}

func  (me *Theme) HasDefaultFont()  {
  panic("TODO: implement")
}

func  (me *Theme) SetDefaultFontSize(font_size int, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetDefaultFontSize()  {
  panic("TODO: implement")
}

func  (me *Theme) HasDefaultFontSize()  {
  panic("TODO: implement")
}

func  (me *Theme) SetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) HasThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RenameThemeItem(data_type ThemeDataType, old_name StringName, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetThemeItemList(data_type ThemeDataType, theme_type String, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetThemeItemTypeList(data_type ThemeDataType, )  {
  panic("TODO: implement")
}

func  (me *Theme) SetTypeVariation(theme_type StringName, base_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) IsTypeVariation(theme_type StringName, base_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) ClearTypeVariation(theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetTypeVariationBase(theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetTypeVariationList(base_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) AddType(theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) RemoveType(theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Theme) GetTypeList()  {
  panic("TODO: implement")
}

func  (me *Theme) MergeWith(other Theme, )  {
  panic("TODO: implement")
}

func  (me *Theme) Clear()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
