# Build Stage
FROM fzxiehui/net_to_uart:1.13 AS build-stage

LABEL app="build-net_to_uart"
LABEL REPO="https://github.com/fzxiehui/net_to_uart"

ENV PROJPATH=/go/src/github.com/fzxiehui/net_to_uart

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/fzxiehui/net_to_uart
WORKDIR /go/src/github.com/fzxiehui/net_to_uart

RUN make build-alpine

# Final Stage
FROM fzxiehui/net_to_uart

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/fzxiehui/net_to_uart"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/net_to_uart/bin

WORKDIR /opt/net_to_uart/bin

COPY --from=build-stage /go/src/github.com/fzxiehui/net_to_uart/bin/net_to_uart /opt/net_to_uart/bin/
RUN chmod +x /opt/net_to_uart/bin/net_to_uart

# Create appuser
RUN adduser -D -g '' net_to_uart
USER net_to_uart

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/net_to_uart/bin/net_to_uart"]
