# Golang base image
FROM golang:bullseye AS builder
#
WORKDIR /app
#
ENV CGO_ENABLED=1
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache
#
COPY . ./
# Install dependencies
RUN --mount=type=cache,target=/gomod-cache \
  go mod tidy
# Compile and save file in /app/bin/main (opcion -o)
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
  go build -o ./bin/main main.go
#
# Use a slim image
#
FROM debian:bullseye-20240211-slim
#
WORKDIR /usr/local/share/ca-certificates/
#
WORKDIR /app/bin
# Copy the binary from builder
COPY --from=builder /app/bin/main ./
# Add security layer creating user
RUN addgroup --system --gid 1001 app
RUN adduser --system --uid 1001 app
# Change owner and permissions
RUN chmod a+x main
RUN chown -R app:app ./main
# Chnage to user app
USER app
# Expose port 3000
EXPOSE 3000
# Run the binary
CMD ["./main"]
