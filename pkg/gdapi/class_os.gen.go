// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OS struct {
  Object
}

func (me *OS) BaseClass() string {
  return "OS"
}



// Enums

type OSRenderingDriver int
const (
  OSRenderingDriverRenderingDriverVulkan OSRenderingDriver = 0
  OSRenderingDriverRenderingDriverOpengl3 OSRenderingDriver = 1
)

type OSSystemDir int
const (
  OSSystemDirSystemDirDesktop OSSystemDir = 0
  OSSystemDirSystemDirDcim OSSystemDir = 1
  OSSystemDirSystemDirDocuments OSSystemDir = 2
  OSSystemDirSystemDirDownloads OSSystemDir = 3
  OSSystemDirSystemDirMovies OSSystemDir = 4
  OSSystemDirSystemDirMusic OSSystemDir = 5
  OSSystemDirSystemDirPictures OSSystemDir = 6
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

func  (me *OS) GetConnectedMidiInputs() PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connected_midi_inputs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) OpenMidiInputs()  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_midi_inputs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) CloseMidiInputs()  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close_midi_inputs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) Alert(text String, title String, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("alert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1783970740) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(title.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) Crash(message String, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("crash")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(message.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) SetLowProcessorUsageMode(enable bool, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_low_processor_usage_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) IsInLowProcessorUsageMode() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_low_processor_usage_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) SetLowProcessorUsageModeSleepUsec(usec int, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_low_processor_usage_mode_sleep_usec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&usec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) GetLowProcessorUsageModeSleepUsec() int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_low_processor_usage_mode_sleep_usec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) SetDeltaSmoothing(delta_smoothing_enabled bool, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delta_smoothing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta_smoothing_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) IsDeltaSmoothingEnabled() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_delta_smoothing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetProcessorCount() int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_processor_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetProcessorName() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_processor_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetSystemFonts() PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_system_fonts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetSystemFontPath(font_name String, weight int, stretch int, italic bool, ) String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_system_font_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 626580860) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_name.AsCTypePtr()), gdc.ConstTypePtr(&weight), gdc.ConstTypePtr(&stretch), gdc.ConstTypePtr(&italic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetSystemFontPathForText(font_name String, text String, locale String, script String, weight int, stretch int, italic bool, ) PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_system_font_path_for_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 197317981) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_name.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(locale.AsCTypePtr()), gdc.ConstTypePtr(script.AsCTypePtr()), gdc.ConstTypePtr(&weight), gdc.ConstTypePtr(&stretch), gdc.ConstTypePtr(&italic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetExecutablePath() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_executable_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) ReadStringFromStdin() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("read_string_from_stdin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) Execute(path String, arguments PackedStringArray, output Array, read_stderr bool, open_console bool, ) int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("execute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1488299882) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(arguments.AsCTypePtr()), gdc.ConstTypePtr(output.AsCTypePtr()), gdc.ConstTypePtr(&read_stderr), gdc.ConstTypePtr(&open_console), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) CreateProcess(path String, arguments PackedStringArray, open_console bool, ) int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_process")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2903767230) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(arguments.AsCTypePtr()), gdc.ConstTypePtr(&open_console), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) CreateInstance(arguments PackedStringArray, ) int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1080601263) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(arguments.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) Kill(pid int, ) Error {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("kill")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pid), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) ShellOpen(uri String, ) Error {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shell_open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(uri.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) ShellShowInFileManager(file_or_dir_path String, open_folder bool, ) Error {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shell_show_in_file_manager")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3565188097) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file_or_dir_path.AsCTypePtr()), gdc.ConstTypePtr(&open_folder), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) IsProcessRunning(pid int, ) bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_process_running")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pid), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetProcessId() int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) HasEnvironment(variable String, ) bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(variable.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetEnvironment(variable String, ) String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(variable.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) SetEnvironment(variable String, value String, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3605043004) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(variable.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) UnsetEnvironment(variable String, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unset_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3089850668) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(variable.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) GetName() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetDistributionName() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distribution_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetVersion() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetCmdlineArgs() PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cmdline_args")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetCmdlineUserArgs() PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cmdline_user_args")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetVideoAdapterDriverInfo() PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_video_adapter_driver_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) SetRestartOnExit(restart bool, arguments PackedStringArray, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_restart_on_exit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3331453935) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&restart), gdc.ConstTypePtr(arguments.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) IsRestartOnExitSet() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_restart_on_exit_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetRestartOnExitArguments() PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_restart_on_exit_arguments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) DelayUsec(usec int, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delay_usec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 998575451) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&usec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) DelayMsec(msec int, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delay_msec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 998575451) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) GetLocale() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_locale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetLocaleLanguage() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_locale_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetModelName() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_model_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) IsUserfsPersistent() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_userfs_persistent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) IsStdoutVerbose() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_stdout_verbose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) IsDebugBuild() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debug_build")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetStaticMemoryUsage() int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_static_memory_usage")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetStaticMemoryPeakUsage() int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_static_memory_peak_usage")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetMemoryInfo() Dictionary {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_memory_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) MoveToTrash(path String, ) Error {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_to_trash")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2113323047) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetUserDataDir() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_user_data_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetSystemDir(dir OSSystemDir, shared_storage bool, ) String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_system_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3073895123) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&dir), gdc.ConstTypePtr(&shared_storage), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetConfigDir() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_config_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetDataDir() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetCacheDir() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cache_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetUniqueId() String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetKeycodeString(code Key, ) String {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keycode_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2261993717) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) IsKeycodeUnicode(code int, ) bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keycode_unicode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) FindKeycodeFromString(string_ String, ) Key {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_keycode_from_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1084858572) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) SetUseFileAccessSaveAndSwap(enabled bool, )  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_file_access_save_and_swap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OS) SetThreadName(name String, ) Error {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_thread_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetThreadCallerId() int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_thread_caller_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetMainThreadId() int {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_main_thread_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) HasFeature(tag_name String, ) bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_feature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tag_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) IsSandboxed() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sandboxed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) RequestPermission(name String, ) bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_permission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) RequestPermissions() bool {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_permissions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) GetGrantedPermissions() PackedStringArray {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_granted_permissions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OS) RevokeGrantedPermissions()  {
  classNameV := StringNameFromStr("OS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("revoke_granted_permissions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
