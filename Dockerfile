FROM golang:alpine AS builder
LABEL stage=builder
RUN apk add --no-cache gcc libc-dev
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a

FROM alpine AS final
WORKDIR /
COPY --from=builder /workspace/goHelm .
EXPOSE 8080
EXPOSE 9110
CMD [ "./goHelm" ]
 