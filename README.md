# wwdcdl

Small program to download WWDC videos and slides (written in Go)

```
Usage:

  ./wwdcdl [options] event

Events:

  tt2013
  wwdc2013
  wwdc2014

Options:

  -hd=false: Prefer videos in HD quality
  -list=false: Only list sessions
  -n=false: Dry run (don't download anything)
  -output="~/Documents/WWDC": Location to store output
  -slides=true: Download slides/PDFs
  -videos=true: Download videos
```

## Installation

- Get it from the [Releases page](https://github.com/sirlantis/wwdcdl/releases) (OS X, Linux)
- **or** run `go get github.com/sirlantis/wwdcdl` if you have `go` installed.

## License

This project is licensed under the GNU General Public License v2.

Copyright (c) 2014 Marcel Jackwerth <marceljackwerth@gmail.com>
