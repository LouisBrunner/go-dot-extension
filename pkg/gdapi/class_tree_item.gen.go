// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TreeItem struct {
  obj gdc.ObjectPtr
}

func (me *TreeItem) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TreeItem) BaseClass() string {
  return "TreeItem"
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

func  (me *TreeItem) SetCellMode(column int, mode TreeItemTreeCellMode, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 289920701) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetCellMode(column int, ) TreeItemTreeCellMode {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3406114978) // FIXME: should cache?
  var ret TreeItemTreeCellMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetEditMultiline(column int, multiline bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edit_multiline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&multiline), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsEditMultiline(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_edit_multiline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetChecked(column int, checked bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&checked), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetIndeterminate(column int, indeterminate bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_indeterminate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&indeterminate), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsChecked(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) IsIndeterminate(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_indeterminate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) PropagateCheck(column int, emit_signal bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("propagate_check")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4023243586) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&emit_signal), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetText(column int, text String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetText(column int, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetTextDirection(column int, direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1707680378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetTextDirection(column int, ) ControlTextDirection {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4235602388) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetAutowrapMode(column int, autowrap_mode TextServerAutowrapMode, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3633006561) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&autowrap_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetAutowrapMode(column int, ) TextServerAutowrapMode {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2902757236) // FIXME: should cache?
  var ret TextServerAutowrapMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetTextOverrunBehavior(column int, overrun_behavior TextServerOverrunBehavior, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1940772195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&overrun_behavior), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetTextOverrunBehavior(column int, ) TextServerOverrunBehavior {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3782727860) // FIXME: should cache?
  var ret TextServerOverrunBehavior
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetStructuredTextBidiOverride(column int, parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 868756907) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&parser), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetStructuredTextBidiOverride(column int, ) TextServerStructuredTextParser {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3377823772) // FIXME: should cache?
  var ret TextServerStructuredTextParser
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetStructuredTextBidiOverrideOptions(column int, args Array, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 537221740) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(args.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetStructuredTextBidiOverrideOptions(column int, ) Array {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetLanguage(column int, language String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetLanguage(column int, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetSuffix(column int, text String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetSuffix(column int, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetIcon(column int, texture Texture2D, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetIcon(column int, ) Texture2D {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetIconRegion(column int, region Rect2, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1356297692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetIconRegion(column int, ) Rect2 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3327874267) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetIconMaxWidth(column int, width int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetIconMaxWidth(column int, ) int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetIconModulate(column int, modulate Color, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetIconModulate(column int, ) Color {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetRange(column int, value float32, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetRange(column int, ) float32 {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetRangeConfig(column int, min float32, max float32, step float32, expr bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1547181014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&min), gdc.ConstTypePtr(&max), gdc.ConstTypePtr(&step), gdc.ConstTypePtr(&expr), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetRangeConfig(column int, ) Dictionary {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3554694381) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetMetadata(column int, meta Variant, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(meta.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetMetadata(column int, ) Variant {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetCustomDraw(column int, object Object, callback StringName, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 272420368) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetCollapsed(enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsCollapsed() bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetCollapsedRecursive(enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collapsed_recursive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsAnyCollapsed(only_visible bool, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_any_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2595650253) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&only_visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetVisible(enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsVisible() bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) UncollapseTree()  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("uncollapse_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetCustomMinimumHeight(height int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_minimum_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetCustomMinimumHeight() int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_minimum_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetSelectable(column int, selectable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&selectable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsSelectable(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) IsSelected(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) Select(column int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) Deselect(column int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetEditable(column int, enabled bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsEditable(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetCustomColor(column int, color Color, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetCustomColor(column int, ) Color {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) ClearCustomColor(column int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetCustomFont(column int, font Font, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2637609184) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetCustomFont(column int, ) Font {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4244553094) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetCustomFontSize(column int, font_size int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetCustomFontSize(column int, ) int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetCustomBgColor(column int, color Color, just_outline bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 894174518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&just_outline), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) ClearCustomBgColor(column int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetCustomBgColor(column int, ) Color {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetCustomAsButton(column int, enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_as_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsCustomSetAsButton(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_custom_set_as_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) AddButton(column int, button Texture2D, id int, disabled bool, tooltip_text String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1507727907) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(button.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&disabled), gdc.ConstTypePtr(tooltip_text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetButtonCount(column int, ) int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetButtonTooltipText(column int, button_index int, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1391810591) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetButtonId(column int, button_index int, ) int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetButtonById(column int, id int, ) int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_by_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetButton(column int, button_index int, ) Texture2D {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2584904275) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetButton(column int, button_index int, button Texture2D, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 176101966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), gdc.ConstTypePtr(button.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) EraseButton(column int, button_index int, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetButtonDisabled(column int, button_index int, disabled bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1383440665) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) SetButtonColor(column int, button_index int, color Color, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3733378741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsButtonDisabled(column int, button_index int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_button_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetTooltipText(column int, tooltip String, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(tooltip.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetTooltipText(column int, ) String {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetTextAlignment(column int, text_alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3276431499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&text_alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetTextAlignment(column int, ) HorizontalAlignment {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4171562184) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetExpandRight(column int, enable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetExpandRight(column int, ) bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_expand_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) SetDisableFolding(disable bool, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_folding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) IsFoldingDisabled() bool {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_folding_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) CreateChild(index int, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 954243986) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) AddChild(child TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(child.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) RemoveChild(child TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(child.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) GetTree() Tree {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2243340556) // FIXME: should cache?
  var ret Tree
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetNext() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetPrev() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_prev")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2768121250) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetParent() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetFirstChild() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_first_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetNextInTree(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_in_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetPrevInTree(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_prev_in_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetNextVisible(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetPrevVisible(wrap bool, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_prev_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1666920593) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetChild(index int, ) TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 306700752) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetChildCount() int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_child_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetChildren() TreeItem {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_children")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) GetIndex() int {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TreeItem) MoveBefore(item TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_before")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(item.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) MoveAfter(item TreeItem, )  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_after")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1819951137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(item.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TreeItem) CallRecursive(method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("TreeItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_recursive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2866548813) // FIXME: should cache?
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(method.AsCTypePtr()), }
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
