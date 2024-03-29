package cli

import (
	"errors"
	"fmt"
)

func Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("cli.Execute: args should not be zero size")
	}
	switch args[0] {
	case HTTPServerRunCommand:
		if err := runHTTPServerCommand(); err != nil {
			return err
		}
	case RunSystemCommand:
		if err := runSystemCommand(); err != nil {
			return err
		}
	case RunGenerateReligionCommand:
		if err := runGenerateReligionCommand(); err != nil {
			return err
		}
	case RunGenerateWorldCommand:
		if err := runGenerateWorldCommand(); err != nil {
			return err
		}
	case RunRefreshDataCommand:
		if err := runRefreshDataCommand(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cli.Execute: not found command = %s", args[0])
	}

	return nil
}
