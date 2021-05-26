FROM golang:alpine3.13

RUN mkdir -p /home/app/
WORKDIR /home/app

COPY . /home/app
RUN cd /home/app/server && go build

EXPOSE 9000

CMD [ "/home/app/server/server" ]