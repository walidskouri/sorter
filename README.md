# sorter
Simple Go application to sort file lines alphabetically

Usage :

```
sortline -h
```

Usage:
  sort [flags]

Flags:
  -f, --file-name string            The input file name (default "file.txt")
  -h, --help                        help for sort
  -s, --sorted-file-suffix string   The output file name suffix (default "sorted")
  -v, --version                     Print the version and exit


example : 

```
sortline -f file1.txt -s sorted
```

output :

file1-sorted.txt
