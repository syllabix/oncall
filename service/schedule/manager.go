package schedule

type Manager interface {
	Prepare() CreateForm
}

func NewManager() Manager {
	return &manager{
		form: newScheduleForm,
	}
}

type manager struct {
	form CreateForm
}

func (m *manager) Prepare() CreateForm {
	return m.form
}
