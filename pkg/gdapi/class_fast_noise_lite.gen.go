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

type ptrsForFastNoiseLiteList struct {
	fnSetNoiseType                   gdc.MethodBindPtr
	fnGetNoiseType                   gdc.MethodBindPtr
	fnSetSeed                        gdc.MethodBindPtr
	fnGetSeed                        gdc.MethodBindPtr
	fnSetFrequency                   gdc.MethodBindPtr
	fnGetFrequency                   gdc.MethodBindPtr
	fnSetOffset                      gdc.MethodBindPtr
	fnGetOffset                      gdc.MethodBindPtr
	fnSetFractalType                 gdc.MethodBindPtr
	fnGetFractalType                 gdc.MethodBindPtr
	fnSetFractalOctaves              gdc.MethodBindPtr
	fnGetFractalOctaves              gdc.MethodBindPtr
	fnSetFractalLacunarity           gdc.MethodBindPtr
	fnGetFractalLacunarity           gdc.MethodBindPtr
	fnSetFractalGain                 gdc.MethodBindPtr
	fnGetFractalGain                 gdc.MethodBindPtr
	fnSetFractalWeightedStrength     gdc.MethodBindPtr
	fnGetFractalWeightedStrength     gdc.MethodBindPtr
	fnSetFractalPingPongStrength     gdc.MethodBindPtr
	fnGetFractalPingPongStrength     gdc.MethodBindPtr
	fnSetCellularDistanceFunction    gdc.MethodBindPtr
	fnGetCellularDistanceFunction    gdc.MethodBindPtr
	fnSetCellularJitter              gdc.MethodBindPtr
	fnGetCellularJitter              gdc.MethodBindPtr
	fnSetCellularReturnType          gdc.MethodBindPtr
	fnGetCellularReturnType          gdc.MethodBindPtr
	fnSetDomainWarpEnabled           gdc.MethodBindPtr
	fnIsDomainWarpEnabled            gdc.MethodBindPtr
	fnSetDomainWarpType              gdc.MethodBindPtr
	fnGetDomainWarpType              gdc.MethodBindPtr
	fnSetDomainWarpAmplitude         gdc.MethodBindPtr
	fnGetDomainWarpAmplitude         gdc.MethodBindPtr
	fnSetDomainWarpFrequency         gdc.MethodBindPtr
	fnGetDomainWarpFrequency         gdc.MethodBindPtr
	fnSetDomainWarpFractalType       gdc.MethodBindPtr
	fnGetDomainWarpFractalType       gdc.MethodBindPtr
	fnSetDomainWarpFractalOctaves    gdc.MethodBindPtr
	fnGetDomainWarpFractalOctaves    gdc.MethodBindPtr
	fnSetDomainWarpFractalLacunarity gdc.MethodBindPtr
	fnGetDomainWarpFractalLacunarity gdc.MethodBindPtr
	fnSetDomainWarpFractalGain       gdc.MethodBindPtr
	fnGetDomainWarpFractalGain       gdc.MethodBindPtr
}

var ptrsForFastNoiseLite ptrsForFastNoiseLiteList

func initFastNoiseLitePtrs(iface gdc.Interface) {

	className := StringNameFromStr("FastNoiseLite")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_noise_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetNoiseType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2624461392))
	}
	{
		methodName := StringNameFromStr("get_noise_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetNoiseType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1458108610))
	}
	{
		methodName := StringNameFromStr("set_seed")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetSeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_seed")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetSeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_frequency")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_frequency")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_fractal_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetFractalType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4132731174))
	}
	{
		methodName := StringNameFromStr("get_fractal_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetFractalType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1036889279))
	}
	{
		methodName := StringNameFromStr("set_fractal_octaves")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetFractalOctaves = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_fractal_octaves")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetFractalOctaves = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_fractal_lacunarity")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetFractalLacunarity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fractal_lacunarity")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetFractalLacunarity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fractal_gain")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetFractalGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fractal_gain")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetFractalGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fractal_weighted_strength")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetFractalWeightedStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fractal_weighted_strength")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetFractalWeightedStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fractal_ping_pong_strength")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetFractalPingPongStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fractal_ping_pong_strength")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetFractalPingPongStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_cellular_distance_function")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetCellularDistanceFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1006013267))
	}
	{
		methodName := StringNameFromStr("get_cellular_distance_function")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetCellularDistanceFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2021274088))
	}
	{
		methodName := StringNameFromStr("set_cellular_jitter")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetCellularJitter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_cellular_jitter")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetCellularJitter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_cellular_return_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetCellularReturnType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2654169698))
	}
	{
		methodName := StringNameFromStr("get_cellular_return_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetCellularReturnType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3699796343))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_enabled")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_domain_warp_enabled")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnIsDomainWarpEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3629692980))
	}
	{
		methodName := StringNameFromStr("get_domain_warp_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetDomainWarpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2980162020))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_amplitude")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpAmplitude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_domain_warp_amplitude")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetDomainWarpAmplitude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_frequency")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_domain_warp_frequency")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetDomainWarpFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_fractal_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpFractalType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3999408287))
	}
	{
		methodName := StringNameFromStr("get_domain_warp_fractal_type")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetDomainWarpFractalType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 407716934))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_fractal_octaves")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpFractalOctaves = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_domain_warp_fractal_octaves")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetDomainWarpFractalOctaves = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_fractal_lacunarity")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpFractalLacunarity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_domain_warp_fractal_lacunarity")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetDomainWarpFractalLacunarity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_domain_warp_fractal_gain")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnSetDomainWarpFractalGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_domain_warp_fractal_gain")
		defer methodName.Destroy()
		ptrsForFastNoiseLite.fnGetDomainWarpFractalGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type FastNoiseLite struct {
	Noise
}

