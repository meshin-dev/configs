GO_VERSION=1.23
export GOROOT="$(brew --prefix go@$GO_VERSION)/libexec"
export GOPATH="/Users/$USER/go$GO_VERSION"
export PATH="/Users/$USER/go$GO_VERSION/bin:$PATH"
export PATH="$(brew --prefix go@$GO_VERSION)/libexec/bin:$PATH"
