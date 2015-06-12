# wwdcdl

Small program to download WWDC videos and slides (written in Go)

```
usage: wwdcdl [<flags>] <event>

Flags:
  --help          Show help (also see --help-long and --help-man).
  --email=sjobs@apple.com
                  Your Apple ID (requires CasperJS)
  --team=TEAM     Your Apple Developer Team ID (requires CasperJS)
  -p, --password  Ask for AppleID password (requires CasperJS)
  --hd            Prefer videos in HD quality
  -n, --dry       Dry run (don't download anything)
  --no-folders    Don't create a separate folder for each event
  --no-videos     Don't download videos
  --no-slides     Don't download slides/PDFs
  -l, --list      Only show list of sessions
  -o, --output="/Users/marcel/Documents/Apple Events"
                  Location to store output
  --version       Show application version.

Args:
  <event>  The event to download videos/slides from

Events:
  tt2013 (Tech-Talks 2013)
  wwdc2012 (WWDC 2012)
  wwdc2013 (WWDC 2013)
  wwdc2014 (WWDC 2014)
  wwdc2015 (WWDC 2015)

Notes:
  CasperJS (http://casperjs.org/) is required for authentication.
```

## Installation

### Recommended (OS X)

```
brew tap sirlantis/wwdcdl
brew install wwdcdl
```

### Alternative Methods

- Get it from the [Releases page](https://github.com/sirlantis/wwdcdl/releases) (OS X, Linux)
- **or** run `go get github.com/sirlantis/wwdcdl` if you have `go` installed.

### Authentication with CasperJS

Some events (like WWDC12) require you to authenticate with your Apple ID.
For this `wwdcdl` requires the headless browser [CasperJS](http://casperjs.org/).
Please install it (for example with Homebrew) if you want to download these videos as well.

## License

This project is licensed under the GNU General Public License v2.

Copyright (c) 2014 Marcel Jackwerth <marceljackwerth@gmail.com>
