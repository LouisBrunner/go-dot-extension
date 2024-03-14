// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CodeHighlighter struct {
  SyntaxHighlighter
}

func (me *CodeHighlighter) BaseClass() string {
  return "CodeHighlighter"
}



// Enums

func (me *CodeHighlighter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CodeHighlighter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CodeHighlighter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CodeHighlighter) AddKeywordColor(keyword String, color Color, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1636512886) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(keyword.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) RemoveKeywordColor(keyword String, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(keyword.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) HasKeywordColor(keyword String, ) bool {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(keyword.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) GetKeywordColor(keyword String, ) Color {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3855908743) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(keyword.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) SetKeywordColors(keywords Dictionary, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keyword_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(keywords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) ClearKeywordColors()  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_keyword_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) GetKeywordColors() Dictionary {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keyword_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) AddMemberKeywordColor(member_keyword String, color Color, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_member_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1636512886) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(member_keyword.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) RemoveMemberKeywordColor(member_keyword String, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_member_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(member_keyword.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) HasMemberKeywordColor(member_keyword String, ) bool {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_member_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(member_keyword.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) GetMemberKeywordColor(member_keyword String, ) Color {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_member_keyword_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3855908743) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(member_keyword.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) SetMemberKeywordColors(member_keyword Dictionary, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_member_keyword_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(member_keyword.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) ClearMemberKeywordColors()  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_member_keyword_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) GetMemberKeywordColors() Dictionary {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_member_keyword_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) AddColorRegion(start_key String, end_key String, color Color, line_only bool, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_color_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2924977451) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), gdc.ConstTypePtr(end_key.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&line_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) RemoveColorRegion(start_key String, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_color_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) HasColorRegion(start_key String, ) bool {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_color_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) SetColorRegions(color_regions Dictionary, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_regions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color_regions.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) ClearColorRegions()  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_color_regions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) GetColorRegions() Dictionary {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_regions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) SetFunctionColor(color Color, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) GetFunctionColor() Color {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) SetNumberColor(color Color, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_number_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) GetNumberColor() Color {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_number_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) SetSymbolColor(color Color, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_symbol_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) GetSymbolColor() Color {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_symbol_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeHighlighter) SetMemberVariableColor(color Color, )  {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_member_variable_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeHighlighter) GetMemberVariableColor() Color {
  classNameV := StringNameFromStr("CodeHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_member_variable_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
