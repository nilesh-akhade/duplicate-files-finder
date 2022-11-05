package types

// DuplicateFilesInfo holds the info about duplicate files as requested by the client
type DuplicateFilesInfo struct {
	Total int
	// Duplicate files that are copies of the original one
	DuplicateFilesCount int
	DuplicateSize       int64
	UniqueFilesCount    int
}
