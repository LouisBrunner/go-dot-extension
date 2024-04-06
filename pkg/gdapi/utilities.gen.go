// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UtilityFunctions interface {
	Sin(angle_rad float64) float64
	Cos(angle_rad float64) float64
	Tan(angle_rad float64) float64
	Sinh(x float64) float64
	Cosh(x float64) float64
	Tanh(x float64) float64
	Asin(x float64) float64
	Acos(x float64) float64
	Atan(x float64) float64
	Atan2(y float64, x float64) float64
	Asinh(x float64) float64
	Acosh(x float64) float64
	Atanh(x float64) float64
	Sqrt(x float64) float64
	Fmod(x float64, y float64) float64
	Fposmod(x float64, y float64) float64
	Posmod(x int64, y int64) int64
	Floor(x Variant) Variant
	Floorf(x float64) float64
	Floori(x float64) int64
	Ceil(x Variant) Variant
	Ceilf(x float64) float64
	Ceili(x float64) int64
	Round(x Variant) Variant
	Roundf(x float64) float64
	Roundi(x float64) int64
	Abs(x Variant) Variant
	Absf(x float64) float64
	Absi(x int64) int64
	Sign(x Variant) Variant
	Signf(x float64) float64
	Signi(x int64) int64
	Snapped(x Variant, step Variant) Variant
	Snappedf(x float64, step float64) float64
	Snappedi(x float64, step int64) int64
	Pow(base float64, exp float64) float64
	Log(x float64) float64
	Exp(x float64) float64
	IsNan(x float64) bool
	IsInf(x float64) bool
	IsEqualApprox(a float64, b float64) bool
	IsZeroApprox(x float64) bool
	IsFinite(x float64) bool
	Ease(x float64, curve float64) float64
	StepDecimals(x float64) int64
	Lerp(from Variant, to Variant, weight Variant) Variant
	Lerpf(from float64, to float64, weight float64) float64
	CubicInterpolate(from float64, to float64, pre float64, post float64, weight float64) float64
	CubicInterpolateAngle(from float64, to float64, pre float64, post float64, weight float64) float64
	CubicInterpolateInTime(from float64, to float64, pre float64, post float64, weight float64, to_t float64, pre_t float64, post_t float64) float64
	CubicInterpolateAngleInTime(from float64, to float64, pre float64, post float64, weight float64, to_t float64, pre_t float64, post_t float64) float64
	BezierInterpolate(start float64, control_1 float64, control_2 float64, end float64, t float64) float64
	BezierDerivative(start float64, control_1 float64, control_2 float64, end float64, t float64) float64
	AngleDifference(from float64, to float64) float64
	LerpAngle(from float64, to float64, weight float64) float64
	InverseLerp(from float64, to float64, weight float64) float64
	Remap(value float64, istart float64, istop float64, ostart float64, ostop float64) float64
	Smoothstep(from float64, to float64, x float64) float64
	MoveToward(from float64, to float64, delta float64) float64
	RotateToward(from float64, to float64, delta float64) float64
	DegToRad(deg float64) float64
	RadToDeg(rad float64) float64
	LinearToDb(lin float64) float64
	DbToLinear(db float64) float64
	Wrap(value Variant, min Variant, max Variant) Variant
	Wrapi(value int64, min int64, max int64) int64
	Wrapf(value float64, min float64, max float64) float64
	Max(args ...Variant) Variant
	Maxi(a int64, b int64) int64
	Maxf(a float64, b float64) float64
	Min(args ...Variant) Variant
	Mini(a int64, b int64) int64
	Minf(a float64, b float64) float64
	Clamp(value Variant, min Variant, max Variant) Variant
	Clampi(value int64, min int64, max int64) int64
	Clampf(value float64, min float64, max float64) float64
	NearestPo2(value int64) int64
	Pingpong(value float64, length float64) float64
	Randomize()
	Randi() int64
	Randf() float64
	RandiRange(from int64, to int64) int64
	RandfRange(from float64, to float64) float64
	Randfn(mean float64, deviation float64) float64
	Seed(base int64)
	RandFromSeed(seed int64) PackedInt64Array
	Weakref(obj Variant) Variant
	Typeof(variable Variant) int64
	TypeConvert(variant Variant, type_ int64) Variant
	Str(args ...Variant) String
	ErrorString(error_ int64) String
	TypeString(type_ int64) String
	Print(args ...Variant)
	PrintRich(args ...Variant)
	Printerr(args ...Variant)
	Printt(args ...Variant)
	Prints(args ...Variant)
	Printraw(args ...Variant)
	PrintVerbose(args ...Variant)
	PushError(args ...Variant)
	PushWarning(args ...Variant)
	VarToStr(variable Variant) String
	StrToVar(string_ String) Variant
	VarToBytes(variable Variant) PackedByteArray
	BytesToVar(bytes PackedByteArray) Variant
	VarToBytesWithObjects(variable Variant) PackedByteArray
	BytesToVarWithObjects(bytes PackedByteArray) Variant
	Hash(variable Variant) int64
	InstanceFromId(instance_id int64) Object
	IsInstanceIdValid(id int64) bool
	IsInstanceValid(instance Variant) bool
	RidAllocateId() int64
	RidFromInt64(base int64) RID
	IsSame(a Variant, b Variant) bool
}

