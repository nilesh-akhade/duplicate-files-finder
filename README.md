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

Following SOLID principles.

Alternate implementation of the [ThoughtWorks Clean Code Workshop Assignment](https://github.com/nilesh-akhade/clean-code-workshop)
