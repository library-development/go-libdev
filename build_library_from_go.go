package libdev

// BuildLibraryFromGo builds a Library from a golang.Package.
// An error is returned if the golang.Package is not compatible with lib.dev.
// If the package imports any other packages, they must be on the GoPackageAllowList.
func BuildLibraryFromGo(golang.Package) (*Library, error) {

}
