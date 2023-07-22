Output images/json.gif
Set Width 1400
Set Height 600
Set Theme "Dracula"
Set FontFamily "Source Code Pro"
Set LetterSpacing 1

Hide
Type "export PS1='$ ' PATH=$PWD/bin:$PATH; mkdir bin"
Enter
Type "go build -o bin/tpl ."
Enter
Type "clear"
Enter

Type 'cat <<EOF | tpl -stdin -env "$(curl https://api.github.com/users/ivanvc/events 2> /dev/null)"'
Enter
Type '{{- define "type" -}}'
Enter
Type '{{ . | replace "Event" "" | kebabcase | replace "-" " " | lower }}'
Enter
Type '{{- end -}}'
Enter
Type '{{- range slice . 0 3 }}'
Enter
Type '* {{ .created_at | toDate "2006-01-02T15:04:05Z07:00" | ago }} ago {{ include "type" .type }}:'
Enter
Type '{{- with .repo }}'
Enter
Type '  {{ .name }}: {{ .url }}'
Enter
Type '{{- end }}'
Enter
Type '{{- end }}'
Enter
Type 'EOF'

Show
Sleep 5s
Enter

Sleep 10s

Hide
Type "rm bin/tpl"
Enter
