FROM golang
WORKDIR /home/udda
COPY . /home/udda
RUN go build