FROM golang:1.15
ADD . ./spt
WORKDIR spt
RUN make
EXPOSE 3000
CMD ["./bin/backend"]
