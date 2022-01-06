package measuretype

import (
	"context"

	"github.com/sofisoft-tech/ms-measureunit/seedwork"
)

type Repository interface {
	FindById(ctx context.Context, id string, receiver interface{}) error
	seedwork.IBaseRepository
}
