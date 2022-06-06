// Code generated by entc, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/syllabix/oncall/datastore/entity/predicate"
	"github.com/syllabix/oncall/datastore/entity/schedule"
	"github.com/syllabix/oncall/datastore/entity/shift"
)

// ScheduleUpdate is the builder for updating Schedule entities.
type ScheduleUpdate struct {
	config
	hooks    []Hook
	mutation *ScheduleMutation
}

// Where appends a list predicates to the ScheduleUpdate builder.
func (su *ScheduleUpdate) Where(ps ...predicate.Schedule) *ScheduleUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSlackChannelID sets the "slack_channel_id" field.
func (su *ScheduleUpdate) SetSlackChannelID(s string) *ScheduleUpdate {
	su.mutation.SetSlackChannelID(s)
	return su
}

// SetTeamSlackID sets the "team_slack_id" field.
func (su *ScheduleUpdate) SetTeamSlackID(s string) *ScheduleUpdate {
	su.mutation.SetTeamSlackID(s)
	return su
}

// SetName sets the "name" field.
func (su *ScheduleUpdate) SetName(s string) *ScheduleUpdate {
	su.mutation.SetName(s)
	return su
}

// SetInterval sets the "interval" field.
func (su *ScheduleUpdate) SetInterval(s schedule.Interval) *ScheduleUpdate {
	su.mutation.SetInterval(s)
	return su
}

// SetIsEnabled sets the "is_enabled" field.
func (su *ScheduleUpdate) SetIsEnabled(b bool) *ScheduleUpdate {
	su.mutation.SetIsEnabled(b)
	return su
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (su *ScheduleUpdate) SetNillableIsEnabled(b *bool) *ScheduleUpdate {
	if b != nil {
		su.SetIsEnabled(*b)
	}
	return su
}

// SetEndTime sets the "end_time" field.
func (su *ScheduleUpdate) SetEndTime(t time.Time) *ScheduleUpdate {
	su.mutation.SetEndTime(t)
	return su
}

// SetStartTime sets the "start_time" field.
func (su *ScheduleUpdate) SetStartTime(t time.Time) *ScheduleUpdate {
	su.mutation.SetStartTime(t)
	return su
}

// SetWeekdaysOnly sets the "weekdays_only" field.
func (su *ScheduleUpdate) SetWeekdaysOnly(b bool) *ScheduleUpdate {
	su.mutation.SetWeekdaysOnly(b)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *ScheduleUpdate) SetUpdatedAt(t time.Time) *ScheduleUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *ScheduleUpdate) SetDeletedAt(t time.Time) *ScheduleUpdate {
	su.mutation.SetDeletedAt(t)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *ScheduleUpdate) SetNillableDeletedAt(t *time.Time) *ScheduleUpdate {
	if t != nil {
		su.SetDeletedAt(*t)
	}
	return su
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (su *ScheduleUpdate) ClearDeletedAt() *ScheduleUpdate {
	su.mutation.ClearDeletedAt()
	return su
}

// AddShiftIDs adds the "shifts" edge to the Shift entity by IDs.
func (su *ScheduleUpdate) AddShiftIDs(ids ...int) *ScheduleUpdate {
	su.mutation.AddShiftIDs(ids...)
	return su
}

// AddShifts adds the "shifts" edges to the Shift entity.
func (su *ScheduleUpdate) AddShifts(s ...*Shift) *ScheduleUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddShiftIDs(ids...)
}

// Mutation returns the ScheduleMutation object of the builder.
func (su *ScheduleUpdate) Mutation() *ScheduleMutation {
	return su.mutation
}

// ClearShifts clears all "shifts" edges to the Shift entity.
func (su *ScheduleUpdate) ClearShifts() *ScheduleUpdate {
	su.mutation.ClearShifts()
	return su
}

