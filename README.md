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
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC,
making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia,
looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature,
discovered the undoubtable source. Lorem Ipsum comes from sections
```

Code
```
normaRes := go_text_export.Create(normalStr).Export(100).SetMoreText(" more..").RemoveWhiteSpace().String()
```

Result
```
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of  more..
```

## ConvertHtmlLink


Text
```
link https://example.com
```

Code
```
go_text_export.Create(linkStr).ConvertHtmlLink().String()
```

Result
```
link  <a href=https://example.com>https://example.com</a>
```

## ConvertHtmlBrTag
## RemoveWhiteSpace
## Prepare
