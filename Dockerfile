FROM golang:1.15
ADD . .
RUN make
EXPOSE 3000
CMD ["./bin/backend"]
