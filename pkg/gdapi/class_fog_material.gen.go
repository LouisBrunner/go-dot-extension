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

type ptrsForFogMaterialList struct {
	fnSetDensity        gdc.MethodBindPtr
	fnGetDensity        gdc.MethodBindPtr
	fnSetAlbedo         gdc.MethodBindPtr
	fnGetAlbedo         gdc.MethodBindPtr
	fnSetEmission       gdc.MethodBindPtr
	fnGetEmission       gdc.MethodBindPtr
	fnSetHeightFalloff  gdc.MethodBindPtr
	fnGetHeightFalloff  gdc.MethodBindPtr
	fnSetEdgeFade       gdc.MethodBindPtr
	fnGetEdgeFade       gdc.MethodBindPtr
	fnSetDensityTexture gdc.MethodBindPtr
	fnGetDensityTexture gdc.MethodBindPtr
}

var ptrsForFogMaterial ptrsForFogMaterialList

func initFogMaterialPtrs(iface gdc.Interface) {

	className := StringNameFromStr("FogMaterial")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_density")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnSetDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_density")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnGetDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_albedo")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnSetAlbedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_albedo")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnGetAlbedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_emission")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnSetEmission = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_emission")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnGetEmission = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_height_falloff")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnSetHeightFalloff = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_height_falloff")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnGetHeightFalloff = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_edge_fade")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnSetEdgeFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_edge_fade")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnGetEdgeFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_density_texture")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnSetDensityTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1188404210))
	}
	{
		methodName := StringNameFromStr("get_density_texture")
		defer methodName.Destroy()
		ptrsForFogMaterial.fnGetDensityTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373985333))
	}

}

type FogMaterial struct {
	Material
}

func (me *FogMaterial) BaseClass() string {
	return "FogMaterial"
}

func NewFogMaterial() *FogMaterial {
	str := StringNameFromStr("FogMaterial") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &FogMaterial{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *FogMaterial) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *FogMaterial) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *FogMaterial) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *FogMaterial) SetDensity(density float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnSetDensity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FogMaterial) GetDensity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnGetDensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FogMaterial) SetAlbedo(albedo Color) {
	cargs := []gdc.ConstTypePtr{albedo.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnSetAlbedo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FogMaterial) GetAlbedo() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnGetAlbedo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FogMaterial) SetEmission(emission Color) {
	cargs := []gdc.ConstTypePtr{emission.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnSetEmission), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FogMaterial) GetEmission() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnGetEmission), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FogMaterial) SetHeightFalloff(height_falloff float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height_falloff)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnSetHeightFalloff), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FogMaterial) GetHeightFalloff() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnGetHeightFalloff), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FogMaterial) SetEdgeFade(edge_fade float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge_fade)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnSetEdgeFade), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FogMaterial) GetEdgeFade() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnGetEdgeFade), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FogMaterial) SetDensityTexture(density_texture Texture3D) {
	cargs := []gdc.ConstTypePtr{density_texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnSetDensityTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FogMaterial) GetDensityTexture() Texture3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogMaterial.fnGetDensityTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
