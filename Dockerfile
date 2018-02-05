FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN  go get github.com/PuerkitoBio/goquery &&  go get github.com/lib/pq && go build -o main . 
ENTRYPOINT ["/app/main"]

EXPOSE 8080
