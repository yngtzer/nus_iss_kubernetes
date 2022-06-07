ARG VERSION=18.2.0

FROM node:${VERSION} AS build

LABEL developer="yt"

WORKDIR /app

# Make sure curl is installed
RUN apt update && apt install curl

ADD package.json .
ADD package-lock.json .
ADD public public
ADD views views
ADD main.js .

RUN npm ci

ENV APP_PORT=3000
EXPOSE ${APP_PORT}

# Does not have to be using curl, could have use ping as well
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD curl -s http://localhost:${APP_PORT}/health || exit 1

# Run the application
ENTRYPOINT [ "node", "main.js" ]