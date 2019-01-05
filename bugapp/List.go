package bugapp

import (
	"fmt"
	"github.com/driusan/bug/bugs"
	"io/ioutil"
	"os"
)

func getBugName(b bugs.Bug, idx int) string {
	if id := b.Identifier(); id != "" {
		return fmt.Sprintf("Issue %s", id)
	}
	return fmt.Sprintf("Issue %d", idx+1)
}
func listTags(files []os.FileInfo, args ArgumentList, config bugs.Config) {
	b := bugs.Bug{}
	for idx := range files {
		b.LoadBug(bugs.Directory(bugs.GetIssuesDir(config) + bugs.Directory(files[idx].Name())))

		for _, tag := range args {
			if b.HasTag(bugs.Tag(tag)) {
				fmt.Printf("%s: %s\n", getBugName(b, idx), b.Title("tags"))
			}
		}
	}
}
func List(args ArgumentList, config bugs.Config) {
	issuesroot := bugs.GetIssuesDir(config)
	issues, _ := ioutil.ReadDir(string(issuesroot))

	var wantTags bool = false
	if args.HasArgument("--tags") {
		wantTags = true
	}

	// No parameters, print a list of all bugs
	if len(args) == 0 || (wantTags && len(args) == 1) {
		//os.Stdout = stdout
		for idx, issue := range issues {
			if issue.IsDir() != true {
				continue
			}
			var dir bugs.Directory = issuesroot + bugs.Directory(issue.Name())
			b := bugs.Bug{Dir: dir}
			name := getBugName(b, idx)
			if wantTags == false {
				fmt.Printf("%s: %s\n", name, b.Title(""))
			} else {
				fmt.Printf("%s: %s\n", name, b.Title("tags"))
			}
		}
		return
	}

	// getAllTags() is defined in Tag.go
	// Get a list of tags, so that when we encounter
	// an error we can check if it's because the user
	// provided a tagname instead of a BugID. If they
	// did, then list bugs matching that tag instead
	// of full descriptions
	tags := getAllTags(config)
	// There were parameters, so show the full description of each
	// of those issues
	for i, length := 0, len(args); i < length; i += 1 {
		b, err := bugs.LoadBugByHeuristic(args[i], config)
		if err != nil {
			for _, tagname := range tags {
				if tagname == args[i] {
					listTags(issues, args, config)
					return
				}
			}
			fmt.Printf("%s\n", err.Error())
			continue
		}

		b.ViewBug()
		if i < length-1 {
			fmt.Printf("\n--\n\n")
		}
	}
	fmt.Printf("\n")
}
