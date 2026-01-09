package recipe

import (
	"context"

	"recipe/api/internal/presenter"
)

func main() {
	srv := presenter.NewServer()
	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}
