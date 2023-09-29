FROM golang:alpine3.17
WORKDIR /app
USER 0 

COPY .  .
RUN go build -o /UserApp
CMD [ "/UserApp" ]
EXPOSE 8443
