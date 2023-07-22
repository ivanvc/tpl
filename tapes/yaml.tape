Output images/yaml.gif
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

Type 'cat <<EOF >main.tf.tpl'
Enter
Type 'module "app" {'
Enter
Type '  source = "../modules/app'
Enter
Type '  name   = "{{ .name }}"'
Enter
Type '}'
Enter
Type 'EOF'
Enter

Show
Sleep 2s

Type 'tpl -env "name: test" main.tf.tpl'
Sleep 1s
Enter

Sleep 10s

Hide
Type "rm bin/tpl main.tf.tpl"
Enter
