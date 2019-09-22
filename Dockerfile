FROM scratch

WORKDIR /labac
COPY ./labac-gin /labac

ENV GIN_MODE=release

EXPOSE 8000
CMD ["./labac"]