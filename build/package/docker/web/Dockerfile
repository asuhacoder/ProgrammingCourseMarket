FROM node:17.3.0-alpine3.12


ENV PATH /web/node_modules/.bin:$PATH

WORKDIR /web
COPY ./web .
RUN yarn install

CMD [ "yarn", "start" ]