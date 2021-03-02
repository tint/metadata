metadata
========

![](https://img.shields.io/github/downloads/tint/metadata/latest/total?style=flat-square)
![](https://img.shields.io/github/license/tint/metadata?style=flat-square)
![](https://img.shields.io/github/directory-file-count/tint/metadata?style=flat-square)
![](https://img.shields.io/github/v/release/tint/metadata?style=flat-square)

A metadata management

Usage
-----

```bash
$ go get github.com/tint/metadata
```

Note: Always vendor your dependencies or fix on a specific version tag.

```go
import (
    "github.com/tint/metadata"
    "context"
)

md := metadata.New(map[string]{
    "Foo": "bar",
    "bar": "baz",
})

md.Get("foo") // bar
md.Get("Foo") // bar
md.Set("baz", "baaa")
md.Delete("foo", "bar")
md.Copy() // new metadata

ctx := context.Background()
// md = metadata.FromContext(ctx)
metadata.NewContext(ctx, md)
metadata.Get(ctx, "foo")
metadata.Set("baz", "zab")
metadata.Delete("baz", "zab")
```

License
-------

See the [LICENSE](LICENSE) file for details.
