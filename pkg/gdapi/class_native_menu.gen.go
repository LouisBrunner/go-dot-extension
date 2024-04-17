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

type ptrsForNativeMenuList struct {
	fnHasFeature               gdc.MethodBindPtr
	fnHasSystemMenu            gdc.MethodBindPtr
	fnGetSystemMenu            gdc.MethodBindPtr
	fnGetSystemMenuName        gdc.MethodBindPtr
	fnCreateMenu               gdc.MethodBindPtr
	fnHasMenu                  gdc.MethodBindPtr
	fnFreeMenu                 gdc.MethodBindPtr
	fnGetSize                  gdc.MethodBindPtr
	fnPopup                    gdc.MethodBindPtr
	fnSetInterfaceDirection    gdc.MethodBindPtr
	fnSetPopupOpenCallback     gdc.MethodBindPtr
	fnGetPopupOpenCallback     gdc.MethodBindPtr
	fnSetPopupCloseCallback    gdc.MethodBindPtr
	fnGetPopupCloseCallback    gdc.MethodBindPtr
	fnSetMinimumWidth          gdc.MethodBindPtr
	fnGetMinimumWidth          gdc.MethodBindPtr
	fnAddSubmenuItem           gdc.MethodBindPtr
	fnAddItem                  gdc.MethodBindPtr
	fnAddCheckItem             gdc.MethodBindPtr
	fnAddIconItem              gdc.MethodBindPtr
	fnAddIconCheckItem         gdc.MethodBindPtr
	fnAddRadioCheckItem        gdc.MethodBindPtr
	fnAddIconRadioCheckItem    gdc.MethodBindPtr
	fnAddMultistateItem        gdc.MethodBindPtr
	fnAddSeparator             gdc.MethodBindPtr
	fnFindItemIndexWithText    gdc.MethodBindPtr
	fnFindItemIndexWithTag     gdc.MethodBindPtr
	fnFindItemIndexWithSubmenu gdc.MethodBindPtr
	fnIsItemChecked            gdc.MethodBindPtr
	fnIsItemCheckable          gdc.MethodBindPtr
	fnIsItemRadioCheckable     gdc.MethodBindPtr
	fnGetItemCallback          gdc.MethodBindPtr
	fnGetItemKeyCallback       gdc.MethodBindPtr
	fnGetItemTag               gdc.MethodBindPtr
	fnGetItemText              gdc.MethodBindPtr
	fnGetItemSubmenu           gdc.MethodBindPtr
	fnGetItemAccelerator       gdc.MethodBindPtr
	fnIsItemDisabled           gdc.MethodBindPtr
	fnIsItemHidden             gdc.MethodBindPtr
	fnGetItemTooltip           gdc.MethodBindPtr
	fnGetItemState             gdc.MethodBindPtr
	fnGetItemMaxStates         gdc.MethodBindPtr
	fnGetItemIcon              gdc.MethodBindPtr
	fnGetItemIndentationLevel  gdc.MethodBindPtr
	fnSetItemChecked           gdc.MethodBindPtr
	fnSetItemCheckable         gdc.MethodBindPtr
	fnSetItemRadioCheckable    gdc.MethodBindPtr
	fnSetItemCallback          gdc.MethodBindPtr
	fnSetItemHoverCallbacks    gdc.MethodBindPtr
	fnSetItemKeyCallback       gdc.MethodBindPtr
	fnSetItemTag               gdc.MethodBindPtr
	fnSetItemText              gdc.MethodBindPtr
	fnSetItemSubmenu           gdc.MethodBindPtr
	fnSetItemAccelerator       gdc.MethodBindPtr
	fnSetItemDisabled          gdc.MethodBindPtr
	fnSetItemHidden            gdc.MethodBindPtr
	fnSetItemTooltip           gdc.MethodBindPtr
	fnSetItemState             gdc.MethodBindPtr
	fnSetItemMaxStates         gdc.MethodBindPtr
	fnSetItemIcon              gdc.MethodBindPtr
	fnSetItemIndentationLevel  gdc.MethodBindPtr
	fnGetItemCount             gdc.MethodBindPtr
	fnIsSystemMenu             gdc.MethodBindPtr
	fnRemoveItem               gdc.MethodBindPtr
	fnClear                    gdc.MethodBindPtr
}

