package gdextension

import (
	"log"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
)

type Signal struct {
	obj  gdapi.Object
	name gdapi.StringName
}

func (me *Signal) Emit(args ...gdapi.Variant) error {
	cerr := me.obj.EmitSignal(me.name, args...)
	if cerr != gdapi.Ok {
		log.Printf("warning: error emitting signal: %v\n", cerr)
		return cerr
	}
	return nil
}
