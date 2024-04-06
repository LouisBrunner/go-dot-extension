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

type ptrsForTreeList struct {
	fnClear                       gdc.MethodBindPtr
	fnCreateItem                  gdc.MethodBindPtr
	fnGetRoot                     gdc.MethodBindPtr
	fnSetColumnCustomMinimumWidth gdc.MethodBindPtr
	fnSetColumnExpand             gdc.MethodBindPtr
	fnSetColumnExpandRatio        gdc.MethodBindPtr
	fnSetColumnClipContent        gdc.MethodBindPtr
	fnIsColumnExpanding           gdc.MethodBindPtr
	fnIsColumnClippingContent     gdc.MethodBindPtr
	fnGetColumnExpandRatio        gdc.MethodBindPtr
	fnGetColumnWidth              gdc.MethodBindPtr
	fnSetHideRoot                 gdc.MethodBindPtr
	fnIsRootHidden                gdc.MethodBindPtr
	fnGetNextSelected             gdc.MethodBindPtr
	fnGetSelected                 gdc.MethodBindPtr
	fnSetSelected                 gdc.MethodBindPtr
	fnGetSelectedColumn           gdc.MethodBindPtr
	fnGetPressedButton            gdc.MethodBindPtr
	fnSetSelectMode               gdc.MethodBindPtr
	fnGetSelectMode               gdc.MethodBindPtr
	fnDeselectAll                 gdc.MethodBindPtr
	fnSetColumns                  gdc.MethodBindPtr
	fnGetColumns                  gdc.MethodBindPtr
	fnGetEdited                   gdc.MethodBindPtr
	fnGetEditedColumn             gdc.MethodBindPtr
	fnEditSelected                gdc.MethodBindPtr
	fnGetCustomPopupRect          gdc.MethodBindPtr
	fnGetItemAreaRect             gdc.MethodBindPtr
	fnGetItemAtPosition           gdc.MethodBindPtr
	fnGetColumnAtPosition         gdc.MethodBindPtr
	fnGetDropSectionAtPosition    gdc.MethodBindPtr
	fnGetButtonIdAtPosition       gdc.MethodBindPtr
	fnEnsureCursorIsVisible       gdc.MethodBindPtr
	fnSetColumnTitlesVisible      gdc.MethodBindPtr
	fnAreColumnTitlesVisible      gdc.MethodBindPtr
	fnSetColumnTitle              gdc.MethodBindPtr
	fnGetColumnTitle              gdc.MethodBindPtr
	fnSetColumnTitleAlignment     gdc.MethodBindPtr
	fnGetColumnTitleAlignment     gdc.MethodBindPtr
	fnSetColumnTitleDirection     gdc.MethodBindPtr
	fnGetColumnTitleDirection     gdc.MethodBindPtr
	fnSetColumnTitleLanguage      gdc.MethodBindPtr
	fnGetColumnTitleLanguage      gdc.MethodBindPtr
	fnGetScroll                   gdc.MethodBindPtr
	fnScrollToItem                gdc.MethodBindPtr
	fnSetHScrollEnabled           gdc.MethodBindPtr
	fnIsHScrollEnabled            gdc.MethodBindPtr
	fnSetVScrollEnabled           gdc.MethodBindPtr
	fnIsVScrollEnabled            gdc.MethodBindPtr
	fnSetHideFolding              gdc.MethodBindPtr
	fnIsFoldingHidden             gdc.MethodBindPtr
	fnSetEnableRecursiveFolding   gdc.MethodBindPtr
	fnIsRecursiveFoldingEnabled   gdc.MethodBindPtr
	fnSetDropModeFlags            gdc.MethodBindPtr
	fnGetDropModeFlags            gdc.MethodBindPtr
	fnSetAllowRmbSelect           gdc.MethodBindPtr
	fnGetAllowRmbSelect           gdc.MethodBindPtr
	fnSetAllowReselect            gdc.MethodBindPtr
	fnGetAllowReselect            gdc.MethodBindPtr
	fnSetAllowSearch              gdc.MethodBindPtr
	fnGetAllowSearch              gdc.MethodBindPtr
}

