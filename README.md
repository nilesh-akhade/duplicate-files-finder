# duplicate-files-finder
Finds duplicate files in a given directory using checksum


## Design

Get dir from cmd args
for file in files {
    // get fileinfo and get checksum
    // decide if file is duplicate or not
    // if yes store the count and file size
}

Total Files: 4
Duplicate Files: 1
Duplicate File Size: (duplicate count x file size) + .. for every file