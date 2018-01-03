FROM golang:alpine

WORKDIR /app

COPY . .

CMD [ "./hello-gitlab-go" ]

EXPOSE 8000
