FROM harbor.central.dev.didevops.com/platform/go:16

COPY demo4 /app/

WORKDIR /app

CMD ["/app/demo4"]