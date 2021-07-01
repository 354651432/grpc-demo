
# Table of Contents

1.  [install grpc](#org442c80c)
2.  [install grpc go plugins](#orge8b208a)
3.  [start server](#org1f531f9)
4.  [start client](#org6a57c08)


<a id="org442c80c"></a>

# install grpc

    yay -S --noconfirm community/grpc


<a id="orge8b208a"></a>

# install grpc go plugins

    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    $ export PATH="$PATH:$(go env GOPATH)/bin"


<a id="org1f531f9"></a>

# start server

    make server


<a id="org6a57c08"></a>

# start client

    make client