type utilities struct {
	iface                              gdc.Interface
	ptrsin                             gdc.PtrUtilityFunction
	ptrcos                             gdc.PtrUtilityFunction
	ptrtan                             gdc.PtrUtilityFunction
	ptrsinh                            gdc.PtrUtilityFunction
	ptrcosh                            gdc.PtrUtilityFunction
	ptrtanh                            gdc.PtrUtilityFunction
	ptrasin                            gdc.PtrUtilityFunction
	ptracos                            gdc.PtrUtilityFunction
	ptratan                            gdc.PtrUtilityFunction
	ptratan2                           gdc.PtrUtilityFunction
	ptrasinh                           gdc.PtrUtilityFunction
	ptracosh                           gdc.PtrUtilityFunction
	ptratanh                           gdc.PtrUtilityFunction
	ptrsqrt                            gdc.PtrUtilityFunction
	ptrfmod                            gdc.PtrUtilityFunction
	ptrfposmod                         gdc.PtrUtilityFunction
	ptrposmod                          gdc.PtrUtilityFunction
	ptrfloor                           gdc.PtrUtilityFunction
	ptrfloorf                          gdc.PtrUtilityFunction
	ptrfloori                          gdc.PtrUtilityFunction
	ptrceil                            gdc.PtrUtilityFunction
	ptrceilf                           gdc.PtrUtilityFunction
	ptrceili                           gdc.PtrUtilityFunction
	ptrround                           gdc.PtrUtilityFunction
	ptrroundf                          gdc.PtrUtilityFunction
	ptrroundi                          gdc.PtrUtilityFunction
	ptrabs                             gdc.PtrUtilityFunction
	ptrabsf                            gdc.PtrUtilityFunction
	ptrabsi                            gdc.PtrUtilityFunction
	ptrsign                            gdc.PtrUtilityFunction
	ptrsignf                           gdc.PtrUtilityFunction
	ptrsigni                           gdc.PtrUtilityFunction
	ptrsnapped                         gdc.PtrUtilityFunction
	ptrsnappedf                        gdc.PtrUtilityFunction
	ptrsnappedi                        gdc.PtrUtilityFunction
	ptrpow                             gdc.PtrUtilityFunction
	ptrlog                             gdc.PtrUtilityFunction
	ptrexp                             gdc.PtrUtilityFunction
	ptris_nan                          gdc.PtrUtilityFunction
	ptris_inf                          gdc.PtrUtilityFunction
	ptris_equal_approx                 gdc.PtrUtilityFunction
	ptris_zero_approx                  gdc.PtrUtilityFunction
	ptris_finite                       gdc.PtrUtilityFunction
	ptrease                            gdc.PtrUtilityFunction
	ptrstep_decimals                   gdc.PtrUtilityFunction
	ptrlerp                            gdc.PtrUtilityFunction
	ptrlerpf                           gdc.PtrUtilityFunction
	ptrcubic_interpolate               gdc.PtrUtilityFunction
	ptrcubic_interpolate_angle         gdc.PtrUtilityFunction
	ptrcubic_interpolate_in_time       gdc.PtrUtilityFunction
	ptrcubic_interpolate_angle_in_time gdc.PtrUtilityFunction
	ptrbezier_interpolate              gdc.PtrUtilityFunction
	ptrbezier_derivative               gdc.PtrUtilityFunction
	ptrangle_difference                gdc.PtrUtilityFunction
	ptrlerp_angle                      gdc.PtrUtilityFunction
	ptrinverse_lerp                    gdc.PtrUtilityFunction
	ptrremap                           gdc.PtrUtilityFunction
	ptrsmoothstep                      gdc.PtrUtilityFunction
	ptrmove_toward                     gdc.PtrUtilityFunction
	ptrrotate_toward                   gdc.PtrUtilityFunction
	ptrdeg_to_rad                      gdc.PtrUtilityFunction
	ptrrad_to_deg                      gdc.PtrUtilityFunction
	ptrlinear_to_db                    gdc.PtrUtilityFunction
	ptrdb_to_linear                    gdc.PtrUtilityFunction
	ptrwrap                            gdc.PtrUtilityFunction
	ptrwrapi                           gdc.PtrUtilityFunction
	ptrwrapf                           gdc.PtrUtilityFunction
	ptrmax                             gdc.PtrUtilityFunction
	ptrmaxi                            gdc.PtrUtilityFunction
	ptrmaxf                            gdc.PtrUtilityFunction
	ptrmin                             gdc.PtrUtilityFunction
	ptrmini                            gdc.PtrUtilityFunction
	ptrminf                            gdc.PtrUtilityFunction
	ptrclamp                           gdc.PtrUtilityFunction
	ptrclampi                          gdc.PtrUtilityFunction
	ptrclampf                          gdc.PtrUtilityFunction
	ptrnearest_po2                     gdc.PtrUtilityFunction
	ptrpingpong                        gdc.PtrUtilityFunction
	ptrrandomize                       gdc.PtrUtilityFunction
	ptrrandi                           gdc.PtrUtilityFunction
	ptrrandf                           gdc.PtrUtilityFunction
	ptrrandi_range                     gdc.PtrUtilityFunction
	ptrrandf_range                     gdc.PtrUtilityFunction
	ptrrandfn                          gdc.PtrUtilityFunction
	ptrseed                            gdc.PtrUtilityFunction
	ptrrand_from_seed                  gdc.PtrUtilityFunction
	ptrweakref                         gdc.PtrUtilityFunction
	ptrtypeof                          gdc.PtrUtilityFunction
	ptrtype_convert                    gdc.PtrUtilityFunction
	ptrstr                             gdc.PtrUtilityFunction
	ptrerror_string                    gdc.PtrUtilityFunction
	ptrtype_string                     gdc.PtrUtilityFunction
	ptrprint                           gdc.PtrUtilityFunction
	ptrprint_rich                      gdc.PtrUtilityFunction
	ptrprinterr                        gdc.PtrUtilityFunction
	ptrprintt                          gdc.PtrUtilityFunction
	ptrprints                          gdc.PtrUtilityFunction
	ptrprintraw                        gdc.PtrUtilityFunction
	ptrprint_verbose                   gdc.PtrUtilityFunction
	ptrpush_error                      gdc.PtrUtilityFunction
	ptrpush_warning                    gdc.PtrUtilityFunction
	ptrvar_to_str                      gdc.PtrUtilityFunction
	ptrstr_to_var                      gdc.PtrUtilityFunction
	ptrvar_to_bytes                    gdc.PtrUtilityFunction
	ptrbytes_to_var                    gdc.PtrUtilityFunction
	ptrvar_to_bytes_with_objects       gdc.PtrUtilityFunction
	ptrbytes_to_var_with_objects       gdc.PtrUtilityFunction
	ptrhash                            gdc.PtrUtilityFunction
	ptrinstance_from_id                gdc.PtrUtilityFunction
	ptris_instance_id_valid            gdc.PtrUtilityFunction
	ptris_instance_valid               gdc.PtrUtilityFunction
	ptrrid_allocate_id                 gdc.PtrUtilityFunction
	ptrrid_from_int64                  gdc.PtrUtilityFunction
	ptris_same                         gdc.PtrUtilityFunction
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
		ptrsin:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strsin.AsCPtr(), 2140049587)),
		ptrcos:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strcos.AsCPtr(), 2140049587)),
		ptrtan:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strtan.AsCPtr(), 2140049587)),
		ptrsinh:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strsinh.AsCPtr(), 2140049587)),
		ptrcosh:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strcosh.AsCPtr(), 2140049587)),
		ptrtanh:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strtanh.AsCPtr(), 2140049587)),
		ptrasin:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strasin.AsCPtr(), 2140049587)),
		ptracos:                            ensurePtr(iface.VariantGetPtrUtilityFunction(stracos.AsCPtr(), 2140049587)),
		ptratan:                            ensurePtr(iface.VariantGetPtrUtilityFunction(stratan.AsCPtr(), 2140049587)),
		ptratan2:                           ensurePtr(iface.VariantGetPtrUtilityFunction(stratan2.AsCPtr(), 92296394)),
		ptrasinh:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strasinh.AsCPtr(), 2140049587)),
		ptracosh:                           ensurePtr(iface.VariantGetPtrUtilityFunction(stracosh.AsCPtr(), 2140049587)),
		ptratanh:                           ensurePtr(iface.VariantGetPtrUtilityFunction(stratanh.AsCPtr(), 2140049587)),
		ptrsqrt:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strsqrt.AsCPtr(), 2140049587)),
		ptrfmod:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strfmod.AsCPtr(), 92296394)),
		ptrfposmod:                         ensurePtr(iface.VariantGetPtrUtilityFunction(strfposmod.AsCPtr(), 92296394)),
		ptrposmod:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strposmod.AsCPtr(), 3133453818)),
		ptrfloor:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strfloor.AsCPtr(), 4776452)),
		ptrfloorf:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strfloorf.AsCPtr(), 2140049587)),
		ptrfloori:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strfloori.AsCPtr(), 2780425386)),
		ptrceil:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strceil.AsCPtr(), 4776452)),
		ptrceilf:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strceilf.AsCPtr(), 2140049587)),
		ptrceili:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strceili.AsCPtr(), 2780425386)),
		ptrround:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strround.AsCPtr(), 4776452)),
		ptrroundf:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strroundf.AsCPtr(), 2140049587)),
		ptrroundi:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strroundi.AsCPtr(), 2780425386)),
		ptrabs:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strabs.AsCPtr(), 4776452)),
		ptrabsf:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strabsf.AsCPtr(), 2140049587)),
		ptrabsi:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strabsi.AsCPtr(), 2157319888)),
		ptrsign:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strsign.AsCPtr(), 4776452)),
		ptrsignf:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strsignf.AsCPtr(), 2140049587)),
		ptrsigni:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strsigni.AsCPtr(), 2157319888)),
		ptrsnapped:                         ensurePtr(iface.VariantGetPtrUtilityFunction(strsnapped.AsCPtr(), 459914704)),
		ptrsnappedf:                        ensurePtr(iface.VariantGetPtrUtilityFunction(strsnappedf.AsCPtr(), 92296394)),
		ptrsnappedi:                        ensurePtr(iface.VariantGetPtrUtilityFunction(strsnappedi.AsCPtr(), 3570758393)),
		ptrpow:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strpow.AsCPtr(), 92296394)),
		ptrlog:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strlog.AsCPtr(), 2140049587)),
		ptrexp:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strexp.AsCPtr(), 2140049587)),
		ptris_nan:                          ensurePtr(iface.VariantGetPtrUtilityFunction(stris_nan.AsCPtr(), 3569215213)),
		ptris_inf:                          ensurePtr(iface.VariantGetPtrUtilityFunction(stris_inf.AsCPtr(), 3569215213)),
		ptris_equal_approx:                 ensurePtr(iface.VariantGetPtrUtilityFunction(stris_equal_approx.AsCPtr(), 1400789633)),
		ptris_zero_approx:                  ensurePtr(iface.VariantGetPtrUtilityFunction(stris_zero_approx.AsCPtr(), 3569215213)),
		ptris_finite:                       ensurePtr(iface.VariantGetPtrUtilityFunction(stris_finite.AsCPtr(), 3569215213)),
		ptrease:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strease.AsCPtr(), 92296394)),
		ptrstep_decimals:                   ensurePtr(iface.VariantGetPtrUtilityFunction(strstep_decimals.AsCPtr(), 2780425386)),
		ptrlerp:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strlerp.AsCPtr(), 3389874542)),
		ptrlerpf:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strlerpf.AsCPtr(), 998901048)),
		ptrcubic_interpolate:               ensurePtr(iface.VariantGetPtrUtilityFunction(strcubic_interpolate.AsCPtr(), 1090965791)),
		ptrcubic_interpolate_angle:         ensurePtr(iface.VariantGetPtrUtilityFunction(strcubic_interpolate_angle.AsCPtr(), 1090965791)),
		ptrcubic_interpolate_in_time:       ensurePtr(iface.VariantGetPtrUtilityFunction(strcubic_interpolate_in_time.AsCPtr(), 388121036)),
		ptrcubic_interpolate_angle_in_time: ensurePtr(iface.VariantGetPtrUtilityFunction(strcubic_interpolate_angle_in_time.AsCPtr(), 388121036)),
		ptrbezier_interpolate:              ensurePtr(iface.VariantGetPtrUtilityFunction(strbezier_interpolate.AsCPtr(), 1090965791)),
		ptrbezier_derivative:               ensurePtr(iface.VariantGetPtrUtilityFunction(strbezier_derivative.AsCPtr(), 1090965791)),
		ptrangle_difference:                ensurePtr(iface.VariantGetPtrUtilityFunction(strangle_difference.AsCPtr(), 92296394)),
		ptrlerp_angle:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strlerp_angle.AsCPtr(), 998901048)),
		ptrinverse_lerp:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strinverse_lerp.AsCPtr(), 998901048)),
		ptrremap:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strremap.AsCPtr(), 1090965791)),
		ptrsmoothstep:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strsmoothstep.AsCPtr(), 998901048)),
		ptrmove_toward:                     ensurePtr(iface.VariantGetPtrUtilityFunction(strmove_toward.AsCPtr(), 998901048)),
		ptrrotate_toward:                   ensurePtr(iface.VariantGetPtrUtilityFunction(strrotate_toward.AsCPtr(), 998901048)),
		ptrdeg_to_rad:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strdeg_to_rad.AsCPtr(), 2140049587)),
		ptrrad_to_deg:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strrad_to_deg.AsCPtr(), 2140049587)),
		ptrlinear_to_db:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strlinear_to_db.AsCPtr(), 2140049587)),
		ptrdb_to_linear:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strdb_to_linear.AsCPtr(), 2140049587)),
		ptrwrap:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strwrap.AsCPtr(), 3389874542)),
		ptrwrapi:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strwrapi.AsCPtr(), 650295447)),
		ptrwrapf:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strwrapf.AsCPtr(), 998901048)),
		ptrmax:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strmax.AsCPtr(), 3896050336)),
		ptrmaxi:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strmaxi.AsCPtr(), 3133453818)),
		ptrmaxf:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strmaxf.AsCPtr(), 92296394)),
		ptrmin:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strmin.AsCPtr(), 3896050336)),
		ptrmini:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strmini.AsCPtr(), 3133453818)),
		ptrminf:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strminf.AsCPtr(), 92296394)),
		ptrclamp:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strclamp.AsCPtr(), 3389874542)),
		ptrclampi:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strclampi.AsCPtr(), 650295447)),
		ptrclampf:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strclampf.AsCPtr(), 998901048)),
		ptrnearest_po2:                     ensurePtr(iface.VariantGetPtrUtilityFunction(strnearest_po2.AsCPtr(), 2157319888)),
		ptrpingpong:                        ensurePtr(iface.VariantGetPtrUtilityFunction(strpingpong.AsCPtr(), 92296394)),
		ptrrandomize:                       ensurePtr(iface.VariantGetPtrUtilityFunction(strrandomize.AsCPtr(), 1691721052)),
		ptrrandi:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strrandi.AsCPtr(), 701202648)),
		ptrrandf:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strrandf.AsCPtr(), 2086227845)),
		ptrrandi_range:                     ensurePtr(iface.VariantGetPtrUtilityFunction(strrandi_range.AsCPtr(), 3133453818)),
		ptrrandf_range:                     ensurePtr(iface.VariantGetPtrUtilityFunction(strrandf_range.AsCPtr(), 92296394)),
		ptrrandfn:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strrandfn.AsCPtr(), 92296394)),
		ptrseed:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strseed.AsCPtr(), 382931173)),
		ptrrand_from_seed:                  ensurePtr(iface.VariantGetPtrUtilityFunction(strrand_from_seed.AsCPtr(), 1391063685)),
		ptrweakref:                         ensurePtr(iface.VariantGetPtrUtilityFunction(strweakref.AsCPtr(), 4776452)),
		ptrtypeof:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strtypeof.AsCPtr(), 326422594)),
		ptrtype_convert:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strtype_convert.AsCPtr(), 2453062746)),
		ptrstr:                             ensurePtr(iface.VariantGetPtrUtilityFunction(strstr.AsCPtr(), 32569176)),
		ptrerror_string:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strerror_string.AsCPtr(), 942708242)),
		ptrtype_string:                     ensurePtr(iface.VariantGetPtrUtilityFunction(strtype_string.AsCPtr(), 942708242)),
		ptrprint:                           ensurePtr(iface.VariantGetPtrUtilityFunction(strprint.AsCPtr(), 2648703342)),
		ptrprint_rich:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strprint_rich.AsCPtr(), 2648703342)),
		ptrprinterr:                        ensurePtr(iface.VariantGetPtrUtilityFunction(strprinterr.AsCPtr(), 2648703342)),
		ptrprintt:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strprintt.AsCPtr(), 2648703342)),
		ptrprints:                          ensurePtr(iface.VariantGetPtrUtilityFunction(strprints.AsCPtr(), 2648703342)),
		ptrprintraw:                        ensurePtr(iface.VariantGetPtrUtilityFunction(strprintraw.AsCPtr(), 2648703342)),
		ptrprint_verbose:                   ensurePtr(iface.VariantGetPtrUtilityFunction(strprint_verbose.AsCPtr(), 2648703342)),
		ptrpush_error:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strpush_error.AsCPtr(), 2648703342)),
		ptrpush_warning:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strpush_warning.AsCPtr(), 2648703342)),
		ptrvar_to_str:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strvar_to_str.AsCPtr(), 866625479)),
		ptrstr_to_var:                      ensurePtr(iface.VariantGetPtrUtilityFunction(strstr_to_var.AsCPtr(), 1891498491)),
		ptrvar_to_bytes:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strvar_to_bytes.AsCPtr(), 2947269930)),
		ptrbytes_to_var:                    ensurePtr(iface.VariantGetPtrUtilityFunction(strbytes_to_var.AsCPtr(), 4249819452)),
		ptrvar_to_bytes_with_objects:       ensurePtr(iface.VariantGetPtrUtilityFunction(strvar_to_bytes_with_objects.AsCPtr(), 2947269930)),
		ptrbytes_to_var_with_objects:       ensurePtr(iface.VariantGetPtrUtilityFunction(strbytes_to_var_with_objects.AsCPtr(), 4249819452)),
		ptrhash:                            ensurePtr(iface.VariantGetPtrUtilityFunction(strhash.AsCPtr(), 326422594)),
		ptrinstance_from_id:                ensurePtr(iface.VariantGetPtrUtilityFunction(strinstance_from_id.AsCPtr(), 1156694636)),
		ptris_instance_id_valid:            ensurePtr(iface.VariantGetPtrUtilityFunction(stris_instance_id_valid.AsCPtr(), 2232439758)),
		ptris_instance_valid:               ensurePtr(iface.VariantGetPtrUtilityFunction(stris_instance_valid.AsCPtr(), 996128841)),
		ptrrid_allocate_id:                 ensurePtr(iface.VariantGetPtrUtilityFunction(strrid_allocate_id.AsCPtr(), 701202648)),
		ptrrid_from_int64:                  ensurePtr(iface.VariantGetPtrUtilityFunction(strrid_from_int64.AsCPtr(), 3426892196)),
		ptris_same:                         ensurePtr(iface.VariantGetPtrUtilityFunction(stris_same.AsCPtr(), 1409423524)),
	}
}

