// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/schema"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/taskinstance"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	itemMixin := schema.Item{}.Mixin()
	itemMixinFields0 := itemMixin[0].Fields()
	_ = itemMixinFields0
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescCreateTime is the schema descriptor for create_time field.
	itemDescCreateTime := itemMixinFields0[0].Descriptor()
	// item.DefaultCreateTime holds the default value on creation for the create_time field.
	item.DefaultCreateTime = itemDescCreateTime.Default.(func() time.Time)
	// itemDescUpdateTime is the schema descriptor for update_time field.
	itemDescUpdateTime := itemMixinFields0[1].Descriptor()
	// item.DefaultUpdateTime holds the default value on creation for the update_time field.
	item.DefaultUpdateTime = itemDescUpdateTime.Default.(func() time.Time)
	// item.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	item.UpdateDefaultUpdateTime = itemDescUpdateTime.UpdateDefault.(func() time.Time)
	// itemDescTimestamp is the schema descriptor for timestamp field.
	itemDescTimestamp := itemFields[1].Descriptor()
	// item.DefaultTimestamp holds the default value on creation for the timestamp field.
	item.DefaultTimestamp = itemDescTimestamp.Default.(func() time.Time)
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreateTime is the schema descriptor for create_time field.
	taskDescCreateTime := taskMixinFields0[0].Descriptor()
	// task.DefaultCreateTime holds the default value on creation for the create_time field.
	task.DefaultCreateTime = taskDescCreateTime.Default.(func() time.Time)
	// taskDescUpdateTime is the schema descriptor for update_time field.
	taskDescUpdateTime := taskMixinFields0[1].Descriptor()
	// task.DefaultUpdateTime holds the default value on creation for the update_time field.
	task.DefaultUpdateTime = taskDescUpdateTime.Default.(func() time.Time)
	// task.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	task.UpdateDefaultUpdateTime = taskDescUpdateTime.UpdateDefault.(func() time.Time)
	// taskDescCode is the schema descriptor for code field.
	taskDescCode := taskFields[1].Descriptor()
	// task.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	task.CodeValidator = taskDescCode.Validators[0].(func(string) error)
	// taskDescTitle is the schema descriptor for title field.
	taskDescTitle := taskFields[2].Descriptor()
	// task.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	task.TitleValidator = taskDescTitle.Validators[0].(func(string) error)
	// taskDescDescription is the schema descriptor for description field.
	taskDescDescription := taskFields[3].Descriptor()
	// task.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	task.DescriptionValidator = taskDescDescription.Validators[0].(func(string) error)
	// taskDescActive is the schema descriptor for active field.
	taskDescActive := taskFields[4].Descriptor()
	// task.DefaultActive holds the default value on creation for the active field.
	task.DefaultActive = taskDescActive.Default.(bool)
	// taskDescDisplay is the schema descriptor for display field.
	taskDescDisplay := taskFields[5].Descriptor()
	// task.DefaultDisplay holds the default value on creation for the display field.
	task.DefaultDisplay = taskDescDisplay.Default.(bool)
	taskinstanceMixin := schema.TaskInstance{}.Mixin()
	taskinstanceMixinFields0 := taskinstanceMixin[0].Fields()
	_ = taskinstanceMixinFields0
	taskinstanceFields := schema.TaskInstance{}.Fields()
	_ = taskinstanceFields
	// taskinstanceDescCreateTime is the schema descriptor for create_time field.
	taskinstanceDescCreateTime := taskinstanceMixinFields0[0].Descriptor()
	// taskinstance.DefaultCreateTime holds the default value on creation for the create_time field.
	taskinstance.DefaultCreateTime = taskinstanceDescCreateTime.Default.(func() time.Time)
	// taskinstanceDescUpdateTime is the schema descriptor for update_time field.
	taskinstanceDescUpdateTime := taskinstanceMixinFields0[1].Descriptor()
	// taskinstance.DefaultUpdateTime holds the default value on creation for the update_time field.
	taskinstance.DefaultUpdateTime = taskinstanceDescUpdateTime.Default.(func() time.Time)
	// taskinstance.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	taskinstance.UpdateDefaultUpdateTime = taskinstanceDescUpdateTime.UpdateDefault.(func() time.Time)
	// taskinstanceDescStartTime is the schema descriptor for start_time field.
	taskinstanceDescStartTime := taskinstanceFields[0].Descriptor()
	// taskinstance.DefaultStartTime holds the default value on creation for the start_time field.
	taskinstance.DefaultStartTime = taskinstanceDescStartTime.Default.(func() time.Time)
	// taskinstanceDescAttempt is the schema descriptor for attempt field.
	taskinstanceDescAttempt := taskinstanceFields[2].Descriptor()
	// taskinstance.DefaultAttempt holds the default value on creation for the attempt field.
	taskinstance.DefaultAttempt = taskinstanceDescAttempt.Default.(int)
	// taskinstanceDescRunning is the schema descriptor for running field.
	taskinstanceDescRunning := taskinstanceFields[4].Descriptor()
	// taskinstance.DefaultRunning holds the default value on creation for the running field.
	taskinstance.DefaultRunning = taskinstanceDescRunning.Default.(bool)
}
