templateFile=$GOPATH/src/github.com/go-easygen/easygen/test/commandlineEnv
[ -s $templateFile.tmpl ] || templateFile=/usr/share/gocode/src/github.com/go-easygen/easygen/test/commandlineEnv
[ -s $templateFile.tmpl ] || templateFile=/usr/share/doc/easygen/examples/commandlineEnv
[ -s $templateFile.tmpl ] || {
  echo No template file found
  exit 1
}

easygen $templateFile mockserver_cli | gofmt > mockserver_opt.go
