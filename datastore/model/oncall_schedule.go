// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OncallSchedule is an object representing the database table.
type OncallSchedule struct {
	ID             string      `db:"id" boil:"id" json:"id" toml:"id" yaml:"id"`
	TeamSlackID    string      `db:"team_slack_id" boil:"team_slack_id" json:"team_slack_id" toml:"team_slack_id" yaml:"team_slack_id"`
	Name           string      `db:"name" boil:"name" json:"name" toml:"name" yaml:"name"`
	Interval       string      `db:"interval" boil:"interval" json:"interval" toml:"interval" yaml:"interval"`
	IsEnabled      bool        `db:"is_enabled" boil:"is_enabled" json:"is_enabled" toml:"is_enabled" yaml:"is_enabled"`
	StartTime      time.Time   `db:"start_time" boil:"start_time" json:"start_time" toml:"start_time" yaml:"start_time"`
	EndTime        time.Time   `db:"end_time" boil:"end_time" json:"end_time" toml:"end_time" yaml:"end_time"`
	ActiveShift    null.String `db:"active_shift" boil:"active_shift" json:"active_shift,omitempty" toml:"active_shift" yaml:"active_shift,omitempty"`
	OverrideShift  null.String `db:"override_shift" boil:"override_shift" json:"override_shift,omitempty" toml:"override_shift" yaml:"override_shift,omitempty"`
	SlackChannelID string      `db:"slack_channel_id" boil:"slack_channel_id" json:"slack_channel_id" toml:"slack_channel_id" yaml:"slack_channel_id"`
	CreatedAt      time.Time   `db:"created_at" boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time   `db:"updated_at" boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt      null.Time   `db:"deleted_at" boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *oncallScheduleR `db:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L oncallScheduleL  `db:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OncallScheduleColumns = struct {
	ID             string
	TeamSlackID    string
	Name           string
	Interval       string
	IsEnabled      string
	StartTime      string
	EndTime        string
	ActiveShift    string
	OverrideShift  string
	SlackChannelID string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
}{
	ID:             "id",
	TeamSlackID:    "team_slack_id",
	Name:           "name",
	Interval:       "interval",
	IsEnabled:      "is_enabled",
	StartTime:      "start_time",
	EndTime:        "end_time",
	ActiveShift:    "active_shift",
	OverrideShift:  "override_shift",
	SlackChannelID: "slack_channel_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var OncallScheduleWhere = struct {
	ID             whereHelperstring
	TeamSlackID    whereHelperstring
	Name           whereHelperstring
	Interval       whereHelperstring
	IsEnabled      whereHelperbool
	StartTime      whereHelpertime_Time
	EndTime        whereHelpertime_Time
	ActiveShift    whereHelpernull_String
	OverrideShift  whereHelpernull_String
	SlackChannelID whereHelperstring
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
	DeletedAt      whereHelpernull_Time
}{
	ID:             whereHelperstring{field: "\"oncall_schedule\".\"id\""},
	TeamSlackID:    whereHelperstring{field: "\"oncall_schedule\".\"team_slack_id\""},
	Name:           whereHelperstring{field: "\"oncall_schedule\".\"name\""},
	Interval:       whereHelperstring{field: "\"oncall_schedule\".\"interval\""},
	IsEnabled:      whereHelperbool{field: "\"oncall_schedule\".\"is_enabled\""},
	StartTime:      whereHelpertime_Time{field: "\"oncall_schedule\".\"start_time\""},
	EndTime:        whereHelpertime_Time{field: "\"oncall_schedule\".\"end_time\""},
	ActiveShift:    whereHelpernull_String{field: "\"oncall_schedule\".\"active_shift\""},
	OverrideShift:  whereHelpernull_String{field: "\"oncall_schedule\".\"override_shift\""},
	SlackChannelID: whereHelperstring{field: "\"oncall_schedule\".\"slack_channel_id\""},
	CreatedAt:      whereHelpertime_Time{field: "\"oncall_schedule\".\"created_at\""},
	UpdatedAt:      whereHelpertime_Time{field: "\"oncall_schedule\".\"updated_at\""},
	DeletedAt:      whereHelpernull_Time{field: "\"oncall_schedule\".\"deleted_at\""},
}

// OncallScheduleRels is where relationship names are stored.
var OncallScheduleRels = struct {
}{}

// oncallScheduleR is where relationships are stored.
type oncallScheduleR struct {
}

// NewStruct creates a new relationship struct
func (*oncallScheduleR) NewStruct() *oncallScheduleR {
	return &oncallScheduleR{}
}

// oncallScheduleL is where Load methods for each relationship are stored.
type oncallScheduleL struct{}

var (
	oncallScheduleAllColumns            = []string{"id", "team_slack_id", "name", "interval", "is_enabled", "start_time", "end_time", "active_shift", "override_shift", "slack_channel_id", "created_at", "updated_at", "deleted_at"}
	oncallScheduleColumnsWithoutDefault = []string{"team_slack_id", "name", "interval", "start_time", "end_time", "active_shift", "override_shift", "slack_channel_id", "deleted_at"}
	oncallScheduleColumnsWithDefault    = []string{"id", "is_enabled", "created_at", "updated_at"}
	oncallSchedulePrimaryKeyColumns     = []string{"id"}
)

type (
	// OncallScheduleSlice is an alias for a slice of pointers to OncallSchedule.
	// This should generally be used opposed to []OncallSchedule.
	OncallScheduleSlice []*OncallSchedule

	oncallScheduleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	oncallScheduleType                 = reflect.TypeOf(&OncallSchedule{})
	oncallScheduleMapping              = queries.MakeStructMapping(oncallScheduleType)
	oncallSchedulePrimaryKeyMapping, _ = queries.BindMapping(oncallScheduleType, oncallScheduleMapping, oncallSchedulePrimaryKeyColumns)
	oncallScheduleInsertCacheMut       sync.RWMutex
	oncallScheduleInsertCache          = make(map[string]insertCache)
	oncallScheduleUpdateCacheMut       sync.RWMutex
	oncallScheduleUpdateCache          = make(map[string]updateCache)
	oncallScheduleUpsertCacheMut       sync.RWMutex
	oncallScheduleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single oncallSchedule record from the query.
func (q oncallScheduleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OncallSchedule, error) {
	o := &OncallSchedule{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for oncall_schedule")
	}

	return o, nil
}

// All returns all OncallSchedule records from the query.
func (q oncallScheduleQuery) All(ctx context.Context, exec boil.ContextExecutor) (OncallScheduleSlice, error) {
	var o []*OncallSchedule

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to OncallSchedule slice")
	}

	return o, nil
}

// Count returns the count of all OncallSchedule records in the query.
func (q oncallScheduleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count oncall_schedule rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q oncallScheduleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if oncall_schedule exists")
	}

	return count > 0, nil
}

// OncallSchedules retrieves all the records using an executor.
func OncallSchedules(mods ...qm.QueryMod) oncallScheduleQuery {
	mods = append(mods, qm.From("\"oncall_schedule\""), qmhelper.WhereIsNull("\"oncall_schedule\".\"deleted_at\""))
	return oncallScheduleQuery{NewQuery(mods...)}
}

// FindOncallSchedule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOncallSchedule(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*OncallSchedule, error) {
	oncallScheduleObj := &OncallSchedule{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"oncall_schedule\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, oncallScheduleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from oncall_schedule")
	}

	return oncallScheduleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OncallSchedule) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no oncall_schedule provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(oncallScheduleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	oncallScheduleInsertCacheMut.RLock()
	cache, cached := oncallScheduleInsertCache[key]
	oncallScheduleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			oncallScheduleAllColumns,
			oncallScheduleColumnsWithDefault,
			oncallScheduleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(oncallScheduleType, oncallScheduleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(oncallScheduleType, oncallScheduleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"oncall_schedule\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"oncall_schedule\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into oncall_schedule")
	}

	if !cached {
		oncallScheduleInsertCacheMut.Lock()
		oncallScheduleInsertCache[key] = cache
		oncallScheduleInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the OncallSchedule.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OncallSchedule) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	oncallScheduleUpdateCacheMut.RLock()
	cache, cached := oncallScheduleUpdateCache[key]
	oncallScheduleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			oncallScheduleAllColumns,
			oncallSchedulePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update oncall_schedule, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"oncall_schedule\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, oncallSchedulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(oncallScheduleType, oncallScheduleMapping, append(wl, oncallSchedulePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update oncall_schedule row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for oncall_schedule")
	}

	if !cached {
		oncallScheduleUpdateCacheMut.Lock()
		oncallScheduleUpdateCache[key] = cache
		oncallScheduleUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q oncallScheduleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for oncall_schedule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for oncall_schedule")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OncallScheduleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oncallSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"oncall_schedule\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, oncallSchedulePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in oncallSchedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all oncallSchedule")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OncallSchedule) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no oncall_schedule provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(oncallScheduleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	oncallScheduleUpsertCacheMut.RLock()
	cache, cached := oncallScheduleUpsertCache[key]
	oncallScheduleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			oncallScheduleAllColumns,
			oncallScheduleColumnsWithDefault,
			oncallScheduleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			oncallScheduleAllColumns,
			oncallSchedulePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("model: unable to upsert oncall_schedule, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(oncallSchedulePrimaryKeyColumns))
			copy(conflict, oncallSchedulePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"oncall_schedule\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(oncallScheduleType, oncallScheduleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(oncallScheduleType, oncallScheduleMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "model: unable to upsert oncall_schedule")
	}

	if !cached {
		oncallScheduleUpsertCacheMut.Lock()
		oncallScheduleUpsertCache[key] = cache
		oncallScheduleUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single OncallSchedule record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OncallSchedule) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no OncallSchedule provided for delete")
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), oncallSchedulePrimaryKeyMapping)
		sql = "DELETE FROM \"oncall_schedule\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"oncall_schedule\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(oncallScheduleType, oncallScheduleMapping, append(wl, oncallSchedulePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from oncall_schedule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for oncall_schedule")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q oncallScheduleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no oncallScheduleQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from oncall_schedule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for oncall_schedule")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OncallScheduleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oncallSchedulePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"oncall_schedule\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, oncallSchedulePrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oncallSchedulePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"oncall_schedule\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, oncallSchedulePrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from oncallSchedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for oncall_schedule")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OncallSchedule) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOncallSchedule(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OncallScheduleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OncallScheduleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oncallSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"oncall_schedule\".* FROM \"oncall_schedule\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, oncallSchedulePrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in OncallScheduleSlice")
	}

	*o = slice

	return nil
}

// OncallScheduleExists checks if the OncallSchedule row exists.
func OncallScheduleExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"oncall_schedule\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if oncall_schedule exists")
	}

	return exists, nil
}