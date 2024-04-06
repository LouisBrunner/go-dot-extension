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

type ptrsForEditorFeatureProfileList struct {
	fnSetDisableClass         gdc.MethodBindPtr
	fnIsClassDisabled         gdc.MethodBindPtr
	fnSetDisableClassEditor   gdc.MethodBindPtr
	fnIsClassEditorDisabled   gdc.MethodBindPtr
	fnSetDisableClassProperty gdc.MethodBindPtr
	fnIsClassPropertyDisabled gdc.MethodBindPtr
	fnSetDisableFeature       gdc.MethodBindPtr
	fnIsFeatureDisabled       gdc.MethodBindPtr
	fnGetFeatureName          gdc.MethodBindPtr
	fnSaveToFile              gdc.MethodBindPtr
	fnLoadFromFile            gdc.MethodBindPtr
}

var ptrsForEditorFeatureProfile ptrsForEditorFeatureProfileList

func initEditorFeatureProfilePtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorFeatureProfile")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_disable_class")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnSetDisableClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2524380260))
	}
	{
		methodName := StringNameFromStr("is_class_disabled")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnIsClassDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("set_disable_class_editor")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnSetDisableClassEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2524380260))
	}
	{
		methodName := StringNameFromStr("is_class_editor_disabled")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnIsClassEditorDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("set_disable_class_property")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnSetDisableClassProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 865197084))
	}
	{
		methodName := StringNameFromStr("is_class_property_disabled")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnIsClassPropertyDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("set_disable_feature")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnSetDisableFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1884871044))
	}
	{
		methodName := StringNameFromStr("is_feature_disabled")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnIsFeatureDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2974403161))
	}
	{
		methodName := StringNameFromStr("get_feature_name")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnGetFeatureName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3401335809))
	}
	{
		methodName := StringNameFromStr("save_to_file")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnSaveToFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("load_from_file")
		defer methodName.Destroy()
		ptrsForEditorFeatureProfile.fnLoadFromFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
}

type EditorFeatureProfile struct {
	RefCounted
}

func (me *EditorFeatureProfile) BaseClass() string {
	return "EditorFeatureProfile"
}

func NewEditorFeatureProfile() *EditorFeatureProfile {
	str := StringNameFromStr("EditorFeatureProfile") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorFeatureProfile{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type EditorFeatureProfileFeature int

const (
	EditorFeatureProfileFeatureFeature3D             EditorFeatureProfileFeature = 0
	EditorFeatureProfileFeatureFeatureScript         EditorFeatureProfileFeature = 1
	EditorFeatureProfileFeatureFeatureAssetLib       EditorFeatureProfileFeature = 2
	EditorFeatureProfileFeatureFeatureSceneTree      EditorFeatureProfileFeature = 3
	EditorFeatureProfileFeatureFeatureNodeDock       EditorFeatureProfileFeature = 4
	EditorFeatureProfileFeatureFeatureFilesystemDock EditorFeatureProfileFeature = 5
	EditorFeatureProfileFeatureFeatureImportDock     EditorFeatureProfileFeature = 6
	EditorFeatureProfileFeatureFeatureHistoryDock    EditorFeatureProfileFeature = 7
	EditorFeatureProfileFeatureFeatureMax            EditorFeatureProfileFeature = 8
)

func (me *EditorFeatureProfile) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorFeatureProfile) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorFeatureProfile) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorFeatureProfile) SetDisableClass(class_name StringName, disable bool) {
	cargs := []gdc.ConstTypePtr{class_name.AsCTypePtr(), gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnSetDisableClass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFeatureProfile) IsClassDisabled(class_name StringName) bool {
	cargs := []gdc.ConstTypePtr{class_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnIsClassDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFeatureProfile) SetDisableClassEditor(class_name StringName, disable bool) {
	cargs := []gdc.ConstTypePtr{class_name.AsCTypePtr(), gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnSetDisableClassEditor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFeatureProfile) IsClassEditorDisabled(class_name StringName) bool {
	cargs := []gdc.ConstTypePtr{class_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnIsClassEditorDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFeatureProfile) SetDisableClassProperty(class_name StringName, property StringName, disable bool) {
	cargs := []gdc.ConstTypePtr{class_name.AsCTypePtr(), property.AsCTypePtr(), gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnSetDisableClassProperty), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFeatureProfile) IsClassPropertyDisabled(class_name StringName, property StringName) bool {
	cargs := []gdc.ConstTypePtr{class_name.AsCTypePtr(), property.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnIsClassPropertyDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFeatureProfile) SetDisableFeature(feature EditorFeatureProfileFeature, disable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature), gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnSetDisableFeature), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFeatureProfile) IsFeatureDisabled(feature EditorFeatureProfileFeature) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&feature)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnIsFeatureDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFeatureProfile) GetFeatureName(feature EditorFeatureProfileFeature) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&feature)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnGetFeatureName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFeatureProfile) SaveToFile(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnSaveToFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *EditorFeatureProfile) LoadFromFile(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFeatureProfile.fnLoadFromFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
