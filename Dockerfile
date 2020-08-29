FROM centos:7

COPY Go-RestfulAPI /app/
COPY conf/config.json /app/conf/

RUN chmod 777 /app/Go-RestfulAPI

WORKDIR /app

EXPOSE 80

ENTRYPOINT ["./Go-RestfulAPI"]