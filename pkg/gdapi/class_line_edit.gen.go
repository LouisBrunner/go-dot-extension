// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type LineEdit struct {
  Control
}

func (me *LineEdit) BaseClass() string {
  return "LineEdit"
}



// Enums

type LineEditMenuItems int
const (
  LineEditMenuItemsMenuCut LineEditMenuItems = 0
  LineEditMenuItemsMenuCopy LineEditMenuItems = 1
  LineEditMenuItemsMenuPaste LineEditMenuItems = 2
  LineEditMenuItemsMenuClear LineEditMenuItems = 3
  LineEditMenuItemsMenuSelectAll LineEditMenuItems = 4
  LineEditMenuItemsMenuUndo LineEditMenuItems = 5
  LineEditMenuItemsMenuRedo LineEditMenuItems = 6
  LineEditMenuItemsMenuSubmenuTextDir LineEditMenuItems = 7
  LineEditMenuItemsMenuDirInherited LineEditMenuItems = 8
  LineEditMenuItemsMenuDirAuto LineEditMenuItems = 9
  LineEditMenuItemsMenuDirLtr LineEditMenuItems = 10
  LineEditMenuItemsMenuDirRtl LineEditMenuItems = 11
  LineEditMenuItemsMenuDisplayUcc LineEditMenuItems = 12
  LineEditMenuItemsMenuSubmenuInsertUcc LineEditMenuItems = 13
  LineEditMenuItemsMenuInsertLrm LineEditMenuItems = 14
  LineEditMenuItemsMenuInsertRlm LineEditMenuItems = 15
  LineEditMenuItemsMenuInsertLre LineEditMenuItems = 16
  LineEditMenuItemsMenuInsertRle LineEditMenuItems = 17
  LineEditMenuItemsMenuInsertLro LineEditMenuItems = 18
  LineEditMenuItemsMenuInsertRlo LineEditMenuItems = 19
  LineEditMenuItemsMenuInsertPdf LineEditMenuItems = 20
  LineEditMenuItemsMenuInsertAlm LineEditMenuItems = 21
  LineEditMenuItemsMenuInsertLri LineEditMenuItems = 22
  LineEditMenuItemsMenuInsertRli LineEditMenuItems = 23
  LineEditMenuItemsMenuInsertFsi LineEditMenuItems = 24
  LineEditMenuItemsMenuInsertPdi LineEditMenuItems = 25
  LineEditMenuItemsMenuInsertZwj LineEditMenuItems = 26
  LineEditMenuItemsMenuInsertZwnj LineEditMenuItems = 27
  LineEditMenuItemsMenuInsertWj LineEditMenuItems = 28
  LineEditMenuItemsMenuInsertShy LineEditMenuItems = 29
  LineEditMenuItemsMenuMax LineEditMenuItems = 30
)

type LineEditVirtualKeyboardType int
const (
  LineEditVirtualKeyboardTypeKeyboardTypeDefault LineEditVirtualKeyboardType = 0
  LineEditVirtualKeyboardTypeKeyboardTypeMultiline LineEditVirtualKeyboardType = 1
  LineEditVirtualKeyboardTypeKeyboardTypeNumber LineEditVirtualKeyboardType = 2
  LineEditVirtualKeyboardTypeKeyboardTypeNumberDecimal LineEditVirtualKeyboardType = 3
  LineEditVirtualKeyboardTypeKeyboardTypePhone LineEditVirtualKeyboardType = 4
  LineEditVirtualKeyboardTypeKeyboardTypeEmailAddress LineEditVirtualKeyboardType = 5
  LineEditVirtualKeyboardTypeKeyboardTypePassword LineEditVirtualKeyboardType = 6
  LineEditVirtualKeyboardTypeKeyboardTypeUrl LineEditVirtualKeyboardType = 7
)

