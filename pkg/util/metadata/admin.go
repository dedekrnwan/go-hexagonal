package metadata

type (
	Admin interface {
	}
	admin struct{}
)

func NewAdmin() Admin {
	return &admin{}
}

// func (m *user) GetAdmin(ctx context.Context) (*entity.Admin, error) {
// 	data, ok := ctx.Value(metadataAdmin).(*entity.Admin)
// 	if !ok {
// 		return nil, fmt.Errorf("error getting metadata user from context")
// 	}
// 	return data, nil
// }

// func (m *user) SetAdmin(ctx context.Context, data *entity.Admin) context.Context {
// 	return context.WithValue(ctx, metadataAdmin, data)
// }
