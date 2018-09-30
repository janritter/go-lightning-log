# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

FROM scratch
ADD bin/go-lightning-log_linux /
CMD ["/go-lightning-log_linux"]