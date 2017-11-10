FROM node:7.8.0

ENV NPM_CONFIG_LOGLEVEL warn

COPY package.json /grade-portal/package.json

WORKDIR /grade-portal

RUN npm install

EXPOSE 5000

CMD ["npm", "start"]
