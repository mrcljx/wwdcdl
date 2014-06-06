# wwdcdl

Small program to download WWDC videos (written in Go)

```
Usage:

  wwdcdl [options] event

Events:

  wwdc2013
  wwdc2014

Options:

  -hd=false: Prefer videos in HD quality
  -n=false: Dry run (don't download anything)
  -output="/Users/marcel/Documents/WWDC": Location to store output
  -slides=true: Download slides/PDFs
  -videos=true: Download videos
```

## Installation

```
go get github.com/sirlantis/wwdcdl
```

or visit the [Releases page](https://github.com/sirlantis/wwdcdl/releases) and download it depending on your OS.

##

This project is licensed under the GNU General Public License v2.

Copyright (c) 2014 Marcel Jackwerth <marceljackwerth@gmail.com>
