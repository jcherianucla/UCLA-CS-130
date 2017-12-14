FROM golang:1.8.5-jessie
# Allow for cross platform compilation
RUN apt-get update && apt-get install -y gcc-multilib g++-multilib
# Check versions
RUN which gcc && which bash && which ls && which grep && which git
# Check container architecture
RUN uname -a
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
