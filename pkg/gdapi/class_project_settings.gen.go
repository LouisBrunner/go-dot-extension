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

type ptrsForProjectSettingsList struct {
	fnHasSetting             gdc.MethodBindPtr
	fnSetSetting             gdc.MethodBindPtr
	fnGetSetting             gdc.MethodBindPtr
	fnGetSettingWithOverride gdc.MethodBindPtr
	fnGetGlobalClassList     gdc.MethodBindPtr
	fnSetOrder               gdc.MethodBindPtr
	fnGetOrder               gdc.MethodBindPtr
	fnSetInitialValue        gdc.MethodBindPtr
	fnSetAsBasic             gdc.MethodBindPtr
	fnSetAsInternal          gdc.MethodBindPtr
	fnAddPropertyInfo        gdc.MethodBindPtr
	fnSetRestartIfChanged    gdc.MethodBindPtr
	fnClear                  gdc.MethodBindPtr
	fnLocalizePath           gdc.MethodBindPtr
	fnGlobalizePath          gdc.MethodBindPtr
	fnSave                   gdc.MethodBindPtr
	fnLoadResourcePack       gdc.MethodBindPtr
	fnSaveCustom             gdc.MethodBindPtr
}

var ptrsForProjectSettings ptrsForProjectSettingsList

func initProjectSettingsPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ProjectSettings")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_setting")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnHasSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("set_setting")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSetSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 402577236))
	}
	{
		methodName := StringNameFromStr("get_setting")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnGetSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 223050753))
	}
	{
		methodName := StringNameFromStr("get_setting_with_override")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnGetSettingWithOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
	{
		methodName := StringNameFromStr("get_global_class_list")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnGetGlobalClassList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("set_order")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSetOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
	}
	{
		methodName := StringNameFromStr("get_order")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnGetOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("set_initial_value")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSetInitialValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 402577236))
	}
	{
		methodName := StringNameFromStr("set_as_basic")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSetAsBasic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
	}
	{
		methodName := StringNameFromStr("set_as_internal")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSetAsInternal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
	}
	{
		methodName := StringNameFromStr("add_property_info")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnAddPropertyInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
	}
	{
		methodName := StringNameFromStr("set_restart_if_changed")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSetRestartIfChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("localize_path")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnLocalizePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
	}
	{
		methodName := StringNameFromStr("globalize_path")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnGlobalizePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
	}
	{
		methodName := StringNameFromStr("save")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("load_resource_pack")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnLoadResourcePack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 708980503))
	}
	{
		methodName := StringNameFromStr("save_custom")
		defer methodName.Destroy()
		ptrsForProjectSettings.fnSaveCustom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}

}

type ProjectSettings struct {
	Object
}

func (me *ProjectSettings) BaseClass() string {
	return "ProjectSettings"
}

func NewProjectSettings() *ProjectSettings {
	str := StringNameFromStr("ProjectSettings") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ProjectSettings{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ProjectSettings) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ProjectSettings) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ProjectSettings) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ProjectSettings) HasSetting(name String) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnHasSetting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ProjectSettings) SetSetting(name String, value Variant) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSetSetting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) GetSetting(name String, default_value Variant) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), default_value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnGetSetting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ProjectSettings) GetSettingWithOverride(name StringName) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnGetSettingWithOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ProjectSettings) GetGlobalClassList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnGetGlobalClassList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *ProjectSettings) SetOrder(name String, position int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSetOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) GetOrder(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnGetOrder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ProjectSettings) SetInitialValue(name String, value Variant) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSetInitialValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) SetAsBasic(name String, basic bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&basic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSetAsBasic), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) SetAsInternal(name String, internal bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&internal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSetAsInternal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) AddPropertyInfo(hint Dictionary) {
	cargs := []gdc.ConstTypePtr{hint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnAddPropertyInfo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) SetRestartIfChanged(name String, restart bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&restart)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSetRestartIfChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) Clear(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ProjectSettings) LocalizePath(path String) String {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnLocalizePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ProjectSettings) GlobalizePath(path String) String {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnGlobalizePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ProjectSettings) Save() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSave), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ProjectSettings) LoadResourcePack(pack String, replace_files bool, offset int64) bool {
	cargs := []gdc.ConstTypePtr{pack.AsCTypePtr(), gdc.ConstTypePtr(&replace_files), gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&replace_files)
	pinner.Pin(&offset)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnLoadResourcePack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ProjectSettings) SaveCustom(file String) Error {
	cargs := []gdc.ConstTypePtr{file.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProjectSettings.fnSaveCustom), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals

type ProjectSettingsSettingsChangedSignalFn func()

func (me *ProjectSettings) ConnectSettingsChanged(subs SignalSubscribers, fn ProjectSettingsSettingsChangedSignalFn) {
	sig := StringNameFromStr("settings_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *ProjectSettings) DisconnectSettingsChanged(subs SignalSubscribers, fn ProjectSettingsSettingsChangedSignalFn) {
	sig := StringNameFromStr("settings_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