func (me *LineEdit) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LineEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LineEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *LineEdit) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetHorizontalAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) Clear()  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) Select(from int, to int, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1328111411) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) SelectAll()  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) Deselect()  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) HasSelection() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) GetSelectedText() String {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) GetSelectionFromColumn() int {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_from_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) GetSelectionToColumn() int {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_to_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetText(text String, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetText() String {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) GetDrawControlChars() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_control_chars")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetDrawControlChars(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_control_chars")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetLanguage() String {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 55961453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385126229) // FIXME: should cache?
  var ret TextServerStructuredTextParser
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetStructuredTextBidiOverrideOptions(args Array, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(args.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetStructuredTextBidiOverrideOptions() Array {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetPlaceholder(text String, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetPlaceholder() String {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetCaretColumn(position int, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetCaretColumn() int {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) GetScrollOffset() float32 {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetExpandToTextLengthEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_to_text_length_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsExpandToTextLengthEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_expand_to_text_length_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetCaretBlinkEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_blink_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsCaretBlinkEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_caret_blink_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetCaretMidGraphemeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_mid_grapheme_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsCaretMidGraphemeEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_caret_mid_grapheme_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetCaretForceDisplayed(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_force_displayed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsCaretForceDisplayed() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_caret_force_displayed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetCaretBlinkInterval(interval float32, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_blink_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetCaretBlinkInterval() float32 {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_blink_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetMaxLength(chars int, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&chars), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetMaxLength() int {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) InsertTextAtCaret(text String, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("insert_text_at_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) DeleteCharAtCaret()  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delete_char_at_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) DeleteText(from_column int, to_column int, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delete_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_column), gdc.ConstTypePtr(&to_column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) SetEditable(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsEditable() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetSecret(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_secret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsSecret() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_secret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetSecretCharacter(character String, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_secret_character")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(character.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetSecretCharacter() String {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_secret_character")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) MenuOption(option int, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("menu_option")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetMenu() PopupMenu {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229722558) // FIXME: should cache?
  var ret PopupMenu
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) IsMenuVisible() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_menu_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetContextMenuEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_context_menu_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsContextMenuEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_context_menu_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetVirtualKeyboardEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_virtual_keyboard_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsVirtualKeyboardEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_virtual_keyboard_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetVirtualKeyboardType(type_ LineEditVirtualKeyboardType, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_virtual_keyboard_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2696893573) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetVirtualKeyboardType() LineEditVirtualKeyboardType {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_virtual_keyboard_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1928699316) // FIXME: should cache?
  var ret LineEditVirtualKeyboardType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetClearButtonEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clear_button_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsClearButtonEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_clear_button_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetShortcutKeysEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_keys_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsShortcutKeysEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shortcut_keys_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetMiddleMousePasteEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_middle_mouse_paste_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsMiddleMousePasteEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_middle_mouse_paste_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetSelectingEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selecting_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsSelectingEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selecting_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetDeselectOnFocusLossEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deselect_on_focus_loss_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsDeselectOnFocusLossEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_deselect_on_focus_loss_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetDragAndDropSelectionEnabled(enable bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_and_drop_selection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsDragAndDropSelectionEnabled() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drag_and_drop_selection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetRightIcon(icon Texture2D, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_right_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(icon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) GetRightIcon() Texture2D {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_right_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 255860311) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetFlat(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsFlat() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LineEdit) SetSelectAllOnFocus(enabled bool, )  {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_select_all_on_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LineEdit) IsSelectAllOnFocus() bool {
  classNameV := StringNameFromStr("LineEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_select_all_on_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type LineEditTextChangedSignalFn func(new_text String, )

func (me *LineEdit) ConnectTextChanged(subs SignalSubscribers, fn LineEditTextChangedSignalFn) {
  sig := StringNameFromStr("text_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *LineEdit) DisconnectTextChanged(subs SignalSubscribers, fn LineEditTextChangedSignalFn) {
  sig := StringNameFromStr("text_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type LineEditTextChangeRejectedSignalFn func(rejected_substring String, )

func (me *LineEdit) ConnectTextChangeRejected(subs SignalSubscribers, fn LineEditTextChangeRejectedSignalFn) {
  sig := StringNameFromStr("text_change_rejected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *LineEdit) DisconnectTextChangeRejected(subs SignalSubscribers, fn LineEditTextChangeRejectedSignalFn) {
  sig := StringNameFromStr("text_change_rejected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type LineEditTextSubmittedSignalFn func(new_text String, )

func (me *LineEdit) ConnectTextSubmitted(subs SignalSubscribers, fn LineEditTextSubmittedSignalFn) {
  sig := StringNameFromStr("text_submitted")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *LineEdit) DisconnectTextSubmitted(subs SignalSubscribers, fn LineEditTextSubmittedSignalFn) {
  sig := StringNameFromStr("text_submitted")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
