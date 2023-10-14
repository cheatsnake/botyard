FROM node:18-alpine as client
WORKDIR /app
COPY ["web/package.json", "web/package-lock.json*", "./"]
RUN npm ci --legacy-peer-deps
COPY ./web .
RUN npm run build

FROM golang:1.21 as server
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
COPY --from=client /app/dist ./web/dist
CMD [ "./main" ]
