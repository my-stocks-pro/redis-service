FROM acoshift/go-alpine

RUN mkdir app_log

VOLUME /app_log

ADD redis-service /

CMD ./redis-service
