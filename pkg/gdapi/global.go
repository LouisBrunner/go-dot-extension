package gdapi

import (
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var giface gdc.Interface

var Singletons *singletons
var Utilities UtilityFunctions

// FIXME: don't like this but it's the only way I found

func Initialize(iface gdc.Interface) {
	giface = iface
}

func InitializeGlobals(iface gdc.Interface) {
	Singletons = newSingletons(iface)
	Utilities = newUtilities(iface)
}
