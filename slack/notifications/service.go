package notifications

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/service/oncall"
	"github.com/syllabix/oncall/service/schedule"
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

	var (
		endhour   = strconv.Itoa(schedule.EndTime.Hour())
		starthour = strconv.Itoa(schedule.StartTime.Hour())

		startExpr string
		endExpr   string
	)

	if schedule.WeekdaysOnly {
		startExpr = fmt.Sprintf("1 %s * * 1-5", starthour)
		endExpr = fmt.Sprintf("1 %s * * 1-5", endhour)
	} else {
		startExpr = fmt.Sprintf("1 %s * * 0-6", starthour)
		endExpr = fmt.Sprintf("1 %s * * 0-6", endhour)
	}

	s.scheduler.Cron(startExpr).
		Tag(fmt.Sprintf("sched-%d-start", schedule.ID)).
		Do(func() { s.startShift(schedule.ID) })
	s.scheduler.Cron(endExpr).
		Tag(fmt.Sprintf("sched-%d-end", schedule.ID)).
		Do(func() { s.endShift(schedule.ID) })
	s.log.Info("on call shifts have been scheduled",
		zap.Int("schedule-id", schedule.ID),
		zap.String("shift-start", startExpr),
		zap.String("shift-end", endExpr),
	)

	return nil
}

func (s *service) Start() {
	s.scheduler.StartAsync()
}

func (s *service) Stop() {
	s.log.Info("stopping scheduled jobs...")
	s.scheduler.Stop()
	s.log.Info("jobs stopped")
}

func (s *service) startShift(scheduleID int) {
	sched, err := s.manager.StartShift(scheduleID)
	if err != nil {
		if errors.Is(err, schedule.ErrNoActiveShift) {
			return
		}

		s.log.Error("an issue occured while starting a shift",
			zap.Error(err),
			zap.Int("schedule-id", scheduleID),
		)
		return
	}

	eod := sched.ActiveShift.SlackHandle
	_, err = s.client.SetTopicOfConversation(sched.ChannelID, sched.Name+": "+eod)
	if err != nil {
		s.log.Error("failed to update channel with shift start",
			zap.Error(err),
			zap.String("channel-id", sched.ChannelID))
	}
}

func (s *service) endShift(scheduleID int) {
	sched, err := s.manager.EndShift(scheduleID)
	if err != nil {
		if errors.Is(err, schedule.ErrNoActiveShift) {
			return
		}
		s.log.Error("an issue occured while starting a shift",
			zap.Error(err),
			zap.Int("schedule-id", scheduleID),
		)
		return
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

	_, err = s.client.SetTopicOfConversation(sched.ChannelID, sched.Name+": :sleeping:")
	if err != nil {
		s.log.Error("failed to update channel with shift end",
			zap.Error(err),
			zap.String("channel-id", sched.ChannelID))
	}
}
