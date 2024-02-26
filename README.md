# is [![GoDoc](https://godoc.org/zuern.dev/is?status.png)](http://godoc.org/zuern.dev/is) [![Go Report Card](https://goreportcard.com/badge/zuern.dev/is)](https://goreportcard.com/report/zuern.dev/is)

Professional lightweight testing mini-framework for Go.

> This is a fork of https://github.com/matryer/is, which adds support for
> supplying a message failed tests as an arg or as a comment.

### Features

- Easy to write and read
- [Beautifully simple API](https://pkg.go.dev/zuern.dev/is) with everything you
  need: `is.Equal`, `is.True`, `is.NoErr`, and `is.Fail`

Failures are very easy to read:

![Examples of failures](https://zuern.dev/is/raw/master/misc/delicious-failures.png)

### Usage

The following code shows a range of useful ways you can use the helper methods:

```go
func Test(t *testing.T) {
	is := is.New(t)
	signedin, err := isSignedIn(ctx)
	is.NoErr(err, "isSignedIn error")
	is.Equal(signedin, true, "must be signed in")
	body := readBody(r)
	is.True(strings.Contains(body, "Hi there")) // Should be in body
}
```

## Color

To turn off the colors, run `go test` with the `-nocolor` flag, or with the env
var [`NO_COLOR` (with any value)](https://no-color.org).

```
go test -nocolor
```

```
NO_COLOR=1 go test
```
