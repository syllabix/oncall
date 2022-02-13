package schedule

import (
	"fmt"

	"github.com/syllabix/oncall/service/schedule/oncall"
)

type Manager interface {
	Create(oncall.Schedule) error
}

func NewManager() Manager {
	return &manager{}
}

type manager struct {
}

func (m *manager) Create(schedule oncall.Schedule) error {
	fmt.Printf("%+v", schedule)
	return nil
}
