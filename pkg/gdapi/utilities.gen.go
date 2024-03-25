// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UtilityFunctions interface {
  Sin(angle_rad float32, ) float32
  Cos(angle_rad float32, ) float32
  Tan(angle_rad float32, ) float32
  Sinh(x float32, ) float32
  Cosh(x float32, ) float32
  Tanh(x float32, ) float32
  Asin(x float32, ) float32
  Acos(x float32, ) float32
  Atan(x float32, ) float32
  Atan2(y float32, x float32, ) float32
  Asinh(x float32, ) float32
  Acosh(x float32, ) float32
  Atanh(x float32, ) float32
  Sqrt(x float32, ) float32
  Fmod(x float32, y float32, ) float32
  Fposmod(x float32, y float32, ) float32
  Posmod(x int, y int, ) int
  Floor(x Variant, ) Variant
  Floorf(x float32, ) float32
  Floori(x float32, ) int
  Ceil(x Variant, ) Variant
  Ceilf(x float32, ) float32
  Ceili(x float32, ) int
  Round(x Variant, ) Variant
  Roundf(x float32, ) float32
  Roundi(x float32, ) int
  Abs(x Variant, ) Variant
  Absf(x float32, ) float32
  Absi(x int, ) int
  Sign(x Variant, ) Variant
  Signf(x float32, ) float32
  Signi(x int, ) int
  Snapped(x Variant, step Variant, ) Variant
  Snappedf(x float32, step float32, ) float32
  Snappedi(x float32, step int, ) int
  Pow(base float32, exp float32, ) float32
  Log(x float32, ) float32
  Exp(x float32, ) float32
  IsNan(x float32, ) bool
  IsInf(x float32, ) bool
  IsEqualApprox(a float32, b float32, ) bool
  IsZeroApprox(x float32, ) bool
  IsFinite(x float32, ) bool
  Ease(x float32, curve float32, ) float32
  StepDecimals(x float32, ) int
  Lerp(from Variant, to Variant, weight Variant, ) Variant
  Lerpf(from float32, to float32, weight float32, ) float32
  CubicInterpolate(from float32, to float32, pre float32, post float32, weight float32, ) float32
  CubicInterpolateAngle(from float32, to float32, pre float32, post float32, weight float32, ) float32
  CubicInterpolateInTime(from float32, to float32, pre float32, post float32, weight float32, to_t float32, pre_t float32, post_t float32, ) float32
  CubicInterpolateAngleInTime(from float32, to float32, pre float32, post float32, weight float32, to_t float32, pre_t float32, post_t float32, ) float32
  BezierInterpolate(start float32, control_1 float32, control_2 float32, end float32, t float32, ) float32
  BezierDerivative(start float32, control_1 float32, control_2 float32, end float32, t float32, ) float32
  AngleDifference(from float32, to float32, ) float32
  LerpAngle(from float32, to float32, weight float32, ) float32
  InverseLerp(from float32, to float32, weight float32, ) float32
  Remap(value float32, istart float32, istop float32, ostart float32, ostop float32, ) float32
  Smoothstep(from float32, to float32, x float32, ) float32
  MoveToward(from float32, to float32, delta float32, ) float32
  RotateToward(from float32, to float32, delta float32, ) float32
  DegToRad(deg float32, ) float32
  RadToDeg(rad float32, ) float32
  LinearToDb(lin float32, ) float32
  DbToLinear(db float32, ) float32
  Wrap(value Variant, min Variant, max Variant, ) Variant
  Wrapi(value int, min int, max int, ) int
  Wrapf(value float32, min float32, max float32, ) float32
  Max(args ...Variant) Variant
  Maxi(a int, b int, ) int
  Maxf(a float32, b float32, ) float32
  Min(args ...Variant) Variant
  Mini(a int, b int, ) int
  Minf(a float32, b float32, ) float32
  Clamp(value Variant, min Variant, max Variant, ) Variant
  Clampi(value int, min int, max int, ) int
  Clampf(value float32, min float32, max float32, ) float32
  NearestPo2(value int, ) int
  Pingpong(value float32, length float32, ) float32
  Randomize() 
  Randi() int
  Randf() float32
  RandiRange(from int, to int, ) int
  RandfRange(from float32, to float32, ) float32
  Randfn(mean float32, deviation float32, ) float32
  Seed(base int, ) 
  RandFromSeed(seed int, ) PackedInt64Array
  Weakref(obj Variant, ) Variant
  Typeof(variable Variant, ) int
  TypeConvert(variant Variant, type_ int, ) Variant
  Str(args ...Variant) String
  ErrorString(error int, ) String
  TypeString(type_ int, ) String
  Print(args ...Variant) 
  PrintRich(args ...Variant) 
  Printerr(args ...Variant) 
  Printt(args ...Variant) 
  Prints(args ...Variant) 
  Printraw(args ...Variant) 
  PrintVerbose(args ...Variant) 
  PushError(args ...Variant) 
  PushWarning(args ...Variant) 
  VarToStr(variable Variant, ) String
  StrToVar(string_ String, ) Variant
  VarToBytes(variable Variant, ) PackedByteArray
  BytesToVar(bytes PackedByteArray, ) Variant
  VarToBytesWithObjects(variable Variant, ) PackedByteArray
  BytesToVarWithObjects(bytes PackedByteArray, ) Variant
  Hash(variable Variant, ) int
  InstanceFromId(instance_id int, ) Object
  IsInstanceIdValid(id int, ) bool
  IsInstanceValid(instance Variant, ) bool
  RidAllocateId() int
  RidFromInt64(base int, ) RID
  IsSame(a Variant, b Variant, ) bool
}

