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

type ptrsForEditorPathsList struct {
	fnGetDataDir            gdc.MethodBindPtr
	fnGetConfigDir          gdc.MethodBindPtr
	fnGetCacheDir           gdc.MethodBindPtr
	fnIsSelfContained       gdc.MethodBindPtr
	fnGetSelfContainedFile  gdc.MethodBindPtr
	fnGetProjectSettingsDir gdc.MethodBindPtr
}

var ptrsForEditorPaths ptrsForEditorPathsList

func initEditorPathsPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorPaths")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_data_dir")
		defer methodName.Destroy()
		ptrsForEditorPaths.fnGetDataDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_config_dir")
		defer methodName.Destroy()
		ptrsForEditorPaths.fnGetConfigDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_cache_dir")
		defer methodName.Destroy()
		ptrsForEditorPaths.fnGetCacheDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_self_contained")
		defer methodName.Destroy()
		ptrsForEditorPaths.fnIsSelfContained = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_self_contained_file")
		defer methodName.Destroy()
		ptrsForEditorPaths.fnGetSelfContainedFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_project_settings_dir")
		defer methodName.Destroy()
		ptrsForEditorPaths.fnGetProjectSettingsDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
}

type EditorPaths struct {
	Object
}

func (me *EditorPaths) BaseClass() string {
	return "EditorPaths"
}

func NewEditorPaths() *EditorPaths {
	str := StringNameFromStr("EditorPaths") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorPaths{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorPaths) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorPaths) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorPaths) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorPaths) GetDataDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPaths.fnGetDataDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPaths) GetConfigDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPaths.fnGetConfigDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPaths) GetCacheDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPaths.fnGetCacheDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPaths) IsSelfContained() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPaths.fnIsSelfContained), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorPaths) GetSelfContainedFile() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPaths.fnGetSelfContainedFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPaths) GetProjectSettingsDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPaths.fnGetProjectSettingsDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
