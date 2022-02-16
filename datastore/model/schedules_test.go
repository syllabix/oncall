// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSchedules(t *testing.T) {
	t.Parallel()

	query := Schedules()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSchedulesSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchedulesQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Schedules().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchedulesSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScheduleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchedulesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchedulesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Schedules().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchedulesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScheduleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchedulesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ScheduleExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Schedule exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ScheduleExists to return true, but got false.")
	}
}

func testSchedulesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	scheduleFound, err := FindSchedule(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if scheduleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSchedulesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Schedules().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSchedulesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Schedules().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSchedulesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	scheduleOne := &Schedule{}
	scheduleTwo := &Schedule{}
	if err = randomize.Struct(seed, scheduleOne, scheduleDBTypes, false, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}
	if err = randomize.Struct(seed, scheduleTwo, scheduleDBTypes, false, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scheduleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scheduleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Schedules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSchedulesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	scheduleOne := &Schedule{}
	scheduleTwo := &Schedule{}
	if err = randomize.Struct(seed, scheduleOne, scheduleDBTypes, false, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}
	if err = randomize.Struct(seed, scheduleTwo, scheduleDBTypes, false, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scheduleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scheduleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testSchedulesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSchedulesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(scheduleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSchedulesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSchedulesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScheduleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSchedulesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Schedules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	scheduleDBTypes = map[string]string{`ID`: `uuid`, `SlackChannelID`: `text`, `TeamSlackID`: `text`, `Name`: `text`, `Interval`: `enum.shift_interval('daily','weekly','bi-weekly','monthly')`, `IsEnabled`: `boolean`, `StartTime`: `time with time zone`, `EndTime`: `time with time zone`, `ActiveShift`: `uuid`, `OverrideShift`: `uuid`, `Shifts`: `ARRAYtext`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testSchedulesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(schedulePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(scheduleAllColumns) == len(schedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scheduleDBTypes, true, schedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSchedulesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(scheduleAllColumns) == len(schedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Schedule{}
	if err = randomize.Struct(seed, o, scheduleDBTypes, true, scheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scheduleDBTypes, true, schedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(scheduleAllColumns, schedulePrimaryKeyColumns) {
		fields = scheduleAllColumns
	} else {
		fields = strmangle.SetComplement(
			scheduleAllColumns,
			schedulePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ScheduleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSchedulesUpsert(t *testing.T) {
	t.Parallel()

	if len(scheduleAllColumns) == len(schedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Schedule{}
	if err = randomize.Struct(seed, &o, scheduleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Schedule: %s", err)
	}

	count, err := Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, scheduleDBTypes, false, schedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Schedule struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Schedule: %s", err)
	}

	count, err = Schedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
