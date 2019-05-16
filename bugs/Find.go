package bugs

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// BugNotFoundError defines a new error.
type BugNotFoundError string

// Error returns a string of the error.
func (b BugNotFoundError) Error() string {
	return string(b)
}

// FindBugsByTag returns an array of tagged issues.
func FindBugsByTag(tags []string, config Config) []Bug {
	root := RootDirer(config)
	issues, _ := ioutil.ReadDir(string(root) + "/issues")

	var bugs []Bug
	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(Directory(root + "/issues/" + Directory(issue.Name())))
			for _, tag := range tags {
				if bug.HasTag(TagBoolTrue(tag)) {
					bugs = append(bugs, bug)
					break
				}
			}
		}
	}
	//fmt.Printf("c %+v\n", bugs)
	sort.Sort(byBug(bugs))
	//fmt.Printf("d %+v\n", bugs)
	return bugs
}

// LoadBugByDirectory returns an issue from the directory name.
func LoadBugByDirectory(dir string, config Config) (*Bug, error) {
	root := RootDirer(config)
	_, err := ioutil.ReadDir(string(root) + "/issues/" + dir)
	if err != nil {
		return nil, BugNotFoundError("Not found " + dir)
	}
	bug := Bug{}
	bug.LoadBug(IssuesDirer(config) + "/" + Directory(dir))
	return &bug, nil
}

// LoadBugByHeuristic returns an issue.
func LoadBugByHeuristic(id string, config Config) (*Bug, error) {
	root := RootDirer(config)
	issues, _ := ioutil.ReadDir(string(root) + "/issues")

	idx, err := strconv.Atoi(id)
	if err == nil { // && idx > 0 && idx <= len(issues) {
		return LoadBugByIndex(idx, config)
	}

	var candidate *Bug
	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(Directory(root + "/issues/" + Directory(issue.Name())))
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

// LoadBugByStringIndex returns an issue from a string index.
func LoadBugByStringIndex(i string, config Config) (*Bug, error) {
	root := RootDirer(config)
	issues, _ := ioutil.ReadDir(string(root) + "/issues")

	idx, err := strconv.Atoi(i)
	if err != nil {
		return nil, BugNotFoundError("Index not a number")
	}
	if idx < 1 || idx > len(issues) {
		return nil, BugNotFoundError("Invalid Index")
	}

	b := Bug{}
	directoryString := fmt.Sprintf("%s%s%s", root, "/issues/", issues[idx-1].Name())
	b.LoadBug(Directory(directoryString))
	return &b, nil
}

// LoadBugByIdentifier returns an issue from a string Identifier
func LoadBugByIdentifier(id string, config Config) (*Bug, error) {
	root := RootDirer(config)
	issues, _ := ioutil.ReadDir(string(root) + "/issues")

	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(Directory(root + "/issues/" + Directory(issue.Name())))
			if bug.Identifier() == id {
				return &bug, nil
			}
		}
	}
	return nil, BugNotFoundError("No bug named " + id)
}

// LoadBugByIndex returns an issue from an int index.
func LoadBugByIndex(idx int, config Config) (*Bug, error) {
	root := RootDirer(config)
	issues, _ := ioutil.ReadDir(string(root) + "/issues")
	if idx < 1 || idx > len(issues) {
		return nil, BugNotFoundError("Invalid bug index")
	}

	b := Bug{}
	directoryString := fmt.Sprintf("%s%s%s", root, "/issues/", issues[idx-1].Name())
	b.LoadBug(Directory(directoryString))
	return &b, nil
}

// GetAllBugs returns an array of all issues.
func GetAllBugs(config Config) []Bug {
	root := RootDirer(config)
	issues, _ := ioutil.ReadDir(string(root) + "/issues")
	//fmt.Printf("%+v\n", issues)

	var bugs []Bug
	for _, issue := range issues { // idx not needed
		if issue.IsDir() == true {
			bug := Bug{}
			bug.LoadBug(Directory(root + "/issues/" + Directory(issue.Name())))
			bugs = append(bugs, bug)
		}
	}

	//fmt.Printf("a %+v\n", bugs)
	sort.Sort(byBug(bugs))
	//fmt.Printf("b %+v\n", bugs)
	return bugs
}
