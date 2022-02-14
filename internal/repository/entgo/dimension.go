package entgo

import (
	"context"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/dimension"
)

func createItemDimensionsInTransaction(ctx context.Context, tx *ent.Tx, i *domain.Item) error {
	for _, d := range i.Dimensions {
		_, err := getOrCreateDimensionInTransaction(ctx, tx, d, i.ID)
		if err != nil {
			// tx already rollbacked here
			return err
		}
	}

	return nil
}

func getOrCreateDimensionInTransaction(ctx context.Context, tx *ent.Tx, d *domain.Dimension, itemID int) (*ent.Dimension, error) {
	// TODO: probably better switch to bulk operation
	dim, err := tx.Dimension.Query().
		Where(dimension.And(
			dimension.TitleEQ(d.Title),
			dimension.ValueEQ(d.Value),
		)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, rollback(tx, err)
		}
		dim, err := tx.Dimension.Create().
			AddItemIDs(itemID).
			SetTitle(d.Title).
			SetValue(d.Value).
			SetMeta(d.Meta).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}

		return dim, nil
	}

	dim, err = dim.Update().AddItemIDs(itemID).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return dim, nil
}
