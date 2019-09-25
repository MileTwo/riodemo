FROM alpine:3.9
# set color from build-ars on docker build command

ARG COLOR
ENV COLOR ${COLOR:-yellow}

EXPOSE 80
COPY app . 
CMD ["./app"]