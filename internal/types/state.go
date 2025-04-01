package types

import (
	. "github.com/dabates/gator/internal/config"
	"github.com/dabates/gator/internal/database"
)

type State struct {
	Config *Config
	Db     *database.Queries
}
