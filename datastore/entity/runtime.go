// Code generated by entc, DO NOT EDIT.

package entity

import (
	"time"

	"github.com/syllabix/oncall/datastore/entity/schedule"
	"github.com/syllabix/oncall/datastore/entity/schema"
	"github.com/syllabix/oncall/datastore/entity/shift"
	"github.com/syllabix/oncall/datastore/entity/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescIsEnabled is the schema descriptor for is_enabled field.
	scheduleDescIsEnabled := scheduleFields[4].Descriptor()
	// schedule.DefaultIsEnabled holds the default value on creation for the is_enabled field.
	schedule.DefaultIsEnabled = scheduleDescIsEnabled.Default.(bool)
	// scheduleDescCreatedAt is the schema descriptor for created_at field.
	scheduleDescCreatedAt := scheduleFields[8].Descriptor()
	// schedule.DefaultCreatedAt holds the default value on creation for the created_at field.
	schedule.DefaultCreatedAt = scheduleDescCreatedAt.Default.(func() time.Time)
	// scheduleDescUpdatedAt is the schema descriptor for updated_at field.
	scheduleDescUpdatedAt := scheduleFields[9].Descriptor()
	// schedule.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	schedule.DefaultUpdatedAt = scheduleDescUpdatedAt.Default.(func() time.Time)
	// schedule.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	schedule.UpdateDefaultUpdatedAt = scheduleDescUpdatedAt.UpdateDefault.(func() time.Time)
	shiftFields := schema.Shift{}.Fields()
	_ = shiftFields
	// shiftDescStartedAt is the schema descriptor for started_at field.
	shiftDescStartedAt := shiftFields[4].Descriptor()
	// shift.DefaultStartedAt holds the default value on creation for the started_at field.
	shift.DefaultStartedAt = shiftDescStartedAt.Default.(func() time.Time)
	// shiftDescCreatedAt is the schema descriptor for created_at field.
	shiftDescCreatedAt := shiftFields[5].Descriptor()
	// shift.DefaultCreatedAt holds the default value on creation for the created_at field.
	shift.DefaultCreatedAt = shiftDescCreatedAt.Default.(func() time.Time)
	// shiftDescUpdatedAt is the schema descriptor for updated_at field.
	shiftDescUpdatedAt := shiftFields[6].Descriptor()
	// shift.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	shift.DefaultUpdatedAt = shiftDescUpdatedAt.Default.(func() time.Time)
	// shift.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	shift.UpdateDefaultUpdatedAt = shiftDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[8].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
