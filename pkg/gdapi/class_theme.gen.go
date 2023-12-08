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

func  (me *Theme) SetIcon(name StringName, theme_type StringName, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetIcon(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasIcon(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RenameIcon(old_name StringName, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearIcon(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetIconList(theme_type String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetIconTypeList() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetStylebox(name StringName, theme_type StringName, texture StyleBox, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetStylebox(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasStylebox(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RenameStylebox(old_name StringName, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearStylebox(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetStyleboxList(theme_type String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetStyleboxTypeList() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetFont(name StringName, theme_type StringName, font Font, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetFont(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasFont(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RenameFont(old_name StringName, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearFont(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetFontList(theme_type String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetFontTypeList() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetFontSize(name StringName, theme_type StringName, font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetFontSize(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasFontSize(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RenameFontSize(old_name StringName, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearFontSize(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetFontSizeList(theme_type String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetFontSizeTypeList() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetColor(name StringName, theme_type StringName, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetColor(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasColor(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RenameColor(old_name StringName, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearColor(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetColorList(theme_type String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetColorTypeList() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetConstant(name StringName, theme_type StringName, constant int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetConstant(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasConstant(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RenameConstant(old_name StringName, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearConstant(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetConstantList(theme_type String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetConstantTypeList() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetDefaultBaseScale(base_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetDefaultBaseScale() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasDefaultBaseScale() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetDefaultFont(font Font, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetDefaultFont() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasDefaultFont() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetDefaultFontSize(font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetDefaultFontSize() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasDefaultFontSize() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) HasThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RenameThemeItem(data_type ThemeDataType, old_name StringName, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetThemeItemList(data_type ThemeDataType, theme_type String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetThemeItemTypeList(data_type ThemeDataType, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) SetTypeVariation(theme_type StringName, base_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) IsTypeVariation(theme_type StringName, base_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) ClearTypeVariation(theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetTypeVariationBase(theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetTypeVariationList(base_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) AddType(theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) RemoveType(theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) GetTypeList() { // TODO: return value
  // TODO: implement
}

func  (me *Theme) MergeWith(other Theme, ) { // TODO: return value
  // TODO: implement
}

func  (me *Theme) Clear() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