func (me *utilities) Sin(angle_rad float64) float64 {
	varg0 := NewFloatFromFloat32(angle_rad)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrsin, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Cos(angle_rad float64) float64 {
	varg0 := NewFloatFromFloat32(angle_rad)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrcos, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Tan(angle_rad float64) float64 {
	varg0 := NewFloatFromFloat32(angle_rad)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrtan, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Sinh(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrsinh, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Cosh(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrcosh, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Tanh(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrtanh, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Asin(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrasin, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Acos(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptracos, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Atan(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptratan, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Atan2(y float64, x float64) float64 {
	varg0 := NewFloatFromFloat32(y)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(x)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptratan2, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Asinh(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrasinh, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Acosh(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptracosh, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Atanh(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptratanh, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Sqrt(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrsqrt, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Fmod(x float64, y float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(y)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrfmod, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Fposmod(x float64, y float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(y)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrfposmod, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Posmod(x int64, y int64) int64 {
	varg0 := NewIntFromInt(x)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(y)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrposmod, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Floor(x Variant) Variant {

	args := []gdc.ConstTypePtr{x.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrfloor, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Floorf(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrfloorf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Floori(x float64) int64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrfloori, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Ceil(x Variant) Variant {

	args := []gdc.ConstTypePtr{x.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrceil, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Ceilf(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrceilf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Ceili(x float64) int64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrceili, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Round(x Variant) Variant {

	args := []gdc.ConstTypePtr{x.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrround, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Roundf(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrroundf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Roundi(x float64) int64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrroundi, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Abs(x Variant) Variant {

	args := []gdc.ConstTypePtr{x.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrabs, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Absf(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrabsf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Absi(x int64) int64 {
	varg0 := NewIntFromInt(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrabsi, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Sign(x Variant) Variant {

	args := []gdc.ConstTypePtr{x.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrsign, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Signf(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrsignf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Signi(x int64) int64 {
	varg0 := NewIntFromInt(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrsigni, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Snapped(x Variant, step Variant) Variant {

	args := []gdc.ConstTypePtr{x.AsCTypePtr(), step.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrsnapped, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Snappedf(x float64, step float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(step)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrsnappedf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Snappedi(x float64, step int64) int64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(step)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrsnappedi, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Pow(base float64, exp float64) float64 {
	varg0 := NewFloatFromFloat32(base)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(exp)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrpow, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Log(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrlog, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Exp(x float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrexp, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) IsNan(x float64) bool {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_nan, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) IsInf(x float64) bool {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_inf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) IsEqualApprox(a float64, b float64) bool {
	varg0 := NewFloatFromFloat32(a)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(b)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_equal_approx, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) IsZeroApprox(x float64) bool {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_zero_approx, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) IsFinite(x float64) bool {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_finite, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Ease(x float64, curve float64) float64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(curve)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrease, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) StepDecimals(x float64) int64 {
	varg0 := NewFloatFromFloat32(x)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrstep_decimals, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Lerp(from Variant, to Variant, weight Variant) Variant {

	args := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), weight.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrlerp, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Lerpf(from float64, to float64, weight float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(weight)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrlerpf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) CubicInterpolate(from float64, to float64, pre float64, post float64, weight float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(pre)
	defer varg2.Destroy()
	varg3 := NewFloatFromFloat32(post)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(weight)
	defer varg4.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) CubicInterpolateAngle(from float64, to float64, pre float64, post float64, weight float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(pre)
	defer varg2.Destroy()
	varg3 := NewFloatFromFloat32(post)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(weight)
	defer varg4.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate_angle, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) CubicInterpolateInTime(from float64, to float64, pre float64, post float64, weight float64, to_t float64, pre_t float64, post_t float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(pre)
	defer varg2.Destroy()
	varg3 := NewFloatFromFloat32(post)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(weight)
	defer varg4.Destroy()
	varg5 := NewFloatFromFloat32(to_t)
	defer varg5.Destroy()
	varg6 := NewFloatFromFloat32(pre_t)
	defer varg6.Destroy()
	varg7 := NewFloatFromFloat32(post_t)
	defer varg7.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr(), varg7.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate_in_time, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) CubicInterpolateAngleInTime(from float64, to float64, pre float64, post float64, weight float64, to_t float64, pre_t float64, post_t float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(pre)
	defer varg2.Destroy()
	varg3 := NewFloatFromFloat32(post)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(weight)
	defer varg4.Destroy()
	varg5 := NewFloatFromFloat32(to_t)
	defer varg5.Destroy()
	varg6 := NewFloatFromFloat32(pre_t)
	defer varg6.Destroy()
	varg7 := NewFloatFromFloat32(post_t)
	defer varg7.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr(), varg7.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrcubic_interpolate_angle_in_time, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) BezierInterpolate(start float64, control_1 float64, control_2 float64, end float64, t float64) float64 {
	varg0 := NewFloatFromFloat32(start)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(control_1)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(control_2)
	defer varg2.Destroy()
	varg3 := NewFloatFromFloat32(end)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(t)
	defer varg4.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrbezier_interpolate, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) BezierDerivative(start float64, control_1 float64, control_2 float64, end float64, t float64) float64 {
	varg0 := NewFloatFromFloat32(start)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(control_1)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(control_2)
	defer varg2.Destroy()
	varg3 := NewFloatFromFloat32(end)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(t)
	defer varg4.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrbezier_derivative, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) AngleDifference(from float64, to float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrangle_difference, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) LerpAngle(from float64, to float64, weight float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(weight)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrlerp_angle, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) InverseLerp(from float64, to float64, weight float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(weight)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrinverse_lerp, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Remap(value float64, istart float64, istop float64, ostart float64, ostop float64) float64 {
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(istart)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(istop)
	defer varg2.Destroy()
	varg3 := NewFloatFromFloat32(ostart)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(ostop)
	defer varg4.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrremap, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Smoothstep(from float64, to float64, x float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(x)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrsmoothstep, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) MoveToward(from float64, to float64, delta float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(delta)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrmove_toward, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) RotateToward(from float64, to float64, delta float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(delta)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrrotate_toward, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) DegToRad(deg float64) float64 {
	varg0 := NewFloatFromFloat32(deg)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrdeg_to_rad, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) RadToDeg(rad float64) float64 {
	varg0 := NewFloatFromFloat32(rad)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrrad_to_deg, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) LinearToDb(lin float64) float64 {
	varg0 := NewFloatFromFloat32(lin)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrlinear_to_db, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) DbToLinear(db float64) float64 {
	varg0 := NewFloatFromFloat32(db)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrdb_to_linear, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Wrap(value Variant, min Variant, max Variant) Variant {

	args := []gdc.ConstTypePtr{value.AsCTypePtr(), min.AsCTypePtr(), max.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrwrap, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Wrapi(value int64, min int64, max int64) int64 {
	varg0 := NewIntFromInt(value)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(min)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(max)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrwrapi, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Wrapf(value float64, min float64, max float64) float64 {
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(min)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(max)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrwrapf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Max(vargs ...Variant) Variant {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrmax, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Maxi(a int64, b int64) int64 {
	varg0 := NewIntFromInt(a)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(b)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrmaxi, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Maxf(a float64, b float64) float64 {
	varg0 := NewFloatFromFloat32(a)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(b)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrmaxf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Min(vargs ...Variant) Variant {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrmin, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Mini(a int64, b int64) int64 {
	varg0 := NewIntFromInt(a)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(b)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrmini, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Minf(a float64, b float64) float64 {
	varg0 := NewFloatFromFloat32(a)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(b)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrminf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Clamp(value Variant, min Variant, max Variant) Variant {

	args := []gdc.ConstTypePtr{value.AsCTypePtr(), min.AsCTypePtr(), max.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrclamp, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Clampi(value int64, min int64, max int64) int64 {
	varg0 := NewIntFromInt(value)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(min)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(max)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrclampi, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Clampf(value float64, min float64, max float64) float64 {
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(min)
	defer varg1.Destroy()
	varg2 := NewFloatFromFloat32(max)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrclampf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) NearestPo2(value int64) int64 {
	varg0 := NewIntFromInt(value)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrnearest_po2, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Pingpong(value float64, length float64) float64 {
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(length)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrpingpong, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Randomize() {
	args := []gdc.ConstTypePtr{}
	me.iface.CallPtrUtilityFunction(me.ptrrandomize, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Randi() int64 {
	args := []gdc.ConstTypePtr{}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrrandi, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Randf() float64 {
	args := []gdc.ConstTypePtr{}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrrandf, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) RandiRange(from int64, to int64) int64 {
	varg0 := NewIntFromInt(from)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(to)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrrandi_range, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) RandfRange(from float64, to float64) float64 {
	varg0 := NewFloatFromFloat32(from)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(to)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrrandf_range, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Randfn(mean float64, deviation float64) float64 {
	varg0 := NewFloatFromFloat32(mean)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(deviation)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewFloat()
	me.iface.CallPtrUtilityFunction(me.ptrrandfn, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) Seed(base int64) {
	varg0 := NewIntFromInt(base)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	me.iface.CallPtrUtilityFunction(me.ptrseed, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) RandFromSeed(seed int64) PackedInt64Array {
	varg0 := NewIntFromInt(seed)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewPackedInt64Array()
	me.iface.CallPtrUtilityFunction(me.ptrrand_from_seed, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Weakref(obj Variant) Variant {

	args := []gdc.ConstTypePtr{obj.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrweakref, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Typeof(variable Variant) int64 {

	args := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrtypeof, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) TypeConvert(variant Variant, type_ int64) Variant {

	varg1 := NewIntFromInt(type_)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{variant.AsCTypePtr(), varg1.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrtype_convert, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Str(vargs ...Variant) String {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	ret := NewString()
	me.iface.CallPtrUtilityFunction(me.ptrstr, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) ErrorString(error_ int64) String {
	varg0 := NewIntFromInt(error_)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewString()
	me.iface.CallPtrUtilityFunction(me.ptrerror_string, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) TypeString(type_ int64) String {
	varg0 := NewIntFromInt(type_)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewString()
	me.iface.CallPtrUtilityFunction(me.ptrtype_string, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Print(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrprint, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PrintRich(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrprint_rich, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Printerr(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrprinterr, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Printt(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrprintt, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Prints(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrprints, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) Printraw(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrprintraw, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PrintVerbose(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrprint_verbose, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PushError(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrpush_error, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) PushWarning(vargs ...Variant) {
	args := make([]gdc.ConstTypePtr, len(vargs))
	for i, arg := range vargs {
		args[i] = arg.AsCTypePtr()
	}

	me.iface.CallPtrUtilityFunction(me.ptrpush_warning, nil, unsafe.SliceData(args), len(args))

}

func (me *utilities) VarToStr(variable Variant) String {

	args := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	ret := NewString()
	me.iface.CallPtrUtilityFunction(me.ptrvar_to_str, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) StrToVar(string_ String) Variant {

	args := []gdc.ConstTypePtr{string_.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrstr_to_var, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) VarToBytes(variable Variant) PackedByteArray {

	args := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	ret := NewPackedByteArray()
	me.iface.CallPtrUtilityFunction(me.ptrvar_to_bytes, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) BytesToVar(bytes PackedByteArray) Variant {

	args := []gdc.ConstTypePtr{bytes.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrbytes_to_var, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) VarToBytesWithObjects(variable Variant) PackedByteArray {

	args := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	ret := NewPackedByteArray()
	me.iface.CallPtrUtilityFunction(me.ptrvar_to_bytes_with_objects, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) BytesToVarWithObjects(bytes PackedByteArray) Variant {

	args := []gdc.ConstTypePtr{bytes.AsCTypePtr()}
	ret := NewVariant()
	me.iface.CallPtrUtilityFunction(me.ptrbytes_to_var_with_objects, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) Hash(variable Variant) int64 {

	args := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrhash, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) InstanceFromId(instance_id int64) Object {
	varg0 := NewIntFromInt(instance_id)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewObject()
	me.iface.CallPtrUtilityFunction(me.ptrinstance_from_id, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) IsInstanceIdValid(id int64) bool {
	varg0 := NewIntFromInt(id)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_instance_id_valid, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) IsInstanceValid(instance Variant) bool {

	args := []gdc.ConstTypePtr{instance.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_instance_valid, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) RidAllocateId() int64 {
	args := []gdc.ConstTypePtr{}
	ret := NewInt()
	me.iface.CallPtrUtilityFunction(me.ptrrid_allocate_id, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}

func (me *utilities) RidFromInt64(base int64) RID {
	varg0 := NewIntFromInt(base)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	ret := NewRID()
	me.iface.CallPtrUtilityFunction(me.ptrrid_from_int64, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return *ret
}

func (me *utilities) IsSame(a Variant, b Variant) bool {

	args := []gdc.ConstTypePtr{a.AsCTypePtr(), b.AsCTypePtr()}
	ret := NewBool()
	me.iface.CallPtrUtilityFunction(me.ptris_same, ret.AsTypePtr(), unsafe.SliceData(args), len(args))
	return ret.Get()
}