type utilities struct {
  iface gdc.Interface
  ptrsin gdc.PtrUtilityFunction
  ptrcos gdc.PtrUtilityFunction
  ptrtan gdc.PtrUtilityFunction
  ptrsinh gdc.PtrUtilityFunction
  ptrcosh gdc.PtrUtilityFunction
  ptrtanh gdc.PtrUtilityFunction
  ptrasin gdc.PtrUtilityFunction
  ptracos gdc.PtrUtilityFunction
  ptratan gdc.PtrUtilityFunction
  ptratan2 gdc.PtrUtilityFunction
  ptrasinh gdc.PtrUtilityFunction
  ptracosh gdc.PtrUtilityFunction
  ptratanh gdc.PtrUtilityFunction
  ptrsqrt gdc.PtrUtilityFunction
  ptrfmod gdc.PtrUtilityFunction
  ptrfposmod gdc.PtrUtilityFunction
  ptrposmod gdc.PtrUtilityFunction
  ptrfloor gdc.PtrUtilityFunction
  ptrfloorf gdc.PtrUtilityFunction
  ptrfloori gdc.PtrUtilityFunction
  ptrceil gdc.PtrUtilityFunction
  ptrceilf gdc.PtrUtilityFunction
  ptrceili gdc.PtrUtilityFunction
  ptrround gdc.PtrUtilityFunction
  ptrroundf gdc.PtrUtilityFunction
  ptrroundi gdc.PtrUtilityFunction
  ptrabs gdc.PtrUtilityFunction
  ptrabsf gdc.PtrUtilityFunction
  ptrabsi gdc.PtrUtilityFunction
  ptrsign gdc.PtrUtilityFunction
  ptrsignf gdc.PtrUtilityFunction
  ptrsigni gdc.PtrUtilityFunction
  ptrsnapped gdc.PtrUtilityFunction
  ptrsnappedf gdc.PtrUtilityFunction
  ptrsnappedi gdc.PtrUtilityFunction
  ptrpow gdc.PtrUtilityFunction
  ptrlog gdc.PtrUtilityFunction
  ptrexp gdc.PtrUtilityFunction
  ptris_nan gdc.PtrUtilityFunction
  ptris_inf gdc.PtrUtilityFunction
  ptris_equal_approx gdc.PtrUtilityFunction
  ptris_zero_approx gdc.PtrUtilityFunction
  ptris_finite gdc.PtrUtilityFunction
  ptrease gdc.PtrUtilityFunction
  ptrstep_decimals gdc.PtrUtilityFunction
  ptrlerp gdc.PtrUtilityFunction
  ptrlerpf gdc.PtrUtilityFunction
  ptrcubic_interpolate gdc.PtrUtilityFunction
  ptrcubic_interpolate_angle gdc.PtrUtilityFunction
  ptrcubic_interpolate_in_time gdc.PtrUtilityFunction
  ptrcubic_interpolate_angle_in_time gdc.PtrUtilityFunction
  ptrbezier_interpolate gdc.PtrUtilityFunction
  ptrbezier_derivative gdc.PtrUtilityFunction
  ptrangle_difference gdc.PtrUtilityFunction
  ptrlerp_angle gdc.PtrUtilityFunction
  ptrinverse_lerp gdc.PtrUtilityFunction
  ptrremap gdc.PtrUtilityFunction
  ptrsmoothstep gdc.PtrUtilityFunction
  ptrmove_toward gdc.PtrUtilityFunction
  ptrrotate_toward gdc.PtrUtilityFunction
  ptrdeg_to_rad gdc.PtrUtilityFunction
  ptrrad_to_deg gdc.PtrUtilityFunction
  ptrlinear_to_db gdc.PtrUtilityFunction
  ptrdb_to_linear gdc.PtrUtilityFunction
  ptrwrap gdc.PtrUtilityFunction
  ptrwrapi gdc.PtrUtilityFunction
  ptrwrapf gdc.PtrUtilityFunction
  ptrmax gdc.PtrUtilityFunction
  ptrmaxi gdc.PtrUtilityFunction
  ptrmaxf gdc.PtrUtilityFunction
  ptrmin gdc.PtrUtilityFunction
  ptrmini gdc.PtrUtilityFunction
  ptrminf gdc.PtrUtilityFunction
  ptrclamp gdc.PtrUtilityFunction
  ptrclampi gdc.PtrUtilityFunction
  ptrclampf gdc.PtrUtilityFunction
  ptrnearest_po2 gdc.PtrUtilityFunction
  ptrpingpong gdc.PtrUtilityFunction
  ptrrandomize gdc.PtrUtilityFunction
  ptrrandi gdc.PtrUtilityFunction
  ptrrandf gdc.PtrUtilityFunction
  ptrrandi_range gdc.PtrUtilityFunction
  ptrrandf_range gdc.PtrUtilityFunction
  ptrrandfn gdc.PtrUtilityFunction
  ptrseed gdc.PtrUtilityFunction
  ptrrand_from_seed gdc.PtrUtilityFunction
  ptrweakref gdc.PtrUtilityFunction
  ptrtypeof gdc.PtrUtilityFunction
  ptrtype_convert gdc.PtrUtilityFunction
  ptrstr gdc.PtrUtilityFunction
  ptrerror_string gdc.PtrUtilityFunction
  ptrtype_string gdc.PtrUtilityFunction
  ptrprint gdc.PtrUtilityFunction
  ptrprint_rich gdc.PtrUtilityFunction
  ptrprinterr gdc.PtrUtilityFunction
  ptrprintt gdc.PtrUtilityFunction
  ptrprints gdc.PtrUtilityFunction
  ptrprintraw gdc.PtrUtilityFunction
  ptrprint_verbose gdc.PtrUtilityFunction
  ptrpush_error gdc.PtrUtilityFunction
  ptrpush_warning gdc.PtrUtilityFunction
  ptrvar_to_str gdc.PtrUtilityFunction
  ptrstr_to_var gdc.PtrUtilityFunction
  ptrvar_to_bytes gdc.PtrUtilityFunction
  ptrbytes_to_var gdc.PtrUtilityFunction
  ptrvar_to_bytes_with_objects gdc.PtrUtilityFunction
  ptrbytes_to_var_with_objects gdc.PtrUtilityFunction
  ptrhash gdc.PtrUtilityFunction
  ptrinstance_from_id gdc.PtrUtilityFunction
  ptris_instance_id_valid gdc.PtrUtilityFunction
  ptris_instance_valid gdc.PtrUtilityFunction
  ptrrid_allocate_id gdc.PtrUtilityFunction
  ptrrid_from_int64 gdc.PtrUtilityFunction
  ptris_same gdc.PtrUtilityFunction
}

