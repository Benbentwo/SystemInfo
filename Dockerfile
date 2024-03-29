FROM golang:1.16.5 AS builder

# Build arguments
ARG binary_name=main
    # See ./sample-data/go-os-arch.csv for a table of OS & Architecture for your base image
ARG target_os=linux
ARG target_arch=amd64

# Build the server Binary
WORKDIR /app
#WORKDIR /go/src/${GIT_SERVER}/${GIT_ORG}/${GIT_REPO}

COPY go.mod .
COPY go.sum .

RUN go mod download

# Seems duplicative, and ideally not needed
COPY . .

RUN rm -rf /app/build
RUN CGO_ENABLED=0 GOOS=${target_os} GOARCH=${target_arch} go build -a -o /app/build/${binary_name} main.go

RUN ls /app

#-----------------------------------------------------------------------------------------------------------------------

FROM centos:7

LABEL author="Benjamin Smith"
RUN echo 'export PS1="[\u@docker] \W # "' >> /root/.bash_profile

COPY --from=builder ./app/build/main /usr/bin/main
RUN ["chmod", "-R", "+x", "/usr/bin/main"]

#ENTRYPOINT ["tail", "-f", "/dev/null"]
ENTRYPOINT ["/bin/bash", "-c"]

