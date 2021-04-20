FROM golang:latest


COPY . ./IoT-Platform
# update the apt
RUN apt update -y && apt upgrade -y
# install the dependencies
RUN apt-get install sqlite3 -y
RUN apt-get install nodejs -y && apt-get install npm -y
RUN npm install -g typescript
# compile the typescript files
RUN tsc ./IoT-Platform/src/frontend/public/script/*.ts
RUN cat ./IoT-Platform/src/sql/initDatabase.sql | sqlite3 ./IoT-Platform/src/sql/iotcameradata.db

# CLEAR SOME TRASH
RUN rm -rf client-for-raspberry-pi && rm -rf docs && rm -rf README.md
RUN rm -rf ./IoT-Platform/src/frontend/public/script/*.ts
RUN apt-get remove nodejs -y && apt-get remove npm -y

WORKDIR /go/IoT-Platform

CMD ["bash","init.sh"]