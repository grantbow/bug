package fitapp

import (
	"fmt"
	bugs "github.com/driusan/bug/bugs"
	"github.com/driusan/bug/scm"
)

// Purge is a subcommand to delete all issues.
func Purge(config bugs.Config) {
	scm, _, err := scm.DetectSCM(make(map[string]bool), config)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	err = scm.Purge(bugs.FitDirer(config))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}
