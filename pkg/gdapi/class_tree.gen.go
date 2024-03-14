// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Tree struct {
  Control
}

func (me *Tree) BaseClass() string {
  return "Tree"
}



// Enums

type TreeSelectMode int
const (
  TreeSelectModeSelectSingle TreeSelectMode = 0
  TreeSelectModeSelectRow TreeSelectMode = 1
  TreeSelectModeSelectMulti TreeSelectMode = 2
)

type TreeDropModeFlags int
const (
  TreeDropModeFlagsDropModeDisabled TreeDropModeFlags = 0
  TreeDropModeFlagsDropModeOnItem TreeDropModeFlags = 1
  TreeDropModeFlagsDropModeInbetween TreeDropModeFlags = 2
)

func (me *Tree) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Tree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Tree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Tree) Clear()  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) CreateItem(parent TreeItem, index int, ) TreeItem {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 528467046) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parent.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetRoot() TreeItem {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetColumnCustomMinimumWidth(column int, min_width int, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_custom_minimum_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&min_width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) SetColumnExpand(column int, expand bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_expand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&expand), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) SetColumnExpandRatio(column int, ratio int, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_expand_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) SetColumnClipContent(column int, enable bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_clip_content")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) IsColumnExpanding(column int, ) bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_column_expanding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) IsColumnClippingContent(column int, ) bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_column_clipping_content")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetColumnExpandRatio(column int, ) int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_column_expand_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetColumnWidth(column int, ) int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_column_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetHideRoot(enable bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hide_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) IsRootHidden() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_root_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetNextSelected(from TreeItem, ) TreeItem {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 873446299) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetSelected() TreeItem {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetSelected(item TreeItem, column int, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2662547442) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(item.AsCTypePtr()), gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetSelectedColumn() int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetPressedButton() int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pressed_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetSelectMode(mode TreeSelectMode, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_select_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3223887270) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetSelectMode() TreeSelectMode {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_select_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 100748571) // FIXME: should cache?
  var ret TreeSelectMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) DeselectAll()  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) SetColumns(amount int, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetColumns() int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetEdited() TreeItem {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1514277247) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetEditedColumn() int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_column")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) EditSelected(force_edit bool, ) bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("edit_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2595650253) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_edit), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetCustomPopupRect() Rect2 {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_popup_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetItemAreaRect(item TreeItem, column int, button_index int, ) Rect2 {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_area_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 47968679) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(item.AsCTypePtr()), gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetItemAtPosition(position Vector2, ) TreeItem {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4193340126) // FIXME: should cache?
  var ret TreeItem
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetColumnAtPosition(position Vector2, ) int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_column_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3820158470) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetDropSectionAtPosition(position Vector2, ) int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drop_section_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3820158470) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetButtonIdAtPosition(position Vector2, ) int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_id_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3820158470) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) EnsureCursorIsVisible()  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ensure_cursor_is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) SetColumnTitlesVisible(visible bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_titles_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) AreColumnTitlesVisible() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_column_titles_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetColumnTitle(column int, title String, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(title.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetColumnTitle(column int, ) String {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_column_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetColumnTitleAlignment(column int, title_alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_title_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3276431499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&title_alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetColumnTitleAlignment(column int, ) HorizontalAlignment {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_column_title_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4171562184) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetColumnTitleDirection(column int, direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_title_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1707680378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetColumnTitleDirection(column int, ) ControlTextDirection {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_column_title_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4235602388) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetColumnTitleLanguage(column int, language String, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_column_title_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetColumnTitleLanguage(column int, ) String {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_column_title_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) GetScroll() Vector2 {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) ScrollToItem(item TreeItem, center_on_item bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scroll_to_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1314737213) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(item.AsCTypePtr()), gdc.ConstTypePtr(&center_on_item), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) SetHScrollEnabled(h_scroll bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_scroll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_scroll), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) IsHScrollEnabled() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_h_scroll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetVScrollEnabled(h_scroll bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_scroll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_scroll), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) IsVScrollEnabled() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_v_scroll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetHideFolding(hide bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hide_folding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hide), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) IsFoldingHidden() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_folding_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetEnableRecursiveFolding(enable bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_recursive_folding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) IsRecursiveFoldingEnabled() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_recursive_folding_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetDropModeFlags(flags int, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drop_mode_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetDropModeFlags() int {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drop_mode_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetAllowRmbSelect(allow bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_rmb_select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetAllowRmbSelect() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_rmb_select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetAllowReselect(allow bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_reselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetAllowReselect() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_reselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tree) SetAllowSearch(allow bool, )  {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_search")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tree) GetAllowSearch() bool {
  classNameV := StringNameFromStr("Tree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_search")
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

type TreeItemSelectedSignalFn func()

func (me *Tree) ConnectItemSelected(subs SignalSubscribers, fn TreeItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemSelected(subs SignalSubscribers, fn TreeItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeCellSelectedSignalFn func()

func (me *Tree) ConnectCellSelected(subs SignalSubscribers, fn TreeCellSelectedSignalFn) {
  sig := StringNameFromStr("cell_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCellSelected(subs SignalSubscribers, fn TreeCellSelectedSignalFn) {
  sig := StringNameFromStr("cell_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeMultiSelectedSignalFn func(item TreeItem, column int, selected bool, )

func (me *Tree) ConnectMultiSelected(subs SignalSubscribers, fn TreeMultiSelectedSignalFn) {
  sig := StringNameFromStr("multi_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectMultiSelected(subs SignalSubscribers, fn TreeMultiSelectedSignalFn) {
  sig := StringNameFromStr("multi_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeItemMouseSelectedSignalFn func(position Vector2, mouse_button_index int, )

func (me *Tree) ConnectItemMouseSelected(subs SignalSubscribers, fn TreeItemMouseSelectedSignalFn) {
  sig := StringNameFromStr("item_mouse_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemMouseSelected(subs SignalSubscribers, fn TreeItemMouseSelectedSignalFn) {
  sig := StringNameFromStr("item_mouse_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeEmptyClickedSignalFn func(position Vector2, mouse_button_index int, )

func (me *Tree) ConnectEmptyClicked(subs SignalSubscribers, fn TreeEmptyClickedSignalFn) {
  sig := StringNameFromStr("empty_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectEmptyClicked(subs SignalSubscribers, fn TreeEmptyClickedSignalFn) {
  sig := StringNameFromStr("empty_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeItemEditedSignalFn func()

func (me *Tree) ConnectItemEdited(subs SignalSubscribers, fn TreeItemEditedSignalFn) {
  sig := StringNameFromStr("item_edited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemEdited(subs SignalSubscribers, fn TreeItemEditedSignalFn) {
  sig := StringNameFromStr("item_edited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeCustomItemClickedSignalFn func(mouse_button_index int, )

func (me *Tree) ConnectCustomItemClicked(subs SignalSubscribers, fn TreeCustomItemClickedSignalFn) {
  sig := StringNameFromStr("custom_item_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCustomItemClicked(subs SignalSubscribers, fn TreeCustomItemClickedSignalFn) {
  sig := StringNameFromStr("custom_item_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeItemIconDoubleClickedSignalFn func()

func (me *Tree) ConnectItemIconDoubleClicked(subs SignalSubscribers, fn TreeItemIconDoubleClickedSignalFn) {
  sig := StringNameFromStr("item_icon_double_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemIconDoubleClicked(subs SignalSubscribers, fn TreeItemIconDoubleClickedSignalFn) {
  sig := StringNameFromStr("item_icon_double_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeItemCollapsedSignalFn func(item TreeItem, )

func (me *Tree) ConnectItemCollapsed(subs SignalSubscribers, fn TreeItemCollapsedSignalFn) {
  sig := StringNameFromStr("item_collapsed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemCollapsed(subs SignalSubscribers, fn TreeItemCollapsedSignalFn) {
  sig := StringNameFromStr("item_collapsed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeCheckPropagatedToItemSignalFn func(item TreeItem, column int, )

func (me *Tree) ConnectCheckPropagatedToItem(subs SignalSubscribers, fn TreeCheckPropagatedToItemSignalFn) {
  sig := StringNameFromStr("check_propagated_to_item")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCheckPropagatedToItem(subs SignalSubscribers, fn TreeCheckPropagatedToItemSignalFn) {
  sig := StringNameFromStr("check_propagated_to_item")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeButtonClickedSignalFn func(item TreeItem, column int, id int, mouse_button_index int, )

func (me *Tree) ConnectButtonClicked(subs SignalSubscribers, fn TreeButtonClickedSignalFn) {
  sig := StringNameFromStr("button_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectButtonClicked(subs SignalSubscribers, fn TreeButtonClickedSignalFn) {
  sig := StringNameFromStr("button_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeCustomPopupEditedSignalFn func(arrow_clicked bool, )

func (me *Tree) ConnectCustomPopupEdited(subs SignalSubscribers, fn TreeCustomPopupEditedSignalFn) {
  sig := StringNameFromStr("custom_popup_edited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCustomPopupEdited(subs SignalSubscribers, fn TreeCustomPopupEditedSignalFn) {
  sig := StringNameFromStr("custom_popup_edited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeItemActivatedSignalFn func()

func (me *Tree) ConnectItemActivated(subs SignalSubscribers, fn TreeItemActivatedSignalFn) {
  sig := StringNameFromStr("item_activated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemActivated(subs SignalSubscribers, fn TreeItemActivatedSignalFn) {
  sig := StringNameFromStr("item_activated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeColumnTitleClickedSignalFn func(column int, mouse_button_index int, )

func (me *Tree) ConnectColumnTitleClicked(subs SignalSubscribers, fn TreeColumnTitleClickedSignalFn) {
  sig := StringNameFromStr("column_title_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectColumnTitleClicked(subs SignalSubscribers, fn TreeColumnTitleClickedSignalFn) {
  sig := StringNameFromStr("column_title_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TreeNothingSelectedSignalFn func()

func (me *Tree) ConnectNothingSelected(subs SignalSubscribers, fn TreeNothingSelectedSignalFn) {
  sig := StringNameFromStr("nothing_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectNothingSelected(subs SignalSubscribers, fn TreeNothingSelectedSignalFn) {
  sig := StringNameFromStr("nothing_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
