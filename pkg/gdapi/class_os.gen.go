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

type ptrsForOSList struct {
	fnGetConnectedMidiInputs            gdc.MethodBindPtr
	fnOpenMidiInputs                    gdc.MethodBindPtr
	fnCloseMidiInputs                   gdc.MethodBindPtr
	fnAlert                             gdc.MethodBindPtr
	fnCrash                             gdc.MethodBindPtr
	fnSetLowProcessorUsageMode          gdc.MethodBindPtr
	fnIsInLowProcessorUsageMode         gdc.MethodBindPtr
	fnSetLowProcessorUsageModeSleepUsec gdc.MethodBindPtr
	fnGetLowProcessorUsageModeSleepUsec gdc.MethodBindPtr
	fnSetDeltaSmoothing                 gdc.MethodBindPtr
	fnIsDeltaSmoothingEnabled           gdc.MethodBindPtr
	fnGetProcessorCount                 gdc.MethodBindPtr
	fnGetProcessorName                  gdc.MethodBindPtr
	fnGetSystemFonts                    gdc.MethodBindPtr
	fnGetSystemFontPath                 gdc.MethodBindPtr
	fnGetSystemFontPathForText          gdc.MethodBindPtr
	fnGetExecutablePath                 gdc.MethodBindPtr
	fnReadStringFromStdin               gdc.MethodBindPtr
	fnExecute                           gdc.MethodBindPtr
	fnCreateProcess                     gdc.MethodBindPtr
	fnCreateInstance                    gdc.MethodBindPtr
	fnKill                              gdc.MethodBindPtr
	fnShellOpen                         gdc.MethodBindPtr
	fnShellShowInFileManager            gdc.MethodBindPtr
	fnIsProcessRunning                  gdc.MethodBindPtr
	fnGetProcessId                      gdc.MethodBindPtr
	fnHasEnvironment                    gdc.MethodBindPtr
	fnGetEnvironment                    gdc.MethodBindPtr
	fnSetEnvironment                    gdc.MethodBindPtr
	fnUnsetEnvironment                  gdc.MethodBindPtr
	fnGetName                           gdc.MethodBindPtr
	fnGetDistributionName               gdc.MethodBindPtr
	fnGetVersion                        gdc.MethodBindPtr
	fnGetCmdlineArgs                    gdc.MethodBindPtr
	fnGetCmdlineUserArgs                gdc.MethodBindPtr
	fnGetVideoAdapterDriverInfo         gdc.MethodBindPtr
	fnSetRestartOnExit                  gdc.MethodBindPtr
	fnIsRestartOnExitSet                gdc.MethodBindPtr
	fnGetRestartOnExitArguments         gdc.MethodBindPtr
	fnDelayUsec                         gdc.MethodBindPtr
	fnDelayMsec                         gdc.MethodBindPtr
	fnGetLocale                         gdc.MethodBindPtr
	fnGetLocaleLanguage                 gdc.MethodBindPtr
	fnGetModelName                      gdc.MethodBindPtr
	fnIsUserfsPersistent                gdc.MethodBindPtr
	fnIsStdoutVerbose                   gdc.MethodBindPtr
	fnIsDebugBuild                      gdc.MethodBindPtr
	fnGetStaticMemoryUsage              gdc.MethodBindPtr
	fnGetStaticMemoryPeakUsage          gdc.MethodBindPtr
	fnGetMemoryInfo                     gdc.MethodBindPtr
	fnMoveToTrash                       gdc.MethodBindPtr
	fnGetUserDataDir                    gdc.MethodBindPtr
	fnGetSystemDir                      gdc.MethodBindPtr
	fnGetConfigDir                      gdc.MethodBindPtr
	fnGetDataDir                        gdc.MethodBindPtr
	fnGetCacheDir                       gdc.MethodBindPtr
	fnGetUniqueId                       gdc.MethodBindPtr
	fnGetKeycodeString                  gdc.MethodBindPtr
	fnIsKeycodeUnicode                  gdc.MethodBindPtr
	fnFindKeycodeFromString             gdc.MethodBindPtr
	fnSetUseFileAccessSaveAndSwap       gdc.MethodBindPtr
	fnSetThreadName                     gdc.MethodBindPtr
	fnGetThreadCallerId                 gdc.MethodBindPtr
	fnGetMainThreadId                   gdc.MethodBindPtr
	fnHasFeature                        gdc.MethodBindPtr
	fnIsSandboxed                       gdc.MethodBindPtr
	fnRequestPermission                 gdc.MethodBindPtr
	fnRequestPermissions                gdc.MethodBindPtr
	fnGetGrantedPermissions             gdc.MethodBindPtr
	fnRevokeGrantedPermissions          gdc.MethodBindPtr
}

