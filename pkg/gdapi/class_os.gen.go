// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OS struct {
  obj gdc.ObjectPtr
}

func (me *OS) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OS) BaseClass() string {
  return "OS"
}

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

func  (me *OS) GetConnectedMidiInputs() { // TODO: return value
  // TODO: implement
}

func  (me *OS) OpenMidiInputs() { // TODO: return value
  // TODO: implement
}

func  (me *OS) CloseMidiInputs() { // TODO: return value
  // TODO: implement
}

func  (me *OS) Alert(text String, title String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) Crash(message String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) SetLowProcessorUsageMode(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsInLowProcessorUsageMode() { // TODO: return value
  // TODO: implement
}

func  (me *OS) SetLowProcessorUsageModeSleepUsec(usec int, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetLowProcessorUsageModeSleepUsec() { // TODO: return value
  // TODO: implement
}

func  (me *OS) SetDeltaSmoothing(delta_smoothing_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsDeltaSmoothingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetProcessorCount() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetProcessorName() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetSystemFonts() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetSystemFontPath(font_name String, weight int, stretch int, italic bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetSystemFontPathForText(font_name String, text String, locale String, script String, weight int, stretch int, italic bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetExecutablePath() { // TODO: return value
  // TODO: implement
}

func  (me *OS) ReadStringFromStdin() { // TODO: return value
  // TODO: implement
}

func  (me *OS) Execute(path String, arguments PackedStringArray, output Array, read_stderr bool, open_console bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) CreateProcess(path String, arguments PackedStringArray, open_console bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) CreateInstance(arguments PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) Kill(pid int, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) ShellOpen(uri String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) ShellShowInFileManager(file_or_dir_path String, open_folder bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsProcessRunning(pid int, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetProcessId() { // TODO: return value
  // TODO: implement
}

func  (me *OS) HasEnvironment(variable String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetEnvironment(variable String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) SetEnvironment(variable String, value String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) UnsetEnvironment(variable String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetName() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetDistributionName() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetVersion() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetCmdlineArgs() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetCmdlineUserArgs() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetVideoAdapterDriverInfo() { // TODO: return value
  // TODO: implement
}

func  (me *OS) SetRestartOnExit(restart bool, arguments PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsRestartOnExitSet() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetRestartOnExitArguments() { // TODO: return value
  // TODO: implement
}

func  (me *OS) DelayUsec(usec int, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) DelayMsec(msec int, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetLocale() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetLocaleLanguage() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetModelName() { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsUserfsPersistent() { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsStdoutVerbose() { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsDebugBuild() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetStaticMemoryUsage() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetStaticMemoryPeakUsage() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetMemoryInfo() { // TODO: return value
  // TODO: implement
}

func  (me *OS) MoveToTrash(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetUserDataDir() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetSystemDir(dir OSSystemDir, shared_storage bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetConfigDir() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetDataDir() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetCacheDir() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetUniqueId() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetKeycodeString(code Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) IsKeycodeUnicode(code int, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) FindKeycodeFromString(string String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) SetUseFileAccessSaveAndSwap(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) SetThreadName(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetThreadCallerId() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetMainThreadId() { // TODO: return value
  // TODO: implement
}

func  (me *OS) HasFeature(tag_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) RequestPermission(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *OS) RequestPermissions() { // TODO: return value
  // TODO: implement
}

func  (me *OS) GetGrantedPermissions() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
