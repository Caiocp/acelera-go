FROM golang:1.19.1

# RUN apt update && apt install -y --no-install-recommends \
#     git \
#     ca-certificates

WORKDIR /home/go/app

CMD [ "sh", "-c", "yarn && tail -f /dev/null" ]

# ENTRYPOINT [ "tail", "-f", "/dev/null" ]