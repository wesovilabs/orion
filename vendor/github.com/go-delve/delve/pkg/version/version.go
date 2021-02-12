package version

import "fmt"

// Version represents the current version of Delve.
type Version struct {
	Major    string
	Minor    string
	Patch    string
	Metadata string
	Build    string
}

var (
	// DelveVersion is the current version of Delve.
	DelveVersion = Version{
		Major: "1", Minor: "6", Patch: "0", Metadata: "",
		Build: "$Id: 8cc9751909843dd55a46e8ea2a561544f70db34d $",
	}
)

func (v Version) String() string {
	ver := fmt.Sprintf("Version: %s.%s.%s", v.Major, v.Minor, v.Patch)
	if v.Metadata != "" {
		ver += "-" + v.Metadata
	}
	return fmt.Sprintf("%s\nBuild: %s", ver, v.Build)
}
