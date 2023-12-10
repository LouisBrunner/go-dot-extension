// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type FastNoiseLite struct {
  obj gdc.ObjectPtr
}

func (me *FastNoiseLite) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FastNoiseLite) BaseClass() string {
  return "FastNoiseLite"
}



// Enums

type FastNoiseLiteNoiseType int
const (
  FastNoiseLiteNoiseTypeTypeValue FastNoiseLiteNoiseType = 5
  FastNoiseLiteNoiseTypeTypeValueCubic FastNoiseLiteNoiseType = 4
  FastNoiseLiteNoiseTypeTypePerlin FastNoiseLiteNoiseType = 3
  FastNoiseLiteNoiseTypeTypeCellular FastNoiseLiteNoiseType = 2
  FastNoiseLiteNoiseTypeTypeSimplex FastNoiseLiteNoiseType = 0
  FastNoiseLiteNoiseTypeTypeSimplexSmooth FastNoiseLiteNoiseType = 1
)

type FastNoiseLiteFractalType int
const (
  FastNoiseLiteFractalTypeFractalNone FastNoiseLiteFractalType = 0
  FastNoiseLiteFractalTypeFractalFbm FastNoiseLiteFractalType = 1
  FastNoiseLiteFractalTypeFractalRidged FastNoiseLiteFractalType = 2
  FastNoiseLiteFractalTypeFractalPingPong FastNoiseLiteFractalType = 3
)

type FastNoiseLiteCellularDistanceFunction int
const (
  FastNoiseLiteCellularDistanceFunctionDistanceEuclidean FastNoiseLiteCellularDistanceFunction = 0
  FastNoiseLiteCellularDistanceFunctionDistanceEuclideanSquared FastNoiseLiteCellularDistanceFunction = 1
  FastNoiseLiteCellularDistanceFunctionDistanceManhattan FastNoiseLiteCellularDistanceFunction = 2
  FastNoiseLiteCellularDistanceFunctionDistanceHybrid FastNoiseLiteCellularDistanceFunction = 3
)

type FastNoiseLiteCellularReturnType int
const (
  FastNoiseLiteCellularReturnTypeReturnCellValue FastNoiseLiteCellularReturnType = 0
  FastNoiseLiteCellularReturnTypeReturnDistance FastNoiseLiteCellularReturnType = 1
  FastNoiseLiteCellularReturnTypeReturnDistance2 FastNoiseLiteCellularReturnType = 2
  FastNoiseLiteCellularReturnTypeReturnDistance2Add FastNoiseLiteCellularReturnType = 3
  FastNoiseLiteCellularReturnTypeReturnDistance2Sub FastNoiseLiteCellularReturnType = 4
  FastNoiseLiteCellularReturnTypeReturnDistance2Mul FastNoiseLiteCellularReturnType = 5
  FastNoiseLiteCellularReturnTypeReturnDistance2Div FastNoiseLiteCellularReturnType = 6
)

type FastNoiseLiteDomainWarpType int
const (
  FastNoiseLiteDomainWarpTypeDomainWarpSimplex FastNoiseLiteDomainWarpType = 0
  FastNoiseLiteDomainWarpTypeDomainWarpSimplexReduced FastNoiseLiteDomainWarpType = 1
  FastNoiseLiteDomainWarpTypeDomainWarpBasicGrid FastNoiseLiteDomainWarpType = 2
)

