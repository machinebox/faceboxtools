# facebox Tools

These are some little tools for [facebox](https://machinebox.io/docs/facebox), that can help you to teach or check based on images on your filesystem. Visit https://machinebox.io to get a key and run facebox on your machine.

## faceboxteach

`faceboxteach` is a tool that walks a directory structure teaching each
image file.

To install it:
```
$ go get github.com/machinebox/faceboxtools/faceboxteach
```

The folder name should be the `name` of the person, and the filenames will become
the `id` in facebox.

```
├── George_Harrison
│   ├── george1.jpg
│   ├── george2.jpg
│   └── george3.jpg
├── John_Lennon
│   ├── john1.jpg
│   ├── john2.jpg
│   └── john3.jpg
├── Paul_McCartney
│   ├── paul1.jpg
│   ├── paul2.jpg
│   ├── paul3.jpg
│   ├── paul4.jpg
│   └── paul5.jpg
└── Ringo_Starr
    ├── ringo1.jpg
    ├── ringo2.jpg
    ├── ringo3.jpg
    └── ringo4.jpg
```

* Underscores are translated into spaces for a better experience

To teach each item in the above structure, you would run:

```
$ faceboxteach -dir=/path/to/directory -images=.jpg -facebox=http://localhost:8080
```

* `dir` - (string) Directory to process
* `images` - (string) Pattern for files to teach (see filepath.Glob)
* `facebox` - (string) Address of running facebox


## faceboxtag

`faceboxtag` is a tool that walks a directory structure and outputs the people that recognise

To install it:
```
$ go get github.com/machinebox/faceboxtools/faceboxtag
```

```
├── photos
│   ├── whoisthis.jpg
│   ├── unknown.jpg
│   └── ????.jpg
```

To find out who is on the photos, will be on the output

```
$ faceboxteach -dir=./photos -images=.jpg -facebox=http://localhost:8080
```

* `dir` - (string) Directory to process
* `images` - (string) Pattern for files to teach (see filepath.Glob)
* `facebox` - (string) Address of running facebox
