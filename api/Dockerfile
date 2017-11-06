FROM golang

# Copy over folder into container go path
COPY . /go/src/github.com/jcherianucla/UCLA-CS-130/api
WORKDIR /go/src/github.com/jcherianucla/UCLA-CS-130/api

# Fetch dependencies
RUN go get ./
# Build
RUN go build

# Run
CMD ["./api"]

EXPOSE 8080
