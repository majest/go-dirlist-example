FROM alpine:3.3
RUN mkdir -p /opt/app
ADD /go-dirlisting-example /opt/app/
EXPOSE 8086
VOLUME /data
CMD ["/opt/app/go-dirlisting-example"]
