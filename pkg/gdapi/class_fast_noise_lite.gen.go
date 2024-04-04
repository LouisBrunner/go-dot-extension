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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetNoiseType() FastNoiseLiteNoiseType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1458108610) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FastNoiseLiteNoiseType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FastNoiseLite) SetSeed(seed int64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetSeed() int64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetFrequency(freq float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freq) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetFrequency() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetOffset(offset Vector3, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetOffset() Vector3 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FastNoiseLite) SetFractalType(type_ FastNoiseLiteFractalType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4132731174) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetFractalType() FastNoiseLiteFractalType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1036889279) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FastNoiseLiteFractalType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FastNoiseLite) SetFractalOctaves(octave_count int64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&octave_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetFractalOctaves() int64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetFractalLacunarity(lacunarity float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lacunarity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetFractalLacunarity() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetFractalGain(gain float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gain) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetFractalGain() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetFractalWeightedStrength(weighted_strength float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_weighted_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&weighted_strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetFractalWeightedStrength() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_weighted_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetFractalPingPongStrength(ping_pong_strength float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractal_ping_pong_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ping_pong_strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetFractalPingPongStrength() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractal_ping_pong_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetCellularDistanceFunction(func_ FastNoiseLiteCellularDistanceFunction, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cellular_distance_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1006013267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetCellularDistanceFunction() FastNoiseLiteCellularDistanceFunction {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cellular_distance_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2021274088) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FastNoiseLiteCellularDistanceFunction

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FastNoiseLite) SetCellularJitter(jitter float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cellular_jitter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&jitter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetCellularJitter() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cellular_jitter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetCellularReturnType(ret FastNoiseLiteCellularReturnType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cellular_return_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2654169698) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ret) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetCellularReturnType() FastNoiseLiteCellularReturnType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cellular_return_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3699796343) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FastNoiseLiteCellularReturnType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpEnabled(domain_warp_enabled bool, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) IsDomainWarpEnabled() bool {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_domain_warp_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetDomainWarpType(domain_warp_type FastNoiseLiteDomainWarpType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3629692980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetDomainWarpType() FastNoiseLiteDomainWarpType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2980162020) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FastNoiseLiteDomainWarpType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpAmplitude(domain_warp_amplitude float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_amplitude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_amplitude) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetDomainWarpAmplitude() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_amplitude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetDomainWarpFrequency(domain_warp_frequency float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_frequency) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetDomainWarpFrequency() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetDomainWarpFractalType(domain_warp_fractal_type FastNoiseLiteDomainWarpFractalType, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3999408287) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_fractal_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetDomainWarpFractalType() FastNoiseLiteDomainWarpFractalType {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 407716934) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FastNoiseLiteDomainWarpFractalType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FastNoiseLite) SetDomainWarpFractalOctaves(domain_warp_octave_count int64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_octave_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetDomainWarpFractalOctaves() int64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_octaves")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetDomainWarpFractalLacunarity(domain_warp_lacunarity float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_lacunarity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetDomainWarpFractalLacunarity() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_lacunarity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FastNoiseLite) SetDomainWarpFractalGain(domain_warp_gain float64, )  {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_domain_warp_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&domain_warp_gain) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FastNoiseLite) GetDomainWarpFractalGain() float64 {
  classNameV := StringNameFromStr("FastNoiseLite")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_domain_warp_fractal_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
