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

type ptrsForEditorSettingsList struct {
	fnHasSetting                  gdc.MethodBindPtr
	fnSetSetting                  gdc.MethodBindPtr
	fnGetSetting                  gdc.MethodBindPtr
	fnErase                       gdc.MethodBindPtr
	fnSetInitialValue             gdc.MethodBindPtr
	fnAddPropertyInfo             gdc.MethodBindPtr
	fnSetProjectMetadata          gdc.MethodBindPtr
	fnGetProjectMetadata          gdc.MethodBindPtr
	fnSetFavorites                gdc.MethodBindPtr
	fnGetFavorites                gdc.MethodBindPtr
	fnSetRecentDirs               gdc.MethodBindPtr
	fnGetRecentDirs               gdc.MethodBindPtr
	fnSetBuiltinActionOverride    gdc.MethodBindPtr
	fnCheckChangedSettingsInGroup gdc.MethodBindPtr
	fnGetChangedSettings          gdc.MethodBindPtr
	fnMarkSettingChanged          gdc.MethodBindPtr
}

var ptrsForEditorSettings ptrsForEditorSettingsList

func initEditorSettingsPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorSettings")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_setting")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnHasSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("set_setting")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnSetSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 402577236))
	}
	{
		methodName := StringNameFromStr("get_setting")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnGetSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1868160156))
	}
	{
		methodName := StringNameFromStr("erase")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnErase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_initial_value")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnSetInitialValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1529169264))
	}
	{
		methodName := StringNameFromStr("add_property_info")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnAddPropertyInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
	}
	{
		methodName := StringNameFromStr("set_project_metadata")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnSetProjectMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2504492430))
	}
	{
		methodName := StringNameFromStr("get_project_metadata")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnGetProjectMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 89809366))
	}
	{
		methodName := StringNameFromStr("set_favorites")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnSetFavorites = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
	}
	{
		methodName := StringNameFromStr("get_favorites")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnGetFavorites = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_recent_dirs")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnSetRecentDirs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
	}
	{
		methodName := StringNameFromStr("get_recent_dirs")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnGetRecentDirs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_builtin_action_override")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnSetBuiltinActionOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1209351045))
	}
	{
		methodName := StringNameFromStr("check_changed_settings_in_group")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnCheckChangedSettingsInGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("get_changed_settings")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnGetChangedSettings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("mark_setting_changed")
		defer methodName.Destroy()
		ptrsForEditorSettings.fnMarkSettingChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}

}

type EditorSettings struct {
	Resource
}

func (me *EditorSettings) BaseClass() string {
	return "EditorSettings"
}

func NewEditorSettings() *EditorSettings {
	str := StringNameFromStr("EditorSettings") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorSettings{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	EditorSettingsNotificationEditorSettingsChanged = 10000
)

// Enums

func (me *EditorSettings) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorSettings) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorSettings) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorSettings) HasSetting(name String) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnHasSetting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorSettings) SetSetting(name String, value Variant) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnSetSetting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) GetSetting(name String) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnGetSetting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorSettings) Erase(property String) {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnErase), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) SetInitialValue(name StringName, value Variant, update_current bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), gdc.ConstTypePtr(&update_current)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnSetInitialValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) AddPropertyInfo(info Dictionary) {
	cargs := []gdc.ConstTypePtr{info.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnAddPropertyInfo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) SetProjectMetadata(section String, key String, data Variant) {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr(), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnSetProjectMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) GetProjectMetadata(section String, key String, default_ Variant) Variant {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr(), default_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnGetProjectMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorSettings) SetFavorites(dirs PackedStringArray) {
	cargs := []gdc.ConstTypePtr{dirs.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnSetFavorites), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) GetFavorites() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnGetFavorites), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorSettings) SetRecentDirs(dirs PackedStringArray) {
	cargs := []gdc.ConstTypePtr{dirs.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnSetRecentDirs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) GetRecentDirs() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnGetRecentDirs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorSettings) SetBuiltinActionOverride(name String, actions_list []InputEvent) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&actions_list)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnSetBuiltinActionOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSettings) CheckChangedSettingsInGroup(setting_prefix String) bool {
	cargs := []gdc.ConstTypePtr{setting_prefix.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnCheckChangedSettingsInGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorSettings) GetChangedSettings() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnGetChangedSettings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorSettings) MarkSettingChanged(setting String) {
	cargs := []gdc.ConstTypePtr{setting.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSettings.fnMarkSettingChanged), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type EditorSettingsSettingsChangedSignalFn func()

func (me *EditorSettings) ConnectSettingsChanged(subs SignalSubscribers, fn EditorSettingsSettingsChangedSignalFn) {
	sig := StringNameFromStr("settings_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorSettings) DisconnectSettingsChanged(subs SignalSubscribers, fn EditorSettingsSettingsChangedSignalFn) {
	sig := StringNameFromStr("settings_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
