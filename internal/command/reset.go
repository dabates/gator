package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/types"
)

func ResetHandler(s *types.State, cmd Command) error {
	err := s.Db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	s.Config.SetUser("")
	fmt.Println("All data deleted")

	return nil
}
