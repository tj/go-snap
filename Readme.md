
# go-snap

 Tiny snappy command-line utility operating over stdio to test compression rates. Supports encoding and decoding of snappy streams.

## Installation

```
$ go get github.com/visionmedia/go-snap
```

## Usage

```

  Usage: snap [options]

  Options:
    -e, --encode  encode stdin (the default)
    -d, --decode  decode stdin

  Help Options:
    -h, --help    Show this help message

```

## Examples

```
$ go-snap < some.txt | go-snap --decode
```

# License

 MIT