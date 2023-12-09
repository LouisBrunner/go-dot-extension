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

func  (me *OS) GetConnectedMidiInputs()  {
  panic("TODO: implement")
}

func  (me *OS) OpenMidiInputs()  {
  panic("TODO: implement")
}

func  (me *OS) CloseMidiInputs()  {
  panic("TODO: implement")
}

func  (me *OS) Alert(text String, title String, )  {
  panic("TODO: implement")
}

func  (me *OS) Crash(message String, )  {
  panic("TODO: implement")
}

func  (me *OS) SetLowProcessorUsageMode(enable bool, )  {
  panic("TODO: implement")
}

func  (me *OS) IsInLowProcessorUsageMode()  {
  panic("TODO: implement")
}

func  (me *OS) SetLowProcessorUsageModeSleepUsec(usec int, )  {
  panic("TODO: implement")
}

func  (me *OS) GetLowProcessorUsageModeSleepUsec()  {
  panic("TODO: implement")
}

func  (me *OS) SetDeltaSmoothing(delta_smoothing_enabled bool, )  {
  panic("TODO: implement")
}

func  (me *OS) IsDeltaSmoothingEnabled()  {
  panic("TODO: implement")
}

func  (me *OS) GetProcessorCount()  {
  panic("TODO: implement")
}

func  (me *OS) GetProcessorName()  {
  panic("TODO: implement")
}

func  (me *OS) GetSystemFonts()  {
  panic("TODO: implement")
}

func  (me *OS) GetSystemFontPath(font_name String, weight int, stretch int, italic bool, )  {
  panic("TODO: implement")
}

func  (me *OS) GetSystemFontPathForText(font_name String, text String, locale String, script String, weight int, stretch int, italic bool, )  {
  panic("TODO: implement")
}

func  (me *OS) GetExecutablePath()  {
  panic("TODO: implement")
}

func  (me *OS) ReadStringFromStdin()  {
  panic("TODO: implement")
}

func  (me *OS) Execute(path String, arguments PackedStringArray, output Array, read_stderr bool, open_console bool, )  {
  panic("TODO: implement")
}

func  (me *OS) CreateProcess(path String, arguments PackedStringArray, open_console bool, )  {
  panic("TODO: implement")
}

func  (me *OS) CreateInstance(arguments PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *OS) Kill(pid int, )  {
  panic("TODO: implement")
}

func  (me *OS) ShellOpen(uri String, )  {
  panic("TODO: implement")
}

func  (me *OS) ShellShowInFileManager(file_or_dir_path String, open_folder bool, )  {
  panic("TODO: implement")
}

func  (me *OS) IsProcessRunning(pid int, )  {
  panic("TODO: implement")
}

func  (me *OS) GetProcessId()  {
  panic("TODO: implement")
}

func  (me *OS) HasEnvironment(variable String, )  {
  panic("TODO: implement")
}

func  (me *OS) GetEnvironment(variable String, )  {
  panic("TODO: implement")
}

func  (me *OS) SetEnvironment(variable String, value String, )  {
  panic("TODO: implement")
}

func  (me *OS) UnsetEnvironment(variable String, )  {
  panic("TODO: implement")
}

func  (me *OS) GetName()  {
  panic("TODO: implement")
}

func  (me *OS) GetDistributionName()  {
  panic("TODO: implement")
}

func  (me *OS) GetVersion()  {
  panic("TODO: implement")
}

func  (me *OS) GetCmdlineArgs()  {
  panic("TODO: implement")
}

func  (me *OS) GetCmdlineUserArgs()  {
  panic("TODO: implement")
}

func  (me *OS) GetVideoAdapterDriverInfo()  {
  panic("TODO: implement")
}

func  (me *OS) SetRestartOnExit(restart bool, arguments PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *OS) IsRestartOnExitSet()  {
  panic("TODO: implement")
}

func  (me *OS) GetRestartOnExitArguments()  {
  panic("TODO: implement")
}

func  (me *OS) DelayUsec(usec int, )  {
  panic("TODO: implement")
}

func  (me *OS) DelayMsec(msec int, )  {
  panic("TODO: implement")
}

func  (me *OS) GetLocale()  {
  panic("TODO: implement")
}

func  (me *OS) GetLocaleLanguage()  {
  panic("TODO: implement")
}

func  (me *OS) GetModelName()  {
  panic("TODO: implement")
}

func  (me *OS) IsUserfsPersistent()  {
  panic("TODO: implement")
}

func  (me *OS) IsStdoutVerbose()  {
  panic("TODO: implement")
}

func  (me *OS) IsDebugBuild()  {
  panic("TODO: implement")
}

func  (me *OS) GetStaticMemoryUsage()  {
  panic("TODO: implement")
}

func  (me *OS) GetStaticMemoryPeakUsage()  {
  panic("TODO: implement")
}

func  (me *OS) GetMemoryInfo()  {
  panic("TODO: implement")
}

func  (me *OS) MoveToTrash(path String, )  {
  panic("TODO: implement")
}

func  (me *OS) GetUserDataDir()  {
  panic("TODO: implement")
}

func  (me *OS) GetSystemDir(dir OSSystemDir, shared_storage bool, )  {
  panic("TODO: implement")
}

func  (me *OS) GetConfigDir()  {
  panic("TODO: implement")
}

func  (me *OS) GetDataDir()  {
  panic("TODO: implement")
}

func  (me *OS) GetCacheDir()  {
  panic("TODO: implement")
}

func  (me *OS) GetUniqueId()  {
  panic("TODO: implement")
}

func  (me *OS) GetKeycodeString(code Key, )  {
  panic("TODO: implement")
}

func  (me *OS) IsKeycodeUnicode(code int, )  {
  panic("TODO: implement")
}

func  (me *OS) FindKeycodeFromString(string_ String, )  {
  panic("TODO: implement")
}

func  (me *OS) SetUseFileAccessSaveAndSwap(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *OS) SetThreadName(name String, )  {
  panic("TODO: implement")
}

func  (me *OS) GetThreadCallerId()  {
  panic("TODO: implement")
}

func  (me *OS) GetMainThreadId()  {
  panic("TODO: implement")
}

func  (me *OS) HasFeature(tag_name String, )  {
  panic("TODO: implement")
}

func  (me *OS) RequestPermission(name String, )  {
  panic("TODO: implement")
}

func  (me *OS) RequestPermissions()  {
  panic("TODO: implement")
}

func  (me *OS) GetGrantedPermissions()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
