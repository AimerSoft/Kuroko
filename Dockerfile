FROM golang:1.20 AS backend
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
WORKDIR /app
COPY . .
RUN go build -o app ./cmd
FROM node:latest AS frontend
ARG DOMAIN
ENV VITE_API_URL_ROOT=$DOMAIN
WORKDIR /app
COPY ./web/app .
RUN git submodule update --init --recursive
RUN yarn install && yarn build

FROM nginx:latest
WORKDIR /app
COPY --from=backend /app/app .
COPY --from=frontend /app/dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["sh", "-c", "./app & nginx -g 'daemon off;'"]
