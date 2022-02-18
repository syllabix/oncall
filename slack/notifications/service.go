package notifications

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/service/schedule"
	"github.com/syllabix/oncall/service/schedule/oncall"
	"go.uber.org/zap"
)

type Service interface {
	Schedule(oncall.Schedule) error
	ScheduleAll() error
	Start()
	Stop()
}

func NewService(
	client *slack.Client,
	manager schedule.Manager,
	log *zap.Logger,
) Service {
	scheduler := gocron.NewScheduler(time.UTC)
	return &service{client, scheduler, manager, log}
}

type service struct {
	client    *slack.Client
	scheduler *gocron.Scheduler
	manager   schedule.Manager
	log       *zap.Logger
}

func (s *service) ScheduleAll() error {
	schedules, err := s.manager.GetAll()
	if err != nil {
		return err
	}

	for _, sched := range schedules {
		s.Schedule(sched)
	}

	return nil
}

func (s *service) Schedule(schedule oncall.Schedule) error {
	s.scheduler.Every(5).Seconds().Tag(schedule.ID).Do(func() {
		err := s.startShift(schedule.ID)
		if err != nil {
			s.log.Error("an issue occured while starting a shift",
				zap.Error(err),
				zap.String("schedule-id", schedule.ID),
			)
		}
	})

	s.scheduler.Every(10).Seconds().WaitForSchedule().Tag(schedule.ID).Do(func() {
		err := s.endShift(schedule.ID)
		if err != nil {
			s.log.Error("an issue occured while starting a shift",
				zap.Error(err),
				zap.String("schedule-id", schedule.ID),
			)
		}
	})

	return nil
}

func (s *service) Start() {
	s.scheduler.StartAsync()
}

func (s *service) Stop() {
	s.scheduler.Stop()
}

func (s *service) startShift(scheduleID string) error {
	sched, err := s.manager.StartShift(scheduleID)
	if err != nil {
		return err
	}

	if sched.ActiveShift == nil {
		return nil
	}

	eod := "eod " + sched.ActiveShift.SlackHandle
	_, err = s.client.SetTopicOfConversation(sched.ChannelID, sched.Name+": "+eod)
	if err != nil {
		s.log.Error("failed to update channel with shift start",
			zap.Error(err),
			zap.String("channel-id", sched.ChannelID))
	}

	return nil
}

func (s *service) endShift(scheduleID string) error {
	sched, err := s.manager.EndShift(scheduleID)
	if err != nil {
		return err
	}

	if sched.ActiveShift == nil {
		return nil
	}

	_, _, err = s.client.PostMessage(
		sched.ChannelID,
		slack.MsgOptionText(
			"on call rotation ending for "+sched.ActiveShift.SlackHandle,
			false,
		),
	)
	if err != nil {
		s.log.Error("failed to update channel with shift end message",
			zap.Error(err),
			zap.String("channel-id", sched.ChannelID))
	}

	_, err = s.client.SetTopicOfConversation(sched.ChannelID, sched.Name+": :no_bell:")
	if err != nil {
		s.log.Error("failed to update channel with shift end",
			zap.Error(err),
			zap.String("channel-id", sched.ChannelID))
	}

	return nil
}