var ptrsForNativeMenu ptrsForNativeMenuList

func initNativeMenuPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NativeMenu")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_feature")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnHasFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1708975490))
	}
	{
		methodName := StringNameFromStr("has_system_menu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnHasSystemMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 718213027))
	}
	{
		methodName := StringNameFromStr("get_system_menu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetSystemMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 469707506))
	}
	{
		methodName := StringNameFromStr("get_system_menu_name")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetSystemMenuName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1281499290))
	}
	{
		methodName := StringNameFromStr("create_menu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnCreateMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("has_menu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnHasMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("free_menu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnFreeMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2440833711))
	}
	{
		methodName := StringNameFromStr("popup")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2450610377))
	}
	{
		methodName := StringNameFromStr("set_interface_direction")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetInterfaceDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("set_popup_open_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetPopupOpenCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
	}
	{
		methodName := StringNameFromStr("get_popup_open_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetPopupOpenCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3170603026))
	}
	{
		methodName := StringNameFromStr("set_popup_close_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetPopupCloseCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
	}
	{
		methodName := StringNameFromStr("get_popup_close_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetPopupCloseCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3170603026))
	}
	{
		methodName := StringNameFromStr("set_minimum_width")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetMinimumWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("get_minimum_width")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetMinimumWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("add_submenu_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddSubmenuItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1002030223))
	}
	{
		methodName := StringNameFromStr("add_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2553375659))
	}
	{
		methodName := StringNameFromStr("add_check_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2553375659))
	}
	{
		methodName := StringNameFromStr("add_icon_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddIconItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2987595282))
	}
	{
		methodName := StringNameFromStr("add_icon_check_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddIconCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2987595282))
	}
	{
		methodName := StringNameFromStr("add_radio_check_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddRadioCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2553375659))
	}
	{
		methodName := StringNameFromStr("add_icon_radio_check_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddIconRadioCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2987595282))
	}
	{
		methodName := StringNameFromStr("add_multistate_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddMultistateItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1558592568))
	}
	{
		methodName := StringNameFromStr("add_separator")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnAddSeparator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 448810126))
	}
	{
		methodName := StringNameFromStr("find_item_index_with_text")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnFindItemIndexWithText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1362438794))
	}
	{
		methodName := StringNameFromStr("find_item_index_with_tag")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnFindItemIndexWithTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1260085030))
	}
	{
		methodName := StringNameFromStr("find_item_index_with_submenu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnFindItemIndexWithSubmenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 893635918))
	}
	{
		methodName := StringNameFromStr("is_item_checked")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnIsItemChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3120086654))
	}
	{
		methodName := StringNameFromStr("is_item_checkable")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnIsItemCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3120086654))
	}
	{
		methodName := StringNameFromStr("is_item_radio_checkable")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnIsItemRadioCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3120086654))
	}
	{
		methodName := StringNameFromStr("get_item_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639989698))
	}
	{
		methodName := StringNameFromStr("get_item_key_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemKeyCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639989698))
	}
	{
		methodName := StringNameFromStr("get_item_tag")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4069510997))
	}
	{
		methodName := StringNameFromStr("get_item_text")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1464764419))
	}
	{
		methodName := StringNameFromStr("get_item_submenu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemSubmenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1066463050))
	}
	{
		methodName := StringNameFromStr("get_item_accelerator")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemAccelerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 316800700))
	}
	{
		methodName := StringNameFromStr("is_item_disabled")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnIsItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3120086654))
	}
	{
		methodName := StringNameFromStr("is_item_hidden")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnIsItemHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3120086654))
	}
	{
		methodName := StringNameFromStr("get_item_tooltip")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1464764419))
	}
	{
		methodName := StringNameFromStr("get_item_state")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("get_item_max_states")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemMaxStates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("get_item_icon")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3391850701))
	}
	{
		methodName := StringNameFromStr("get_item_indentation_level")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemIndentationLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("set_item_checked")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("set_item_checkable")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("set_item_radio_checkable")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemRadioCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("set_item_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2779810226))
	}
	{
		methodName := StringNameFromStr("set_item_hover_callbacks")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemHoverCallbacks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2779810226))
	}
	{
		methodName := StringNameFromStr("set_item_key_callback")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemKeyCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2779810226))
	}
	{
		methodName := StringNameFromStr("set_item_tag")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2706844827))
	}
	{
		methodName := StringNameFromStr("set_item_text")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4153150897))
	}
	{
		methodName := StringNameFromStr("set_item_submenu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemSubmenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
	}
	{
		methodName := StringNameFromStr("set_item_accelerator")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemAccelerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 786300043))
	}
	{
		methodName := StringNameFromStr("set_item_disabled")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("set_item_hidden")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("set_item_tooltip")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4153150897))
	}
	{
		methodName := StringNameFromStr("set_item_state")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
	}
	{
		methodName := StringNameFromStr("set_item_max_states")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemMaxStates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
	}
	{
		methodName := StringNameFromStr("set_item_icon")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1388763257))
	}
	{
		methodName := StringNameFromStr("set_item_indentation_level")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnSetItemIndentationLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
	}
	{
		methodName := StringNameFromStr("get_item_count")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnGetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("is_system_menu")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnIsSystemMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("remove_item")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnRemoveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForNativeMenu.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}

}