func newUtilities(iface gdc.Interface) UtilityFunctions {
  strsin := StringNameFromStr("sin")
  defer strsin.Destroy()
  strcos := StringNameFromStr("cos")
  defer strcos.Destroy()
  strtan := StringNameFromStr("tan")
  defer strtan.Destroy()
  strsinh := StringNameFromStr("sinh")
  defer strsinh.Destroy()
  strcosh := StringNameFromStr("cosh")
  defer strcosh.Destroy()
  strtanh := StringNameFromStr("tanh")
  defer strtanh.Destroy()
  strasin := StringNameFromStr("asin")
  defer strasin.Destroy()
  stracos := StringNameFromStr("acos")
  defer stracos.Destroy()
  stratan := StringNameFromStr("atan")
  defer stratan.Destroy()
  stratan2 := StringNameFromStr("atan2")
  defer stratan2.Destroy()
  strasinh := StringNameFromStr("asinh")
  defer strasinh.Destroy()
  stracosh := StringNameFromStr("acosh")
  defer stracosh.Destroy()
  stratanh := StringNameFromStr("atanh")
  defer stratanh.Destroy()
  strsqrt := StringNameFromStr("sqrt")
  defer strsqrt.Destroy()
  strfmod := StringNameFromStr("fmod")
  defer strfmod.Destroy()
  strfposmod := StringNameFromStr("fposmod")
  defer strfposmod.Destroy()
  strposmod := StringNameFromStr("posmod")
  defer strposmod.Destroy()
  strfloor := StringNameFromStr("floor")
  defer strfloor.Destroy()
  strfloorf := StringNameFromStr("floorf")
  defer strfloorf.Destroy()
  strfloori := StringNameFromStr("floori")
  defer strfloori.Destroy()
  strceil := StringNameFromStr("ceil")
  defer strceil.Destroy()
  strceilf := StringNameFromStr("ceilf")
  defer strceilf.Destroy()
  strceili := StringNameFromStr("ceili")
  defer strceili.Destroy()
  strround := StringNameFromStr("round")
  defer strround.Destroy()
  strroundf := StringNameFromStr("roundf")
  defer strroundf.Destroy()
  strroundi := StringNameFromStr("roundi")
  defer strroundi.Destroy()
  strabs := StringNameFromStr("abs")
  defer strabs.Destroy()
  strabsf := StringNameFromStr("absf")
  defer strabsf.Destroy()
  strabsi := StringNameFromStr("absi")
  defer strabsi.Destroy()
  strsign := StringNameFromStr("sign")
  defer strsign.Destroy()
  strsignf := StringNameFromStr("signf")
  defer strsignf.Destroy()
  strsigni := StringNameFromStr("signi")
  defer strsigni.Destroy()
  strsnapped := StringNameFromStr("snapped")
  defer strsnapped.Destroy()
  strsnappedf := StringNameFromStr("snappedf")
  defer strsnappedf.Destroy()
  strsnappedi := StringNameFromStr("snappedi")
  defer strsnappedi.Destroy()
  strpow := StringNameFromStr("pow")
  defer strpow.Destroy()
  strlog := StringNameFromStr("log")
  defer strlog.Destroy()
  strexp := StringNameFromStr("exp")
  defer strexp.Destroy()
  stris_nan := StringNameFromStr("is_nan")
  defer stris_nan.Destroy()
  stris_inf := StringNameFromStr("is_inf")
  defer stris_inf.Destroy()
  stris_equal_approx := StringNameFromStr("is_equal_approx")
  defer stris_equal_approx.Destroy()
  stris_zero_approx := StringNameFromStr("is_zero_approx")
  defer stris_zero_approx.Destroy()
  stris_finite := StringNameFromStr("is_finite")
  defer stris_finite.Destroy()
  strease := StringNameFromStr("ease")
  defer strease.Destroy()
  strstep_decimals := StringNameFromStr("step_decimals")
  defer strstep_decimals.Destroy()
  strlerp := StringNameFromStr("lerp")
  defer strlerp.Destroy()
  strlerpf := StringNameFromStr("lerpf")
  defer strlerpf.Destroy()
  strcubic_interpolate := StringNameFromStr("cubic_interpolate")
  defer strcubic_interpolate.Destroy()
  strcubic_interpolate_angle := StringNameFromStr("cubic_interpolate_angle")
  defer strcubic_interpolate_angle.Destroy()
  strcubic_interpolate_in_time := StringNameFromStr("cubic_interpolate_in_time")
  defer strcubic_interpolate_in_time.Destroy()
  strcubic_interpolate_angle_in_time := StringNameFromStr("cubic_interpolate_angle_in_time")
  defer strcubic_interpolate_angle_in_time.Destroy()
  strbezier_interpolate := StringNameFromStr("bezier_interpolate")
  defer strbezier_interpolate.Destroy()
  strbezier_derivative := StringNameFromStr("bezier_derivative")
  defer strbezier_derivative.Destroy()
  strangle_difference := StringNameFromStr("angle_difference")
  defer strangle_difference.Destroy()
  strlerp_angle := StringNameFromStr("lerp_angle")
  defer strlerp_angle.Destroy()
  strinverse_lerp := StringNameFromStr("inverse_lerp")
  defer strinverse_lerp.Destroy()
  strremap := StringNameFromStr("remap")
  defer strremap.Destroy()
  strsmoothstep := StringNameFromStr("smoothstep")
  defer strsmoothstep.Destroy()
  strmove_toward := StringNameFromStr("move_toward")
  defer strmove_toward.Destroy()
  strrotate_toward := StringNameFromStr("rotate_toward")
  defer strrotate_toward.Destroy()
  strdeg_to_rad := StringNameFromStr("deg_to_rad")
  defer strdeg_to_rad.Destroy()
  strrad_to_deg := StringNameFromStr("rad_to_deg")
  defer strrad_to_deg.Destroy()
  strlinear_to_db := StringNameFromStr("linear_to_db")
  defer strlinear_to_db.Destroy()
  strdb_to_linear := StringNameFromStr("db_to_linear")
  defer strdb_to_linear.Destroy()
  strwrap := StringNameFromStr("wrap")
  defer strwrap.Destroy()
  strwrapi := StringNameFromStr("wrapi")
  defer strwrapi.Destroy()
  strwrapf := StringNameFromStr("wrapf")
  defer strwrapf.Destroy()
  strmax := StringNameFromStr("max")
  defer strmax.Destroy()
  strmaxi := StringNameFromStr("maxi")
  defer strmaxi.Destroy()
  strmaxf := StringNameFromStr("maxf")
  defer strmaxf.Destroy()
  strmin := StringNameFromStr("min")
  defer strmin.Destroy()
  strmini := StringNameFromStr("mini")
  defer strmini.Destroy()
  strminf := StringNameFromStr("minf")
  defer strminf.Destroy()
  strclamp := StringNameFromStr("clamp")
  defer strclamp.Destroy()
  strclampi := StringNameFromStr("clampi")
  defer strclampi.Destroy()
  strclampf := StringNameFromStr("clampf")
  defer strclampf.Destroy()
  strnearest_po2 := StringNameFromStr("nearest_po2")
  defer strnearest_po2.Destroy()
  strpingpong := StringNameFromStr("pingpong")
  defer strpingpong.Destroy()
  strrandomize := StringNameFromStr("randomize")
  defer strrandomize.Destroy()
  strrandi := StringNameFromStr("randi")
  defer strrandi.Destroy()
  strrandf := StringNameFromStr("randf")
  defer strrandf.Destroy()
  strrandi_range := StringNameFromStr("randi_range")
  defer strrandi_range.Destroy()
  strrandf_range := StringNameFromStr("randf_range")
  defer strrandf_range.Destroy()
  strrandfn := StringNameFromStr("randfn")
  defer strrandfn.Destroy()
  strseed := StringNameFromStr("seed")
  defer strseed.Destroy()
  strrand_from_seed := StringNameFromStr("rand_from_seed")
  defer strrand_from_seed.Destroy()
  strweakref := StringNameFromStr("weakref")
  defer strweakref.Destroy()
  strtypeof := StringNameFromStr("typeof")
  defer strtypeof.Destroy()
  strtype_convert := StringNameFromStr("type_convert")
  defer strtype_convert.Destroy()
  strstr := StringNameFromStr("str")
  defer strstr.Destroy()
  strerror_string := StringNameFromStr("error_string")
  defer strerror_string.Destroy()
  strtype_string := StringNameFromStr("type_string")
  defer strtype_string.Destroy()
  strprint := StringNameFromStr("print")
  defer strprint.Destroy()
  strprint_rich := StringNameFromStr("print_rich")
  defer strprint_rich.Destroy()
  strprinterr := StringNameFromStr("printerr")
  defer strprinterr.Destroy()
  strprintt := StringNameFromStr("printt")
  defer strprintt.Destroy()
  strprints := StringNameFromStr("prints")
  defer strprints.Destroy()
  strprintraw := StringNameFromStr("printraw")
  defer strprintraw.Destroy()
  strprint_verbose := StringNameFromStr("print_verbose")
  defer strprint_verbose.Destroy()
  strpush_error := StringNameFromStr("push_error")
  defer strpush_error.Destroy()
  strpush_warning := StringNameFromStr("push_warning")
  defer strpush_warning.Destroy()
  strvar_to_str := StringNameFromStr("var_to_str")
  defer strvar_to_str.Destroy()
  strstr_to_var := StringNameFromStr("str_to_var")
  defer strstr_to_var.Destroy()
  strvar_to_bytes := StringNameFromStr("var_to_bytes")
  defer strvar_to_bytes.Destroy()
  strbytes_to_var := StringNameFromStr("bytes_to_var")
  defer strbytes_to_var.Destroy()
  strvar_to_bytes_with_objects := StringNameFromStr("var_to_bytes_with_objects")
  defer strvar_to_bytes_with_objects.Destroy()
  strbytes_to_var_with_objects := StringNameFromStr("bytes_to_var_with_objects")
  defer strbytes_to_var_with_objects.Destroy()
  strhash := StringNameFromStr("hash")
  defer strhash.Destroy()
  strinstance_from_id := StringNameFromStr("instance_from_id")
  defer strinstance_from_id.Destroy()
  stris_instance_id_valid := StringNameFromStr("is_instance_id_valid")
  defer stris_instance_id_valid.Destroy()
  stris_instance_valid := StringNameFromStr("is_instance_valid")
  defer stris_instance_valid.Destroy()
  strrid_allocate_id := StringNameFromStr("rid_allocate_id")
  defer strrid_allocate_id.Destroy()
  strrid_from_int64 := StringNameFromStr("rid_from_int64")
  defer strrid_from_int64.Destroy()
  stris_same := StringNameFromStr("is_same")
  defer stris_same.Destroy()
  return &utilities{
    ptrsin: iface.VariantGetPtrUtilityFunction(strsin.AsCPtr(), 2140049587),
    ptrcos: iface.VariantGetPtrUtilityFunction(strcos.AsCPtr(), 2140049587),
    ptrtan: iface.VariantGetPtrUtilityFunction(strtan.AsCPtr(), 2140049587),
    ptrsinh: iface.VariantGetPtrUtilityFunction(strsinh.AsCPtr(), 2140049587),
    ptrcosh: iface.VariantGetPtrUtilityFunction(strcosh.AsCPtr(), 2140049587),
    ptrtanh: iface.VariantGetPtrUtilityFunction(strtanh.AsCPtr(), 2140049587),
    ptrasin: iface.VariantGetPtrUtilityFunction(strasin.AsCPtr(), 2140049587),
    ptracos: iface.VariantGetPtrUtilityFunction(stracos.AsCPtr(), 2140049587),
    ptratan: iface.VariantGetPtrUtilityFunction(stratan.AsCPtr(), 2140049587),
    ptratan2: iface.VariantGetPtrUtilityFunction(stratan2.AsCPtr(), 92296394),
    ptrasinh: iface.VariantGetPtrUtilityFunction(strasinh.AsCPtr(), 2140049587),
    ptracosh: iface.VariantGetPtrUtilityFunction(stracosh.AsCPtr(), 2140049587),
    ptratanh: iface.VariantGetPtrUtilityFunction(stratanh.AsCPtr(), 2140049587),
    ptrsqrt: iface.VariantGetPtrUtilityFunction(strsqrt.AsCPtr(), 2140049587),
    ptrfmod: iface.VariantGetPtrUtilityFunction(strfmod.AsCPtr(), 92296394),
    ptrfposmod: iface.VariantGetPtrUtilityFunction(strfposmod.AsCPtr(), 92296394),
    ptrposmod: iface.VariantGetPtrUtilityFunction(strposmod.AsCPtr(), 3133453818),
    ptrfloor: iface.VariantGetPtrUtilityFunction(strfloor.AsCPtr(), 4776452),
    ptrfloorf: iface.VariantGetPtrUtilityFunction(strfloorf.AsCPtr(), 2140049587),
    ptrfloori: iface.VariantGetPtrUtilityFunction(strfloori.AsCPtr(), 2780425386),
    ptrceil: iface.VariantGetPtrUtilityFunction(strceil.AsCPtr(), 4776452),
    ptrceilf: iface.VariantGetPtrUtilityFunction(strceilf.AsCPtr(), 2140049587),
    ptrceili: iface.VariantGetPtrUtilityFunction(strceili.AsCPtr(), 2780425386),
    ptrround: iface.VariantGetPtrUtilityFunction(strround.AsCPtr(), 4776452),
    ptrroundf: iface.VariantGetPtrUtilityFunction(strroundf.AsCPtr(), 2140049587),
    ptrroundi: iface.VariantGetPtrUtilityFunction(strroundi.AsCPtr(), 2780425386),
    ptrabs: iface.VariantGetPtrUtilityFunction(strabs.AsCPtr(), 4776452),
    ptrabsf: iface.VariantGetPtrUtilityFunction(strabsf.AsCPtr(), 2140049587),
    ptrabsi: iface.VariantGetPtrUtilityFunction(strabsi.AsCPtr(), 2157319888),
    ptrsign: iface.VariantGetPtrUtilityFunction(strsign.AsCPtr(), 4776452),
    ptrsignf: iface.VariantGetPtrUtilityFunction(strsignf.AsCPtr(), 2140049587),
    ptrsigni: iface.VariantGetPtrUtilityFunction(strsigni.AsCPtr(), 2157319888),
    ptrsnapped: iface.VariantGetPtrUtilityFunction(strsnapped.AsCPtr(), 459914704),
    ptrsnappedf: iface.VariantGetPtrUtilityFunction(strsnappedf.AsCPtr(), 92296394),
    ptrsnappedi: iface.VariantGetPtrUtilityFunction(strsnappedi.AsCPtr(), 3570758393),
    ptrpow: iface.VariantGetPtrUtilityFunction(strpow.AsCPtr(), 92296394),
    ptrlog: iface.VariantGetPtrUtilityFunction(strlog.AsCPtr(), 2140049587),
    ptrexp: iface.VariantGetPtrUtilityFunction(strexp.AsCPtr(), 2140049587),
    ptris_nan: iface.VariantGetPtrUtilityFunction(stris_nan.AsCPtr(), 3569215213),
    ptris_inf: iface.VariantGetPtrUtilityFunction(stris_inf.AsCPtr(), 3569215213),
    ptris_equal_approx: iface.VariantGetPtrUtilityFunction(stris_equal_approx.AsCPtr(), 1400789633),
    ptris_zero_approx: iface.VariantGetPtrUtilityFunction(stris_zero_approx.AsCPtr(), 3569215213),
    ptris_finite: iface.VariantGetPtrUtilityFunction(stris_finite.AsCPtr(), 3569215213),
    ptrease: iface.VariantGetPtrUtilityFunction(strease.AsCPtr(), 92296394),
    ptrstep_decimals: iface.VariantGetPtrUtilityFunction(strstep_decimals.AsCPtr(), 2780425386),
    ptrlerp: iface.VariantGetPtrUtilityFunction(strlerp.AsCPtr(), 3389874542),
    ptrlerpf: iface.VariantGetPtrUtilityFunction(strlerpf.AsCPtr(), 998901048),
    ptrcubic_interpolate: iface.VariantGetPtrUtilityFunction(strcubic_interpolate.AsCPtr(), 1090965791),
    ptrcubic_interpolate_angle: iface.VariantGetPtrUtilityFunction(strcubic_interpolate_angle.AsCPtr(), 1090965791),
    ptrcubic_interpolate_in_time: iface.VariantGetPtrUtilityFunction(strcubic_interpolate_in_time.AsCPtr(), 388121036),
    ptrcubic_interpolate_angle_in_time: iface.VariantGetPtrUtilityFunction(strcubic_interpolate_angle_in_time.AsCPtr(), 388121036),
    ptrbezier_interpolate: iface.VariantGetPtrUtilityFunction(strbezier_interpolate.AsCPtr(), 1090965791),
    ptrbezier_derivative: iface.VariantGetPtrUtilityFunction(strbezier_derivative.AsCPtr(), 1090965791),
    ptrangle_difference: iface.VariantGetPtrUtilityFunction(strangle_difference.AsCPtr(), 92296394),
    ptrlerp_angle: iface.VariantGetPtrUtilityFunction(strlerp_angle.AsCPtr(), 998901048),
    ptrinverse_lerp: iface.VariantGetPtrUtilityFunction(strinverse_lerp.AsCPtr(), 998901048),
    ptrremap: iface.VariantGetPtrUtilityFunction(strremap.AsCPtr(), 1090965791),
    ptrsmoothstep: iface.VariantGetPtrUtilityFunction(strsmoothstep.AsCPtr(), 998901048),
    ptrmove_toward: iface.VariantGetPtrUtilityFunction(strmove_toward.AsCPtr(), 998901048),
    ptrrotate_toward: iface.VariantGetPtrUtilityFunction(strrotate_toward.AsCPtr(), 998901048),
    ptrdeg_to_rad: iface.VariantGetPtrUtilityFunction(strdeg_to_rad.AsCPtr(), 2140049587),
    ptrrad_to_deg: iface.VariantGetPtrUtilityFunction(strrad_to_deg.AsCPtr(), 2140049587),
    ptrlinear_to_db: iface.VariantGetPtrUtilityFunction(strlinear_to_db.AsCPtr(), 2140049587),
    ptrdb_to_linear: iface.VariantGetPtrUtilityFunction(strdb_to_linear.AsCPtr(), 2140049587),
    ptrwrap: iface.VariantGetPtrUtilityFunction(strwrap.AsCPtr(), 3389874542),
    ptrwrapi: iface.VariantGetPtrUtilityFunction(strwrapi.AsCPtr(), 650295447),
    ptrwrapf: iface.VariantGetPtrUtilityFunction(strwrapf.AsCPtr(), 998901048),
    ptrmax: iface.VariantGetPtrUtilityFunction(strmax.AsCPtr(), 3896050336),
    ptrmaxi: iface.VariantGetPtrUtilityFunction(strmaxi.AsCPtr(), 3133453818),
    ptrmaxf: iface.VariantGetPtrUtilityFunction(strmaxf.AsCPtr(), 92296394),
    ptrmin: iface.VariantGetPtrUtilityFunction(strmin.AsCPtr(), 3896050336),
    ptrmini: iface.VariantGetPtrUtilityFunction(strmini.AsCPtr(), 3133453818),
    ptrminf: iface.VariantGetPtrUtilityFunction(strminf.AsCPtr(), 92296394),
    ptrclamp: iface.VariantGetPtrUtilityFunction(strclamp.AsCPtr(), 3389874542),
    ptrclampi: iface.VariantGetPtrUtilityFunction(strclampi.AsCPtr(), 650295447),
    ptrclampf: iface.VariantGetPtrUtilityFunction(strclampf.AsCPtr(), 998901048),
    ptrnearest_po2: iface.VariantGetPtrUtilityFunction(strnearest_po2.AsCPtr(), 2157319888),
    ptrpingpong: iface.VariantGetPtrUtilityFunction(strpingpong.AsCPtr(), 92296394),
    ptrrandomize: iface.VariantGetPtrUtilityFunction(strrandomize.AsCPtr(), 1691721052),
    ptrrandi: iface.VariantGetPtrUtilityFunction(strrandi.AsCPtr(), 701202648),
    ptrrandf: iface.VariantGetPtrUtilityFunction(strrandf.AsCPtr(), 2086227845),
    ptrrandi_range: iface.VariantGetPtrUtilityFunction(strrandi_range.AsCPtr(), 3133453818),
    ptrrandf_range: iface.VariantGetPtrUtilityFunction(strrandf_range.AsCPtr(), 92296394),
    ptrrandfn: iface.VariantGetPtrUtilityFunction(strrandfn.AsCPtr(), 92296394),
    ptrseed: iface.VariantGetPtrUtilityFunction(strseed.AsCPtr(), 382931173),
    ptrrand_from_seed: iface.VariantGetPtrUtilityFunction(strrand_from_seed.AsCPtr(), 1391063685),
    ptrweakref: iface.VariantGetPtrUtilityFunction(strweakref.AsCPtr(), 4776452),
    ptrtypeof: iface.VariantGetPtrUtilityFunction(strtypeof.AsCPtr(), 326422594),
    ptrtype_convert: iface.VariantGetPtrUtilityFunction(strtype_convert.AsCPtr(), 2453062746),
    ptrstr: iface.VariantGetPtrUtilityFunction(strstr.AsCPtr(), 32569176),
    ptrerror_string: iface.VariantGetPtrUtilityFunction(strerror_string.AsCPtr(), 942708242),
    ptrtype_string: iface.VariantGetPtrUtilityFunction(strtype_string.AsCPtr(), 942708242),
    ptrprint: iface.VariantGetPtrUtilityFunction(strprint.AsCPtr(), 2648703342),
    ptrprint_rich: iface.VariantGetPtrUtilityFunction(strprint_rich.AsCPtr(), 2648703342),
    ptrprinterr: iface.VariantGetPtrUtilityFunction(strprinterr.AsCPtr(), 2648703342),
    ptrprintt: iface.VariantGetPtrUtilityFunction(strprintt.AsCPtr(), 2648703342),
    ptrprints: iface.VariantGetPtrUtilityFunction(strprints.AsCPtr(), 2648703342),
    ptrprintraw: iface.VariantGetPtrUtilityFunction(strprintraw.AsCPtr(), 2648703342),
    ptrprint_verbose: iface.VariantGetPtrUtilityFunction(strprint_verbose.AsCPtr(), 2648703342),
    ptrpush_error: iface.VariantGetPtrUtilityFunction(strpush_error.AsCPtr(), 2648703342),
    ptrpush_warning: iface.VariantGetPtrUtilityFunction(strpush_warning.AsCPtr(), 2648703342),
    ptrvar_to_str: iface.VariantGetPtrUtilityFunction(strvar_to_str.AsCPtr(), 866625479),
    ptrstr_to_var: iface.VariantGetPtrUtilityFunction(strstr_to_var.AsCPtr(), 1891498491),
    ptrvar_to_bytes: iface.VariantGetPtrUtilityFunction(strvar_to_bytes.AsCPtr(), 2947269930),
    ptrbytes_to_var: iface.VariantGetPtrUtilityFunction(strbytes_to_var.AsCPtr(), 4249819452),
    ptrvar_to_bytes_with_objects: iface.VariantGetPtrUtilityFunction(strvar_to_bytes_with_objects.AsCPtr(), 2947269930),
    ptrbytes_to_var_with_objects: iface.VariantGetPtrUtilityFunction(strbytes_to_var_with_objects.AsCPtr(), 4249819452),
    ptrhash: iface.VariantGetPtrUtilityFunction(strhash.AsCPtr(), 326422594),
    ptrinstance_from_id: iface.VariantGetPtrUtilityFunction(strinstance_from_id.AsCPtr(), 1156694636),
    ptris_instance_id_valid: iface.VariantGetPtrUtilityFunction(stris_instance_id_valid.AsCPtr(), 2232439758),
    ptris_instance_valid: iface.VariantGetPtrUtilityFunction(stris_instance_valid.AsCPtr(), 996128841),
    ptrrid_allocate_id: iface.VariantGetPtrUtilityFunction(strrid_allocate_id.AsCPtr(), 701202648),
    ptrrid_from_int64: iface.VariantGetPtrUtilityFunction(strrid_from_int64.AsCPtr(), 3426892196),
    ptris_same: iface.VariantGetPtrUtilityFunction(stris_same.AsCPtr(), 1409423524),
  }
}



