FROM node:7.8.0

ENV NPM_CONFIG_LOGLEVEL warn

RUN useradd --user-group --create-home --shell /bin/false appuser

ENV HOME=/home/appuser

COPY package.json npm-shrinkwrap.json $HOME/
RUN chown -R appuser:appuser $HOME/*

USER appuser

WORKDIR $HOME

RUN npm install

EXPOSE 5000

CMD ["npm", "start"]
