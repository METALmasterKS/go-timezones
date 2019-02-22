package go_timezones

import "sync"

type repository struct {
	d Timezones
	m sync.Mutex
}

func NewRepository() repository {
	return repository{
		d: GetTimezones(),
	}

}

func (r *repository) GetTimezones() Timezones {
	r.m.Lock()
	defer r.m.Unlock()
	return r.d
}
