package entgo

import (
	"context"

	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/tasktag"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
)

func (r *EntgoRepository) GetOrCreateTaskTag(ctx context.Context, t *domain.TaskTag) (*domain.TaskTag, error) {
	tag, err := r.client.TaskTag.Query().
		Where(tasktag.TitleEQ(t.Title)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		tag, err := r.client.TaskTag.Create().
			SetTitle(t.Title).
			SetDescription(t.Description).
			SetMeta(t.Meta).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return entToDomainTaskTag(tag), nil
	}

	return entToDomainTaskTag(tag), nil
}

func entToDomainTaskTag(t *ent.TaskTag) *domain.TaskTag {
	return &domain.TaskTag{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Display:     t.Display,
		Meta:        t.Meta,
	}
}
