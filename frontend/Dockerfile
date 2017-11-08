FROM node:7.8.0

ENV NPM_CONFIG_LOGLEVEL warn

WORKDIR /grade-portal

COPY package.json npm-shrinkwrap.json /grade-portal/

RUN npm install

EXPOSE 5000

CMD ["npm", "start"]
