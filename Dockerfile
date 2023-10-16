# Use the official Golang image to create a build artifact.
FROM golang:1.19-alpine

WORKDIR /app
COPY . .
RUN go build -o resumedawamr

EXPOSE 4000
CMD ["./resumedawamr"]
