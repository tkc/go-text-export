# go-text-export
go-text-export is text export and convert html library for Golang applications.

# installation

If you have Go installed, you can just do:

```shell
go get github.com/tkc/go-text-export
```

# usage

## Export

```
Contrary to popular belief, Lorem Ipsum is not simply random text.
It has roots in a piece of classical Latin literature from 45 BC,

==>Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of  more..
```


Code
```go
go_text_export.Create(Str).SetMoreText(" more..").Export(100).RemoveWhiteSpace().String()
```

## SetMoreUrl

```
Contrary to popular belief, Lorem Ipsum is not simply random text.
It has roots in a piece of classical Latin literature from 45 BC,

==>Contrary to popular belie <a href='https://example.com/1'>more..</a>
```

Code
```go
go_text_export.Create(str).SetMoreText("more..").Export(30).SetMoreUrl("https://example.com/1")
```

## ConvertHtmlLink

```
link https://example.com

==>link  <a href=https://example.com>https://example.com</a>
```

Code
```go
go_text_export.Create(Str).ConvertHtmlLink().String()
```

## More Sample
```go
go_text_export.Create(Str).ConvertHtmlBrTag().String()
go_text_export.Create(Str).RemoveWhiteSpace().String()
go_text_export.Create(Str).Prepare(100).String()

```

## Method List
```go
Export
Prepare
ConvertHtmlLink
ConvertHtmlBrTag
RemoveWhiteSpace
RemoveNewLine
SetMoreText
SetMoreUrl
String
```
