FROM docker.io/ruby:3.4-alpine

COPY . /app
WORKDIR /app

RUN gem install rack rackup webrick

EXPOSE 3000

CMD ["rackup", "--port", "3000", "--host", "0.0.0.0"]
