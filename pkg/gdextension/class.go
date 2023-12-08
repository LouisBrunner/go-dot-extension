package gdextension

type classEntry struct {
	class       Class
	constructor ClassConstructor

	instances map[uint64]struct{}
}

func (me *classEntry) addInstance(id uint64) {
	me.instances[id] = struct{}{}
}

func (me *classEntry) deleteInstance(id uint64) {
	delete(me.instances, id)
}
