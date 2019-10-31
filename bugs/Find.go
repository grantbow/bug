package bugs

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

//var dops = Directory(os.PathSeparator)
//var sops = string(os.PathSeparator)

// BugNotFoundError defines a new error.
type BugNotFoundError string

// Error returns a string of the error.
func (b BugNotFoundError) Error() string {
	return string(b)
}

//var dops = Directory(os.PathSeparator)
//var sops = string(os.PathSeparator)

// FindBugsByTag returns an array of tagged issues.
func FindBugsByTag(tags []string, config Config) []Bug {
	root := RootDirer(&config)
	issuesroot := IssuesDirer(config)
	issues := readIssues(string(issuesroot))
	sort.Sort(byDir(issues))

	var bugs []Bug
	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(root+dops+Directory(config.IssuesDirName)+dops+Directory(issue.Name()), config)
			for _, tag := range tags {
				if bug.HasTag(TagBoolTrue(tag)) {
					bugs = append(bugs, bug)
					break
				}
			}
		}
	}
	//fmt.Printf("c %+v\n", bugs)
	sort.Sort(byBug(bugs)) // allowed by byBug defined in Bug.go
	//fmt.Printf("d %+v\n", bugs)
	return bugs
}

// LoadBugByDirectory returns an issue from the directory name.
func LoadBugByDirectory(dir string, config Config) (*Bug, error) {
	root := RootDirer(&config)
	_, err := ioutil.ReadDir(string(root) + sops + config.IssuesDirName + sops + dir)
	if err != nil {
		return nil, BugNotFoundError("Not found " + dir)
	}
	bug := Bug{}
	bug.LoadBug(IssuesDirer(config)+dops+Directory(dir), config)
	return &bug, nil
}

// LoadBugByHeuristic returns an issue.
func LoadBugByHeuristic(id string, config Config) (*Bug, error) {
	root := RootDirer(&config)
	issuesroot := IssuesDirer(config)
	issues := readIssues(string(issuesroot))
	sort.Sort(byDir(issues))
	//fmt.Printf("debug Found %v files\n", len(issues)) //, strings.Join(issues, ", "))
	//for _, issue := range issues {                    // idx not needed
	//	fmt.Printf("debug %v\n", issue.Name())
	//}

	idx, err := strconv.Atoi(id)
	if err == nil { // && idx > 0 && idx <= len(issues) {
		return LoadBugByIndex(idx, config)
	}

	var candidate *Bug
	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(root+dops+Directory(config.IssuesDirName)+dops+Directory(issue.Name()), config)
			if bugid := bug.Identifier(); bugid == id {
				return &bug, nil
			} else if strings.Index(bugid, id) >= 0 {
				candidate = &bug
			}

		}
	}
	if candidate != nil {
		return candidate, nil
	}
	return nil, BugNotFoundError("Not found " + id)
}

//// LoadBugByStringIndex returns an issue from a string index.
//func LoadBugByStringIndex(i string, config Config) (*Bug, error) {
//	root := RootDirer(&config)
//	issues, _ := ioutil.ReadDir(string(root) + sops + config.IssuesDirName)
//
//	idx, err := strconv.Atoi(i)
//	if err != nil {
//		return nil, BugNotFoundError("Index not a number")
//	}
//	if idx < 1 || idx > len(issues) {
//		return nil, BugNotFoundError("Invalid Index")
//	}
//
//	b := Bug{}
//	directoryString := fmt.Sprintf("%s%s%s%s%s", root, sops, config.IssuesDirName, sops, issues[idx-1].Name())
//	b.LoadBug(Directory(directoryString), config)
//	return &b, nil
//}

// LoadBugByIdentifier returns an issue from a string Identifier
func LoadBugByIdentifier(id string, config Config) (*Bug, error) {
	root := RootDirer(&config)
	issuesroot := IssuesDirer(config)
	issues := readIssues(string(issuesroot))
	sort.Sort(byDir(issues))

	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(root+dops+Directory(config.IssuesDirName)+dops+Directory(issue.Name()), config)
			if bug.Identifier() == id {
				return &bug, nil
			}
		}
	}
	return nil, BugNotFoundError("No bug named " + id)
}

// LoadBugByIndex returns an issue from an int index.
func LoadBugByIndex(idx int, config Config) (*Bug, error) {
	root := RootDirer(&config)
	issuesroot := IssuesDirer(config)
	issues := readIssues(string(issuesroot))
	sort.Sort(byDir(issues))
	if idx < 1 || idx > len(issues) {
		return nil, BugNotFoundError("Invalid bug index")
	}

	b := Bug{}
	directoryString := fmt.Sprintf("%s%s%s%s%s", root, sops, config.IssuesDirName, sops, issues[idx-1].Name())
	// TODO: fix, can return files that are non-issues located in the issues directory
	b.LoadBug(Directory(directoryString), config)
	return &b, nil
}

// GetAllBugs returns an array of all issues.
func GetAllBugs(config Config) []Bug {
	root := RootDirer(&config)
	issuesroot := IssuesDirer(config)
	issues := readIssues(string(issuesroot))
	sort.Sort(byDir(issues))
	//fmt.Printf("%+v\n", issues)

	var bugs []Bug
	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(root+dops+Directory(config.IssuesDirName)+dops+Directory(issue.Name()), config)
			bugs = append(bugs, bug)
		}
	}

	//fmt.Printf("a %+v\n", bugs)
	sort.Sort(byBug(bugs)) // allowed by byBug defined in Bug.go
	//fmt.Printf("b %+v\n", bugs)
	return bugs
}
