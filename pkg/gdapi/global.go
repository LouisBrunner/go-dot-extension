package gdapi

import (
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var giface gdc.Interface

var Singletons *singletons
var Utilities UtilityFunctions

// FIXME: don't like this but it's the only way I found
// without requiring the user to pass the extension to EVERYTHING

func Initialize(iface gdc.Interface) {
	giface = iface
	initBClasses(iface)
	Utilities = newUtilities(iface)
}

func InitializeSingletons(iface gdc.Interface) {
	Singletons = newSingletons(iface)
}

func InitializeClasses(iface gdc.Interface) {
	initClasses(iface)
}