func (me *FastNoiseLite) BaseClass() string {
	return "FastNoiseLite"
}

func NewFastNoiseLite() *FastNoiseLite {
	str := StringNameFromStr("FastNoiseLite") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &FastNoiseLite{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type FastNoiseLiteNoiseType int

const (
	FastNoiseLiteNoiseTypeTypeValue         FastNoiseLiteNoiseType = 5
	FastNoiseLiteNoiseTypeTypeValueCubic    FastNoiseLiteNoiseType = 4
	FastNoiseLiteNoiseTypeTypePerlin        FastNoiseLiteNoiseType = 3
	FastNoiseLiteNoiseTypeTypeCellular      FastNoiseLiteNoiseType = 2
	FastNoiseLiteNoiseTypeTypeSimplex       FastNoiseLiteNoiseType = 0
	FastNoiseLiteNoiseTypeTypeSimplexSmooth FastNoiseLiteNoiseType = 1
)

type FastNoiseLiteFractalType int

const (
	FastNoiseLiteFractalTypeFractalNone     FastNoiseLiteFractalType = 0
	FastNoiseLiteFractalTypeFractalFbm      FastNoiseLiteFractalType = 1
	FastNoiseLiteFractalTypeFractalRidged   FastNoiseLiteFractalType = 2
	FastNoiseLiteFractalTypeFractalPingPong FastNoiseLiteFractalType = 3
)

type FastNoiseLiteCellularDistanceFunction int

const (
	FastNoiseLiteCellularDistanceFunctionDistanceEuclidean        FastNoiseLiteCellularDistanceFunction = 0
	FastNoiseLiteCellularDistanceFunctionDistanceEuclideanSquared FastNoiseLiteCellularDistanceFunction = 1
	FastNoiseLiteCellularDistanceFunctionDistanceManhattan        FastNoiseLiteCellularDistanceFunction = 2
	FastNoiseLiteCellularDistanceFunctionDistanceHybrid           FastNoiseLiteCellularDistanceFunction = 3
)

type FastNoiseLiteCellularReturnType int

const (
	FastNoiseLiteCellularReturnTypeReturnCellValue    FastNoiseLiteCellularReturnType = 0
	FastNoiseLiteCellularReturnTypeReturnDistance     FastNoiseLiteCellularReturnType = 1
	FastNoiseLiteCellularReturnTypeReturnDistance2    FastNoiseLiteCellularReturnType = 2
	FastNoiseLiteCellularReturnTypeReturnDistance2Add FastNoiseLiteCellularReturnType = 3
	FastNoiseLiteCellularReturnTypeReturnDistance2Sub FastNoiseLiteCellularReturnType = 4
	FastNoiseLiteCellularReturnTypeReturnDistance2Mul FastNoiseLiteCellularReturnType = 5
	FastNoiseLiteCellularReturnTypeReturnDistance2Div FastNoiseLiteCellularReturnType = 6
)

type FastNoiseLiteDomainWarpType int

const (
	FastNoiseLiteDomainWarpTypeDomainWarpSimplex        FastNoiseLiteDomainWarpType = 0
	FastNoiseLiteDomainWarpTypeDomainWarpSimplexReduced FastNoiseLiteDomainWarpType = 1
	FastNoiseLiteDomainWarpTypeDomainWarpBasicGrid      FastNoiseLiteDomainWarpType = 2
)

type FastNoiseLiteDomainWarpFractalType int

const (
	FastNoiseLiteDomainWarpFractalTypeDomainWarpFractalNone        FastNoiseLiteDomainWarpFractalType = 0
	FastNoiseLiteDomainWarpFractalTypeDomainWarpFractalProgressive FastNoiseLiteDomainWarpFractalType = 1
	FastNoiseLiteDomainWarpFractalTypeDomainWarpFractalIndependent FastNoiseLiteDomainWarpFractalType = 2
)

func (me *FastNoiseLite) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *FastNoiseLite) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *FastNoiseLite) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *FastNoiseLite) SetNoiseType(type_ FastNoiseLiteNoiseType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetNoiseType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetNoiseType() FastNoiseLiteNoiseType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret FastNoiseLiteNoiseType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetNoiseType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FastNoiseLite) SetSeed(seed int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetSeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetSeed() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetSeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetFrequency(freq float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freq)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetFrequency), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetFrequency() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetFrequency), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetOffset(offset Vector3) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetOffset() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FastNoiseLite) SetFractalType(type_ FastNoiseLiteFractalType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetFractalType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetFractalType() FastNoiseLiteFractalType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret FastNoiseLiteFractalType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetFractalType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FastNoiseLite) SetFractalOctaves(octave_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&octave_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetFractalOctaves), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetFractalOctaves() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetFractalOctaves), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetFractalLacunarity(lacunarity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lacunarity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetFractalLacunarity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetFractalLacunarity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetFractalLacunarity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetFractalGain(gain float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gain)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetFractalGain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetFractalGain() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetFractalGain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetFractalWeightedStrength(weighted_strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&weighted_strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetFractalWeightedStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetFractalWeightedStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetFractalWeightedStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetFractalPingPongStrength(ping_pong_strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ping_pong_strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetFractalPingPongStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetFractalPingPongStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetFractalPingPongStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetCellularDistanceFunction(func_ FastNoiseLiteCellularDistanceFunction) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetCellularDistanceFunction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetCellularDistanceFunction() FastNoiseLiteCellularDistanceFunction {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret FastNoiseLiteCellularDistanceFunction

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetCellularDistanceFunction), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FastNoiseLite) SetCellularJitter(jitter float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&jitter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetCellularJitter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetCellularJitter() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetCellularJitter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetCellularReturnType(ret FastNoiseLiteCellularReturnType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ret)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetCellularReturnType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetCellularReturnType() FastNoiseLiteCellularReturnType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret FastNoiseLiteCellularReturnType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetCellularReturnType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FastNoiseLite) SetDomainWarpEnabled(domain_warp_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) IsDomainWarpEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnIsDomainWarpEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetDomainWarpType(domain_warp_type FastNoiseLiteDomainWarpType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetDomainWarpType() FastNoiseLiteDomainWarpType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret FastNoiseLiteDomainWarpType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetDomainWarpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FastNoiseLite) SetDomainWarpAmplitude(domain_warp_amplitude float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_amplitude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpAmplitude), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetDomainWarpAmplitude() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetDomainWarpAmplitude), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetDomainWarpFrequency(domain_warp_frequency float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_frequency)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpFrequency), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetDomainWarpFrequency() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetDomainWarpFrequency), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetDomainWarpFractalType(domain_warp_fractal_type FastNoiseLiteDomainWarpFractalType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_fractal_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpFractalType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetDomainWarpFractalType() FastNoiseLiteDomainWarpFractalType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret FastNoiseLiteDomainWarpFractalType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetDomainWarpFractalType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FastNoiseLite) SetDomainWarpFractalOctaves(domain_warp_octave_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_octave_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpFractalOctaves), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetDomainWarpFractalOctaves() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetDomainWarpFractalOctaves), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetDomainWarpFractalLacunarity(domain_warp_lacunarity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_lacunarity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpFractalLacunarity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetDomainWarpFractalLacunarity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetDomainWarpFractalLacunarity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FastNoiseLite) SetDomainWarpFractalGain(domain_warp_gain float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_gain)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnSetDomainWarpFractalGain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FastNoiseLite) GetDomainWarpFractalGain() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFastNoiseLite.fnGetDomainWarpFractalGain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
