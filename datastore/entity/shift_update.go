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
	"github.com/syllabix/oncall/datastore/entity/user"
)

// ShiftUpdate is the builder for updating Shift entities.
type ShiftUpdate struct {
	config
	hooks    []Hook
	mutation *ShiftMutation
}

// Where appends a list predicates to the ShiftUpdate builder.
func (su *ShiftUpdate) Where(ps ...predicate.Shift) *ShiftUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSequenceID sets the "sequence_id" field.
func (su *ShiftUpdate) SetSequenceID(i int) *ShiftUpdate {
	su.mutation.ResetSequenceID()
	su.mutation.SetSequenceID(i)
	return su
}

// SetNillableSequenceID sets the "sequence_id" field if the given value is not nil.
func (su *ShiftUpdate) SetNillableSequenceID(i *int) *ShiftUpdate {
	if i != nil {
		su.SetSequenceID(*i)
	}
	return su
}

// AddSequenceID adds i to the "sequence_id" field.
func (su *ShiftUpdate) AddSequenceID(i int) *ShiftUpdate {
	su.mutation.AddSequenceID(i)
	return su
}

// ClearSequenceID clears the value of the "sequence_id" field.
func (su *ShiftUpdate) ClearSequenceID() *ShiftUpdate {
	su.mutation.ClearSequenceID()
	return su
}

// SetStatus sets the "status" field.
func (su *ShiftUpdate) SetStatus(s shift.Status) *ShiftUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *ShiftUpdate) SetNillableStatus(s *shift.Status) *ShiftUpdate {
	if s != nil {
		su.SetStatus(*s)
	}
	return su
}

// ClearStatus clears the value of the "status" field.
func (su *ShiftUpdate) ClearStatus() *ShiftUpdate {
	su.mutation.ClearStatus()
	return su
}

// SetUserID sets the "user_id" field.
func (su *ShiftUpdate) SetUserID(i int) *ShiftUpdate {
	su.mutation.SetUserID(i)
	return su
}

// SetScheduleID sets the "schedule_id" field.
func (su *ShiftUpdate) SetScheduleID(i int) *ShiftUpdate {
	su.mutation.SetScheduleID(i)
	return su
}

