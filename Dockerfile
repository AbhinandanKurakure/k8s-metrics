FROM alpine

WORKDIR /bin
RUN mkdir /bin/api
COPY ListPodsAPI /bin/api
EXPOSE 9090

CMD ["./api/ListPodsAPI"]