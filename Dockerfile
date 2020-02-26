FROM scratch
EXPOSE 8080
ENTRYPOINT ["/jrfoo141"]
COPY ./build/linux /