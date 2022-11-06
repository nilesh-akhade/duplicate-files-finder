package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New("/tmp", true))
}

func TestFind_DirNotFound(t *testing.T) {
	dff := New("/unavailable-dir", true)
	_, err := dff.Find()
	assert.Error(t, err)
}

func TestFind(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			dir       string
			recursive bool
		}
		want struct {
			dfi *DuplicateFilesInfo
			err error
		}
	}{
		{
			name: "Sample dir with 2 dups",
			args: struct {
				dir       string
				recursive bool
			}{
				dir:       "./testdata/2dupes",
				recursive: false,
			},
			want: struct {
				dfi *DuplicateFilesInfo
				err error
			}{
				dfi: &DuplicateFilesInfo{
					Total:               4,
					UniqueFilesCount:    3,
					DuplicateFilesCount: 1,
					DuplicateSize:       11,
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		dff := New(tt.args.dir, tt.args.recursive)
		dfi, err := dff.Find()
		assert.Equal(t, tt.want.err, err)
		assert.Equal(t, tt.want.dfi, dfi)
	}
}
