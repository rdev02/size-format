This is a simple go size formatting library, that would allow you to format `int64` to a `string` like `"1.2GB"` or inversely, parse a string like `"5GB"` into an `"int64"`.

## Usage

add a dependency in your `go.mod`
```
require (
	github.com/rdev02/size-format v0.1.0
)
```

the you can import
```
import sizeFormat "github.com/rdev02/size-format" 
```
and carry out conversions:
```
// int64 -> string
fmt.Println(sizeFormat.ToString(5 * sizeFormat.GB))

// string -> int64
sizeBytes, err := sizeFormat.ToNum("5.5GB")
```

supported sizes are B, KB, MB, GB, TB, PB
