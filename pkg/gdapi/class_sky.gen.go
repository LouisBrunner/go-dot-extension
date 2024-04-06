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

type ptrsForSkyList struct {
	fnSetRadianceSize gdc.MethodBindPtr
	fnGetRadianceSize gdc.MethodBindPtr
	fnSetProcessMode  gdc.MethodBindPtr
	fnGetProcessMode  gdc.MethodBindPtr
	fnSetMaterial     gdc.MethodBindPtr
	fnGetMaterial     gdc.MethodBindPtr
}

var ptrsForSky ptrsForSkyList

func initSkyPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Sky")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radiance_size")
		defer methodName.Destroy()
		ptrsForSky.fnSetRadianceSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1512957179))
	}
	{
		methodName := StringNameFromStr("get_radiance_size")
		defer methodName.Destroy()
		ptrsForSky.fnGetRadianceSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2708733976))
	}
	{
		methodName := StringNameFromStr("set_process_mode")
		defer methodName.Destroy()
		ptrsForSky.fnSetProcessMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 875986769))
	}
	{
		methodName := StringNameFromStr("get_process_mode")
		defer methodName.Destroy()
		ptrsForSky.fnGetProcessMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 731245043))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForSky.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForSky.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
}

type Sky struct {
	Resource
}

func (me *Sky) BaseClass() string {
	return "Sky"
}

func NewSky() *Sky {
	str := StringNameFromStr("Sky") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Sky{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type SkyRadianceSize int

const (
	SkyRadianceSizeRadianceSize32   SkyRadianceSize = 0
	SkyRadianceSizeRadianceSize64   SkyRadianceSize = 1
	SkyRadianceSizeRadianceSize128  SkyRadianceSize = 2
	SkyRadianceSizeRadianceSize256  SkyRadianceSize = 3
	SkyRadianceSizeRadianceSize512  SkyRadianceSize = 4
	SkyRadianceSizeRadianceSize1024 SkyRadianceSize = 5
	SkyRadianceSizeRadianceSize2048 SkyRadianceSize = 6
	SkyRadianceSizeRadianceSizeMax  SkyRadianceSize = 7
)

type SkyProcessMode int

const (
	SkyProcessModeProcessModeAutomatic   SkyProcessMode = 0
	SkyProcessModeProcessModeQuality     SkyProcessMode = 1
	SkyProcessModeProcessModeIncremental SkyProcessMode = 2
	SkyProcessModeProcessModeRealtime    SkyProcessMode = 3
)

func (me *Sky) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Sky) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Sky) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Sky) SetRadianceSize(size SkyRadianceSize) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSky.fnSetRadianceSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Sky) GetRadianceSize() SkyRadianceSize {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret SkyRadianceSize

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSky.fnGetRadianceSize), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Sky) SetProcessMode(mode SkyProcessMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSky.fnSetProcessMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Sky) GetProcessMode() SkyProcessMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret SkyProcessMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSky.fnGetProcessMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Sky) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSky.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Sky) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSky.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
