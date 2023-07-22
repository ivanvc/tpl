Output images/toml.gif
Set Width 1400
Set Height 640
Set Theme "Dracula"
Set FontFamily "Source Code Pro"
Set LetterSpacing 1

Hide
Type "export PS1='$ ' PATH=$PWD/bin:$PATH; mkdir bin"
Enter
Type "go build -o bin/tpl ."
Enter

Type 'cat <<EOF >index.html.tpl'
Enter
Type '<!DOCTYPE html>'
Enter
Type '<html lang="en">'
Enter
Type '  <head>'
Enter
Type '    <meta charset="utf-8">'
Enter
Type '    <title>Index</title>'
Enter
Type '  </head>'
Enter
Type '  <body>'
Enter
Type '    <h1>Index of {{ .dir }}</h1>'
Enter
Type '    <ul>'
Enter
Type '      {{- range .files }}'
Enter
Type '      <li><a href="{{ . }}">{{ . }}</a></li>'
Enter
Type '      {{- end }}'
Enter
Type '    </ul>'
Enter
Type '  </body>'
Enter
Type '</html>'
Enter
Type 'EOF'
Enter

Type "clear"
Enter

Show

Type 'cat index.html.tpl'
Sleep 1s
Enter
Sleep 5s

Type "tpl -env @<(echo -e "
Type '"'
Type "dir='$(basename $PWD)'\nfiles=['$(echo * | sed s/[[:space:]]/\',\'/g)']"
Type '") index.html.tpl | more'
Sleep 3s
Enter

Sleep 5s
Enter
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 1s
Type ' '
Sleep 10s

Hide
Type 'q'
Enter
Type "rm index.html.tpl bin/tpl"
Enter
