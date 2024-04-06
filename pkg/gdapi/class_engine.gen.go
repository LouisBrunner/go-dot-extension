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

type ptrsForEngineList struct {
	fnSetPhysicsTicksPerSecond        gdc.MethodBindPtr
	fnGetPhysicsTicksPerSecond        gdc.MethodBindPtr
	fnSetMaxPhysicsStepsPerFrame      gdc.MethodBindPtr
	fnGetMaxPhysicsStepsPerFrame      gdc.MethodBindPtr
	fnSetPhysicsJitterFix             gdc.MethodBindPtr
	fnGetPhysicsJitterFix             gdc.MethodBindPtr
	fnGetPhysicsInterpolationFraction gdc.MethodBindPtr
	fnSetMaxFps                       gdc.MethodBindPtr
	fnGetMaxFps                       gdc.MethodBindPtr
	fnSetTimeScale                    gdc.MethodBindPtr
	fnGetTimeScale                    gdc.MethodBindPtr
	fnGetFramesDrawn                  gdc.MethodBindPtr
	fnGetFramesPerSecond              gdc.MethodBindPtr
	fnGetPhysicsFrames                gdc.MethodBindPtr
	fnGetProcessFrames                gdc.MethodBindPtr
	fnGetMainLoop                     gdc.MethodBindPtr
	fnGetVersionInfo                  gdc.MethodBindPtr
	fnGetAuthorInfo                   gdc.MethodBindPtr
	fnGetCopyrightInfo                gdc.MethodBindPtr
	fnGetDonorInfo                    gdc.MethodBindPtr
	fnGetLicenseInfo                  gdc.MethodBindPtr
	fnGetLicenseText                  gdc.MethodBindPtr
	fnGetArchitectureName             gdc.MethodBindPtr
	fnIsInPhysicsFrame                gdc.MethodBindPtr
	fnHasSingleton                    gdc.MethodBindPtr
	fnGetSingleton                    gdc.MethodBindPtr
	fnRegisterSingleton               gdc.MethodBindPtr
	fnUnregisterSingleton             gdc.MethodBindPtr
	fnGetSingletonList                gdc.MethodBindPtr
	fnRegisterScriptLanguage          gdc.MethodBindPtr
	fnUnregisterScriptLanguage        gdc.MethodBindPtr
	fnGetScriptLanguageCount          gdc.MethodBindPtr
	fnGetScriptLanguage               gdc.MethodBindPtr
	fnIsEditorHint                    gdc.MethodBindPtr
	fnGetWriteMoviePath               gdc.MethodBindPtr
	fnSetPrintErrorMessages           gdc.MethodBindPtr
	fnIsPrintingErrorMessages         gdc.MethodBindPtr
}

var ptrsForEngine ptrsForEngineList

func initEnginePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Engine")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_physics_ticks_per_second")
		defer methodName.Destroy()
		ptrsForEngine.fnSetPhysicsTicksPerSecond = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_physics_ticks_per_second")
		defer methodName.Destroy()
		ptrsForEngine.fnGetPhysicsTicksPerSecond = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_physics_steps_per_frame")
		defer methodName.Destroy()
		ptrsForEngine.fnSetMaxPhysicsStepsPerFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_physics_steps_per_frame")
		defer methodName.Destroy()
		ptrsForEngine.fnGetMaxPhysicsStepsPerFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_physics_jitter_fix")
		defer methodName.Destroy()
		ptrsForEngine.fnSetPhysicsJitterFix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_physics_jitter_fix")
		defer methodName.Destroy()
		ptrsForEngine.fnGetPhysicsJitterFix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_physics_interpolation_fraction")
		defer methodName.Destroy()
		ptrsForEngine.fnGetPhysicsInterpolationFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_fps")
		defer methodName.Destroy()
		ptrsForEngine.fnSetMaxFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_fps")
		defer methodName.Destroy()
		ptrsForEngine.fnGetMaxFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_time_scale")
		defer methodName.Destroy()
		ptrsForEngine.fnSetTimeScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_time_scale")
		defer methodName.Destroy()
		ptrsForEngine.fnGetTimeScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("get_frames_drawn")
		defer methodName.Destroy()
		ptrsForEngine.fnGetFramesDrawn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_frames_per_second")
		defer methodName.Destroy()
		ptrsForEngine.fnGetFramesPerSecond = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_physics_frames")
		defer methodName.Destroy()
		ptrsForEngine.fnGetPhysicsFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_process_frames")
		defer methodName.Destroy()
		ptrsForEngine.fnGetProcessFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_main_loop")
		defer methodName.Destroy()
		ptrsForEngine.fnGetMainLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1016888095))
	}
	{
		methodName := StringNameFromStr("get_version_info")
		defer methodName.Destroy()
		ptrsForEngine.fnGetVersionInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_author_info")
		defer methodName.Destroy()
		ptrsForEngine.fnGetAuthorInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_copyright_info")
		defer methodName.Destroy()
		ptrsForEngine.fnGetCopyrightInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_donor_info")
		defer methodName.Destroy()
		ptrsForEngine.fnGetDonorInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_license_info")
		defer methodName.Destroy()
		ptrsForEngine.fnGetLicenseInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_license_text")
		defer methodName.Destroy()
		ptrsForEngine.fnGetLicenseText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_architecture_name")
		defer methodName.Destroy()
		ptrsForEngine.fnGetArchitectureName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_in_physics_frame")
		defer methodName.Destroy()
		ptrsForEngine.fnIsInPhysicsFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("has_singleton")
		defer methodName.Destroy()
		ptrsForEngine.fnHasSingleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_singleton")
		defer methodName.Destroy()
		ptrsForEngine.fnGetSingleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1371597918))
	}
	{
		methodName := StringNameFromStr("register_singleton")
		defer methodName.Destroy()
		ptrsForEngine.fnRegisterSingleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 965313290))
	}
	{
		methodName := StringNameFromStr("unregister_singleton")
		defer methodName.Destroy()
		ptrsForEngine.fnUnregisterSingleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_singleton_list")
		defer methodName.Destroy()
		ptrsForEngine.fnGetSingletonList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("register_script_language")
		defer methodName.Destroy()
		ptrsForEngine.fnRegisterScriptLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1850254898))
	}
	{
		methodName := StringNameFromStr("unregister_script_language")
		defer methodName.Destroy()
		ptrsForEngine.fnUnregisterScriptLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1850254898))
	}
	{
		methodName := StringNameFromStr("get_script_language_count")
		defer methodName.Destroy()
		ptrsForEngine.fnGetScriptLanguageCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_script_language")
		defer methodName.Destroy()
		ptrsForEngine.fnGetScriptLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2151255799))
	}
	{
		methodName := StringNameFromStr("is_editor_hint")
		defer methodName.Destroy()
		ptrsForEngine.fnIsEditorHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_write_movie_path")
		defer methodName.Destroy()
		ptrsForEngine.fnGetWriteMoviePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_print_error_messages")
		defer methodName.Destroy()
		ptrsForEngine.fnSetPrintErrorMessages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_printing_error_messages")
		defer methodName.Destroy()
		ptrsForEngine.fnIsPrintingErrorMessages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type Engine struct {
	Object
}

func (me *Engine) BaseClass() string {
	return "Engine"
}

func NewEngine() *Engine {
	str := StringNameFromStr("Engine") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Engine{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Engine) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Engine) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Engine) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Engine) SetPhysicsTicksPerSecond(physics_ticks_per_second int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&physics_ticks_per_second)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnSetPhysicsTicksPerSecond), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) GetPhysicsTicksPerSecond() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetPhysicsTicksPerSecond), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) SetMaxPhysicsStepsPerFrame(max_physics_steps int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_physics_steps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnSetMaxPhysicsStepsPerFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) GetMaxPhysicsStepsPerFrame() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetMaxPhysicsStepsPerFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) SetPhysicsJitterFix(physics_jitter_fix float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&physics_jitter_fix)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnSetPhysicsJitterFix), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) GetPhysicsJitterFix() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetPhysicsJitterFix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetPhysicsInterpolationFraction() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetPhysicsInterpolationFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) SetMaxFps(max_fps int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_fps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnSetMaxFps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) GetMaxFps() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetMaxFps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) SetTimeScale(time_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnSetTimeScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) GetTimeScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetTimeScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetFramesDrawn() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetFramesDrawn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetFramesPerSecond() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetFramesPerSecond), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetPhysicsFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetPhysicsFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetProcessFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetProcessFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetMainLoop() MainLoop {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMainLoop()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetMainLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) GetVersionInfo() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetVersionInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) GetAuthorInfo() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetAuthorInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) GetCopyrightInfo() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetCopyrightInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Engine) GetDonorInfo() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetDonorInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) GetLicenseInfo() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetLicenseInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) GetLicenseText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetLicenseText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) GetArchitectureName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetArchitectureName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) IsInPhysicsFrame() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnIsInPhysicsFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) HasSingleton(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnHasSingleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetSingleton(name StringName) Object {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetSingleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) RegisterSingleton(name StringName, instance Object) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), instance.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnRegisterSingleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) UnregisterSingleton(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnUnregisterSingleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) GetSingletonList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetSingletonList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) RegisterScriptLanguage(language ScriptLanguage) Error {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnRegisterScriptLanguage), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Engine) UnregisterScriptLanguage(language ScriptLanguage) Error {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnUnregisterScriptLanguage), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Engine) GetScriptLanguageCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetScriptLanguageCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetScriptLanguage(index int64) ScriptLanguage {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewScriptLanguage()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetScriptLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) IsEditorHint() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnIsEditorHint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Engine) GetWriteMoviePath() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnGetWriteMoviePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Engine) SetPrintErrorMessages(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnSetPrintErrorMessages), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Engine) IsPrintingErrorMessages() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngine.fnIsPrintingErrorMessages), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
