package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

// DuplicateFilesInfo holds the info about duplicate files
type DuplicateFilesInfo struct {
	Total            int
	UniqueFilesCount int

	// If a file is duplicated(any no. of times) then it is counted once
	DuplicateFilesCount int
	// The original file size * (no. of duplicate copies - 1)
	DuplicateSize int64
}

type fileEntry struct {
	Path                 string
	IsChecksumCalculated bool
}

func New(dir string, recursive bool) *duplicateFilesFinder {
	return &duplicateFilesFinder{
		dir:          dir,
		checksumCalc: NewSHA1ChecksumCalc(),
		recursive:    recursive,
	}
}

type duplicateFilesFinder struct {
	dir          string
	recursive    bool
	checksumCalc ChecksumCalc

	fileSzToEntry map[int64]fileEntry
	checksums     map[string]bool

	dupInfo *DuplicateFilesInfo
}

func (d *duplicateFilesFinder) Find() (*DuplicateFilesInfo, error) {
	d.fileSzToEntry = make(map[int64]fileEntry)
	d.checksums = make(map[string]bool)
	d.dupInfo = &DuplicateFilesInfo{}
	var err error
	if d.recursive {
		err = filepath.WalkDir(d.dir, d.walkDirEntries)
	} else {
		err = d.findDuplicatesNonRecursive()
	}
	return d.dupInfo, err
}

func (d *duplicateFilesFinder) walkDirEntries(path string, de os.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if de.IsDir() || !de.Type().IsRegular() {
		return nil
	}
	fileInfo, err := de.Info()
	if err != nil {
		return err
	}
	return d.processFileInfo(path, fileInfo)
}

func (d *duplicateFilesFinder) findDuplicatesNonRecursive() error {
	dirEntries, err := os.ReadDir(d.dir)
	if err != nil {
		return err
	}

	var path string
	var fileInfo fs.FileInfo
	for _, de := range dirEntries {
		path = filepath.Join(d.dir, de.Name())
		if de.IsDir() || !de.Type().IsRegular() {
			continue
		}
		fileInfo, err = de.Info()
		if err != nil {
			return err
		}
		err = d.processFileInfo(path, fileInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *duplicateFilesFinder) processFileInfo(path string, info fs.FileInfo) error {
	// TODO: This function has multiple responsibilities; refactor
	// Responsibilities:
	// - Update entries in file size to fileEntry map
	// - Update entries in checksums map
	// TODO: Logic is not readable; refactor with meaningful func names
	d.dupInfo.Total++
	currentFileSz := info.Size()
	if fc, found := d.fileSzToEntry[currentFileSz]; found {
		if !fc.IsChecksumCalculated {
			oldChecksum, err := d.checksumCalc.Calculate(fc.Path)
			if err != nil {
				return err
			}
			fc.IsChecksumCalculated = true
			d.checksums[oldChecksum] = true
		}
		fileChecksum, err := d.checksumCalc.Calculate(path)
		if err != nil {
			return err
		}
		if _, found := d.checksums[fileChecksum]; found {
			d.dupInfo.DuplicateFilesCount++
			d.dupInfo.DuplicateSize += currentFileSz
			return nil
		}
		d.checksums[fileChecksum] = true
	} else {
		d.fileSzToEntry[currentFileSz] = fileEntry{Path: path}
	}
	d.dupInfo.UniqueFilesCount++
	return nil
}
