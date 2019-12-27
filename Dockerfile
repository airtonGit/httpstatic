from scratch

WORKDIR /app

ADD httpstatic /app/httpstatic

ENTRYPOINT [ "./httpstatic" ]