type FastNoiseLiteDomainWarpFractalType int
const (
  FastNoiseLiteDomainWarpFractalTypeDomainWarpFractalNone FastNoiseLiteDomainWarpFractalType = 0
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

func  (me *FastNoiseLite) SetNoiseType(type_ FastNoiseLiteNoiseType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_noise_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2624461392) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetNoiseType() FastNoiseLiteNoiseType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1458108610) // FIXME: should cache?
  var ret FastNoiseLiteNoiseType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetSeed(seed int, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetSeed() int {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetFrequency(freq float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freq), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetFrequency() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetOffset(offset Vector3, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetOffset() Vector3 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetFractalType(type_ FastNoiseLiteFractalType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4132731174) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetFractalType() FastNoiseLiteFractalType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1036889279) // FIXME: should cache?
  var ret FastNoiseLiteFractalType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetFractalOctaves(octave_count int, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&octave_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetFractalOctaves() int {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetFractalLacunarity(lacunarity float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lacunarity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetFractalLacunarity() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetFractalGain(gain float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gain), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetFractalGain() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetFractalWeightedStrength(weighted_strength float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_weighted_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&weighted_strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetFractalWeightedStrength() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_weighted_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetFractalPingPongStrength(ping_pong_strength float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_ping_pong_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ping_pong_strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetFractalPingPongStrength() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_ping_pong_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetCellularDistanceFunction(func_ FastNoiseLiteCellularDistanceFunction, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cellular_distance_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1006013267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetCellularDistanceFunction() FastNoiseLiteCellularDistanceFunction {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cellular_distance_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2021274088) // FIXME: should cache?
  var ret FastNoiseLiteCellularDistanceFunction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetCellularJitter(jitter float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cellular_jitter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&jitter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetCellularJitter() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cellular_jitter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetCellularReturnType(ret FastNoiseLiteCellularReturnType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cellular_return_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2654169698) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ret), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetCellularReturnType() FastNoiseLiteCellularReturnType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cellular_return_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3699796343) // FIXME: should cache?
  var ret FastNoiseLiteCellularReturnType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpEnabled(domain_warp_enabled bool, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) IsDomainWarpEnabled() bool {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_domain_warp_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpType(domain_warp_type FastNoiseLiteDomainWarpType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3629692980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetDomainWarpType() FastNoiseLiteDomainWarpType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2980162020) // FIXME: should cache?
  var ret FastNoiseLiteDomainWarpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpAmplitude(domain_warp_amplitude float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_amplitude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_amplitude), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetDomainWarpAmplitude() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_amplitude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpFrequency(domain_warp_frequency float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_frequency), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetDomainWarpFrequency() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpFractalType(domain_warp_fractal_type FastNoiseLiteDomainWarpFractalType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3999408287) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_fractal_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetDomainWarpFractalType() FastNoiseLiteDomainWarpFractalType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 407716934) // FIXME: should cache?
  var ret FastNoiseLiteDomainWarpFractalType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpFractalOctaves(domain_warp_octave_count int, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_octave_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetDomainWarpFractalOctaves() int {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpFractalLacunarity(domain_warp_lacunarity float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_lacunarity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetDomainWarpFractalLacunarity() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpFractalGain(domain_warp_gain float32, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_gain), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FastNoiseLite) GetDomainWarpFractalGain() float32 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *FastNoiseLite) GetPropNoiseType() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropNoiseType(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropSeed() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropSeed(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropFrequency() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropFrequency(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropOffset() Vector3 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropOffset(value Vector3) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropFractalType() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropFractalType(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropFractalOctaves() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropFractalOctaves(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropFractalLacunarity() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropFractalLacunarity(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropFractalGain() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropFractalGain(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropFractalWeightedStrength() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropFractalWeightedStrength(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropFractalPingPongStrength() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropFractalPingPongStrength(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropCellularDistanceFunction() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropCellularDistanceFunction(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropCellularJitter() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropCellularJitter(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropCellularReturnType() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropCellularReturnType(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpEnabled() bool {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpEnabled(value bool) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpType() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpType(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpAmplitude() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpAmplitude(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpFrequency() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpFrequency(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpFractalType() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpFractalType(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpFractalOctaves() int {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpFractalOctaves(value int) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpFractalLacunarity() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpFractalLacunarity(value float32) {
  panic("TODO: implement")
}

func (me *FastNoiseLite) GetPropDomainWarpFractalGain() float32 {
  panic("TODO: implement")
}

func (me *FastNoiseLite) SetPropDomainWarpFractalGain(value float32) {
  panic("TODO: implement")
}