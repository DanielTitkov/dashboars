// Code generated by entc, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// Display applies equality check predicate on the "display" field. It's identical to DisplayEQ.
func Display(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplay), v))
	})
}

// Schedule applies equality check predicate on the "schedule" field. It's identical to ScheduleEQ.
func Schedule(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchedule), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// DisplayEQ applies the EQ predicate on the "display" field.
func DisplayEQ(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplay), v))
	})
}

// DisplayNEQ applies the NEQ predicate on the "display" field.
func DisplayNEQ(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisplay), v))
	})
}

// ScheduleEQ applies the EQ predicate on the "schedule" field.
func ScheduleEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchedule), v))
	})
}

// ScheduleNEQ applies the NEQ predicate on the "schedule" field.
func ScheduleNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSchedule), v))
	})
}

// ScheduleIn applies the In predicate on the "schedule" field.
func ScheduleIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSchedule), v...))
	})
}

// ScheduleNotIn applies the NotIn predicate on the "schedule" field.
func ScheduleNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSchedule), v...))
	})
}

// ScheduleGT applies the GT predicate on the "schedule" field.
func ScheduleGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSchedule), v))
	})
}

// ScheduleGTE applies the GTE predicate on the "schedule" field.
func ScheduleGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSchedule), v))
	})
}

// ScheduleLT applies the LT predicate on the "schedule" field.
func ScheduleLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSchedule), v))
	})
}

// ScheduleLTE applies the LTE predicate on the "schedule" field.
func ScheduleLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSchedule), v))
	})
}

// ScheduleContains applies the Contains predicate on the "schedule" field.
func ScheduleContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSchedule), v))
	})
}

// ScheduleHasPrefix applies the HasPrefix predicate on the "schedule" field.
func ScheduleHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSchedule), v))
	})
}

// ScheduleHasSuffix applies the HasSuffix predicate on the "schedule" field.
func ScheduleHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSchedule), v))
	})
}

// ScheduleIsNil applies the IsNil predicate on the "schedule" field.
func ScheduleIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSchedule)))
	})
}

// ScheduleNotNil applies the NotNil predicate on the "schedule" field.
func ScheduleNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSchedule)))
	})
}

// ScheduleEqualFold applies the EqualFold predicate on the "schedule" field.
func ScheduleEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSchedule), v))
	})
}

// ScheduleContainsFold applies the ContainsFold predicate on the "schedule" field.
func ScheduleContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSchedule), v))
	})
}

// ArgsIsNil applies the IsNil predicate on the "args" field.
func ArgsIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldArgs)))
	})
}

// ArgsNotNil applies the NotNil predicate on the "args" field.
func ArgsNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldArgs)))
	})
}

// HasInstances applies the HasEdge predicate on the "instances" edge.
func HasInstances() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstancesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstancesWith applies the HasEdge predicate on the "instances" edge with a given conditions (other predicates).
func HasInstancesWith(preds ...predicate.TaskInstance) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstancesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetrics applies the HasEdge predicate on the "metrics" edge.
func HasMetrics() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetricsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetricsTable, MetricsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetricsWith applies the HasEdge predicate on the "metrics" edge with a given conditions (other predicates).
func HasMetricsWith(preds ...predicate.Metric) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetricsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetricsTable, MetricsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.TaskCategory) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.TaskTag) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