func (me *utilities) Sin(angle_rad float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&angle_rad),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsin, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Cos(angle_rad float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&angle_rad),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrcos, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Tan(angle_rad float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&angle_rad),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrtan, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Sinh(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsinh, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Cosh(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrcosh, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Tanh(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrtanh, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Asin(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrasin, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Acos(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptracos, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Atan(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptratan, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Atan2(y float32, x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&y),
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptratan2, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Asinh(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrasinh, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Acosh(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptracosh, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Atanh(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptratanh, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Sqrt(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsqrt, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Fmod(x float32, y float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
    gdc.ConstTypePtr(&y),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrfmod, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Fposmod(x float32, y float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
    gdc.ConstTypePtr(&y),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrfposmod, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Posmod(x int, y int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
    gdc.ConstTypePtr(&y),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrposmod, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Floor(x Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    x.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrfloor, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Floorf(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrfloorf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Floori(x float32, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrfloori, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Ceil(x Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    x.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrceil, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Ceilf(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrceilf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Ceili(x float32, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrceili, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Round(x Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    x.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrround, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Roundf(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrroundf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Roundi(x float32, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrroundi, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Abs(x Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    x.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrabs, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Absf(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrabsf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Absi(x int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrabsi, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Sign(x Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    x.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsign, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Signf(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsignf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Signi(x int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsigni, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Snapped(x Variant, step Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    x.AsCTypePtr(),
    step.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsnapped, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Snappedf(x float32, step float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
    gdc.ConstTypePtr(&step),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsnappedf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Snappedi(x float32, step int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
    gdc.ConstTypePtr(&step),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsnappedi, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Pow(base float32, exp float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&base),
    gdc.ConstTypePtr(&exp),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrpow, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Log(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrlog, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Exp(x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrexp, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsNan(x float32, ) bool {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_nan, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsInf(x float32, ) bool {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_inf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsEqualApprox(a float32, b float32, ) bool {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&a),
    gdc.ConstTypePtr(&b),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_equal_approx, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsZeroApprox(x float32, ) bool {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_zero_approx, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsFinite(x float32, ) bool {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_finite, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Ease(x float32, curve float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
    gdc.ConstTypePtr(&curve),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrease, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) StepDecimals(x float32, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&x),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrstep_decimals, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Lerp(from Variant, to Variant, weight Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    from.AsCTypePtr(),
    to.AsCTypePtr(),
    weight.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrlerp, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Lerpf(from float32, to float32, weight float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&weight),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrlerpf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) CubicInterpolate(from float32, to float32, pre float32, post float32, weight float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&pre),
    gdc.ConstTypePtr(&post),
    gdc.ConstTypePtr(&weight),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) CubicInterpolateAngle(from float32, to float32, pre float32, post float32, weight float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&pre),
    gdc.ConstTypePtr(&post),
    gdc.ConstTypePtr(&weight),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate_angle, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) CubicInterpolateInTime(from float32, to float32, pre float32, post float32, weight float32, to_t float32, pre_t float32, post_t float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&pre),
    gdc.ConstTypePtr(&post),
    gdc.ConstTypePtr(&weight),
    gdc.ConstTypePtr(&to_t),
    gdc.ConstTypePtr(&pre_t),
    gdc.ConstTypePtr(&post_t),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate_in_time, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) CubicInterpolateAngleInTime(from float32, to float32, pre float32, post float32, weight float32, to_t float32, pre_t float32, post_t float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&pre),
    gdc.ConstTypePtr(&post),
    gdc.ConstTypePtr(&weight),
    gdc.ConstTypePtr(&to_t),
    gdc.ConstTypePtr(&pre_t),
    gdc.ConstTypePtr(&post_t),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate_angle_in_time, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) BezierInterpolate(start float32, control_1 float32, control_2 float32, end float32, t float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&start),
    gdc.ConstTypePtr(&control_1),
    gdc.ConstTypePtr(&control_2),
    gdc.ConstTypePtr(&end),
    gdc.ConstTypePtr(&t),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrbezier_interpolate, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) BezierDerivative(start float32, control_1 float32, control_2 float32, end float32, t float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&start),
    gdc.ConstTypePtr(&control_1),
    gdc.ConstTypePtr(&control_2),
    gdc.ConstTypePtr(&end),
    gdc.ConstTypePtr(&t),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrbezier_derivative, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) AngleDifference(from float32, to float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrangle_difference, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) LerpAngle(from float32, to float32, weight float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&weight),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrlerp_angle, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) InverseLerp(from float32, to float32, weight float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&weight),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrinverse_lerp, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Remap(value float32, istart float32, istop float32, ostart float32, ostop float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&value),
    gdc.ConstTypePtr(&istart),
    gdc.ConstTypePtr(&istop),
    gdc.ConstTypePtr(&ostart),
    gdc.ConstTypePtr(&ostop),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrremap, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Smoothstep(from float32, to float32, x float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&x),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrsmoothstep, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) MoveToward(from float32, to float32, delta float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&delta),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrmove_toward, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) RotateToward(from float32, to float32, delta float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
    gdc.ConstTypePtr(&delta),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrotate_toward, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) DegToRad(deg float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&deg),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrdeg_to_rad, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) RadToDeg(rad float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&rad),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrad_to_deg, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) LinearToDb(lin float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&lin),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrlinear_to_db, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) DbToLinear(db float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&db),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrdb_to_linear, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Wrap(value Variant, min Variant, max Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    value.AsCTypePtr(),
    min.AsCTypePtr(),
    max.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrwrap, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Wrapi(value int, min int, max int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&value),
    gdc.ConstTypePtr(&min),
    gdc.ConstTypePtr(&max),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrwrapi, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Wrapf(value float32, min float32, max float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&value),
    gdc.ConstTypePtr(&min),
    gdc.ConstTypePtr(&max),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrwrapf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Max(vargs ...Variant) Variant {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }

  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrmax, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Maxi(a int, b int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&a),
    gdc.ConstTypePtr(&b),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrmaxi, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Maxf(a float32, b float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&a),
    gdc.ConstTypePtr(&b),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrmaxf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Min(vargs ...Variant) Variant {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }

  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrmin, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Mini(a int, b int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&a),
    gdc.ConstTypePtr(&b),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrmini, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Minf(a float32, b float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&a),
    gdc.ConstTypePtr(&b),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrminf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Clamp(value Variant, min Variant, max Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    value.AsCTypePtr(),
    min.AsCTypePtr(),
    max.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrclamp, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Clampi(value int, min int, max int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&value),
    gdc.ConstTypePtr(&min),
    gdc.ConstTypePtr(&max),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrclampi, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Clampf(value float32, min float32, max float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&value),
    gdc.ConstTypePtr(&min),
    gdc.ConstTypePtr(&max),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrclampf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) NearestPo2(value int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&value),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrnearest_po2, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Pingpong(value float32, length float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&value),
    gdc.ConstTypePtr(&length),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrpingpong, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Randomize()  {
  args := []gdc.ConstTypePtr{
  }

  me.iface.CallPtrUtilityFunction(me.ptrrandomize, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Randi() int {
  args := []gdc.ConstTypePtr{
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrandi, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Randf() float32 {
  args := []gdc.ConstTypePtr{
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrandf, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) RandiRange(from int, to int, ) int {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrandi_range, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) RandfRange(from float32, to float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&from),
    gdc.ConstTypePtr(&to),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrandf_range, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Randfn(mean float32, deviation float32, ) float32 {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&mean),
    gdc.ConstTypePtr(&deviation),
  }
  var ret float32
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrandfn, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Seed(base int, )  {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&base),
  }

  me.iface.CallPtrUtilityFunction(me.ptrseed, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) RandFromSeed(seed int, ) PackedInt64Array {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&seed),
  }
  var ret PackedInt64Array
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrand_from_seed, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Weakref(obj Variant, ) Variant {
  args := []gdc.ConstTypePtr{
    obj.AsCTypePtr(),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrweakref, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Typeof(variable Variant, ) int {
  args := []gdc.ConstTypePtr{
    variable.AsCTypePtr(),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrtypeof, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) TypeConvert(variant Variant, type_ int, ) Variant {
  args := []gdc.ConstTypePtr{
    variant.AsCTypePtr(),
    gdc.ConstTypePtr(&type_),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrtype_convert, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Str(vargs ...Variant) String {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }

  var ret String
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrstr, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) ErrorString(error int, ) String {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&error),
  }
  var ret String
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrerror_string, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) TypeString(type_ int, ) String {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&type_),
  }
  var ret String
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrtype_string, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Print(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrprint, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PrintRich(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrprint_rich, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Printerr(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrprinterr, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Printt(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrprintt, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Prints(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrprints, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Printraw(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrprintraw, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PrintVerbose(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrprint_verbose, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PushError(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrpush_error, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PushWarning(vargs ...Variant)  {
  args := make([]gdc.ConstTypePtr, len(vargs))
  for i, arg := range vargs {
    args[i] = arg.AsCTypePtr()
  }


  me.iface.CallPtrUtilityFunction(me.ptrpush_warning, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) VarToStr(variable Variant, ) String {
  args := []gdc.ConstTypePtr{
    variable.AsCTypePtr(),
  }
  var ret String
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrvar_to_str, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) StrToVar(string_ String, ) Variant {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&string_),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrstr_to_var, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) VarToBytes(variable Variant, ) PackedByteArray {
  args := []gdc.ConstTypePtr{
    variable.AsCTypePtr(),
  }
  var ret PackedByteArray
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrvar_to_bytes, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) BytesToVar(bytes PackedByteArray, ) Variant {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&bytes),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrbytes_to_var, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) VarToBytesWithObjects(variable Variant, ) PackedByteArray {
  args := []gdc.ConstTypePtr{
    variable.AsCTypePtr(),
  }
  var ret PackedByteArray
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrvar_to_bytes_with_objects, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) BytesToVarWithObjects(bytes PackedByteArray, ) Variant {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&bytes),
  }
  var ret Variant
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrbytes_to_var_with_objects, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) Hash(variable Variant, ) int {
  args := []gdc.ConstTypePtr{
    variable.AsCTypePtr(),
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrhash, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) InstanceFromId(instance_id int, ) Object {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&instance_id),
  }
  var ret Object
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrinstance_from_id, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsInstanceIdValid(id int, ) bool {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&id),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_instance_id_valid, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsInstanceValid(instance Variant, ) bool {
  args := []gdc.ConstTypePtr{
    instance.AsCTypePtr(),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_instance_valid, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) RidAllocateId() int {
  args := []gdc.ConstTypePtr{
  }
  var ret int
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrid_allocate_id, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) RidFromInt64(base int, ) RID {
  args := []gdc.ConstTypePtr{
    gdc.ConstTypePtr(&base),
  }
  var ret RID
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptrrid_from_int64, retPtr, unsafe.SliceData(args), len(args))
  return ret
}

func (me *utilities) IsSame(a Variant, b Variant, ) bool {
  args := []gdc.ConstTypePtr{
    a.AsCTypePtr(),
    b.AsCTypePtr(),
  }
  var ret bool
  retPtr := gdc.TypePtr(&ret)
  me.iface.CallPtrUtilityFunction(me.ptris_same, retPtr, unsafe.SliceData(args), len(args))
  return ret
}