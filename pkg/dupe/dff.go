package dupe

import (
	"os"
	"path/filepath"

	"github.com/nilesh-akhade/duplicate-files-finder/pkg/checksum"
	"github.com/nilesh-akhade/duplicate-files-finder/pkg/types"
)

// DuplicateFilesFinder find the duplicate files and returns the
// info about duplicate files
type DuplicateFilesFinder interface {
	Find() (*types.DuplicateFilesInfo, error)
}

func New(dir string, recursive bool) DuplicateFilesFinder {
	return &duplicateFilesFinder{
		dir:          dir,
		checksumCalc: checksum.NewSHA1Checksum(),
	}
}

type duplicateFilesFinder struct {
	dir          string
	checksumCalc checksum.ChecksumCalc
}

func (d *duplicateFilesFinder) Find() (*types.DuplicateFilesInfo, error) {
	type FileChecksum struct {
		Path           string
		IsChecksumCalc bool
	}
	knownFileSizes := make(map[int64]FileChecksum)
	knownChecksums := make(map[string]int64)
	dup := &types.DuplicateFilesInfo{}
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
			if fc, found := knownFileSizes[info.Size()]; found {
				if !fc.IsChecksumCalc {
					oldChecksum, err := d.checksumCalc.Calculate(fc.Path)
					if err != nil {
						return err
					}
					knownChecksums[oldChecksum] = info.Size()
				}
				currentFileChecksum, err := d.checksumCalc.Calculate(fc.Path)
				if err != nil {
					return err
				}
				if size, found := knownChecksums[currentFileChecksum]; found {
					dup.DuplicateFilesCount++
					dup.DuplicateSize += size
					return nil
				}
			} else {
				knownFileSizes[info.Size()] = FileChecksum{Path: path}
			}
			dup.UniqueFilesCount++
			return nil
		})
	if err != nil {
		return nil, err
	}

	return dup, nil
}
