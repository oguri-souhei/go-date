# go-date

`go-date` is a simple Go package for handling and manipulating calendar dates. It encapsulates the `time.Time` object internally and provides convenient methods for date operations without time information.

## Features

- Create a new date with year, month, and day
- Retrieve year, month, and day components
- Compare dates
- Add years, months, and days to a date
- Format and parse dates

## Installation

To install the package, use `go get`:

```sh
go get github.com/oguri-souhei/go-date
```

## Usage

### Creating a Date

```go
package main

import (
"fmt"
"github.com/oguri-souhei/go-date"
"time"
)

func main() {
  d := date.New(2023, time.July, 13)
  fmt.Println(d) // Output: 2023-07-13
}
```

### Retrieving Date Components

```go
d := date.New(2023, time.July, 13)
fmt.Println(d.Year()) // Output: 2023
fmt.Println(d.Month()) // Output: July
fmt.Println(d.Day()) // Output: 13
```

### Comparing Dates

```go
d1 := date.New(2023, time.July, 13)
d2 := date.New(2023, time.August, 13)

fmt.Println(d1.Before(d2)) // Output: true
fmt.Println(d1.After(d2)) // Output: false
fmt.Println(d1.Equal(d2)) // Output: false
```

### Adding to Dates

```go
d := date.New(2023, time.July, 13)
d2 := d.AddYear(1)
d3 := d.AddMonth(1)
d4 := d.AddDay(1)

fmt.Println(d2) // Output: 2024-07-13
fmt.Println(d3) // Output: 2023-08-13
fmt.Println(d4) // Output: 2023-07-14
```

### Formatting Dates

```go
d := date.New(2023, time.July, 13)
fmt.Println(d.Format("02-01-2006")) // Output: 13-07-2023
fmt.Println(d.String()) // Output: 2023-07-13
```

### Parsing Dates

```go
d, err := date.Parse("13-07-2023")
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(d) // Output: 2023-07-13
}
```

## Documentation

For detailed documentation, visit the [GoDoc page](https://pkg.go.dev/github.com/oguri-souhei/go-date).

## License

This project is licensed under the [MIT License](http://opensource.org/license/MIT). See the LICENSE file for details.
