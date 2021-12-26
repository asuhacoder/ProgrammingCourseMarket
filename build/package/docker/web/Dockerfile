FROM node:17.3.0-alpine3.12


ENV PATH /web/node_modules/.bin:$PATH

RUN file="$(ls -1)" && echo $file

WORKDIR /web
COPY ./web .
RUN yarn install
RUN file="$(ls -1)" && echo $file

CMD [ "yarn", "start" ]