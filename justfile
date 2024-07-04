# list available recipes
default:
  just --list

# build on first install and when there are new completers or actions
[no-cd]
build:
  go generate ./...
  go install

# build completer
[no-cd]
build-completer:
  go install -ldflags '-s -w'
