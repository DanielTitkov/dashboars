// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DimensionsColumns holds the columns for the "dimensions" table.
	DimensionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "display_title", Type: field.TypeJSON, Nullable: true},
		{Name: "display_value", Type: field.TypeJSON, Nullable: true},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
	}
	// DimensionsTable holds the schema information for the "dimensions" table.
	DimensionsTable = &schema.Table{
		Name:       "dimensions",
		Columns:    DimensionsColumns,
		PrimaryKey: []*schema.Column{DimensionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "dimension_title_value",
				Unique:  true,
				Columns: []*schema.Column{DimensionsColumns[2], DimensionsColumns[3]},
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "value", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "metric_items", Type: field.TypeInt, Nullable: true},
		{Name: "task_instance_items", Type: field.TypeInt, Nullable: true},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "items_metrics_items",
				Columns:    []*schema.Column{ItemsColumns[6]},
				RefColumns: []*schema.Column{MetricsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "items_task_instances_items",
				Columns:    []*schema.Column{ItemsColumns[7]},
				RefColumns: []*schema.Column{TaskInstancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MetricsColumns holds the columns for the "metrics" table.
	MetricsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "display_title", Type: field.TypeJSON, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 280},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "task_metrics", Type: field.TypeInt, Nullable: true},
	}
	// MetricsTable holds the schema information for the "metrics" table.
	MetricsTable = &schema.Table{
		Name:       "metrics",
		Columns:    MetricsColumns,
		PrimaryKey: []*schema.Column{MetricsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "metrics_tasks_metrics",
				Columns:    []*schema.Column{MetricsColumns[7]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "metric_title_task_metrics",
				Unique:  true,
				Columns: []*schema.Column{MetricsColumns[3], MetricsColumns[7]},
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"random", "failing", "parser"}},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 280},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "display", Type: field.TypeBool, Default: false},
		{Name: "schedule", Type: field.TypeString, Nullable: true},
		{Name: "args", Type: field.TypeJSON, Nullable: true},
		{Name: "task_category_tasks", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_task_categories_tasks",
				Columns:    []*schema.Column{TasksColumns[11]},
				RefColumns: []*schema.Column{TaskCategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TaskCategoriesColumns holds the columns for the "task_categories" table.
	TaskCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 280},
		{Name: "display", Type: field.TypeBool, Default: false},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
	}
	// TaskCategoriesTable holds the schema information for the "task_categories" table.
	TaskCategoriesTable = &schema.Table{
		Name:       "task_categories",
		Columns:    TaskCategoriesColumns,
		PrimaryKey: []*schema.Column{TaskCategoriesColumns[0]},
	}
	// TaskInstancesColumns holds the columns for the "task_instances" table.
	TaskInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "attempt", Type: field.TypeInt, Default: 0},
		{Name: "success", Type: field.TypeBool, Nullable: true},
		{Name: "running", Type: field.TypeBool, Default: false},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "task_instances", Type: field.TypeInt, Nullable: true},
	}
	// TaskInstancesTable holds the schema information for the "task_instances" table.
	TaskInstancesTable = &schema.Table{
		Name:       "task_instances",
		Columns:    TaskInstancesColumns,
		PrimaryKey: []*schema.Column{TaskInstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_instances_tasks_instances",
				Columns:    []*schema.Column{TaskInstancesColumns[10]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TaskTagsColumns holds the columns for the "task_tags" table.
	TaskTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 280},
		{Name: "display", Type: field.TypeBool, Default: false},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
	}
	// TaskTagsTable holds the schema information for the "task_tags" table.
	TaskTagsTable = &schema.Table{
		Name:       "task_tags",
		Columns:    TaskTagsColumns,
		PrimaryKey: []*schema.Column{TaskTagsColumns[0]},
	}
	// ItemDimensionsColumns holds the columns for the "item_dimensions" table.
	ItemDimensionsColumns = []*schema.Column{
		{Name: "item_id", Type: field.TypeInt},
		{Name: "dimension_id", Type: field.TypeInt},
	}
	// ItemDimensionsTable holds the schema information for the "item_dimensions" table.
	ItemDimensionsTable = &schema.Table{
		Name:       "item_dimensions",
		Columns:    ItemDimensionsColumns,
		PrimaryKey: []*schema.Column{ItemDimensionsColumns[0], ItemDimensionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "item_dimensions_item_id",
				Columns:    []*schema.Column{ItemDimensionsColumns[0]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "item_dimensions_dimension_id",
				Columns:    []*schema.Column{ItemDimensionsColumns[1]},
				RefColumns: []*schema.Column{DimensionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TaskTagTasksColumns holds the columns for the "task_tag_tasks" table.
	TaskTagTasksColumns = []*schema.Column{
		{Name: "task_tag_id", Type: field.TypeInt},
		{Name: "task_id", Type: field.TypeInt},
	}
	// TaskTagTasksTable holds the schema information for the "task_tag_tasks" table.
	TaskTagTasksTable = &schema.Table{
		Name:       "task_tag_tasks",
		Columns:    TaskTagTasksColumns,
		PrimaryKey: []*schema.Column{TaskTagTasksColumns[0], TaskTagTasksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_tag_tasks_task_tag_id",
				Columns:    []*schema.Column{TaskTagTasksColumns[0]},
				RefColumns: []*schema.Column{TaskTagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "task_tag_tasks_task_id",
				Columns:    []*schema.Column{TaskTagTasksColumns[1]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DimensionsTable,
		ItemsTable,
		MetricsTable,
		TasksTable,
		TaskCategoriesTable,
		TaskInstancesTable,
		TaskTagsTable,
		ItemDimensionsTable,
		TaskTagTasksTable,
	}
)

func init() {
	ItemsTable.ForeignKeys[0].RefTable = MetricsTable
	ItemsTable.ForeignKeys[1].RefTable = TaskInstancesTable
	MetricsTable.ForeignKeys[0].RefTable = TasksTable
	TasksTable.ForeignKeys[0].RefTable = TaskCategoriesTable
	TaskInstancesTable.ForeignKeys[0].RefTable = TasksTable
	ItemDimensionsTable.ForeignKeys[0].RefTable = ItemsTable
	ItemDimensionsTable.ForeignKeys[1].RefTable = DimensionsTable
	TaskTagTasksTable.ForeignKeys[0].RefTable = TaskTagsTable
	TaskTagTasksTable.ForeignKeys[1].RefTable = TasksTable
}
