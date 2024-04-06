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

type ptrsForPhysicalSkyMaterialList struct {
	fnSetRayleighCoefficient gdc.MethodBindPtr
	fnGetRayleighCoefficient gdc.MethodBindPtr
	fnSetRayleighColor       gdc.MethodBindPtr
	fnGetRayleighColor       gdc.MethodBindPtr
	fnSetMieCoefficient      gdc.MethodBindPtr
	fnGetMieCoefficient      gdc.MethodBindPtr
	fnSetMieEccentricity     gdc.MethodBindPtr
	fnGetMieEccentricity     gdc.MethodBindPtr
	fnSetMieColor            gdc.MethodBindPtr
	fnGetMieColor            gdc.MethodBindPtr
	fnSetTurbidity           gdc.MethodBindPtr
	fnGetTurbidity           gdc.MethodBindPtr
	fnSetSunDiskScale        gdc.MethodBindPtr
	fnGetSunDiskScale        gdc.MethodBindPtr
	fnSetGroundColor         gdc.MethodBindPtr
	fnGetGroundColor         gdc.MethodBindPtr
	fnSetEnergyMultiplier    gdc.MethodBindPtr
	fnGetEnergyMultiplier    gdc.MethodBindPtr
	fnSetUseDebanding        gdc.MethodBindPtr
	fnGetUseDebanding        gdc.MethodBindPtr
	fnSetNightSky            gdc.MethodBindPtr
	fnGetNightSky            gdc.MethodBindPtr
}

var ptrsForPhysicalSkyMaterial ptrsForPhysicalSkyMaterialList

func initPhysicalSkyMaterialPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicalSkyMaterial")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_rayleigh_coefficient")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetRayleighCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_rayleigh_coefficient")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetRayleighCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_rayleigh_color")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetRayleighColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_rayleigh_color")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetRayleighColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_mie_coefficient")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetMieCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mie_coefficient")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetMieCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_mie_eccentricity")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetMieEccentricity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mie_eccentricity")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetMieEccentricity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_mie_color")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetMieColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_mie_color")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetMieColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_turbidity")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetTurbidity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_turbidity")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetTurbidity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sun_disk_scale")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetSunDiskScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_sun_disk_scale")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetSunDiskScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_ground_color")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetGroundColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_ground_color")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetGroundColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_energy_multiplier")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_energy_multiplier")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_use_debanding")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetUseDebanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_debanding")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetUseDebanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_night_sky")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnSetNightSky = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_night_sky")
		defer methodName.Destroy()
		ptrsForPhysicalSkyMaterial.fnGetNightSky = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
}

type PhysicalSkyMaterial struct {
	Material
}

func (me *PhysicalSkyMaterial) BaseClass() string {
	return "PhysicalSkyMaterial"
}

func NewPhysicalSkyMaterial() *PhysicalSkyMaterial {
	str := StringNameFromStr("PhysicalSkyMaterial") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicalSkyMaterial{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicalSkyMaterial) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicalSkyMaterial) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicalSkyMaterial) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicalSkyMaterial) SetRayleighCoefficient(rayleigh float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rayleigh)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetRayleighCoefficient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetRayleighCoefficient() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetRayleighCoefficient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalSkyMaterial) SetRayleighColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetRayleighColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetRayleighColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetRayleighColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalSkyMaterial) SetMieCoefficient(mie float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mie)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetMieCoefficient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetMieCoefficient() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetMieCoefficient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalSkyMaterial) SetMieEccentricity(eccentricity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&eccentricity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetMieEccentricity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetMieEccentricity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetMieEccentricity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalSkyMaterial) SetMieColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetMieColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetMieColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetMieColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalSkyMaterial) SetTurbidity(turbidity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbidity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetTurbidity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetTurbidity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetTurbidity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalSkyMaterial) SetSunDiskScale(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetSunDiskScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetSunDiskScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetSunDiskScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalSkyMaterial) SetGroundColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetGroundColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetGroundColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetGroundColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalSkyMaterial) SetEnergyMultiplier(multiplier float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetEnergyMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetEnergyMultiplier() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetEnergyMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalSkyMaterial) SetUseDebanding(use_debanding bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_debanding)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetUseDebanding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetUseDebanding() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetUseDebanding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalSkyMaterial) SetNightSky(night_sky Texture2D) {
	cargs := []gdc.ConstTypePtr{night_sky.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnSetNightSky), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalSkyMaterial) GetNightSky() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalSkyMaterial.fnGetNightSky), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
