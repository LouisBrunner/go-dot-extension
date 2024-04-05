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

type ptrsForCodeHighlighterList struct {
  fnAddKeywordColor gdc.MethodBindPtr
  fnRemoveKeywordColor gdc.MethodBindPtr
  fnHasKeywordColor gdc.MethodBindPtr
  fnGetKeywordColor gdc.MethodBindPtr
  fnSetKeywordColors gdc.MethodBindPtr
  fnClearKeywordColors gdc.MethodBindPtr
  fnGetKeywordColors gdc.MethodBindPtr
  fnAddMemberKeywordColor gdc.MethodBindPtr
  fnRemoveMemberKeywordColor gdc.MethodBindPtr
  fnHasMemberKeywordColor gdc.MethodBindPtr
  fnGetMemberKeywordColor gdc.MethodBindPtr
  fnSetMemberKeywordColors gdc.MethodBindPtr
  fnClearMemberKeywordColors gdc.MethodBindPtr
  fnGetMemberKeywordColors gdc.MethodBindPtr
  fnAddColorRegion gdc.MethodBindPtr
  fnRemoveColorRegion gdc.MethodBindPtr
  fnHasColorRegion gdc.MethodBindPtr
  fnSetColorRegions gdc.MethodBindPtr
  fnClearColorRegions gdc.MethodBindPtr
  fnGetColorRegions gdc.MethodBindPtr
  fnSetFunctionColor gdc.MethodBindPtr
  fnGetFunctionColor gdc.MethodBindPtr
  fnSetNumberColor gdc.MethodBindPtr
  fnGetNumberColor gdc.MethodBindPtr
  fnSetSymbolColor gdc.MethodBindPtr
  fnGetSymbolColor gdc.MethodBindPtr
  fnSetMemberVariableColor gdc.MethodBindPtr
  fnGetMemberVariableColor gdc.MethodBindPtr
}

var ptrsForCodeHighlighter ptrsForCodeHighlighterList

func initCodeHighlighterPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CodeHighlighter")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnAddKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1636512886))
  }
  {
    methodName := StringNameFromStr("remove_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnRemoveKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("has_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnHasKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("get_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3855908743))
  }
  {
    methodName := StringNameFromStr("set_keyword_colors")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnSetKeywordColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("clear_keyword_colors")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnClearKeywordColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_keyword_colors")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetKeywordColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
  }
  {
    methodName := StringNameFromStr("add_member_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnAddMemberKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1636512886))
  }
  {
    methodName := StringNameFromStr("remove_member_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnRemoveMemberKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("has_member_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnHasMemberKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("get_member_keyword_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetMemberKeywordColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3855908743))
  }
  {
    methodName := StringNameFromStr("set_member_keyword_colors")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnSetMemberKeywordColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("clear_member_keyword_colors")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnClearMemberKeywordColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_member_keyword_colors")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetMemberKeywordColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
  }
  {
    methodName := StringNameFromStr("add_color_region")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnAddColorRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2924977451))
  }
  {
    methodName := StringNameFromStr("remove_color_region")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnRemoveColorRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("has_color_region")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnHasColorRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("set_color_regions")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnSetColorRegions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("clear_color_regions")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnClearColorRegions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_color_regions")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetColorRegions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
  }
  {
    methodName := StringNameFromStr("set_function_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnSetFunctionColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_function_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetFunctionColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_number_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnSetNumberColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_number_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetNumberColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_symbol_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnSetSymbolColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_symbol_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetSymbolColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_member_variable_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnSetMemberVariableColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_member_variable_color")
    defer methodName.Destroy()
    ptrsForCodeHighlighter.fnGetMemberVariableColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
}

type CodeHighlighter struct {
  SyntaxHighlighter
}

func (me *CodeHighlighter) BaseClass() string {
  return "CodeHighlighter"
}

func NewCodeHighlighter() *CodeHighlighter {
  str := StringNameFromStr("CodeHighlighter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CodeHighlighter{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{keyword.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnAddKeywordColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) RemoveKeywordColor(keyword String, )  {
  cargs := []gdc.ConstTypePtr{keyword.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnRemoveKeywordColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) HasKeywordColor(keyword String, ) bool {
  cargs := []gdc.ConstTypePtr{keyword.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnHasKeywordColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeHighlighter) GetKeywordColor(keyword String, ) Color {
  cargs := []gdc.ConstTypePtr{keyword.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetKeywordColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) SetKeywordColors(keywords Dictionary, )  {
  cargs := []gdc.ConstTypePtr{keywords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnSetKeywordColors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) ClearKeywordColors()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnClearKeywordColors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) GetKeywordColors() Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetKeywordColors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) AddMemberKeywordColor(member_keyword String, color Color, )  {
  cargs := []gdc.ConstTypePtr{member_keyword.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnAddMemberKeywordColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) RemoveMemberKeywordColor(member_keyword String, )  {
  cargs := []gdc.ConstTypePtr{member_keyword.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnRemoveMemberKeywordColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) HasMemberKeywordColor(member_keyword String, ) bool {
  cargs := []gdc.ConstTypePtr{member_keyword.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnHasMemberKeywordColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeHighlighter) GetMemberKeywordColor(member_keyword String, ) Color {
  cargs := []gdc.ConstTypePtr{member_keyword.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetMemberKeywordColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) SetMemberKeywordColors(member_keyword Dictionary, )  {
  cargs := []gdc.ConstTypePtr{member_keyword.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnSetMemberKeywordColors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) ClearMemberKeywordColors()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnClearMemberKeywordColors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) GetMemberKeywordColors() Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetMemberKeywordColors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) AddColorRegion(start_key String, end_key String, color Color, line_only bool, )  {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), end_key.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&line_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnAddColorRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) RemoveColorRegion(start_key String, )  {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnRemoveColorRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) HasColorRegion(start_key String, ) bool {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnHasColorRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeHighlighter) SetColorRegions(color_regions Dictionary, )  {
  cargs := []gdc.ConstTypePtr{color_regions.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnSetColorRegions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) ClearColorRegions()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnClearColorRegions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) GetColorRegions() Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetColorRegions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) SetFunctionColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnSetFunctionColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) GetFunctionColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetFunctionColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) SetNumberColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnSetNumberColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) GetNumberColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetNumberColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) SetSymbolColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnSetSymbolColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) GetSymbolColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetSymbolColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeHighlighter) SetMemberVariableColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnSetMemberVariableColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeHighlighter) GetMemberVariableColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeHighlighter.fnGetMemberVariableColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
