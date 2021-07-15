
FROM node:14.15.0-alpine

WORKDIR /vue/src
COPY ./src /vue/src

RUN apk update && npm install -g @vue/cli
COPY package.json .
RUN npm install
