package schedule

import (
	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/service/schedule/oncall"
	"go.uber.org/zap"
)

type Manager interface {
	Create(oncall.Schedule) (oncall.Schedule, error)
}

func NewManager(db schedule.Store, log *zap.Logger) Manager {
	return &manager{db, log}
}

type manager struct {
	db  schedule.Store
	log *zap.Logger
}

func (m *manager) Create(schedule oncall.Schedule) (oncall.Schedule, error) {
	result, err := m.db.Create(asModel(schedule))
	if err != nil {
		m.log.Error("failed to create schedule", zap.Error(err))
		return oncall.Schedule{}, err
	}
	return asSchedule(result), nil
}

func asModel(schedule oncall.Schedule) model.OncallSchedule {
	return model.OncallSchedule{
		Name:           schedule.Name,
		TeamSlackID:    schedule.TeamID,
		Interval:       string(schedule.Interval),
		SlackChannelID: schedule.ChannelID,
		StartTime:      schedule.StartTime,
		EndTime:        schedule.EndTime,
	}
}

func asSchedule(model model.OncallSchedule) oncall.Schedule {
	return oncall.Schedule{
		ID:        model.ID,
		Name:      model.Name,
		TeamID:    model.TeamSlackID,
		Interval:  oncall.Interval(model.Interval),
		ChannelID: model.SlackChannelID,
		StartTime: model.StartTime,
		EndTime:   model.EndTime,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
