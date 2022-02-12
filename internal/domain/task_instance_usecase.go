package domain

func (ti *TaskInstance) WithSuccess(items []*Item) *TaskInstance {
	ti.Error = nil
	ti.Success = true
	ti.Items = items
	return ti
}

func (ti *TaskInstance) WithError(err error) *TaskInstance {
	ti.Error = err
	ti.Success = false
	return ti
}
