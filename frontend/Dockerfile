FROM node:13.12.0-alpine3.11 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN npm ci
RUN npm run export

FROM nginx:1.19.1-alpine
COPY --from=builder /build/__sapper__/export/ /usr/share/nginx/html/
COPY nginx.conf /conf/nginx.conf
CMD nginx -c /conf/nginx.conf
