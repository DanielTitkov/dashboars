package entgo

import (
	"context"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/taskcategory"
)

func (r *EntgoRepository) GetOrCreateTaskCategory(ctx context.Context, c *domain.TaskCategory) (*domain.TaskCategory, error) {
	cat, err := r.client.TaskCategory.Query().
		Where(taskcategory.TitleEQ(c.Title)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		cat, err := r.client.TaskCategory.Create().
			SetTitle(c.Title).
			SetDescription(c.Description).
			SetMeta(c.Meta).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return entToDomainTaskCategory(cat), nil
	}

	return entToDomainTaskCategory(cat), nil
}

func entToDomainTaskCategory(t *ent.TaskCategory) *domain.TaskCategory {
	return &domain.TaskCategory{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Display:     t.Display,
		Meta:        t.Meta,
	}
}
