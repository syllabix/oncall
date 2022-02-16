// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Schedules", testSchedules)
	t.Run("Shifts", testShifts)
	t.Run("TeamMembers", testTeamMembers)
	t.Run("Teams", testTeams)
	t.Run("Users", testUsers)
}

func TestSoftDelete(t *testing.T) {
	t.Run("Schedules", testSchedulesSoftDelete)
	t.Run("Shifts", testShiftsSoftDelete)
	t.Run("Teams", testTeamsSoftDelete)
	t.Run("Users", testUsersSoftDelete)
}

func TestQuerySoftDeleteAll(t *testing.T) {
	t.Run("Schedules", testSchedulesQuerySoftDeleteAll)
	t.Run("Shifts", testShiftsQuerySoftDeleteAll)
	t.Run("Teams", testTeamsQuerySoftDeleteAll)
	t.Run("Users", testUsersQuerySoftDeleteAll)
}

func TestSliceSoftDeleteAll(t *testing.T) {
	t.Run("Schedules", testSchedulesSliceSoftDeleteAll)
	t.Run("Shifts", testShiftsSliceSoftDeleteAll)
	t.Run("Teams", testTeamsSliceSoftDeleteAll)
	t.Run("Users", testUsersSliceSoftDeleteAll)
}

func TestDelete(t *testing.T) {
	t.Run("Schedules", testSchedulesDelete)
	t.Run("Shifts", testShiftsDelete)
	t.Run("TeamMembers", testTeamMembersDelete)
	t.Run("Teams", testTeamsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Schedules", testSchedulesQueryDeleteAll)
	t.Run("Shifts", testShiftsQueryDeleteAll)
	t.Run("TeamMembers", testTeamMembersQueryDeleteAll)
	t.Run("Teams", testTeamsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Schedules", testSchedulesSliceDeleteAll)
	t.Run("Shifts", testShiftsSliceDeleteAll)
	t.Run("TeamMembers", testTeamMembersSliceDeleteAll)
	t.Run("Teams", testTeamsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Schedules", testSchedulesExists)
	t.Run("Shifts", testShiftsExists)
	t.Run("TeamMembers", testTeamMembersExists)
	t.Run("Teams", testTeamsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Schedules", testSchedulesFind)
	t.Run("Shifts", testShiftsFind)
	t.Run("TeamMembers", testTeamMembersFind)
	t.Run("Teams", testTeamsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Schedules", testSchedulesBind)
	t.Run("Shifts", testShiftsBind)
	t.Run("TeamMembers", testTeamMembersBind)
	t.Run("Teams", testTeamsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Schedules", testSchedulesOne)
	t.Run("Shifts", testShiftsOne)
	t.Run("TeamMembers", testTeamMembersOne)
	t.Run("Teams", testTeamsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Schedules", testSchedulesAll)
	t.Run("Shifts", testShiftsAll)
	t.Run("TeamMembers", testTeamMembersAll)
	t.Run("Teams", testTeamsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Schedules", testSchedulesCount)
	t.Run("Shifts", testShiftsCount)
	t.Run("TeamMembers", testTeamMembersCount)
	t.Run("Teams", testTeamsCount)
	t.Run("Users", testUsersCount)
}

func TestInsert(t *testing.T) {
	t.Run("Schedules", testSchedulesInsert)
	t.Run("Schedules", testSchedulesInsertWhitelist)
	t.Run("Shifts", testShiftsInsert)
	t.Run("Shifts", testShiftsInsertWhitelist)
	t.Run("TeamMembers", testTeamMembersInsert)
	t.Run("TeamMembers", testTeamMembersInsertWhitelist)
	t.Run("Teams", testTeamsInsert)
	t.Run("Teams", testTeamsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Schedules", testSchedulesReload)
	t.Run("Shifts", testShiftsReload)
	t.Run("TeamMembers", testTeamMembersReload)
	t.Run("Teams", testTeamsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Schedules", testSchedulesReloadAll)
	t.Run("Shifts", testShiftsReloadAll)
	t.Run("TeamMembers", testTeamMembersReloadAll)
	t.Run("Teams", testTeamsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Schedules", testSchedulesSelect)
	t.Run("Shifts", testShiftsSelect)
	t.Run("TeamMembers", testTeamMembersSelect)
	t.Run("Teams", testTeamsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Schedules", testSchedulesUpdate)
	t.Run("Shifts", testShiftsUpdate)
	t.Run("TeamMembers", testTeamMembersUpdate)
	t.Run("Teams", testTeamsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Schedules", testSchedulesSliceUpdateAll)
	t.Run("Shifts", testShiftsSliceUpdateAll)
	t.Run("TeamMembers", testTeamMembersSliceUpdateAll)
	t.Run("Teams", testTeamsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
