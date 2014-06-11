# wwdcdl

Small program to download WWDC videos and slides (written in Go)

```
Usage:

  ./wwdcdl [options] event

Events:

  tt2013
  wwdc2012
  wwdc2013
  wwdc2014

Options:

  -hd=false: Prefer videos in HD quality
  -list=false: Only list sessions
  -n=false: Dry run (don't download anything)
  -output="/Users/marcel/Documents/WWDC": Location to store output
  -password=false: Ask for AppleID password (requires CasperJS)
  -slides=true: Download slides/PDFs
  -team="": Apple Developer Team ID (requires CasperJS)
  -username="": AppleID username (requires CasperJS)
  -videos=true: Download videos

Notes:

CasperJS (http://casperjs.org/) is required for authentication.
```

## Installation

- Get it from the [Releases page](https://github.com/sirlantis/wwdcdl/releases) (OS X, Linux)
- **or** run `go get github.com/sirlantis/wwdcdl` if you have `go` installed.

### Authentication with CasperJS

Some events (like WWDC12) require you to authenticate with your Apple ID.
For this `wwdcdl` requires the headless browser [CasperJS](http://casperjs.org/).
Please install it (for example with Homebrew) if you want to download these videos as well.

## License

This project is licensed under the GNU General Public License v2.

Copyright (c) 2014 Marcel Jackwerth <marceljackwerth@gmail.com>
