FROM scratch
MAINTAINER Scott Miller <scott.miller171@gmail.com>

ADD go-redirect go-redirect

EXPOSE 80

ENTRYPOINT [ "/go-redirect" ]
