package flyweight

import (
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/list"
)

// Item represents item
type Item interface {
	IsActive() bool
}

// Pool represents factory
type Pool struct {
	actives *list.List
	pool    *list.List
	ite     *Iterator
}

// NewPool creates new Pool
func NewPool() *Pool {
	f := &Pool{}
	f.actives = list.NewList()
	f.pool = list.NewList()
	f.ite = &Iterator{}

	return f
}

// AddToPool adds resusable item
func (f *Pool) AddToPool(item unsafe.Pointer) {
	o := &Object{}
	o.data = item
	o.isActive = false
	ptr := unsafe.Pointer(o)
	elem := list.NewElement(ptr)
	o.elem = elem
	f.pool.AddElement(elem)
}

// GetIterator adds resusable item
func (f *Pool) GetIterator() *Iterator {
	ite := f.ite
	ite.current = f.actives.GetFirstElement()

	return ite
}

// CreateFromPool adds resusable item
func (f *Pool) CreateFromPool() unsafe.Pointer {
	e := f.pool.GetFirstElement()
	if e == nil {
		return nil
	}
	f.pool.RemoveElement(e)
	f.actives.AddElement(e)
	o := (*Object)(e.GetValue())
	o.isActive = true
	return o.GetData()
}

// Sweep remove non active objects from active list
func (f *Pool) Sweep() {
	for ite := f.actives.GetIterator(); ite.HasNext(); {
		elem := ite.Next()
		o := (*Object)(elem.GetValue())
		if o.isActive == false {
			f.actives.RemoveElement(elem)
			f.pool.AddElement(elem)
		}
	}
}
