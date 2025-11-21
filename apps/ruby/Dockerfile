FROM docker.io/ruby:3.4-alpine

COPY Gemfile /app/
WORKDIR /app
RUN bundle install

COPY app.rb config.ru /app/

EXPOSE 3000

CMD ["rackup", "--host", "0.0.0.0", "--port", "3000"]
