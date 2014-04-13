jconfig
=======

[![Build Status](https://travis-ci.org/GroundSix/jconfig.svg?branch=master)](https://travis-ci.org/GroundSix/jconfig)

This is a Go package to parse a configuration file using JSON.

It was ogirinally developed by [Stathat](https://github.com/stathat/jconfig) but we
are maintaining a fork of it here due to it being a dependency to other projects
we maintain.

### Installation

```bash
$ go get github.com/stathat/jconfig
```

### Usage

```go
import (
    "github.com/groundsix/jconfig"
)
```

#### Example

Here is our example JSON file:

```json
{
    "name" : "Harry",
    "age"  : 22,
    "pets" : ["Dog", "Cat"]
}
```

And to make use of this data:

```go
package main

import (
    "fmt"
    "github.com/groundsix/jconfig"
)

func main() {
    config := jconfig.LoadConfig("path/to/your/config.json")

    name := config.GetString("name")
    age  := config.GetInt("age")
    pets := config.GetArray("pets")

    fmt.Println(pets[0]) // Dog
}
```

See `config.go` for full API and types that can be pulled from
your config file.

### Running Tests

```bash
$ make test
```

### About

Originally written by Patrick Crosby at [StatHat](http://www.stathat.com).
[(@stat_hat)](http://twitter.com/stat_hat)

Fork currently being maintained by [Ground Six](http://groundsix.com).
[(@groundsix)](http://twitter.com/groundsix)

### License

[MIT](https://github.com/GroundSix/jconfig/blob/master/LICENSE)
