package types

// DuplicateFilesInfo holds the info about duplicate files
type DuplicateFilesInfo struct {
	Total int
	// No. of files which as at least one duplicate somewhere
	DuplicateFilesCount int
	// The original file size * (no. of duplicate copies - 1)
	DuplicateSize int64
	// No. of unique files, which will be left after deleting duplicates
	UniqueFilesCount int
}