// RemoveShiftIDs removes the "shifts" edge to Shift entities by IDs.
func (su *ScheduleUpdate) RemoveShiftIDs(ids ...int) *ScheduleUpdate {
	su.mutation.RemoveShiftIDs(ids...)
	return su
}

// RemoveShifts removes "shifts" edges to Shift entities.
func (su *ScheduleUpdate) RemoveShifts(s ...*Shift) *ScheduleUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveShiftIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScheduleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("entity: uninitialized hook (forgotten import entity/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScheduleUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScheduleUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScheduleUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ScheduleUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := schedule.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ScheduleUpdate) check() error {
	if v, ok := su.mutation.Interval(); ok {
		if err := schedule.IntervalValidator(v); err != nil {
			return &ValidationError{Name: "interval", err: fmt.Errorf(`entity: validator failed for field "Schedule.interval": %w`, err)}
		}
	}
	return nil
}

func (su *ScheduleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schedule.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.SlackChannelID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldSlackChannelID,
		})
	}
	if value, ok := su.mutation.TeamSlackID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldTeamSlackID,
		})
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldName,
		})
	}
	if value, ok := su.mutation.Interval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: schedule.FieldInterval,
		})
	}
	if value, ok := su.mutation.IsEnabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: schedule.FieldIsEnabled,
		})
	}
	if value, ok := su.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldEndTime,
		})
	}
	if value, ok := su.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldStartTime,
		})
	}
	if value, ok := su.mutation.WeekdaysOnly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: schedule.FieldWeekdaysOnly,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldDeletedAt,
		})
	}
	if su.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: schedule.FieldDeletedAt,
		})
	}
	if su.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ShiftsTable,
			Columns: []string{schedule.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shift.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedShiftsIDs(); len(nodes) > 0 && !su.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ShiftsTable,
			Columns: []string{schedule.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shift.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ShiftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ShiftsTable,
			Columns: []string{schedule.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shift.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ScheduleUpdateOne is the builder for updating a single Schedule entity.
type ScheduleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScheduleMutation
}

// SetSlackChannelID sets the "slack_channel_id" field.
func (suo *ScheduleUpdateOne) SetSlackChannelID(s string) *ScheduleUpdateOne {
	suo.mutation.SetSlackChannelID(s)
	return suo
}

// SetTeamSlackID sets the "team_slack_id" field.
func (suo *ScheduleUpdateOne) SetTeamSlackID(s string) *ScheduleUpdateOne {
	suo.mutation.SetTeamSlackID(s)
	return suo
}

// SetName sets the "name" field.
func (suo *ScheduleUpdateOne) SetName(s string) *ScheduleUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetInterval sets the "interval" field.
func (suo *ScheduleUpdateOne) SetInterval(s schedule.Interval) *ScheduleUpdateOne {
	suo.mutation.SetInterval(s)
	return suo
}

// SetIsEnabled sets the "is_enabled" field.
func (suo *ScheduleUpdateOne) SetIsEnabled(b bool) *ScheduleUpdateOne {
	suo.mutation.SetIsEnabled(b)
	return suo
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (suo *ScheduleUpdateOne) SetNillableIsEnabled(b *bool) *ScheduleUpdateOne {
	if b != nil {
		suo.SetIsEnabled(*b)
	}
	return suo
}

// SetEndTime sets the "end_time" field.
func (suo *ScheduleUpdateOne) SetEndTime(t time.Time) *ScheduleUpdateOne {
	suo.mutation.SetEndTime(t)
	return suo
}

// SetStartTime sets the "start_time" field.
func (suo *ScheduleUpdateOne) SetStartTime(t time.Time) *ScheduleUpdateOne {
	suo.mutation.SetStartTime(t)
	return suo
}

// SetWeekdaysOnly sets the "weekdays_only" field.
func (suo *ScheduleUpdateOne) SetWeekdaysOnly(b bool) *ScheduleUpdateOne {
	suo.mutation.SetWeekdaysOnly(b)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *ScheduleUpdateOne) SetUpdatedAt(t time.Time) *ScheduleUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *ScheduleUpdateOne) SetDeletedAt(t time.Time) *ScheduleUpdateOne {
	suo.mutation.SetDeletedAt(t)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *ScheduleUpdateOne) SetNillableDeletedAt(t *time.Time) *ScheduleUpdateOne {
	if t != nil {
		suo.SetDeletedAt(*t)
	}
	return suo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suo *ScheduleUpdateOne) ClearDeletedAt() *ScheduleUpdateOne {
	suo.mutation.ClearDeletedAt()
	return suo
}

// AddShiftIDs adds the "shifts" edge to the Shift entity by IDs.
func (suo *ScheduleUpdateOne) AddShiftIDs(ids ...int) *ScheduleUpdateOne {
	suo.mutation.AddShiftIDs(ids...)
	return suo
}

// AddShifts adds the "shifts" edges to the Shift entity.
func (suo *ScheduleUpdateOne) AddShifts(s ...*Shift) *ScheduleUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddShiftIDs(ids...)
}

// Mutation returns the ScheduleMutation object of the builder.
func (suo *ScheduleUpdateOne) Mutation() *ScheduleMutation {
	return suo.mutation
}

// ClearShifts clears all "shifts" edges to the Shift entity.
func (suo *ScheduleUpdateOne) ClearShifts() *ScheduleUpdateOne {
	suo.mutation.ClearShifts()
	return suo
}

// RemoveShiftIDs removes the "shifts" edge to Shift entities by IDs.
func (suo *ScheduleUpdateOne) RemoveShiftIDs(ids ...int) *ScheduleUpdateOne {
	suo.mutation.RemoveShiftIDs(ids...)
	return suo
}

// RemoveShifts removes "shifts" edges to Shift entities.
func (suo *ScheduleUpdateOne) RemoveShifts(s ...*Shift) *ScheduleUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveShiftIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScheduleUpdateOne) Select(field string, fields ...string) *ScheduleUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Schedule entity.
func (suo *ScheduleUpdateOne) Save(ctx context.Context) (*Schedule, error) {
	var (
		err  error
		node *Schedule
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("entity: uninitialized hook (forgotten import entity/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScheduleUpdateOne) SaveX(ctx context.Context) *Schedule {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScheduleUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScheduleUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ScheduleUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := schedule.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ScheduleUpdateOne) check() error {
	if v, ok := suo.mutation.Interval(); ok {
		if err := schedule.IntervalValidator(v); err != nil {
			return &ValidationError{Name: "interval", err: fmt.Errorf(`entity: validator failed for field "Schedule.interval": %w`, err)}
		}
	}
	return nil
}

func (suo *ScheduleUpdateOne) sqlSave(ctx context.Context) (_node *Schedule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schedule.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entity: missing "Schedule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schedule.FieldID)
		for _, f := range fields {
			if !schedule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
			}
			if f != schedule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.SlackChannelID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldSlackChannelID,
		})
	}
	if value, ok := suo.mutation.TeamSlackID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldTeamSlackID,
		})
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldName,
		})
	}
	if value, ok := suo.mutation.Interval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: schedule.FieldInterval,
		})
	}
	if value, ok := suo.mutation.IsEnabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: schedule.FieldIsEnabled,
		})
	}
	if value, ok := suo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldEndTime,
		})
	}
	if value, ok := suo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldStartTime,
		})
	}
	if value, ok := suo.mutation.WeekdaysOnly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: schedule.FieldWeekdaysOnly,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldDeletedAt,
		})
	}
	if suo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: schedule.FieldDeletedAt,
		})
	}
	if suo.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ShiftsTable,
			Columns: []string{schedule.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shift.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedShiftsIDs(); len(nodes) > 0 && !suo.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ShiftsTable,
			Columns: []string{schedule.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shift.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ShiftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ShiftsTable,
			Columns: []string{schedule.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shift.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Schedule{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
