
FROM node:14.15.0-alpine

WORKDIR /vue/src
COPY ./src /vue/src

RUN apk update && npm install -g @vue/cli
COPY package.json .
RUN npm install
FROM node:14.15.0-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

CMD ["npm", "run", "serve"]