type NativeMenu struct {
	Object
}

func (me *NativeMenu) BaseClass() string {
	return "NativeMenu"
}

func NewNativeMenu() *NativeMenu {
	str := StringNameFromStr("NativeMenu") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NativeMenu{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type NativeMenuFeature int

const (
	NativeMenuFeatureFeatureGlobalMenu        NativeMenuFeature = 0
	NativeMenuFeatureFeaturePopupMenu         NativeMenuFeature = 1
	NativeMenuFeatureFeatureOpenCloseCallback NativeMenuFeature = 2
	NativeMenuFeatureFeatureHoverCallback     NativeMenuFeature = 3
	NativeMenuFeatureFeatureKeyCallback       NativeMenuFeature = 4
)

type NativeMenuSystemMenus int

const (
	NativeMenuSystemMenusInvalidMenuId     NativeMenuSystemMenus = 0
	NativeMenuSystemMenusMainMenuId        NativeMenuSystemMenus = 1
	NativeMenuSystemMenusApplicationMenuId NativeMenuSystemMenus = 2
	NativeMenuSystemMenusWindowMenuId      NativeMenuSystemMenus = 3
	NativeMenuSystemMenusHelpMenuId        NativeMenuSystemMenus = 4
	NativeMenuSystemMenusDockMenuId        NativeMenuSystemMenus = 5
)

func (me *NativeMenu) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NativeMenu) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NativeMenu) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NativeMenu) HasFeature(feature NativeMenuFeature) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&feature)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnHasFeature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) HasSystemMenu(menu_id NativeMenuSystemMenus) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&menu_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnHasSystemMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) GetSystemMenu(menu_id NativeMenuSystemMenus) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&menu_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetSystemMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetSystemMenuName(menu_id NativeMenuSystemMenus) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&menu_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetSystemMenuName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) CreateMenu() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnCreateMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) HasMenu(rid RID) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnHasMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) FreeMenu(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnFreeMenu), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) GetSize(rid RID) Vector2 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) Popup(rid RID, position Vector2i) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnPopup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetInterfaceDirection(rid RID, is_rtl bool) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&is_rtl)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetInterfaceDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetPopupOpenCallback(rid RID, callback Callable) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetPopupOpenCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) GetPopupOpenCallback(rid RID) Callable {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetPopupOpenCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) SetPopupCloseCallback(rid RID, callback Callable) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetPopupCloseCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) GetPopupCloseCallback(rid RID) Callable {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetPopupCloseCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) SetMinimumWidth(rid RID, width float64) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetMinimumWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) GetMinimumWidth(rid RID) float64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetMinimumWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddSubmenuItem(rid RID, label String, submenu_rid RID, tag Variant, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), label.AsCTypePtr(), submenu_rid.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddSubmenuItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddItem(rid RID, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddCheckItem(rid RID, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddIconItem(rid RID, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), icon.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddIconItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddIconCheckItem(rid RID, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), icon.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddIconCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddRadioCheckItem(rid RID, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddRadioCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddIconRadioCheckItem(rid RID, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), icon.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddIconRadioCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddMultistateItem(rid RID, label String, max_states int64, default_state int64, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), label.AsCTypePtr(), gdc.ConstTypePtr(&max_states), gdc.ConstTypePtr(&default_state), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&max_states)
	pinner.Pin(&default_state)
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddMultistateItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) AddSeparator(rid RID, index int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnAddSeparator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) FindItemIndexWithText(rid RID, text String) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnFindItemIndexWithText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) FindItemIndexWithTag(rid RID, tag Variant) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), tag.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnFindItemIndexWithTag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) FindItemIndexWithSubmenu(rid RID, submenu_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), submenu_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnFindItemIndexWithSubmenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) IsItemChecked(rid RID, idx int64) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnIsItemChecked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) IsItemCheckable(rid RID, idx int64) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnIsItemCheckable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) IsItemRadioCheckable(rid RID, idx int64) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnIsItemRadioCheckable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) GetItemCallback(rid RID, idx int64) Callable {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetItemKeyCallback(rid RID, idx int64) Callable {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemKeyCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetItemTag(rid RID, idx int64) Variant {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemTag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetItemText(rid RID, idx int64) String {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetItemSubmenu(rid RID, idx int64) RID {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemSubmenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetItemAccelerator(rid RID, idx int64) Key {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemAccelerator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NativeMenu) IsItemDisabled(rid RID, idx int64) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnIsItemDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) IsItemHidden(rid RID, idx int64) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnIsItemHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) GetItemTooltip(rid RID, idx int64) String {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetItemState(rid RID, idx int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) GetItemMaxStates(rid RID, idx int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemMaxStates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) GetItemIcon(rid RID, idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NativeMenu) GetItemIndentationLevel(rid RID, idx int64) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemIndentationLevel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) SetItemChecked(rid RID, idx int64, checked bool) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checked)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemChecked), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemCheckable(rid RID, idx int64, checkable bool) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checkable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemCheckable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemRadioCheckable(rid RID, idx int64, checkable bool) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checkable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemRadioCheckable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemCallback(rid RID, idx int64, callback Callable) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemHoverCallbacks(rid RID, idx int64, callback Callable) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemHoverCallbacks), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemKeyCallback(rid RID, idx int64, key_callback Callable) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), key_callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemKeyCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemTag(rid RID, idx int64, tag Variant) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), tag.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemTag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemText(rid RID, idx int64, text String) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemSubmenu(rid RID, idx int64, submenu_rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), submenu_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemSubmenu), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemAccelerator(rid RID, idx int64, keycode Key) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemAccelerator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemDisabled(rid RID, idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemHidden(rid RID, idx int64, hidden bool) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemHidden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemTooltip(rid RID, idx int64, tooltip String) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemState(rid RID, idx int64, state int64) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemMaxStates(rid RID, idx int64, max_states int64) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&max_states)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemMaxStates), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemIcon(rid RID, idx int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) SetItemIndentationLevel(rid RID, idx int64, level int64) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&level)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnSetItemIndentationLevel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) GetItemCount(rid RID) int64 {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnGetItemCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) IsSystemMenu(rid RID) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnIsSystemMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NativeMenu) RemoveItem(rid RID, idx int64) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnRemoveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NativeMenu) Clear(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNativeMenu.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
