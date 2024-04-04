// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type TextEdit struct {
  Control
}

func (me *TextEdit) BaseClass() string {
  return "TextEdit"
}

func NewTextEdit() *TextEdit {
  str := StringNameFromStr("TextEdit") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextEdit{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TextEditMenuItems int
const (
  TextEditMenuItemsMenuCut TextEditMenuItems = 0
  TextEditMenuItemsMenuCopy TextEditMenuItems = 1
  TextEditMenuItemsMenuPaste TextEditMenuItems = 2
  TextEditMenuItemsMenuClear TextEditMenuItems = 3
  TextEditMenuItemsMenuSelectAll TextEditMenuItems = 4
  TextEditMenuItemsMenuUndo TextEditMenuItems = 5
  TextEditMenuItemsMenuRedo TextEditMenuItems = 6
  TextEditMenuItemsMenuSubmenuTextDir TextEditMenuItems = 7
  TextEditMenuItemsMenuDirInherited TextEditMenuItems = 8
  TextEditMenuItemsMenuDirAuto TextEditMenuItems = 9
  TextEditMenuItemsMenuDirLtr TextEditMenuItems = 10
  TextEditMenuItemsMenuDirRtl TextEditMenuItems = 11
  TextEditMenuItemsMenuDisplayUcc TextEditMenuItems = 12
  TextEditMenuItemsMenuSubmenuInsertUcc TextEditMenuItems = 13
  TextEditMenuItemsMenuInsertLrm TextEditMenuItems = 14
  TextEditMenuItemsMenuInsertRlm TextEditMenuItems = 15
  TextEditMenuItemsMenuInsertLre TextEditMenuItems = 16
  TextEditMenuItemsMenuInsertRle TextEditMenuItems = 17
  TextEditMenuItemsMenuInsertLro TextEditMenuItems = 18
  TextEditMenuItemsMenuInsertRlo TextEditMenuItems = 19
  TextEditMenuItemsMenuInsertPdf TextEditMenuItems = 20
  TextEditMenuItemsMenuInsertAlm TextEditMenuItems = 21
  TextEditMenuItemsMenuInsertLri TextEditMenuItems = 22
  TextEditMenuItemsMenuInsertRli TextEditMenuItems = 23
  TextEditMenuItemsMenuInsertFsi TextEditMenuItems = 24
  TextEditMenuItemsMenuInsertPdi TextEditMenuItems = 25
  TextEditMenuItemsMenuInsertZwj TextEditMenuItems = 26
  TextEditMenuItemsMenuInsertZwnj TextEditMenuItems = 27
  TextEditMenuItemsMenuInsertWj TextEditMenuItems = 28
  TextEditMenuItemsMenuInsertShy TextEditMenuItems = 29
  TextEditMenuItemsMenuMax TextEditMenuItems = 30
)

type TextEditEditAction int
const (
  TextEditEditActionActionNone TextEditEditAction = 0
  TextEditEditActionActionTyping TextEditEditAction = 1
  TextEditEditActionActionBackspace TextEditEditAction = 2
  TextEditEditActionActionDelete TextEditEditAction = 3
)

type TextEditSearchFlags int
const (
  TextEditSearchFlagsSearchMatchCase TextEditSearchFlags = 1
  TextEditSearchFlagsSearchWholeWords TextEditSearchFlags = 2
  TextEditSearchFlagsSearchBackwards TextEditSearchFlags = 4
)

type TextEditCaretType int
const (
  TextEditCaretTypeCaretTypeLine TextEditCaretType = 0
  TextEditCaretTypeCaretTypeBlock TextEditCaretType = 1
)

type TextEditSelectionMode int
const (
  TextEditSelectionModeSelectionModeNone TextEditSelectionMode = 0
  TextEditSelectionModeSelectionModeShift TextEditSelectionMode = 1
  TextEditSelectionModeSelectionModePointer TextEditSelectionMode = 2
  TextEditSelectionModeSelectionModeWord TextEditSelectionMode = 3
  TextEditSelectionModeSelectionModeLine TextEditSelectionMode = 4
)

type TextEditLineWrappingMode int
const (
  TextEditLineWrappingModeLineWrappingNone TextEditLineWrappingMode = 0
  TextEditLineWrappingModeLineWrappingBoundary TextEditLineWrappingMode = 1
)

type TextEditGutterType int
const (
  TextEditGutterTypeGutterTypeString TextEditGutterType = 0
  TextEditGutterTypeGutterTypeIcon TextEditGutterType = 1
  TextEditGutterTypeGutterTypeCustom TextEditGutterType = 2
)

func (me *TextEdit) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextEdit) HasImeText() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_ime_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetEditable(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsEditable() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ControlTextDirection

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextEdit) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLanguage() String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 55961453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385126229) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerStructuredTextParser

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextEdit) SetStructuredTextBidiOverrideOptions(args Array, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{args.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetStructuredTextBidiOverrideOptions() Array {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetTabSize(size int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetTabSize() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetOvertypeModeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_overtype_mode_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsOvertypeModeEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_overtype_mode_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetContextMenuEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_context_menu_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsContextMenuEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_context_menu_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetShortcutKeysEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_keys_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsShortcutKeysEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shortcut_keys_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetVirtualKeyboardEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_virtual_keyboard_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsVirtualKeyboardEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_virtual_keyboard_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetMiddleMousePasteEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_middle_mouse_paste_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsMiddleMousePasteEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_middle_mouse_paste_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) Clear()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SetText(text String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetText() String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetLineCount() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetPlaceholder(text String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetPlaceholder() String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetLine(line int64, new_text String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , new_text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLine(line int64, ) String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetLineWidth(line int64, wrap_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 688195400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&wrap_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)
  pinner.Pin(&wrap_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetLineHeight() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetIndentLevel(line int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_indent_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetFirstNonWhitespaceColumn(line int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_first_non_whitespace_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SwapLines(from_line int64, to_line int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("swap_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&to_line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) InsertLineAt(line int64, text String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("insert_line_at")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) InsertTextAtCaret(text String, caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("insert_text_at_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2697778442) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) RemoveText(from_line int64, from_column int64, to_line int64, to_column int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4275841770) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&from_column) , gdc.ConstTypePtr(&to_line) , gdc.ConstTypePtr(&to_column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLastUnhiddenLine() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_unhidden_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetNextVisibleLineOffsetFrom(line int64, visible_amount int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_visible_line_offset_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&visible_amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)
  pinner.Pin(&visible_amount)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetNextVisibleLineIndexOffsetFrom(line int64, wrap_index int64, visible_amount int64, ) Vector2i {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_visible_line_index_offset_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3386475622) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&wrap_index) , gdc.ConstTypePtr(&visible_amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&line)
  pinner.Pin(&wrap_index)
  pinner.Pin(&visible_amount)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) Backspace(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("backspace")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) Cut(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) Copy(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("copy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) Paste(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("paste")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) PastePrimaryClipboard(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("paste_primary_clipboard")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) StartAction(action TextEditEditAction, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2834827583) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&action) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) EndAction()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("end_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) BeginComplexOperation()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("begin_complex_operation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) EndComplexOperation()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("end_complex_operation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) HasUndo() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_undo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) HasRedo() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_redo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) Undo()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("undo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) Redo()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("redo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) ClearUndoHistory()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_undo_history")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) TagSavedVersion()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tag_saved_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetVersion() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetSavedVersion() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_saved_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetSearchText(search_text String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_search_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{search_text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SetSearchFlags(flags int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_search_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) Search(text String, flags int64, from_line int64, from_colum int64, ) Vector2i {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("search")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1203739136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), gdc.ConstTypePtr(&flags) , gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&from_colum) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&flags)
  pinner.Pin(&from_line)
  pinner.Pin(&from_colum)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetTooltipRequestFunc(callback Callable, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tooltip_request_func")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLocalMousePos() Vector2 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_mouse_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetWordAtPos(position Vector2, ) String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_word_at_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3674420000) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetLineColumnAtPos(position Vector2i, allow_out_of_bounds bool, ) Vector2i {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_column_at_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 239517838) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&allow_out_of_bounds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&allow_out_of_bounds)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetPosAtLineColumn(line int64, column int64, ) Vector2i {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pos_at_line_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410388347) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetRectAtLineColumn(line int64, column int64, ) Rect2i {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rect_at_line_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3256618057) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2i()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetMinimapLineAtPos(position Vector2i, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimap_line_at_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) IsDraggingCursor() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dragging_cursor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) IsMouseOverSelection(edges bool, caret_index int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_mouse_over_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1840282309) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edges) , gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&edges)
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetCaretType(type_ TextEditCaretType, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1211596914) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetCaretType() TextEditCaretType {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2830252959) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextEditCaretType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextEdit) SetCaretBlinkEnabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_blink_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsCaretBlinkEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_caret_blink_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetCaretBlinkInterval(interval float64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_blink_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetCaretBlinkInterval() float64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_blink_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetDrawCaretWhenEditableDisabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_caret_when_editable_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsDrawingCaretWhenEditableDisabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_caret_when_editable_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetMoveCaretOnRightClickEnabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_move_caret_on_right_click_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsMoveCaretOnRightClickEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_move_caret_on_right_click_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetCaretMidGraphemeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_mid_grapheme_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsCaretMidGraphemeEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_caret_mid_grapheme_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetMultipleCaretsEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiple_carets_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsMultipleCaretsEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_multiple_carets_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) AddCaret(line int64, col int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 50157827) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&col) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)
  pinner.Pin(&col)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) RemoveCaret(caret int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) RemoveSecondaryCarets()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_secondary_carets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) MergeOverlappingCarets()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("merge_overlapping_carets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetCaretCount() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) AddCaretAtCarets(below bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_caret_at_carets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&below) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetCaretIndexEditOrder() PackedInt32Array {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_index_edit_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) AdjustCaretsAfterEdit(caret int64, from_line int64, from_col int64, to_line int64, to_col int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("adjust_carets_after_edit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1770277138) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret) , gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&from_col) , gdc.ConstTypePtr(&to_line) , gdc.ConstTypePtr(&to_col) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsCaretVisible(caret_index int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_caret_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1051549951) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetCaretDrawPos(caret_index int64, ) Vector2 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_draw_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 478253731) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetCaretLine(line int64, adjust_viewport bool, can_be_hidden bool, wrap_index int64, caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1302582944) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&adjust_viewport) , gdc.ConstTypePtr(&can_be_hidden) , gdc.ConstTypePtr(&wrap_index) , gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetCaretLine(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetCaretColumn(column int64, adjust_viewport bool, caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_caret_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3796796178) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&adjust_viewport) , gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetCaretColumn(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetCaretWrapIndex(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_caret_wrap_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetWordUnderCaret(caret_index int64, ) String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_word_under_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3929349208) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetSelectingEnabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selecting_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsSelectingEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selecting_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetDeselectOnFocusLossEnabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deselect_on_focus_loss_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsDeselectOnFocusLossEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_deselect_on_focus_loss_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetDragAndDropSelectionEnabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_and_drop_selection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsDragAndDropSelectionEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drag_and_drop_selection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetSelectionMode(mode TextEditSelectionMode, line int64, column int64, caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selection_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1443345937) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetSelectionMode() TextEditSelectionMode {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3750106938) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextEditSelectionMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextEdit) SelectAll()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SelectWordUnderCaret(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_word_under_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) AddSelectionForNextOccurrence()  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_selection_for_next_occurrence")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) Select(from_line int64, from_column int64, to_line int64, to_column int64, caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2560984452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&from_column) , gdc.ConstTypePtr(&to_line) , gdc.ConstTypePtr(&to_column) , gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) HasSelection(caret_index int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2824505868) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetSelectedText(caret_index int64, ) String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2309358862) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetSelectionLine(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetSelectionColumn(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetSelectionFromLine(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_from_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetSelectionFromColumn(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_from_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetSelectionToLine(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_to_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetSelectionToColumn(caret_index int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_to_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&caret_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) Deselect(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) DeleteSelection(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delete_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SetLineWrappingMode(mode TextEditLineWrappingMode, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_wrapping_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2525115309) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLineWrappingMode() TextEditLineWrappingMode {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_wrapping_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3562716114) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextEditLineWrappingMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextEdit) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289138044) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetAutowrapMode() TextServerAutowrapMode {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549071663) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerAutowrapMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextEdit) IsLineWrapped(line int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_wrapped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetLineWrapCount(line int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_wrap_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetLineWrapIndexAtColumn(line int64, column int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_wrap_index_at_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetLineWrappedText(line int64, ) PackedStringArray {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_wrapped_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 647634434) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetSmoothScrollEnabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_smooth_scroll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsSmoothScrollEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_smooth_scroll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetVScrollBar() VScrollBar {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_scroll_bar")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3226026593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVScrollBar()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) GetHScrollBar() HScrollBar {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_scroll_bar")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3774687988) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewHScrollBar()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetVScroll(value float64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetVScroll() float64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetHScroll(value int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetHScroll() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetScrollPastEndOfFileEnabled(enable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_past_end_of_file_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsScrollPastEndOfFileEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_scroll_past_end_of_file_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetVScrollSpeed(speed float64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_scroll_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetVScrollSpeed() float64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_scroll_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetFitContentHeightEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fit_content_height_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsFitContentHeightEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_fit_content_height_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetScrollPosForLine(line int64, wrap_index int64, ) float64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_pos_for_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3929084198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&wrap_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&line)
  pinner.Pin(&wrap_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetLineAsFirstVisible(line int64, wrap_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_as_first_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2230941749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&wrap_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetFirstVisibleLine() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_first_visible_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetLineAsCenterVisible(line int64, wrap_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_as_center_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2230941749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&wrap_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SetLineAsLastVisible(line int64, wrap_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_as_last_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2230941749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&wrap_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLastFullVisibleLine() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_full_visible_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetLastFullVisibleLineWrapIndex() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_full_visible_line_wrap_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetVisibleLineCount() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetVisibleLineCountInRange(from_line int64, to_line int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_line_count_in_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&to_line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&from_line)
  pinner.Pin(&to_line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetTotalVisibleLineCount() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_visible_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) AdjustViewportToCaret(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("adjust_viewport_to_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) CenterViewportToCaret(caret_index int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("center_viewport_to_caret")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SetDrawMinimap(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_minimap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsDrawingMinimap() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_minimap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetMinimapWidth(width int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minimap_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetMinimapWidth() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimap_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetMinimapVisibleLines() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimap_visible_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) AddGutter(at int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) RemoveGutter(gutter int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetGutterCount() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gutter_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetGutterName(gutter int64, name String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gutter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetGutterName(gutter int64, ) String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gutter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetGutterType(gutter int64, type_ TextEditGutterType, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gutter_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1088959071) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetGutterType(gutter int64, ) TextEditGutterType {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gutter_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1159699127) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextEditGutterType
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextEdit) SetGutterWidth(gutter int64, width int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gutter_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetGutterWidth(gutter int64, ) int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gutter_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetGutterDraw(gutter int64, draw bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gutter_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , gdc.ConstTypePtr(&draw) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsGutterDrawn(gutter int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_gutter_drawn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetGutterClickable(gutter int64, clickable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gutter_clickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , gdc.ConstTypePtr(&clickable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsGutterClickable(gutter int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_gutter_clickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetGutterOverwritable(gutter int64, overwritable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gutter_overwritable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , gdc.ConstTypePtr(&overwritable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsGutterOverwritable(gutter int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_gutter_overwritable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) MergeGutters(from_line int64, to_line int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("merge_gutters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&to_line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SetGutterCustomDraw(column int64, draw_callback Callable, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gutter_custom_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 957362965) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column) , draw_callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetTotalGutterWidth() int64 {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_gutter_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetLineGutterMetadata(line int64, gutter int64, metadata Variant, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_gutter_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2060538656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , metadata.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLineGutterMetadata(line int64, gutter int64, ) Variant {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_gutter_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 678354945) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&line)
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetLineGutterText(line int64, gutter int64, text String, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_gutter_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285447957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLineGutterText(line int64, gutter int64, ) String {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_gutter_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1391810591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&line)
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetLineGutterIcon(line int64, gutter int64, icon Texture2D, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_gutter_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 176101966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , icon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLineGutterIcon(line int64, gutter int64, ) Texture2D {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_gutter_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2584904275) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&line)
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetLineGutterItemColor(line int64, gutter int64, color Color, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_gutter_item_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3733378741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLineGutterItemColor(line int64, gutter int64, ) Color {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_gutter_item_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2165839948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&line)
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetLineGutterClickable(line int64, gutter int64, clickable bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_gutter_clickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1383440665) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , gdc.ConstTypePtr(&clickable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsLineGutterClickable(line int64, gutter int64, ) bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_gutter_clickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&gutter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)
  pinner.Pin(&gutter)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetLineBackgroundColor(line int64, color Color, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_background_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetLineBackgroundColor(line int64, ) Color {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_background_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetSyntaxHighlighter(syntax_highlighter SyntaxHighlighter, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_syntax_highlighter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2765644541) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{syntax_highlighter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) GetSyntaxHighlighter() SyntaxHighlighter {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_syntax_highlighter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2721131626) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSyntaxHighlighter()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) SetHighlightCurrentLine(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_highlight_current_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsHighlightCurrentLineEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_highlight_current_line_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetHighlightAllOccurrences(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_highlight_all_occurrences")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsHighlightAllOccurrencesEnabled() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_highlight_all_occurrences_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetDrawControlChars() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_control_chars")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetDrawControlChars(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_control_chars")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) SetDrawTabs(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_tabs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsDrawingTabs() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_tabs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) SetDrawSpaces(enabled bool, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_spaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextEdit) IsDrawingSpaces() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_spaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) GetMenu() PopupMenu {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229722558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPopupMenu()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextEdit) IsMenuVisible() bool {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_menu_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextEdit) MenuOption(option int64, )  {
  classNameV := StringNameFromStr("TextEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("menu_option")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TextEditTextSetSignalFn func()

func (me *TextEdit) ConnectTextSet(subs SignalSubscribers, fn TextEditTextSetSignalFn) {
  sig := StringNameFromStr("text_set")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectTextSet(subs SignalSubscribers, fn TextEditTextSetSignalFn) {
  sig := StringNameFromStr("text_set")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditTextChangedSignalFn func()

func (me *TextEdit) ConnectTextChanged(subs SignalSubscribers, fn TextEditTextChangedSignalFn) {
  sig := StringNameFromStr("text_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectTextChanged(subs SignalSubscribers, fn TextEditTextChangedSignalFn) {
  sig := StringNameFromStr("text_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditLinesEditedFromSignalFn func(from_line int, to_line int, )

func (me *TextEdit) ConnectLinesEditedFrom(subs SignalSubscribers, fn TextEditLinesEditedFromSignalFn) {
  sig := StringNameFromStr("lines_edited_from")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectLinesEditedFrom(subs SignalSubscribers, fn TextEditLinesEditedFromSignalFn) {
  sig := StringNameFromStr("lines_edited_from")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditCaretChangedSignalFn func()

func (me *TextEdit) ConnectCaretChanged(subs SignalSubscribers, fn TextEditCaretChangedSignalFn) {
  sig := StringNameFromStr("caret_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectCaretChanged(subs SignalSubscribers, fn TextEditCaretChangedSignalFn) {
  sig := StringNameFromStr("caret_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditGutterClickedSignalFn func(line int, gutter int, )

func (me *TextEdit) ConnectGutterClicked(subs SignalSubscribers, fn TextEditGutterClickedSignalFn) {
  sig := StringNameFromStr("gutter_clicked")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectGutterClicked(subs SignalSubscribers, fn TextEditGutterClickedSignalFn) {
  sig := StringNameFromStr("gutter_clicked")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditGutterAddedSignalFn func()

func (me *TextEdit) ConnectGutterAdded(subs SignalSubscribers, fn TextEditGutterAddedSignalFn) {
  sig := StringNameFromStr("gutter_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectGutterAdded(subs SignalSubscribers, fn TextEditGutterAddedSignalFn) {
  sig := StringNameFromStr("gutter_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditGutterRemovedSignalFn func()

func (me *TextEdit) ConnectGutterRemoved(subs SignalSubscribers, fn TextEditGutterRemovedSignalFn) {
  sig := StringNameFromStr("gutter_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectGutterRemoved(subs SignalSubscribers, fn TextEditGutterRemovedSignalFn) {
  sig := StringNameFromStr("gutter_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
