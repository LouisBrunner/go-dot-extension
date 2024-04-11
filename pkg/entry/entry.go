package entry

// #cgo CFLAGS: -I${SRCDIR}/../gdraw
/*
#include <gdextension_interface.h>

GDExtensionBool go_dot_gdextension_entry(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization* r_initialization);
*/
import "C"

import (
	"log"
	"os"
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
	gext = ext
	ext.Logf(gdextension.LogLevelDebug, "entry point called")
	ext.Register(globalRegisteredClasses...)
	res := C.GDExtensionBool(ext.Initialize(
		gdc.InitializationRawFromUnsafe(unsafe.Pointer(rInitialization)),
		gdc.InitializationInitializeFn(gdc.Callbacks.GetInitializationInitializeCallback()),
		gdc.InitializationDeinitializeFn(gdc.Callbacks.GetInitializationDeinitializeCallback()),
	))
	gdc.Callbacks.SetInitializationDeinitializeHandler(func(userdata unsafe.Pointer, pLevel gdc.InitializationLevel) {
		ext.Deinitialize(userdata, pLevel)
		if pLevel == gdc.InitializationCore {
			exit()
		}
	})
	return res
}

func exit() {
	// TODO: no way to avoid this???
	os.Exit(0)
}
