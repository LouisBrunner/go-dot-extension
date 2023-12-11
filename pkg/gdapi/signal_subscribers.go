package gdapi

import (
	"log"
	"reflect"

	"github.com/google/uuid"
)

type Subcriber interface{}

type subscriber struct {
	fn   Subcriber
	call Callable
}

type SignalSubscribers struct {
	Object
	dispatchName StringName
	subs         map[string]subscriber
	subsRef      map[Subcriber]string
}

func NewSignalSubscribers() SignalSubscribers {
	return SignalSubscribers{
		subs:    make(map[string]subscriber),
		subsRef: make(map[Subcriber]string),
	}
}

func (me *SignalSubscribers) XInit() {
	me.dispatchName = StringNameFromStr("dispatch")
}

func (me *SignalSubscribers) Destroy() {
	me.dispatchName.Destroy()
	for _, sub := range me.subs {
		sub.call.Destroy()
	}
	me.subs = nil
	me.subsRef = nil
}

func (me *SignalSubscribers) add(sub interface{}) Callable {
	callable := NewCallableFromObjectStringName(me.Object, me.dispatchName)
	defer callable.Destroy()
	fnID := uuid.New().String()
	fnIDVar := NewVariantFrom(StringFromStr(fnID))
	defer fnIDVar.Destroy()
	boundCallable := callable.Bind(fnIDVar)
	me.subs[fnID] = subscriber{
		fn:   sub,
		call: boundCallable,
	}
	return boundCallable
}

func (me *SignalSubscribers) Dispatch(fnIDVar Variant, cargs ...interface{}) {
	fnID := fnIDVar.String()
	sub, found := me.subs[fnID]
	if !found {
		log.Printf("SignalSubscribers.Dispatch: fn is not a valid pointer")
		return
	}
	args := make([]reflect.Value, len(cargs))
	for i, arg := range cargs {
		args[i] = reflect.ValueOf(arg)
	}
	reflect.ValueOf(sub.fn).Call(args)
}

func (me *SignalSubscribers) remove(sub interface{}) *Callable {
	fnID, found := me.subsRef[sub]
	if !found {
		return nil
	}
	record, found := me.subs[fnID]
	if !found {
		return nil
	}
	delete(me.subsRef, sub)
	delete(me.subs, fnID)
	return &record.call
}
