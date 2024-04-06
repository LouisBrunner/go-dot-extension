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

type ptrsForThemeList struct {
	fnSetIcon              gdc.MethodBindPtr
	fnGetIcon              gdc.MethodBindPtr
	fnHasIcon              gdc.MethodBindPtr
	fnRenameIcon           gdc.MethodBindPtr
	fnClearIcon            gdc.MethodBindPtr
	fnGetIconList          gdc.MethodBindPtr
	fnGetIconTypeList      gdc.MethodBindPtr
	fnSetStylebox          gdc.MethodBindPtr
	fnGetStylebox          gdc.MethodBindPtr
	fnHasStylebox          gdc.MethodBindPtr
	fnRenameStylebox       gdc.MethodBindPtr
	fnClearStylebox        gdc.MethodBindPtr
	fnGetStyleboxList      gdc.MethodBindPtr
	fnGetStyleboxTypeList  gdc.MethodBindPtr
	fnSetFont              gdc.MethodBindPtr
	fnGetFont              gdc.MethodBindPtr
	fnHasFont              gdc.MethodBindPtr
	fnRenameFont           gdc.MethodBindPtr
	fnClearFont            gdc.MethodBindPtr
	fnGetFontList          gdc.MethodBindPtr
	fnGetFontTypeList      gdc.MethodBindPtr
	fnSetFontSize          gdc.MethodBindPtr
	fnGetFontSize          gdc.MethodBindPtr
	fnHasFontSize          gdc.MethodBindPtr
	fnRenameFontSize       gdc.MethodBindPtr
	fnClearFontSize        gdc.MethodBindPtr
	fnGetFontSizeList      gdc.MethodBindPtr
	fnGetFontSizeTypeList  gdc.MethodBindPtr
	fnSetColor             gdc.MethodBindPtr
	fnGetColor             gdc.MethodBindPtr
	fnHasColor             gdc.MethodBindPtr
	fnRenameColor          gdc.MethodBindPtr
	fnClearColor           gdc.MethodBindPtr
	fnGetColorList         gdc.MethodBindPtr
	fnGetColorTypeList     gdc.MethodBindPtr
	fnSetConstant          gdc.MethodBindPtr
	fnGetConstant          gdc.MethodBindPtr
	fnHasConstant          gdc.MethodBindPtr
	fnRenameConstant       gdc.MethodBindPtr
	fnClearConstant        gdc.MethodBindPtr
	fnGetConstantList      gdc.MethodBindPtr
	fnGetConstantTypeList  gdc.MethodBindPtr
	fnSetDefaultBaseScale  gdc.MethodBindPtr
	fnGetDefaultBaseScale  gdc.MethodBindPtr
	fnHasDefaultBaseScale  gdc.MethodBindPtr
	fnSetDefaultFont       gdc.MethodBindPtr
	fnGetDefaultFont       gdc.MethodBindPtr
	fnHasDefaultFont       gdc.MethodBindPtr
	fnSetDefaultFontSize   gdc.MethodBindPtr
	fnGetDefaultFontSize   gdc.MethodBindPtr
	fnHasDefaultFontSize   gdc.MethodBindPtr
	fnSetThemeItem         gdc.MethodBindPtr
	fnGetThemeItem         gdc.MethodBindPtr
	fnHasThemeItem         gdc.MethodBindPtr
	fnRenameThemeItem      gdc.MethodBindPtr
	fnClearThemeItem       gdc.MethodBindPtr
	fnGetThemeItemList     gdc.MethodBindPtr
	fnGetThemeItemTypeList gdc.MethodBindPtr
	fnSetTypeVariation     gdc.MethodBindPtr
	fnIsTypeVariation      gdc.MethodBindPtr
	fnClearTypeVariation   gdc.MethodBindPtr
	fnGetTypeVariationBase gdc.MethodBindPtr
	fnGetTypeVariationList gdc.MethodBindPtr
	fnAddType              gdc.MethodBindPtr
	fnRemoveType           gdc.MethodBindPtr
	fnGetTypeList          gdc.MethodBindPtr
	fnMergeWith            gdc.MethodBindPtr
	fnClear                gdc.MethodBindPtr
}

