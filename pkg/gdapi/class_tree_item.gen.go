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

type TreeItem struct {
  Object
}

func (me *TreeItem) BaseClass() string {
  return "TreeItem"
}

func NewTreeItem() *TreeItem {
  str := StringNameFromStr("TreeItem") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TreeItem{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TreeItemTreeCellMode int
const (
  TreeItemTreeCellModeCellModeString TreeItemTreeCellMode = 0
  TreeItemTreeCellModeCellModeCheck TreeItemTreeCellMode = 1
  TreeItemTreeCellModeCellModeRange TreeItemTreeCellMode = 2
  TreeItemTreeCellModeCellModeIcon TreeItemTreeCellMode = 3
  TreeItemTreeCellModeCellModeCustom TreeItemTreeCellMode = 4
)

func (me *TreeItem) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TreeItem) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TreeItem) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TreeItem) SetCellMode(column int64, mode TreeItemTreeCellMode, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 289920701) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetCellMode(column int64, ) TreeItemTreeCellMode {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3406114978) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TreeItemTreeCellMode
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TreeItem) SetEditMultiline(column int64, multiline bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edit_multiline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&multiline) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsEditMultiline(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_edit_multiline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetChecked(column int64, checked bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&checked) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetIndeterminate(column int64, indeterminate bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_indeterminate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&indeterminate) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsChecked(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) IsIndeterminate(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_indeterminate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) PropagateCheck(column int64, emit_signal bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("propagate_check")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 972357352) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&emit_signal) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetText(column int64, text String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetText(column int64, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetTextDirection(column int64, direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1707680378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetTextDirection(column int64, ) ControlTextDirection {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4235602388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ControlTextDirection
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TreeItem) SetAutowrapMode(column int64, autowrap_mode TextServerAutowrapMode, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3633006561) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&autowrap_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetAutowrapMode(column int64, ) TextServerAutowrapMode {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2902757236) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerAutowrapMode
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TreeItem) SetTextOverrunBehavior(column int64, overrun_behavior TextServerOverrunBehavior, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1940772195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&overrun_behavior) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetTextOverrunBehavior(column int64, ) TextServerOverrunBehavior {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3782727860) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerOverrunBehavior
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TreeItem) SetStructuredTextBidiOverride(column int64, parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 868756907) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&parser) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetStructuredTextBidiOverride(column int64, ) TextServerStructuredTextParser {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3377823772) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerStructuredTextParser
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TreeItem) SetStructuredTextBidiOverrideOptions(column int64, args Array, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 537221740) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , args.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetStructuredTextBidiOverrideOptions(column int64, ) Array {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetLanguage(column int64, language String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetLanguage(column int64, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetSuffix(column int64, text String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetSuffix(column int64, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetIcon(column int64, texture Texture2D, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetIcon(column int64, ) Texture2D {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetIconRegion(column int64, region Rect2, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1356297692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetIconRegion(column int64, ) Rect2 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3327874267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetIconMaxWidth(column int64, width int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetIconMaxWidth(column int64, ) int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetIconModulate(column int64, modulate Color, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetIconModulate(column int64, ) Color {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetRange(column int64, value float64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetRange(column int64, ) float64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetRangeConfig(column int64, min float64, max float64, step float64, expr bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1547181014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&min) , gdc.ConstTypePtr(&max) , gdc.ConstTypePtr(&step) , gdc.ConstTypePtr(&expr) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetRangeConfig(column int64, ) Dictionary {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3554694381) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetMetadata(column int64, meta Variant, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , meta.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetMetadata(column int64, ) Variant {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetCustomDraw(column int64, object Object, callback StringName, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 272420368) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , object.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetCollapsed(enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsCollapsed() bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetCollapsedRecursive(enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collapsed_recursive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsAnyCollapsed(only_visible bool, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_any_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2595650253) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&only_visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&only_visible)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetVisible(enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsVisible() bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) UncollapseTree()  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("uncollapse_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetCustomMinimumHeight(height int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_minimum_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetCustomMinimumHeight() int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_minimum_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetSelectable(column int64, selectable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&selectable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsSelectable(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) IsSelected(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) Select(column int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) Deselect(column int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetEditable(column int64, enabled bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsEditable(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetCustomColor(column int64, color Color, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetCustomColor(column int64, ) Color {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) ClearCustomColor(column int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetCustomFont(column int64, font Font, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2637609184) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , font.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetCustomFont(column int64, ) Font {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4244553094) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFont()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetCustomFontSize(column int64, font_size int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&font_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetCustomFontSize(column int64, ) int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetCustomBgColor(column int64, color Color, just_outline bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 894174518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , color.AsCTypePtr(), gdc.ConstTypePtr(&just_outline) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) ClearCustomBgColor(column int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetCustomBgColor(column int64, ) Color {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetCustomAsButton(column int64, enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_as_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsCustomSetAsButton(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_custom_set_as_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) AddButton(column int64, button Texture2D, id int64, disabled bool, tooltip_text String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1688223362) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , button.AsCTypePtr(), gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&disabled) , tooltip_text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetButtonCount(column int64, ) int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) GetButtonTooltipText(column int64, button_index int64, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1391810591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&column)
  pinner.Pin(&button_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetButtonId(column int64, button_index int64, ) int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&column)
  pinner.Pin(&button_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) GetButtonById(column int64, id int64, ) int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_by_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&column)
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) GetButton(column int64, button_index int64, ) Texture2D {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2584904275) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&column)
  pinner.Pin(&button_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetButtonTooltipText(column int64, button_index int64, tooltip String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285447957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , tooltip.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetButton(column int64, button_index int64, button Texture2D, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 176101966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , button.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) EraseButton(column int64, button_index int64, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetButtonDisabled(column int64, button_index int64, disabled bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1383440665) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) SetButtonColor(column int64, button_index int64, color Color, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3733378741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsButtonDisabled(column int64, button_index int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_button_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&button_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)
  pinner.Pin(&button_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetTooltipText(column int64, tooltip String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , tooltip.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetTooltipText(column int64, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) SetTextAlignment(column int64, text_alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3276431499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&text_alignment) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetTextAlignment(column int64, ) HorizontalAlignment {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4171562184) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret HorizontalAlignment
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TreeItem) SetExpandRight(column int64, enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetExpandRight(column int64, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_expand_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) SetDisableFolding(disable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_folding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) IsFoldingDisabled() bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_folding_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) CreateChild(index int64, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 954243986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) AddChild(child TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{child.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) RemoveChild(child TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{child.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) GetTree() Tree {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2243340556) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTree()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetNext() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetPrev() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_prev")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2768121250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetParent() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetFirstChild() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_first_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetNextInTree(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_in_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()
  pinner.Pin(&wrap)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetPrevInTree(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_prev_in_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()
  pinner.Pin(&wrap)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetNextVisible(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()
  pinner.Pin(&wrap)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetPrevVisible(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_prev_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()
  pinner.Pin(&wrap)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetChild(index int64, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 306700752) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTreeItem()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TreeItem) GetChildCount() int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_child_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) GetChildren() []TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_children")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[TreeItem](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *TreeItem) GetIndex() int64 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TreeItem) MoveBefore(item TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_before")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) MoveAfter(item TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_after")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TreeItem) CallRecursive(method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_recursive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2866548813) // FIXME: should cache?
  cargs := make([]gdc.ConstVariantPtr, 0, 1 + len(varargs))
  var0 := method.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }

  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return
  }

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
