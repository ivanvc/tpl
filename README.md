# tpl

<svg xmlns="http://www.w3.org/2000/svg" version="1.1" align="right" width="81px" height="43px" style="shape-rendering:geometricPrecision; text-rendering:geometricPrecision; image-rendering:optimizeQuality; fill-rule:evenodd; clip-rule:evenodd" xmlns:xlink="http://www.w3.org/1999/xlink">
<g>
<path style="opacity:1" fill="#9f9cfa" d="M 9.5,2.5 C 11.5,2.5 13.5,2.5 15.5,2.5C 15.5,5.16667 15.5,7.83333 15.5,10.5C 17.8333,10.5 20.1667,10.5 22.5,10.5C 22.5,12.8333 22.5,15.1667 22.5,17.5C 20.1667,17.5 17.8333,17.5 15.5,17.5C 15.5,20.1667 15.5,22.8333 15.5,25.5C 17.8333,25.5 20.1667,25.5 22.5,25.5C 22.5,27.5 22.5,29.5 22.5,31.5C 18.1667,31.5 13.8333,31.5 9.5,31.5C 9.5,26.8333 9.5,22.1667 9.5,17.5C 7.16667,17.5 4.83333,17.5 2.5,17.5C 2.5,15.1667 2.5,12.8333 2.5,10.5C 4.83333,10.5 7.16667,10.5 9.5,10.5C 9.5,7.83333 9.5,5.16667 9.5,2.5 Z"/>
<path style="opacity:1" fill="#9f9cfa" d="M 58.5,2.5 C 60.5,2.5 62.5,2.5 64.5,2.5C 64.5,10.1667 64.5,17.8333 64.5,25.5C 66.8333,25.5 69.1667,25.5 71.5,25.5C 71.5,27.5 71.5,29.5 71.5,31.5C 67.1667,31.5 62.8333,31.5 58.5,31.5C 58.5,21.8333 58.5,12.1667 58.5,2.5 Z"/>
<path style="opacity:1" fill="#9f9cfa" d="M 30.5,10.5 C 37.1667,10.5 43.8333,10.5 50.5,10.5C 50.5,17.5 50.5,24.5 50.5,31.5C 45.8333,31.5 41.1667,31.5 36.5,31.5C 36.5,34.1667 36.5,36.8333 36.5,39.5C 34.5,39.5 32.5,39.5 30.5,39.5C 30.5,29.8333 30.5,20.1667 30.5,10.5 Z M 36.5,17.5 C 39.1667,17.5 41.8333,17.5 44.5,17.5C 44.5,20.1667 44.5,22.8333 44.5,25.5C 41.8333,25.5 39.1667,25.5 36.5,25.5C 36.5,22.8333 36.5,20.1667 36.5,17.5 Z"/>
</g>
</svg>

`tpl` is a simple Go templates wrapper. It enables using them in the command
line by using the `tpl` binary.

## Installation

1. Download the binary from the [releases page], or
2. Download with go
    ```bash
    $ go install github.com/ivanvc/tpl@latest
    ```

## Usage

```bash
$ tpl -h
Usage: tpl [options] [template input]:

Executes the template from the input applying the environment passed in options.

If you specify the input template via an option flag (-input), then it will
read it inline.
If you specify that the template comes from the stdin by setting the option flag
(-stdin), then it will read it inline from the stdin.
If you set it as the first argument, it assumes that it is a file.

For the environment (-env) it will expect it as inline data. However, if you
start it with @, it will assume it is a file.

The output is sent to stdout.

Options:
  -debug
        Set log level to debug.
  -env string
        The environment for the template (YAML, JSON or TOML).
  -input string
        The template input to process.
  -stdin
        Read template from stdin.
```
### YAML

Assuming you have a template `main.tf.tpl` with:

```hcl
module "app" {
    source = "../modules/app"
    name   = "{{ .name }}"
}
```

Executing with the following YAML environment:

```YAML
name: test
```

```bash
$
cat <<EOF | tpl -env 'name: test' -stdin
module "app" {
    source = "../modules/app"
    name   = "{{ .name }}"
}
EOF
```

Returns the rendered output:

```hcl
module "app" {
    source = "../modules/app"
    name   = "test"
}
```

### <a name="JSON"></a>JSON

It works with JSON environments:

```bash
$ cat <<EOF | tpl -env "$(curl https://api.github.com/users/ivanvc/events 2> /dev/null)"
{{- define "type" -}}
{{ . | replace "Event" "" | lower }}
{{- end -}}
{{- range . }}
{{ .created_at | toDate "2006-01-02T15:04:05Z07:00" | ago }} ago {{ include "type" .type }}:
{{- with .repo }}
  {{ .name }}: {{ .url }}
{{- end }}
{{- end }}
EOF
```

Generates the output:

```
18h51m49s ago push:
  ivanvc/tpl: https://api.github.com/repos/ivanvc/tpl
18h58m44s ago push:
  ivanvc/tpl: https://api.github.com/repos/ivanvc/tpl
35h53m29s ago push:
  ivanvc/tpl: https://api.github.com/repos/ivanvc/tpl
35h59m5s ago create:
  ivanvc/tpl: https://api.github.com/repos/ivanvc/tpl
36h0m1s ago create:
  ivanvc/tpl: https://api.github.com/repos/ivanvc/tpl
95h44m45s ago push:
  ivanvc/aports: https://api.github.com/repos/ivanvc/aports
95h52m22s ago push:
  ivanvc/aports: https://api.github.com/repos/ivanvc/aports
...
```

### TOML

It also supports TOML. Generating index pages is easy, given the template:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Index</title>
  </head>
  <body>
    <h1>Index of {{ .dir }}</h1>
    <ul>
      {{- range .files }}
      <li><a href="{{ . }}">{{ . }}</a></li>
      {{- end }}
    </ul>
  </body>
</html>
```

Running it with:

```bash
$ tpl -env @<(echo -e "dir='$(basename $PWD)'\nfiles=['$(echo * | sed s/[[:space:]]/\',\'/g)']") index.html
```

Generates the following output:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Index</title>
  </head>
  <body>
    <h1>Index of tpl</h1>
    <ul>
      <li><a href="Dockerfile">Dockerfile</a></li>
      <li><a href="go.mod">go.mod</a></li>
      <li><a href="go.sum">go.sum</a></li>
      <li><a href="internal">internal</a></li>
      <li><a href="main.go">main.go</a></li>
      <li><a href="pkg">pkg</a></li>
      <li><a href="README.md">README.md</a></li>
    </ul>
  </body>
</html>
```

### Template functions

`tpl` includes Sprig functions, refer to their
[documentation](https://masterminds.github.io/sprig/).

It also defines the `include` function as defined by Helm. So you can include
templates defined before, see the [JSON][#JSON] example to see its usage.

## License

See [LICENSE](LICENSE) © [Ivan Valdes](https://github.com/ivanvc/)

[releases page]: https://github.com/ivanvc/tpl/releases

<!--
▄█▄ ▄▄▄ █
 █▄ █▄█ █▄
    ▀
-->
