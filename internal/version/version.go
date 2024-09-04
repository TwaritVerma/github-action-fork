package version

const versionPrefix = "golang-test-"

//go:generate go run gen.go
func GetVersion() string {
	return versionPrefix + version
}
