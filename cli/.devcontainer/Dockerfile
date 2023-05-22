FROM golang:1.20

WORKDIR /app

COPY . .

ENV GO111MODULE=off

# Install required Go tools
RUN go get -u github.com/cweill/gotests/...
RUN go get -u github.com/fatih/gomodifytags
RUN go get -u github.com/josharian/impl
RUN go get -u github.com/haya14busa/goplay/cmd/goplay
RUN go get -u golang.org/x/tools/cmd/goimports
# RUN go get -u golang.org/x/tools/gopls
RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u honnef.co/go/tools/cmd/staticcheck
RUN go get -u github.com/godoctor/godoctor