# go-text-export
go-text-export is text export and convert html library for Golang applications.

# installation

If you have Go installed, you can just do:

```shell
go get github.com/tkc/go-text-export
```

# usage

## Export

Text
```
Contrary to popular belief, Lorem Ipsum is not simply random text.
It has roots in a piece of classical Latin literature from 45 BC,
```

Code
```go
go_text_export.Create(Str).Export(100).SetMoreText(" more..").RemoveWhiteSpace().String()
```

Result
```
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of  more..
```

## SetMoreUrl

Text
```
Contrary to popular belief, Lorem Ipsum is not simply random text.
It has roots in a piece of classical Latin literature from 45 BC,
```

Code
```go
go_text_export.Create(str).Export(30).SetMoreText("more..").SetMoreUrl("https://example.com/1")
```

Result
```
Contrary to popular belie <a href='https://example.com/1'>more..</>
```

## ConvertHtmlLink


Text
```
link https://example.com
```

Code
```go
go_text_export.Create(Str).ConvertHtmlLink().String()
```

Result
```
link  <a href=https://example.com>https://example.com</a>
```

## More
```go
go_text_export.Create(Str).ConvertHtmlBrTag().String()
go_text_export.Create(Str).RemoveWhiteSpace().String()
go_text_export.Create(Str).Prepare(100).String()

```
