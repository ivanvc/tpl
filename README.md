# tpl

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
