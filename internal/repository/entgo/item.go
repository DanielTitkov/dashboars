package entgo

import "context"

func (r *EntgoRepository) GetItemCount(ctx context.Context) (int, error) {
	return r.client.Item.Query().Count(ctx)
}
