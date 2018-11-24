package bugapp

import (
	"fmt"
	"os"
	"strings"
)

func Import(Args ArgumentList) {
	if (len(Args) != 2 && Args[0] == "--github") ||
	   (len(Args) != 1 && Args[0] == "--be") {
        fmt.Fprintf(os.Stderr, "Usage: %s import --github user/repo\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "       %s import --be\n", os.Args[0])
        fmt.Fprintf(os.Stderr, `
Use this tool to import an external bug database into the local
issues/ directory.

Either "--github user/repo" is required to import GitHub issues,
from GitHub, or "--be" is required to import local BugsEverywhere issues.

`)
        return
    }
	if githubRepo := Args.GetArgument("--github", ""); githubRepo != "" {
		if strings.Count(githubRepo, "/") != 1 {
			fmt.Fprintf(os.Stderr, "Invalid GitHub repo: %s\n", githubRepo)
			return
		}
		pieces := strings.Split(githubRepo, "/")
		githubImport(pieces[0], pieces[1])
    } else if Args.GetArgument("--be", "") != "" {
	    // if x < len(Args) {
		beImport()
    } 
}

