FROM golang:alpine as build-env
WORKDIR /app
ADD . /app
ENV GOOS=linux 
ENV GOARCH=amd64 
ENV CGO_ENABLED=0
RUN cd /app && go build -o builtapp

FROM scratch

WORKDIR /app
COPY --from=build-env /app/builtapp /httpstatic
#ADD httpstatic /app/httpstatic

ENTRYPOINT [ "./httpstatic" ]