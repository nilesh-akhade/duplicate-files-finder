package main

import (
	"os"
	"path/filepath"
)

// DuplicateFilesFinder find the duplicate files and returns the
// info about duplicate files
type DuplicateFilesFinder interface {
	Find() (*DuplicateFilesInfo, error)
}

func New(dir string, recursive bool) DuplicateFilesFinder {
	return &duplicateFilesFinder{
		dir:          dir,
		checksumCalc: NewSHA1ChecksumCalc(),
	}
}

type duplicateFilesFinder struct {
	dir          string
	checksumCalc ChecksumCalc
}

type FileEntry struct {
	Path                 string
	IsChecksumCalculated bool
}

func (d *duplicateFilesFinder) Find() (*DuplicateFilesInfo, error) {
	fileSzToPath := make(map[int64]FileEntry)
	checksums := make(map[string]bool)
	dup := &DuplicateFilesInfo{}
	err := filepath.WalkDir(d.dir,
		func(path string, de os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if de.IsDir() || !de.Type().IsRegular() {
				return nil
			}
			info, err := de.Info()
			if err != nil {
				return err
			}
			dup.Total++
			if fc, found := fileSzToPath[info.Size()]; found {
				if !fc.IsChecksumCalculated {
					oldChecksum, err := d.checksumCalc.Calculate(fc.Path)
					if err != nil {
						return err
					}
					fc.IsChecksumCalculated = true
					checksums[oldChecksum] = true
				}
				fileChecksum, err := d.checksumCalc.Calculate(path)
				if err != nil {
					return err
				}
				if _, found := checksums[fileChecksum]; found {
					dup.DuplicateFilesCount++
					dup.DuplicateSize += info.Size()
					return nil
				}
				checksums[fileChecksum] = true
			} else {
				fileSzToPath[info.Size()] = FileEntry{Path: path}
			}
			dup.UniqueFilesCount++
			return nil
		})
	if err != nil {
		return nil, err
	}

	return dup, nil
}