var ptrsForOS ptrsForOSList

func initOSPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OS")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_connected_midi_inputs")
		defer methodName.Destroy()
		ptrsForOS.fnGetConnectedMidiInputs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
	}
	{
		methodName := StringNameFromStr("open_midi_inputs")
		defer methodName.Destroy()
		ptrsForOS.fnOpenMidiInputs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("close_midi_inputs")
		defer methodName.Destroy()
		ptrsForOS.fnCloseMidiInputs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("alert")
		defer methodName.Destroy()
		ptrsForOS.fnAlert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1783970740))
	}
	{
		methodName := StringNameFromStr("crash")
		defer methodName.Destroy()
		ptrsForOS.fnCrash = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_low_processor_usage_mode")
		defer methodName.Destroy()
		ptrsForOS.fnSetLowProcessorUsageMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_in_low_processor_usage_mode")
		defer methodName.Destroy()
		ptrsForOS.fnIsInLowProcessorUsageMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_low_processor_usage_mode_sleep_usec")
		defer methodName.Destroy()
		ptrsForOS.fnSetLowProcessorUsageModeSleepUsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_low_processor_usage_mode_sleep_usec")
		defer methodName.Destroy()
		ptrsForOS.fnGetLowProcessorUsageModeSleepUsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_delta_smoothing")
		defer methodName.Destroy()
		ptrsForOS.fnSetDeltaSmoothing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_delta_smoothing_enabled")
		defer methodName.Destroy()
		ptrsForOS.fnIsDeltaSmoothingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_processor_count")
		defer methodName.Destroy()
		ptrsForOS.fnGetProcessorCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_processor_name")
		defer methodName.Destroy()
		ptrsForOS.fnGetProcessorName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_system_fonts")
		defer methodName.Destroy()
		ptrsForOS.fnGetSystemFonts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_system_font_path")
		defer methodName.Destroy()
		ptrsForOS.fnGetSystemFontPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 626580860))
	}
	{
		methodName := StringNameFromStr("get_system_font_path_for_text")
		defer methodName.Destroy()
		ptrsForOS.fnGetSystemFontPathForText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 197317981))
	}
	{
		methodName := StringNameFromStr("get_executable_path")
		defer methodName.Destroy()
		ptrsForOS.fnGetExecutablePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("read_string_from_stdin")
		defer methodName.Destroy()
		ptrsForOS.fnReadStringFromStdin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("execute")
		defer methodName.Destroy()
		ptrsForOS.fnExecute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1488299882))
	}
	{
		methodName := StringNameFromStr("create_process")
		defer methodName.Destroy()
		ptrsForOS.fnCreateProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2903767230))
	}
	{
		methodName := StringNameFromStr("create_instance")
		defer methodName.Destroy()
		ptrsForOS.fnCreateInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1080601263))
	}
	{
		methodName := StringNameFromStr("kill")
		defer methodName.Destroy()
		ptrsForOS.fnKill = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844576869))
	}
	{
		methodName := StringNameFromStr("shell_open")
		defer methodName.Destroy()
		ptrsForOS.fnShellOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("shell_show_in_file_manager")
		defer methodName.Destroy()
		ptrsForOS.fnShellShowInFileManager = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3565188097))
	}
	{
		methodName := StringNameFromStr("is_process_running")
		defer methodName.Destroy()
		ptrsForOS.fnIsProcessRunning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_process_id")
		defer methodName.Destroy()
		ptrsForOS.fnGetProcessId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("has_environment")
		defer methodName.Destroy()
		ptrsForOS.fnHasEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("get_environment")
		defer methodName.Destroy()
		ptrsForOS.fnGetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
	}
	{
		methodName := StringNameFromStr("set_environment")
		defer methodName.Destroy()
		ptrsForOS.fnSetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3605043004))
	}
	{
		methodName := StringNameFromStr("unset_environment")
		defer methodName.Destroy()
		ptrsForOS.fnUnsetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3089850668))
	}
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForOS.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_distribution_name")
		defer methodName.Destroy()
		ptrsForOS.fnGetDistributionName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_version")
		defer methodName.Destroy()
		ptrsForOS.fnGetVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_cmdline_args")
		defer methodName.Destroy()
		ptrsForOS.fnGetCmdlineArgs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
	}
	{
		methodName := StringNameFromStr("get_cmdline_user_args")
		defer methodName.Destroy()
		ptrsForOS.fnGetCmdlineUserArgs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
	}
	{
		methodName := StringNameFromStr("get_video_adapter_driver_info")
		defer methodName.Destroy()
		ptrsForOS.fnGetVideoAdapterDriverInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_restart_on_exit")
		defer methodName.Destroy()
		ptrsForOS.fnSetRestartOnExit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3331453935))
	}
	{
		methodName := StringNameFromStr("is_restart_on_exit_set")
		defer methodName.Destroy()
		ptrsForOS.fnIsRestartOnExitSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_restart_on_exit_arguments")
		defer methodName.Destroy()
		ptrsForOS.fnGetRestartOnExitArguments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("delay_usec")
		defer methodName.Destroy()
		ptrsForOS.fnDelayUsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 998575451))
	}
	{
		methodName := StringNameFromStr("delay_msec")
		defer methodName.Destroy()
		ptrsForOS.fnDelayMsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 998575451))
	}
	{
		methodName := StringNameFromStr("get_locale")
		defer methodName.Destroy()
		ptrsForOS.fnGetLocale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_locale_language")
		defer methodName.Destroy()
		ptrsForOS.fnGetLocaleLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_model_name")
		defer methodName.Destroy()
		ptrsForOS.fnGetModelName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_userfs_persistent")
		defer methodName.Destroy()
		ptrsForOS.fnIsUserfsPersistent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_stdout_verbose")
		defer methodName.Destroy()
		ptrsForOS.fnIsStdoutVerbose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_debug_build")
		defer methodName.Destroy()
		ptrsForOS.fnIsDebugBuild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_static_memory_usage")
		defer methodName.Destroy()
		ptrsForOS.fnGetStaticMemoryUsage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_static_memory_peak_usage")
		defer methodName.Destroy()
		ptrsForOS.fnGetStaticMemoryPeakUsage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_memory_info")
		defer methodName.Destroy()
		ptrsForOS.fnGetMemoryInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("move_to_trash")
		defer methodName.Destroy()
		ptrsForOS.fnMoveToTrash = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2113323047))
	}
	{
		methodName := StringNameFromStr("get_user_data_dir")
		defer methodName.Destroy()
		ptrsForOS.fnGetUserDataDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_system_dir")
		defer methodName.Destroy()
		ptrsForOS.fnGetSystemDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3073895123))
	}
	{
		methodName := StringNameFromStr("get_config_dir")
		defer methodName.Destroy()
		ptrsForOS.fnGetConfigDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_data_dir")
		defer methodName.Destroy()
		ptrsForOS.fnGetDataDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_cache_dir")
		defer methodName.Destroy()
		ptrsForOS.fnGetCacheDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_unique_id")
		defer methodName.Destroy()
		ptrsForOS.fnGetUniqueId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_keycode_string")
		defer methodName.Destroy()
		ptrsForOS.fnGetKeycodeString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2261993717))
	}
	{
		methodName := StringNameFromStr("is_keycode_unicode")
		defer methodName.Destroy()
		ptrsForOS.fnIsKeycodeUnicode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("find_keycode_from_string")
		defer methodName.Destroy()
		ptrsForOS.fnFindKeycodeFromString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1084858572))
	}
	{
		methodName := StringNameFromStr("set_use_file_access_save_and_swap")
		defer methodName.Destroy()
		ptrsForOS.fnSetUseFileAccessSaveAndSwap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_thread_name")
		defer methodName.Destroy()
		ptrsForOS.fnSetThreadName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("get_thread_caller_id")
		defer methodName.Destroy()
		ptrsForOS.fnGetThreadCallerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_main_thread_id")
		defer methodName.Destroy()
		ptrsForOS.fnGetMainThreadId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("has_feature")
		defer methodName.Destroy()
		ptrsForOS.fnHasFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("is_sandboxed")
		defer methodName.Destroy()
		ptrsForOS.fnIsSandboxed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("request_permission")
		defer methodName.Destroy()
		ptrsForOS.fnRequestPermission = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
	}
	{
		methodName := StringNameFromStr("request_permissions")
		defer methodName.Destroy()
		ptrsForOS.fnRequestPermissions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_granted_permissions")
		defer methodName.Destroy()
		ptrsForOS.fnGetGrantedPermissions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("revoke_granted_permissions")
		defer methodName.Destroy()
		ptrsForOS.fnRevokeGrantedPermissions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type OS struct {
	Object
}

func (me *OS) BaseClass() string {
	return "OS"
}

func NewOS() *OS {
	str := StringNameFromStr("OS") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OS{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type OSRenderingDriver int

const (
	OSRenderingDriverRenderingDriverVulkan  OSRenderingDriver = 0
	OSRenderingDriverRenderingDriverOpengl3 OSRenderingDriver = 1
)

type OSSystemDir int

const (
	OSSystemDirSystemDirDesktop   OSSystemDir = 0
	OSSystemDirSystemDirDcim      OSSystemDir = 1
	OSSystemDirSystemDirDocuments OSSystemDir = 2
	OSSystemDirSystemDirDownloads OSSystemDir = 3
	OSSystemDirSystemDirMovies    OSSystemDir = 4
	OSSystemDirSystemDirMusic     OSSystemDir = 5
	OSSystemDirSystemDirPictures  OSSystemDir = 6
	OSSystemDirSystemDirRingtones OSSystemDir = 7
)

func (me *OS) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OS) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OS) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OS) GetConnectedMidiInputs() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetConnectedMidiInputs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) OpenMidiInputs() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnOpenMidiInputs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) CloseMidiInputs() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnCloseMidiInputs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) Alert(text String, title String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnAlert), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) Crash(message String) {
	cargs := []gdc.ConstTypePtr{message.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnCrash), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) SetLowProcessorUsageMode(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnSetLowProcessorUsageMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) IsInLowProcessorUsageMode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsInLowProcessorUsageMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) SetLowProcessorUsageModeSleepUsec(usec int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&usec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnSetLowProcessorUsageModeSleepUsec), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) GetLowProcessorUsageModeSleepUsec() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetLowProcessorUsageModeSleepUsec), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) SetDeltaSmoothing(delta_smoothing_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta_smoothing_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnSetDeltaSmoothing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) IsDeltaSmoothingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsDeltaSmoothingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetProcessorCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetProcessorCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetProcessorName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetProcessorName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetSystemFonts() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetSystemFonts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetSystemFontPath(font_name String, weight int64, stretch int64, italic bool) String {
	cargs := []gdc.ConstTypePtr{font_name.AsCTypePtr(), gdc.ConstTypePtr(&weight), gdc.ConstTypePtr(&stretch), gdc.ConstTypePtr(&italic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&weight)
	pinner.Pin(&stretch)
	pinner.Pin(&italic)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetSystemFontPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetSystemFontPathForText(font_name String, text String, locale String, script String, weight int64, stretch int64, italic bool) PackedStringArray {
	cargs := []gdc.ConstTypePtr{font_name.AsCTypePtr(), text.AsCTypePtr(), locale.AsCTypePtr(), script.AsCTypePtr(), gdc.ConstTypePtr(&weight), gdc.ConstTypePtr(&stretch), gdc.ConstTypePtr(&italic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&weight)
	pinner.Pin(&stretch)
	pinner.Pin(&italic)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetSystemFontPathForText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetExecutablePath() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetExecutablePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) ReadStringFromStdin() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnReadStringFromStdin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) Execute(path String, arguments PackedStringArray, output Array, read_stderr bool, open_console bool) int64 {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), arguments.AsCTypePtr(), output.AsCTypePtr(), gdc.ConstTypePtr(&read_stderr), gdc.ConstTypePtr(&open_console)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&read_stderr)
	pinner.Pin(&open_console)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnExecute), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) CreateProcess(path String, arguments PackedStringArray, open_console bool) int64 {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), arguments.AsCTypePtr(), gdc.ConstTypePtr(&open_console)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&open_console)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnCreateProcess), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) CreateInstance(arguments PackedStringArray) int64 {
	cargs := []gdc.ConstTypePtr{arguments.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnCreateInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) Kill(pid int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pid)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&pid)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnKill), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OS) ShellOpen(uri String) Error {
	cargs := []gdc.ConstTypePtr{uri.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnShellOpen), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OS) ShellShowInFileManager(file_or_dir_path String, open_folder bool) Error {
	cargs := []gdc.ConstTypePtr{file_or_dir_path.AsCTypePtr(), gdc.ConstTypePtr(&open_folder)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&open_folder)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnShellShowInFileManager), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OS) IsProcessRunning(pid int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pid)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&pid)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsProcessRunning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetProcessId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetProcessId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) HasEnvironment(variable String) bool {
	cargs := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnHasEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetEnvironment(variable String) String {
	cargs := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) SetEnvironment(variable String, value String) {
	cargs := []gdc.ConstTypePtr{variable.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnSetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) UnsetEnvironment(variable String) {
	cargs := []gdc.ConstTypePtr{variable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnUnsetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) GetName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetDistributionName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetDistributionName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetVersion() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetCmdlineArgs() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetCmdlineArgs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetCmdlineUserArgs() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetCmdlineUserArgs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetVideoAdapterDriverInfo() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetVideoAdapterDriverInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) SetRestartOnExit(restart bool, arguments PackedStringArray) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&restart), arguments.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnSetRestartOnExit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) IsRestartOnExitSet() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsRestartOnExitSet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetRestartOnExitArguments() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetRestartOnExitArguments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) DelayUsec(usec int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&usec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnDelayUsec), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) DelayMsec(msec int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnDelayMsec), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) GetLocale() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetLocale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetLocaleLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetLocaleLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetModelName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetModelName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) IsUserfsPersistent() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsUserfsPersistent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) IsStdoutVerbose() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsStdoutVerbose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) IsDebugBuild() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsDebugBuild), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetStaticMemoryUsage() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetStaticMemoryUsage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetStaticMemoryPeakUsage() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetStaticMemoryPeakUsage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetMemoryInfo() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetMemoryInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) MoveToTrash(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnMoveToTrash), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OS) GetUserDataDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetUserDataDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetSystemDir(dir OSSystemDir, shared_storage bool) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&dir), gdc.ConstTypePtr(&shared_storage)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&dir)
	pinner.Pin(&shared_storage)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetSystemDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetConfigDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetConfigDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetDataDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetDataDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetCacheDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetCacheDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetUniqueId() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetUniqueId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) GetKeycodeString(code Key) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&code)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetKeycodeString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) IsKeycodeUnicode(code int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&code)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsKeycodeUnicode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) FindKeycodeFromString(string_ String) Key {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnFindKeycodeFromString), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OS) SetUseFileAccessSaveAndSwap(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnSetUseFileAccessSaveAndSwap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OS) SetThreadName(name String) Error {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnSetThreadName), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OS) GetThreadCallerId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetThreadCallerId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetMainThreadId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetMainThreadId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) HasFeature(tag_name String) bool {
	cargs := []gdc.ConstTypePtr{tag_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnHasFeature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) IsSandboxed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnIsSandboxed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) RequestPermission(name String) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnRequestPermission), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) RequestPermissions() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnRequestPermissions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OS) GetGrantedPermissions() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnGetGrantedPermissions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OS) RevokeGrantedPermissions() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOS.fnRevokeGrantedPermissions), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
