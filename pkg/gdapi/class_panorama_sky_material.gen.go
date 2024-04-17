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

type ptrsForPanoramaSkyMaterialList struct {
	fnSetPanorama         gdc.MethodBindPtr
	fnGetPanorama         gdc.MethodBindPtr
	fnSetFilteringEnabled gdc.MethodBindPtr
	fnIsFilteringEnabled  gdc.MethodBindPtr
	fnSetEnergyMultiplier gdc.MethodBindPtr
	fnGetEnergyMultiplier gdc.MethodBindPtr
}

var ptrsForPanoramaSkyMaterial ptrsForPanoramaSkyMaterialList

func initPanoramaSkyMaterialPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PanoramaSkyMaterial")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_panorama")
		defer methodName.Destroy()
		ptrsForPanoramaSkyMaterial.fnSetPanorama = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_panorama")
		defer methodName.Destroy()
		ptrsForPanoramaSkyMaterial.fnGetPanorama = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_filtering_enabled")
		defer methodName.Destroy()
		ptrsForPanoramaSkyMaterial.fnSetFilteringEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_filtering_enabled")
		defer methodName.Destroy()
		ptrsForPanoramaSkyMaterial.fnIsFilteringEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_energy_multiplier")
		defer methodName.Destroy()
		ptrsForPanoramaSkyMaterial.fnSetEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_energy_multiplier")
		defer methodName.Destroy()
		ptrsForPanoramaSkyMaterial.fnGetEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type PanoramaSkyMaterial struct {
	Material
}

func (me *PanoramaSkyMaterial) BaseClass() string {
	return "PanoramaSkyMaterial"
}

func NewPanoramaSkyMaterial() *PanoramaSkyMaterial {
	str := StringNameFromStr("PanoramaSkyMaterial") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PanoramaSkyMaterial{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PanoramaSkyMaterial) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PanoramaSkyMaterial) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PanoramaSkyMaterial) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PanoramaSkyMaterial) SetPanorama(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPanoramaSkyMaterial.fnSetPanorama), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PanoramaSkyMaterial) GetPanorama() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPanoramaSkyMaterial.fnGetPanorama), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PanoramaSkyMaterial) SetFilteringEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPanoramaSkyMaterial.fnSetFilteringEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PanoramaSkyMaterial) IsFilteringEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPanoramaSkyMaterial.fnIsFilteringEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PanoramaSkyMaterial) SetEnergyMultiplier(multiplier float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPanoramaSkyMaterial.fnSetEnergyMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PanoramaSkyMaterial) GetEnergyMultiplier() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPanoramaSkyMaterial.fnGetEnergyMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
