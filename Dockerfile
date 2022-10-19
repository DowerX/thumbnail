FROM golang:alpine AS build

RUN apk --no-cache add git ca-certificates

RUN mkdir build
WORKDIR /build
COPY ./* /build/
RUN CGO_ENABLED=0 \
    go build \
    -installsuffix "static" \
    -o thumbnail

FROM scratch AS final
COPY --from=build /build/thumbnail /thumbnail
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/thumbnail"]