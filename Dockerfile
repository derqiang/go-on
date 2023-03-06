FROM alpine

WORKDIR /usr/local/bin

COPY ./chf .

CMD ["chf"]