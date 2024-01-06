package health

import (
	"context"
	"net/http"
	"time"

	healthcheck "github.com/hellofresh/health-go/v5"
	"github.com/jessedearing/service-catalog/internal/vars"
)

func NewHealthHandler() (http.Handler, error) {
	h, err := healthcheck.New()
	if err != nil {
		return nil, err
	}

	err = h.Register(healthcheck.Config{
		Name:      "postgres (via jackc/pgx/v5/pgpool.Ping)",
		Timeout:   2 * time.Second,
		SkipOnErr: false,
		Check: func(ctx context.Context) error {
			db, err := vars.GetDBFromContext(ctx)
			if err != nil {
				return err
			}

			return db.DB.Ping(ctx)
		},
	})

	if err != nil {
		return nil, err
	}

	return h.Handler(), nil
}