var ptrsForTree ptrsForTreeList

func initTreePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Tree")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForTree.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("create_item")
		defer methodName.Destroy()
		ptrsForTree.fnCreateItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 528467046))
	}
	{
		methodName := StringNameFromStr("get_root")
		defer methodName.Destroy()
		ptrsForTree.fnGetRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1514277247))
	}
	{
		methodName := StringNameFromStr("set_column_custom_minimum_width")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnCustomMinimumWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_column_expand")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnExpand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_column_expand_ratio")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnExpandRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_column_clip_content")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnClipContent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_column_expanding")
		defer methodName.Destroy()
		ptrsForTree.fnIsColumnExpanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("is_column_clipping_content")
		defer methodName.Destroy()
		ptrsForTree.fnIsColumnClippingContent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_column_expand_ratio")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumnExpandRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_column_width")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumnWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_hide_root")
		defer methodName.Destroy()
		ptrsForTree.fnSetHideRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_root_hidden")
		defer methodName.Destroy()
		ptrsForTree.fnIsRootHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_next_selected")
		defer methodName.Destroy()
		ptrsForTree.fnGetNextSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 873446299))
	}
	{
		methodName := StringNameFromStr("get_selected")
		defer methodName.Destroy()
		ptrsForTree.fnGetSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1514277247))
	}
	{
		methodName := StringNameFromStr("set_selected")
		defer methodName.Destroy()
		ptrsForTree.fnSetSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2662547442))
	}
	{
		methodName := StringNameFromStr("get_selected_column")
		defer methodName.Destroy()
		ptrsForTree.fnGetSelectedColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_pressed_button")
		defer methodName.Destroy()
		ptrsForTree.fnGetPressedButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_select_mode")
		defer methodName.Destroy()
		ptrsForTree.fnSetSelectMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3223887270))
	}
	{
		methodName := StringNameFromStr("get_select_mode")
		defer methodName.Destroy()
		ptrsForTree.fnGetSelectMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 100748571))
	}
	{
		methodName := StringNameFromStr("deselect_all")
		defer methodName.Destroy()
		ptrsForTree.fnDeselectAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_columns")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_columns")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_edited")
		defer methodName.Destroy()
		ptrsForTree.fnGetEdited = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1514277247))
	}
	{
		methodName := StringNameFromStr("get_edited_column")
		defer methodName.Destroy()
		ptrsForTree.fnGetEditedColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("edit_selected")
		defer methodName.Destroy()
		ptrsForTree.fnEditSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2595650253))
	}
	{
		methodName := StringNameFromStr("get_custom_popup_rect")
		defer methodName.Destroy()
		ptrsForTree.fnGetCustomPopupRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("get_item_area_rect")
		defer methodName.Destroy()
		ptrsForTree.fnGetItemAreaRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 47968679))
	}
	{
		methodName := StringNameFromStr("get_item_at_position")
		defer methodName.Destroy()
		ptrsForTree.fnGetItemAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4193340126))
	}
	{
		methodName := StringNameFromStr("get_column_at_position")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumnAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3820158470))
	}
	{
		methodName := StringNameFromStr("get_drop_section_at_position")
		defer methodName.Destroy()
		ptrsForTree.fnGetDropSectionAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3820158470))
	}
	{
		methodName := StringNameFromStr("get_button_id_at_position")
		defer methodName.Destroy()
		ptrsForTree.fnGetButtonIdAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3820158470))
	}
	{
		methodName := StringNameFromStr("ensure_cursor_is_visible")
		defer methodName.Destroy()
		ptrsForTree.fnEnsureCursorIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_column_titles_visible")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnTitlesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("are_column_titles_visible")
		defer methodName.Destroy()
		ptrsForTree.fnAreColumnTitlesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_column_title")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_column_title")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumnTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_column_title_alignment")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnTitleAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3276431499))
	}
	{
		methodName := StringNameFromStr("get_column_title_alignment")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumnTitleAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4171562184))
	}
	{
		methodName := StringNameFromStr("set_column_title_direction")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnTitleDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1707680378))
	}
	{
		methodName := StringNameFromStr("get_column_title_direction")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumnTitleDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4235602388))
	}
	{
		methodName := StringNameFromStr("set_column_title_language")
		defer methodName.Destroy()
		ptrsForTree.fnSetColumnTitleLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_column_title_language")
		defer methodName.Destroy()
		ptrsForTree.fnGetColumnTitleLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_scroll")
		defer methodName.Destroy()
		ptrsForTree.fnGetScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("scroll_to_item")
		defer methodName.Destroy()
		ptrsForTree.fnScrollToItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1314737213))
	}
	{
		methodName := StringNameFromStr("set_h_scroll_enabled")
		defer methodName.Destroy()
		ptrsForTree.fnSetHScrollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_h_scroll_enabled")
		defer methodName.Destroy()
		ptrsForTree.fnIsHScrollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_v_scroll_enabled")
		defer methodName.Destroy()
		ptrsForTree.fnSetVScrollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_v_scroll_enabled")
		defer methodName.Destroy()
		ptrsForTree.fnIsVScrollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hide_folding")
		defer methodName.Destroy()
		ptrsForTree.fnSetHideFolding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_folding_hidden")
		defer methodName.Destroy()
		ptrsForTree.fnIsFoldingHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_enable_recursive_folding")
		defer methodName.Destroy()
		ptrsForTree.fnSetEnableRecursiveFolding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_recursive_folding_enabled")
		defer methodName.Destroy()
		ptrsForTree.fnIsRecursiveFoldingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_drop_mode_flags")
		defer methodName.Destroy()
		ptrsForTree.fnSetDropModeFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_drop_mode_flags")
		defer methodName.Destroy()
		ptrsForTree.fnGetDropModeFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_allow_rmb_select")
		defer methodName.Destroy()
		ptrsForTree.fnSetAllowRmbSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_allow_rmb_select")
		defer methodName.Destroy()
		ptrsForTree.fnGetAllowRmbSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_reselect")
		defer methodName.Destroy()
		ptrsForTree.fnSetAllowReselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_allow_reselect")
		defer methodName.Destroy()
		ptrsForTree.fnGetAllowReselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_search")
		defer methodName.Destroy()
		ptrsForTree.fnSetAllowSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_allow_search")
		defer methodName.Destroy()
		ptrsForTree.fnGetAllowSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type Tree struct {
	Control
}

func (me *Tree) BaseClass() string {
	return "Tree"
}

func NewTree() *Tree {
	str := StringNameFromStr("Tree") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Tree{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TreeSelectMode int

const (
	TreeSelectModeSelectSingle TreeSelectMode = 0
	TreeSelectModeSelectRow    TreeSelectMode = 1
	TreeSelectModeSelectMulti  TreeSelectMode = 2
)

type TreeDropModeFlags int

const (
	TreeDropModeFlagsDropModeDisabled  TreeDropModeFlags = 0
	TreeDropModeFlagsDropModeOnItem    TreeDropModeFlags = 1
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

func (me *Tree) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) CreateItem(parent TreeItem, index int64) TreeItem {
	cargs := []gdc.ConstTypePtr{parent.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnCreateItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) GetRoot() TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetRoot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) SetColumnCustomMinimumWidth(column int64, min_width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&min_width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnCustomMinimumWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) SetColumnExpand(column int64, expand bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&expand)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnExpand), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) SetColumnExpandRatio(column int64, ratio int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnExpandRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) SetColumnClipContent(column int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnClipContent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) IsColumnExpanding(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnIsColumnExpanding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) IsColumnClippingContent(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnIsColumnClippingContent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetColumnExpandRatio(column int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumnExpandRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetColumnWidth(column int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumnWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetHideRoot(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetHideRoot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) IsRootHidden() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnIsRootHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetNextSelected(from TreeItem) TreeItem {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetNextSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) GetSelected() TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) SetSelected(item TreeItem, column int64) {
	cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetSelected), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetSelectedColumn() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetSelectedColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetPressedButton() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetPressedButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetSelectMode(mode TreeSelectMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetSelectMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetSelectMode() TreeSelectMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TreeSelectMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetSelectMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Tree) DeselectAll() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnDeselectAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) SetColumns(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumns), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetColumns() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumns), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetEdited() TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetEdited), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) GetEditedColumn() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetEditedColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) EditSelected(force_edit bool) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_edit)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&force_edit)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnEditSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetCustomPopupRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetCustomPopupRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) GetItemAreaRect(item TreeItem, column int64, button_index int64) Rect2 {
	cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&column)
	pinner.Pin(&button_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetItemAreaRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) GetItemAtPosition(position Vector2) TreeItem {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetItemAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) GetColumnAtPosition(position Vector2) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumnAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetDropSectionAtPosition(position Vector2) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetDropSectionAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) GetButtonIdAtPosition(position Vector2) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetButtonIdAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) EnsureCursorIsVisible() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnEnsureCursorIsVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) SetColumnTitlesVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnTitlesVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) AreColumnTitlesVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnAreColumnTitlesVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetColumnTitle(column int64, title String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetColumnTitle(column int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumnTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) SetColumnTitleAlignment(column int64, title_alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&title_alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnTitleAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetColumnTitleAlignment(column int64) HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumnTitleAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Tree) SetColumnTitleDirection(column int64, direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnTitleDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetColumnTitleDirection(column int64) ControlTextDirection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumnTitleDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Tree) SetColumnTitleLanguage(column int64, language String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetColumnTitleLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetColumnTitleLanguage(column int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetColumnTitleLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) GetScroll() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetScroll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tree) ScrollToItem(item TreeItem, center_on_item bool) {
	cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&center_on_item)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnScrollToItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) SetHScrollEnabled(h_scroll bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_scroll)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetHScrollEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) IsHScrollEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnIsHScrollEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetVScrollEnabled(h_scroll bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_scroll)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetVScrollEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) IsVScrollEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnIsVScrollEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetHideFolding(hide bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hide)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetHideFolding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) IsFoldingHidden() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnIsFoldingHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetEnableRecursiveFolding(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetEnableRecursiveFolding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) IsRecursiveFoldingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnIsRecursiveFoldingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetDropModeFlags(flags int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetDropModeFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetDropModeFlags() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetDropModeFlags), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetAllowRmbSelect(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetAllowRmbSelect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetAllowRmbSelect() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetAllowRmbSelect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetAllowReselect(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetAllowReselect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetAllowReselect() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetAllowReselect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tree) SetAllowSearch(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnSetAllowSearch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tree) GetAllowSearch() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTree.fnGetAllowSearch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TreeItemSelectedSignalFn func()

func (me *Tree) ConnectItemSelected(subs SignalSubscribers, fn TreeItemSelectedSignalFn) {
	sig := StringNameFromStr("item_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemSelected(subs SignalSubscribers, fn TreeItemSelectedSignalFn) {
	sig := StringNameFromStr("item_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeCellSelectedSignalFn func()

func (me *Tree) ConnectCellSelected(subs SignalSubscribers, fn TreeCellSelectedSignalFn) {
	sig := StringNameFromStr("cell_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCellSelected(subs SignalSubscribers, fn TreeCellSelectedSignalFn) {
	sig := StringNameFromStr("cell_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeMultiSelectedSignalFn func(item TreeItem, column int, selected bool)

func (me *Tree) ConnectMultiSelected(subs SignalSubscribers, fn TreeMultiSelectedSignalFn) {
	sig := StringNameFromStr("multi_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectMultiSelected(subs SignalSubscribers, fn TreeMultiSelectedSignalFn) {
	sig := StringNameFromStr("multi_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeItemMouseSelectedSignalFn func(position Vector2, mouse_button_index int)

func (me *Tree) ConnectItemMouseSelected(subs SignalSubscribers, fn TreeItemMouseSelectedSignalFn) {
	sig := StringNameFromStr("item_mouse_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemMouseSelected(subs SignalSubscribers, fn TreeItemMouseSelectedSignalFn) {
	sig := StringNameFromStr("item_mouse_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeEmptyClickedSignalFn func(position Vector2, mouse_button_index int)

func (me *Tree) ConnectEmptyClicked(subs SignalSubscribers, fn TreeEmptyClickedSignalFn) {
	sig := StringNameFromStr("empty_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectEmptyClicked(subs SignalSubscribers, fn TreeEmptyClickedSignalFn) {
	sig := StringNameFromStr("empty_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeItemEditedSignalFn func()

func (me *Tree) ConnectItemEdited(subs SignalSubscribers, fn TreeItemEditedSignalFn) {
	sig := StringNameFromStr("item_edited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemEdited(subs SignalSubscribers, fn TreeItemEditedSignalFn) {
	sig := StringNameFromStr("item_edited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeCustomItemClickedSignalFn func(mouse_button_index int)

func (me *Tree) ConnectCustomItemClicked(subs SignalSubscribers, fn TreeCustomItemClickedSignalFn) {
	sig := StringNameFromStr("custom_item_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCustomItemClicked(subs SignalSubscribers, fn TreeCustomItemClickedSignalFn) {
	sig := StringNameFromStr("custom_item_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeItemIconDoubleClickedSignalFn func()

func (me *Tree) ConnectItemIconDoubleClicked(subs SignalSubscribers, fn TreeItemIconDoubleClickedSignalFn) {
	sig := StringNameFromStr("item_icon_double_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemIconDoubleClicked(subs SignalSubscribers, fn TreeItemIconDoubleClickedSignalFn) {
	sig := StringNameFromStr("item_icon_double_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeItemCollapsedSignalFn func(item TreeItem)

func (me *Tree) ConnectItemCollapsed(subs SignalSubscribers, fn TreeItemCollapsedSignalFn) {
	sig := StringNameFromStr("item_collapsed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemCollapsed(subs SignalSubscribers, fn TreeItemCollapsedSignalFn) {
	sig := StringNameFromStr("item_collapsed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeCheckPropagatedToItemSignalFn func(item TreeItem, column int)

func (me *Tree) ConnectCheckPropagatedToItem(subs SignalSubscribers, fn TreeCheckPropagatedToItemSignalFn) {
	sig := StringNameFromStr("check_propagated_to_item")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCheckPropagatedToItem(subs SignalSubscribers, fn TreeCheckPropagatedToItemSignalFn) {
	sig := StringNameFromStr("check_propagated_to_item")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeButtonClickedSignalFn func(item TreeItem, column int, id int, mouse_button_index int)

func (me *Tree) ConnectButtonClicked(subs SignalSubscribers, fn TreeButtonClickedSignalFn) {
	sig := StringNameFromStr("button_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectButtonClicked(subs SignalSubscribers, fn TreeButtonClickedSignalFn) {
	sig := StringNameFromStr("button_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeCustomPopupEditedSignalFn func(arrow_clicked bool)

func (me *Tree) ConnectCustomPopupEdited(subs SignalSubscribers, fn TreeCustomPopupEditedSignalFn) {
	sig := StringNameFromStr("custom_popup_edited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectCustomPopupEdited(subs SignalSubscribers, fn TreeCustomPopupEditedSignalFn) {
	sig := StringNameFromStr("custom_popup_edited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeItemActivatedSignalFn func()

func (me *Tree) ConnectItemActivated(subs SignalSubscribers, fn TreeItemActivatedSignalFn) {
	sig := StringNameFromStr("item_activated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectItemActivated(subs SignalSubscribers, fn TreeItemActivatedSignalFn) {
	sig := StringNameFromStr("item_activated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeColumnTitleClickedSignalFn func(column int, mouse_button_index int)

func (me *Tree) ConnectColumnTitleClicked(subs SignalSubscribers, fn TreeColumnTitleClickedSignalFn) {
	sig := StringNameFromStr("column_title_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectColumnTitleClicked(subs SignalSubscribers, fn TreeColumnTitleClickedSignalFn) {
	sig := StringNameFromStr("column_title_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TreeNothingSelectedSignalFn func()

func (me *Tree) ConnectNothingSelected(subs SignalSubscribers, fn TreeNothingSelectedSignalFn) {
	sig := StringNameFromStr("nothing_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tree) DisconnectNothingSelected(subs SignalSubscribers, fn TreeNothingSelectedSignalFn) {
	sig := StringNameFromStr("nothing_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
