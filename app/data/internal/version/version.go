package version

import _ "embed"

//go:embed VERSION
var Version string

var GoVersion string

var CommitID string

var CommitDate string

var OS string

var Arch string
