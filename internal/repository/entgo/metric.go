package entgo

import (
	"context"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/metric"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
)

func (r *EntgoRepository) GetMetricCount(ctx context.Context) (int, error) {
	return r.client.Metric.Query().Count(ctx)
}

func getOrCreateMetricInTransaction(ctx context.Context, tx *ent.Tx, m *domain.Metric) (*ent.Metric, error) {
	metr, err := tx.Metric.Query().
		Where(metric.And(
			metric.TitleEQ(m.Title)),
			metric.HasTaskWith(task.IDEQ(m.TaskID)),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, rollback(tx, err)
		}
		metr, err := tx.Metric.Create().
			SetTaskID(m.TaskID).
			SetTitle(m.Title).
			SetDescription(m.Description).
			SetMeta(m.Meta).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}

		return metr, nil
	}

	return metr, nil
}
