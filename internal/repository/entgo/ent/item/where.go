// Code generated by entc, DO NOT EDIT.

package item

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v float64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v float64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v float64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...float64) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...float64) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v float64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v float64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v float64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v float64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// MetaIsNil applies the IsNil predicate on the "meta" field.
func MetaIsNil() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMeta)))
	})
}

// MetaNotNil applies the NotNil predicate on the "meta" field.
func MetaNotNil() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMeta)))
	})
}

// HasDimensions applies the HasEdge predicate on the "dimensions" edge.
func HasDimensions() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DimensionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, DimensionsTable, DimensionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDimensionsWith applies the HasEdge predicate on the "dimensions" edge with a given conditions (other predicates).
func HasDimensionsWith(preds ...predicate.Dimension) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DimensionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, DimensionsTable, DimensionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaskInstance applies the HasEdge predicate on the "task_instance" edge.
func HasTaskInstance() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskInstanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskInstanceTable, TaskInstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskInstanceWith applies the HasEdge predicate on the "task_instance" edge with a given conditions (other predicates).
func HasTaskInstanceWith(preds ...predicate.TaskInstance) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskInstanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskInstanceTable, TaskInstanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetric applies the HasEdge predicate on the "metric" edge.
func HasMetric() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetricTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetricTable, MetricColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetricWith applies the HasEdge predicate on the "metric" edge with a given conditions (other predicates).
func HasMetricWith(preds ...predicate.Metric) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetricInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetricTable, MetricColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
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
func Not(p predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		p(s.Not())
	})
}