var ptrsForTheme ptrsForThemeList

func initThemePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Theme")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_icon")
		defer methodName.Destroy()
		ptrsForTheme.fnSetIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2188371082))
	}
	{
		methodName := StringNameFromStr("get_icon")
		defer methodName.Destroy()
		ptrsForTheme.fnGetIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 934555193))
	}
	{
		methodName := StringNameFromStr("has_icon")
		defer methodName.Destroy()
		ptrsForTheme.fnHasIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("rename_icon")
		defer methodName.Destroy()
		ptrsForTheme.fnRenameIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642128662))
	}
	{
		methodName := StringNameFromStr("clear_icon")
		defer methodName.Destroy()
		ptrsForTheme.fnClearIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("get_icon_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetIconList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("get_icon_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetIconTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_stylebox")
		defer methodName.Destroy()
		ptrsForTheme.fnSetStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2075907568))
	}
	{
		methodName := StringNameFromStr("get_stylebox")
		defer methodName.Destroy()
		ptrsForTheme.fnGetStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3405608165))
	}
	{
		methodName := StringNameFromStr("has_stylebox")
		defer methodName.Destroy()
		ptrsForTheme.fnHasStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("rename_stylebox")
		defer methodName.Destroy()
		ptrsForTheme.fnRenameStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642128662))
	}
	{
		methodName := StringNameFromStr("clear_stylebox")
		defer methodName.Destroy()
		ptrsForTheme.fnClearStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("get_stylebox_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetStyleboxList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("get_stylebox_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetStyleboxTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_font")
		defer methodName.Destroy()
		ptrsForTheme.fnSetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 177292320))
	}
	{
		methodName := StringNameFromStr("get_font")
		defer methodName.Destroy()
		ptrsForTheme.fnGetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3445063586))
	}
	{
		methodName := StringNameFromStr("has_font")
		defer methodName.Destroy()
		ptrsForTheme.fnHasFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("rename_font")
		defer methodName.Destroy()
		ptrsForTheme.fnRenameFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642128662))
	}
	{
		methodName := StringNameFromStr("clear_font")
		defer methodName.Destroy()
		ptrsForTheme.fnClearFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("get_font_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetFontList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("get_font_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetFontTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnSetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 281601298))
	}
	{
		methodName := StringNameFromStr("get_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnGetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2419549490))
	}
	{
		methodName := StringNameFromStr("has_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnHasFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("rename_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnRenameFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642128662))
	}
	{
		methodName := StringNameFromStr("clear_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnClearFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("get_font_size_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetFontSizeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("get_font_size_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetFontSizeTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_color")
		defer methodName.Destroy()
		ptrsForTheme.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4111215154))
	}
	{
		methodName := StringNameFromStr("get_color")
		defer methodName.Destroy()
		ptrsForTheme.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2015923404))
	}
	{
		methodName := StringNameFromStr("has_color")
		defer methodName.Destroy()
		ptrsForTheme.fnHasColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("rename_color")
		defer methodName.Destroy()
		ptrsForTheme.fnRenameColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642128662))
	}
	{
		methodName := StringNameFromStr("clear_color")
		defer methodName.Destroy()
		ptrsForTheme.fnClearColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("get_color_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetColorList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("get_color_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetColorTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_constant")
		defer methodName.Destroy()
		ptrsForTheme.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 281601298))
	}
	{
		methodName := StringNameFromStr("get_constant")
		defer methodName.Destroy()
		ptrsForTheme.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2419549490))
	}
	{
		methodName := StringNameFromStr("has_constant")
		defer methodName.Destroy()
		ptrsForTheme.fnHasConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("rename_constant")
		defer methodName.Destroy()
		ptrsForTheme.fnRenameConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642128662))
	}
	{
		methodName := StringNameFromStr("clear_constant")
		defer methodName.Destroy()
		ptrsForTheme.fnClearConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("get_constant_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetConstantList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("get_constant_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetConstantTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_default_base_scale")
		defer methodName.Destroy()
		ptrsForTheme.fnSetDefaultBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_default_base_scale")
		defer methodName.Destroy()
		ptrsForTheme.fnGetDefaultBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("has_default_base_scale")
		defer methodName.Destroy()
		ptrsForTheme.fnHasDefaultBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_font")
		defer methodName.Destroy()
		ptrsForTheme.fnSetDefaultFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262170328))
	}
	{
		methodName := StringNameFromStr("get_default_font")
		defer methodName.Destroy()
		ptrsForTheme.fnGetDefaultFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229501585))
	}
	{
		methodName := StringNameFromStr("has_default_font")
		defer methodName.Destroy()
		ptrsForTheme.fnHasDefaultFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnSetDefaultFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_default_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnGetDefaultFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("has_default_font_size")
		defer methodName.Destroy()
		ptrsForTheme.fnHasDefaultFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_theme_item")
		defer methodName.Destroy()
		ptrsForTheme.fnSetThemeItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2492983623))
	}
	{
		methodName := StringNameFromStr("get_theme_item")
		defer methodName.Destroy()
		ptrsForTheme.fnGetThemeItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2191024021))
	}
	{
		methodName := StringNameFromStr("has_theme_item")
		defer methodName.Destroy()
		ptrsForTheme.fnHasThemeItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1739311056))
	}
	{
		methodName := StringNameFromStr("rename_theme_item")
		defer methodName.Destroy()
		ptrsForTheme.fnRenameThemeItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3900867553))
	}
	{
		methodName := StringNameFromStr("clear_theme_item")
		defer methodName.Destroy()
		ptrsForTheme.fnClearThemeItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2965505587))
	}
	{
		methodName := StringNameFromStr("get_theme_item_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetThemeItemList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3726716710))
	}
	{
		methodName := StringNameFromStr("get_theme_item_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetThemeItemTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1316004935))
	}
	{
		methodName := StringNameFromStr("set_type_variation")
		defer methodName.Destroy()
		ptrsForTheme.fnSetTypeVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("is_type_variation")
		defer methodName.Destroy()
		ptrsForTheme.fnIsTypeVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("clear_type_variation")
		defer methodName.Destroy()
		ptrsForTheme.fnClearTypeVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_type_variation_base")
		defer methodName.Destroy()
		ptrsForTheme.fnGetTypeVariationBase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965194235))
	}
	{
		methodName := StringNameFromStr("get_type_variation_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetTypeVariationList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1761182771))
	}
	{
		methodName := StringNameFromStr("add_type")
		defer methodName.Destroy()
		ptrsForTheme.fnAddType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_type")
		defer methodName.Destroy()
		ptrsForTheme.fnRemoveType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_type_list")
		defer methodName.Destroy()
		ptrsForTheme.fnGetTypeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("merge_with")
		defer methodName.Destroy()
		ptrsForTheme.fnMergeWith = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2326690814))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForTheme.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type Theme struct {
	Resource
}

