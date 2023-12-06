package entry

// #cgo CFLAGS: -I${SRCDIR}/../gdraw
/*
#include <gdextension_interface.h>

GDExtensionBool go_dot_gdextension_entry(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization* r_initialization);
void go_dot_gdextension_initialize_module(void* userdata, GDExtensionInitializationLevel p_level);
void go_dot_gdextension_uninitialize_module(void* userdata, GDExtensionInitializationLevel p_level);
*/
import "C"

import (
	"log"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdextension"
)

var DebugMode = false

//export go_dot_gdextension_entry
func go_dot_gdextension_entry(pGetProcAddress C.GDExtensionInterfaceGetProcAddress, pLibrary C.GDExtensionClassLibraryPtr, rInitialization *C.GDExtensionInitialization) C.GDExtensionBool {
	level := gdextension.LogLevelWarning
	if DebugMode {
		level = gdextension.LogLevelDebug
	}

	ext, err := gdextension.New(
		gdc.InterfaceGetProcAddress(pGetProcAddress),
		gdc.ClassLibraryPtr(pLibrary),
		level,
	)
	if err != nil {
		log.Printf("fatal error: %s", err.Error())
		return C.GDExtensionBool(0)
	}
	ext.Logf(gdextension.LogLevelDebug, "entry point called")
	ext.Register(globalRegisteredClasses...)
	return C.GDExtensionBool(ext.Initialize(
		gdc.InitializationRawFromUnsafe(unsafe.Pointer(rInitialization)),
		gdc.InitializationInitializeFn(C.go_dot_gdextension_initialize_module),
		gdc.InitializationDeinitializeFn(C.go_dot_gdextension_uninitialize_module),
	))
}

//export go_dot_gdextension_initialize_module
func go_dot_gdextension_initialize_module(userData *C.void, pLevel C.GDExtensionInitializationLevel) {
	ext, err := gdextension.RestoreFromC(unsafe.Pointer(userData))
	if err != nil {
		log.Printf("fatal error: %s", err.Error())
		return
	}
	ext.Logf(gdextension.LogLevelDebug, "initializing module (level=%v)", pLevel)
	err = ext.OnInit(gdc.InitializationLevel(pLevel))
	if err != nil {
		ext.Logf(gdextension.LogLevelError, "error: %s", err.Error())
	}
}

//export go_dot_gdextension_uninitialize_module
func go_dot_gdextension_uninitialize_module(userData *C.void, pLevel C.GDExtensionInitializationLevel) {
	ext, err := gdextension.RestoreFromC(unsafe.Pointer(userData))
	if err != nil {
		log.Printf("fatal error: %s", err.Error())
		return
	}
	ext.Logf(gdextension.LogLevelDebug, "uninitializing module (level=%v)", pLevel)
	err = ext.OnFini(gdc.InitializationLevel(pLevel))
	if err != nil {
		ext.Logf(gdextension.LogLevelError, "error: %s", err.Error())
	}
}