// SetStartedAt sets the "started_at" field.
func (su *ShiftUpdate) SetStartedAt(t time.Time) *ShiftUpdate {
	su.mutation.SetStartedAt(t)
	return su
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (su *ShiftUpdate) SetNillableStartedAt(t *time.Time) *ShiftUpdate {
	if t != nil {
		su.SetStartedAt(*t)
	}
	return su
}

// ClearStartedAt clears the value of the "started_at" field.
func (su *ShiftUpdate) ClearStartedAt() *ShiftUpdate {
	su.mutation.ClearStartedAt()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *ShiftUpdate) SetUpdatedAt(t time.Time) *ShiftUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *ShiftUpdate) SetUser(u *User) *ShiftUpdate {
	return su.SetUserID(u.ID)
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (su *ShiftUpdate) SetSchedule(s *Schedule) *ShiftUpdate {
	return su.SetScheduleID(s.ID)
}

// Mutation returns the ShiftMutation object of the builder.
func (su *ShiftUpdate) Mutation() *ShiftMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *ShiftUpdate) ClearUser() *ShiftUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearSchedule clears the "schedule" edge to the Schedule entity.
func (su *ShiftUpdate) ClearSchedule() *ShiftUpdate {
	su.mutation.ClearSchedule()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ShiftUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*ShiftMutation)
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
func (su *ShiftUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShiftUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShiftUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ShiftUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := shift.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ShiftUpdate) check() error {
	if v, ok := su.mutation.Status(); ok {
		if err := shift.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entity: validator failed for field "Shift.status": %w`, err)}
		}
	}
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New(`entity: clearing a required unique edge "Shift.user"`)
	}
	if _, ok := su.mutation.ScheduleID(); su.mutation.ScheduleCleared() && !ok {
		return errors.New(`entity: clearing a required unique edge "Shift.schedule"`)
	}
	return nil
}

func (su *ShiftUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shift.Table,
			Columns: shift.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shift.FieldID,
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
	if value, ok := su.mutation.SequenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shift.FieldSequenceID,
		})
	}
	if value, ok := su.mutation.AddedSequenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shift.FieldSequenceID,
		})
	}
	if su.mutation.SequenceIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: shift.FieldSequenceID,
		})
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: shift.FieldStatus,
		})
	}
	if su.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: shift.FieldStatus,
		})
	}
	if value, ok := su.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldStartedAt,
		})
	}
	if su.mutation.StartedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: shift.FieldStartedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldUpdatedAt,
		})
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.UserTable,
			Columns: []string{shift.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.UserTable,
			Columns: []string{shift.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.ScheduleTable,
			Columns: []string{shift.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.ScheduleTable,
			Columns: []string{shift.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
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
			err = &NotFoundError{shift.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ShiftUpdateOne is the builder for updating a single Shift entity.
type ShiftUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShiftMutation
}

// SetSequenceID sets the "sequence_id" field.
func (suo *ShiftUpdateOne) SetSequenceID(i int) *ShiftUpdateOne {
	suo.mutation.ResetSequenceID()
	suo.mutation.SetSequenceID(i)
	return suo
}

// SetNillableSequenceID sets the "sequence_id" field if the given value is not nil.
func (suo *ShiftUpdateOne) SetNillableSequenceID(i *int) *ShiftUpdateOne {
	if i != nil {
		suo.SetSequenceID(*i)
	}
	return suo
}

// AddSequenceID adds i to the "sequence_id" field.
func (suo *ShiftUpdateOne) AddSequenceID(i int) *ShiftUpdateOne {
	suo.mutation.AddSequenceID(i)
	return suo
}

// ClearSequenceID clears the value of the "sequence_id" field.
func (suo *ShiftUpdateOne) ClearSequenceID() *ShiftUpdateOne {
	suo.mutation.ClearSequenceID()
	return suo
}

// SetStatus sets the "status" field.
func (suo *ShiftUpdateOne) SetStatus(s shift.Status) *ShiftUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *ShiftUpdateOne) SetNillableStatus(s *shift.Status) *ShiftUpdateOne {
	if s != nil {
		suo.SetStatus(*s)
	}
	return suo
}

// ClearStatus clears the value of the "status" field.
func (suo *ShiftUpdateOne) ClearStatus() *ShiftUpdateOne {
	suo.mutation.ClearStatus()
	return suo
}

// SetUserID sets the "user_id" field.
func (suo *ShiftUpdateOne) SetUserID(i int) *ShiftUpdateOne {
	suo.mutation.SetUserID(i)
	return suo
}

// SetScheduleID sets the "schedule_id" field.
func (suo *ShiftUpdateOne) SetScheduleID(i int) *ShiftUpdateOne {
	suo.mutation.SetScheduleID(i)
	return suo
}

// SetStartedAt sets the "started_at" field.
func (suo *ShiftUpdateOne) SetStartedAt(t time.Time) *ShiftUpdateOne {
	suo.mutation.SetStartedAt(t)
	return suo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (suo *ShiftUpdateOne) SetNillableStartedAt(t *time.Time) *ShiftUpdateOne {
	if t != nil {
		suo.SetStartedAt(*t)
	}
	return suo
}

// ClearStartedAt clears the value of the "started_at" field.
func (suo *ShiftUpdateOne) ClearStartedAt() *ShiftUpdateOne {
	suo.mutation.ClearStartedAt()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *ShiftUpdateOne) SetUpdatedAt(t time.Time) *ShiftUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *ShiftUpdateOne) SetUser(u *User) *ShiftUpdateOne {
	return suo.SetUserID(u.ID)
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (suo *ShiftUpdateOne) SetSchedule(s *Schedule) *ShiftUpdateOne {
	return suo.SetScheduleID(s.ID)
}

// Mutation returns the ShiftMutation object of the builder.
func (suo *ShiftUpdateOne) Mutation() *ShiftMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *ShiftUpdateOne) ClearUser() *ShiftUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearSchedule clears the "schedule" edge to the Schedule entity.
func (suo *ShiftUpdateOne) ClearSchedule() *ShiftUpdateOne {
	suo.mutation.ClearSchedule()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ShiftUpdateOne) Select(field string, fields ...string) *ShiftUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Shift entity.
func (suo *ShiftUpdateOne) Save(ctx context.Context) (*Shift, error) {
	var (
		err  error
		node *Shift
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShiftMutation)
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
func (suo *ShiftUpdateOne) SaveX(ctx context.Context) *Shift {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ShiftUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShiftUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ShiftUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := shift.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ShiftUpdateOne) check() error {
	if v, ok := suo.mutation.Status(); ok {
		if err := shift.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entity: validator failed for field "Shift.status": %w`, err)}
		}
	}
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New(`entity: clearing a required unique edge "Shift.user"`)
	}
	if _, ok := suo.mutation.ScheduleID(); suo.mutation.ScheduleCleared() && !ok {
		return errors.New(`entity: clearing a required unique edge "Shift.schedule"`)
	}
	return nil
}

func (suo *ShiftUpdateOne) sqlSave(ctx context.Context) (_node *Shift, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shift.Table,
			Columns: shift.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shift.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entity: missing "Shift.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shift.FieldID)
		for _, f := range fields {
			if !shift.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
			}
			if f != shift.FieldID {
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
	if value, ok := suo.mutation.SequenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shift.FieldSequenceID,
		})
	}
	if value, ok := suo.mutation.AddedSequenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shift.FieldSequenceID,
		})
	}
	if suo.mutation.SequenceIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: shift.FieldSequenceID,
		})
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: shift.FieldStatus,
		})
	}
	if suo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: shift.FieldStatus,
		})
	}
	if value, ok := suo.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldStartedAt,
		})
	}
	if suo.mutation.StartedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: shift.FieldStartedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldUpdatedAt,
		})
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.UserTable,
			Columns: []string{shift.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.UserTable,
			Columns: []string{shift.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.ScheduleTable,
			Columns: []string{shift.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shift.ScheduleTable,
			Columns: []string{shift.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Shift{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shift.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}