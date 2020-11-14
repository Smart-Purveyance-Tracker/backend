FROM golang:1.15-alpine as build
ADD . /backend
WORKDIR backend
RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash git openssh make cmake && cd /backend && make

FROM scratch
COPY --from=build /backend/bin/backend .
EXPOSE 3000
CMD ["./backend"]
