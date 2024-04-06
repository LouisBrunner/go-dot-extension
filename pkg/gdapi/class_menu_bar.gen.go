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

type ptrsForMenuBarList struct {
	fnSetSwitchOnHover    gdc.MethodBindPtr
	fnIsSwitchOnHover     gdc.MethodBindPtr
	fnSetDisableShortcuts gdc.MethodBindPtr
	fnSetPreferGlobalMenu gdc.MethodBindPtr
	fnIsPreferGlobalMenu  gdc.MethodBindPtr
	fnIsNativeMenu        gdc.MethodBindPtr
	fnGetMenuCount        gdc.MethodBindPtr
	fnSetTextDirection    gdc.MethodBindPtr
	fnGetTextDirection    gdc.MethodBindPtr
	fnSetLanguage         gdc.MethodBindPtr
	fnGetLanguage         gdc.MethodBindPtr
	fnSetFlat             gdc.MethodBindPtr
	fnIsFlat              gdc.MethodBindPtr
	fnSetStartIndex       gdc.MethodBindPtr
	fnGetStartIndex       gdc.MethodBindPtr
	fnSetMenuTitle        gdc.MethodBindPtr
	fnGetMenuTitle        gdc.MethodBindPtr
	fnSetMenuTooltip      gdc.MethodBindPtr
	fnGetMenuTooltip      gdc.MethodBindPtr
	fnSetMenuDisabled     gdc.MethodBindPtr
	fnIsMenuDisabled      gdc.MethodBindPtr
	fnSetMenuHidden       gdc.MethodBindPtr
	fnIsMenuHidden        gdc.MethodBindPtr
	fnGetMenuPopup        gdc.MethodBindPtr
}

var ptrsForMenuBar ptrsForMenuBarList

func initMenuBarPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MenuBar")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_switch_on_hover")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetSwitchOnHover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_switch_on_hover")
		defer methodName.Destroy()
		ptrsForMenuBar.fnIsSwitchOnHover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_disable_shortcuts")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetDisableShortcuts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_prefer_global_menu")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetPreferGlobalMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_prefer_global_menu")
		defer methodName.Destroy()
		ptrsForMenuBar.fnIsPreferGlobalMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_native_menu")
		defer methodName.Destroy()
		ptrsForMenuBar.fnIsNativeMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_menu_count")
		defer methodName.Destroy()
		ptrsForMenuBar.fnGetMenuCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 119160795))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForMenuBar.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797257663))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForMenuBar.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_flat")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_flat")
		defer methodName.Destroy()
		ptrsForMenuBar.fnIsFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_start_index")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetStartIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_start_index")
		defer methodName.Destroy()
		ptrsForMenuBar.fnGetStartIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_menu_title")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetMenuTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_menu_title")
		defer methodName.Destroy()
		ptrsForMenuBar.fnGetMenuTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_menu_tooltip")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetMenuTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_menu_tooltip")
		defer methodName.Destroy()
		ptrsForMenuBar.fnGetMenuTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_menu_disabled")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetMenuDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_menu_disabled")
		defer methodName.Destroy()
		ptrsForMenuBar.fnIsMenuDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_menu_hidden")
		defer methodName.Destroy()
		ptrsForMenuBar.fnSetMenuHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_menu_hidden")
		defer methodName.Destroy()
		ptrsForMenuBar.fnIsMenuHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_menu_popup")
		defer methodName.Destroy()
		ptrsForMenuBar.fnGetMenuPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2100501353))
	}

}

type MenuBar struct {
	Control
}

func (me *MenuBar) BaseClass() string {
	return "MenuBar"
}

func NewMenuBar() *MenuBar {
	str := StringNameFromStr("MenuBar") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MenuBar{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MenuBar) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MenuBar) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MenuBar) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MenuBar) SetSwitchOnHover(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetSwitchOnHover), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) IsSwitchOnHover() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnIsSwitchOnHover), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) SetDisableShortcuts(disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetDisableShortcuts), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) SetPreferGlobalMenu(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetPreferGlobalMenu), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) IsPreferGlobalMenu() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnIsPreferGlobalMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) IsNativeMenu() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnIsNativeMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) GetMenuCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnGetMenuCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) SetTextDirection(direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) GetTextDirection() ControlTextDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MenuBar) SetLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) GetLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MenuBar) SetFlat(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetFlat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) IsFlat() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnIsFlat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) SetStartIndex(enabled int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetStartIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) GetStartIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnGetStartIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) SetMenuTitle(menu int64, title String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetMenuTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) GetMenuTitle(menu int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&menu)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnGetMenuTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MenuBar) SetMenuTooltip(menu int64, tooltip String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetMenuTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) GetMenuTooltip(menu int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&menu)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnGetMenuTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MenuBar) SetMenuDisabled(menu int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetMenuDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) IsMenuDisabled(menu int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&menu)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnIsMenuDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) SetMenuHidden(menu int64, hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnSetMenuHidden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MenuBar) IsMenuHidden(menu int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&menu)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnIsMenuHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MenuBar) GetMenuPopup(menu int64) PopupMenu {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPopupMenu()
	pinner.Pin(&menu)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuBar.fnGetMenuPopup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
