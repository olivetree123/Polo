FROM alpine

WORKDIR /app

COPY polo /app
RUN mkdir /etc/polo && chmod +x /app/polo
COPY config.toml /etc/polo/
CMD /app/polo