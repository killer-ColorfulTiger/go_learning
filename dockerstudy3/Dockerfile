FROM ubuntu
RUN echo '这是一个基于ubuntu且应用为httpserver的镜像' && mkdir -p /dockerstudy/httpserver
COPY go/gohttpserver /dockerstudy/httpserver
ENTRYPOINT ["/dockerstudy/httpserver/gohttpserver"]
