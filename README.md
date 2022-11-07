# duplicate-files-finder

Finds duplicate files in a given directory using file size and sha1 checksum.

```console
    foo@bar:~$ go run github.com/nilesh-akhade/duplicate-files-finder@latest -r --dir "$HOME/Downloads"
    Finding duplicate files...
    --------------------------------------------------
    Directory                    : /home/foo/Downloads
    Recursive                    : true
    Total files                  : 140087
    Unique files                 : 67500
    Files which are duplicated   : 72587
    Space taken by the duplicates: 192.83MB
    --------------------------------------------------
```

## Design

This is an alternate implementation of the [ThoughtWorks Clean Code Workshop Assignment](https://github.com/nilesh-akhade/clean-code-workshop)

I tried to use SOLID, clean code principles. Feel free to open an issue, if you see any of the principle violated.

## Next?

To test extensiblity and robustness of this design, a dummy customer will keep on changing the requirements as follows. A good design will have less code change, least refactoring to existing code.

- We need names of the duplicate files
- Program fails when it detects, non readable or system files
- Our directory has 1 million files, this program gets out of memory
- Program takes long time we want pause and resume functionality
- Program takes long time can we use faster checksum algorithm instead of sha1
- The program is still slow; can you use better data structure?
- The program shows nothing, provide functionality to show progress.
- Support google drive in addition to the local filesystem
- We need similar files not exactly duplicates. Like file containing "Hello World" is 50% similar to file containing "World"