func (me *Theme) BaseClass() string {
	return "Theme"
}

func NewTheme() *Theme {
	str := StringNameFromStr("Theme") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Theme{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type ThemeDataType int

const (
	ThemeDataTypeDataTypeColor    ThemeDataType = 0
	ThemeDataTypeDataTypeConstant ThemeDataType = 1
	ThemeDataTypeDataTypeFont     ThemeDataType = 2
	ThemeDataTypeDataTypeFontSize ThemeDataType = 3
	ThemeDataTypeDataTypeIcon     ThemeDataType = 4
	ThemeDataTypeDataTypeStylebox ThemeDataType = 5
	ThemeDataTypeDataTypeMax      ThemeDataType = 6
)

func (me *Theme) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Theme) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Theme) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Theme) SetIcon(name StringName, theme_type StringName, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetIcon(name StringName, theme_type StringName) Texture2D {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) HasIcon(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) RenameIcon(old_name StringName, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRenameIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) ClearIcon(name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetIconList(theme_type String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetIconList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetIconTypeList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetIconTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) SetStylebox(name StringName, theme_type StringName, texture StyleBox) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetStylebox), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetStylebox(name StringName, theme_type StringName) StyleBox {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStyleBox()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) HasStylebox(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) RenameStylebox(old_name StringName, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRenameStylebox), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) ClearStylebox(name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearStylebox), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetStyleboxList(theme_type String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetStyleboxList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetStyleboxTypeList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetStyleboxTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) SetFont(name StringName, theme_type StringName, font Font) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetFont(name StringName, theme_type StringName) Font {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) HasFont(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) RenameFont(old_name StringName, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRenameFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) ClearFont(name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetFontList(theme_type String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetFontList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetFontTypeList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetFontTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) SetFontSize(name StringName, theme_type StringName, font_size int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetFontSize(name StringName, theme_type StringName) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) HasFontSize(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) RenameFontSize(old_name StringName, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRenameFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) ClearFontSize(name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetFontSizeList(theme_type String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetFontSizeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetFontSizeTypeList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetFontSizeTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) SetColor(name StringName, theme_type StringName, color Color) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetColor(name StringName, theme_type StringName) Color {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) HasColor(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) RenameColor(old_name StringName, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRenameColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) ClearColor(name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetColorList(theme_type String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetColorList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetColorTypeList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetColorTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) SetConstant(name StringName, theme_type StringName, constant int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr(), gdc.ConstTypePtr(&constant)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetConstant(name StringName, theme_type StringName) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) HasConstant(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) RenameConstant(old_name StringName, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRenameConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) ClearConstant(name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetConstantList(theme_type String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetConstantList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetConstantTypeList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetConstantTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) SetDefaultBaseScale(base_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&base_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetDefaultBaseScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetDefaultBaseScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetDefaultBaseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) HasDefaultBaseScale() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasDefaultBaseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) SetDefaultFont(font Font) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetDefaultFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetDefaultFont() Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetDefaultFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) HasDefaultFont() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasDefaultFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) SetDefaultFontSize(font_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetDefaultFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetDefaultFontSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetDefaultFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) HasDefaultFontSize() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasDefaultFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) SetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type), name.AsCTypePtr(), theme_type.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetThemeItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetThemeItem(data_type ThemeDataType, name StringName, theme_type StringName) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&data_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetThemeItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) HasThemeItem(data_type ThemeDataType, name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&data_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnHasThemeItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) RenameThemeItem(data_type ThemeDataType, old_name StringName, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type), old_name.AsCTypePtr(), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRenameThemeItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) ClearThemeItem(data_type ThemeDataType, name StringName, theme_type StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type), name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearThemeItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetThemeItemList(data_type ThemeDataType, theme_type String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&data_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetThemeItemList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetThemeItemTypeList(data_type ThemeDataType) PackedStringArray {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&data_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetThemeItemTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) SetTypeVariation(theme_type StringName, base_type StringName) {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), base_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnSetTypeVariation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) IsTypeVariation(theme_type StringName, base_type StringName) bool {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr(), base_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnIsTypeVariation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Theme) ClearTypeVariation(theme_type StringName) {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClearTypeVariation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetTypeVariationBase(theme_type StringName) StringName {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetTypeVariationBase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) GetTypeVariationList(base_type StringName) PackedStringArray {
	cargs := []gdc.ConstTypePtr{base_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetTypeVariationList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) AddType(theme_type StringName) {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnAddType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) RemoveType(theme_type StringName) {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnRemoveType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) GetTypeList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnGetTypeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Theme) MergeWith(other Theme) {
	cargs := []gdc.ConstTypePtr{other.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnMergeWith), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Theme) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTheme.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
