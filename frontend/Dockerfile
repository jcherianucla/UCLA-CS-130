FROM node:7.8.0

ENV NPM_CONFIG_LOGLEVEL warn

COPY package.json npm-shrinkwrap.json /grade-portal/

WORKDIR /grade-portal

RUN npm install

EXPOSE 5000

RUN useradd --user-group --shell /bin/false admin
RUN chown -R admin:admin /grade-portal/

CMD ["npm", "start"]

ENV HOME=/home/appuser
