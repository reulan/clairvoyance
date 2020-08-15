FROM golang:alpine as build
WORKDIR /truesight
COPY . .
RUN go build -v -o TruesightAura

FROM alpine as aura
WORKDIR /truesight

COPY --from=build /truesight/TruesightAura .
CMD ["./TrueSightAura"]
