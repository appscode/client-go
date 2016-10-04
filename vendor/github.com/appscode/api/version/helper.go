package version

import "fmt"

func (m *Version) Print() {
	fmt.Printf("Name = %v\n", m.Name)
	fmt.Printf("Version = %v\n", m.Version)
	fmt.Printf("Os = %v\n", m.Os)
	fmt.Printf("Arch = %v\n", m.Arch)

	fmt.Printf("CommitHash = %v\n", m.CommitHash)
	fmt.Printf("GitBranch = %v\n", m.GitBranch)
	fmt.Printf("GitTag = %v\n", m.GitTag)
	fmt.Printf("CommitTimestamp = %v\n", m.CommitTimestamp)

	fmt.Printf("BuildTimestamp = %v\n", m.BuildTimestamp)
	fmt.Printf("BuildHost = %v\n", m.BuildHost)
	fmt.Printf("BuildHostOs = %v\n", m.BuildHostOs)
	fmt.Printf("BuildHostArch = %v\n", m.BuildHostArch)
}
