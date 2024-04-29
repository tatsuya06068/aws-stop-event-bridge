FROM amazon/aws-cli:latest

RUN apk -uv add --no-cache groff less python3 py3-pip jq bash curl && \
    pip3 install --no-cache --upgrade pip && \
    pip3 install --no-cache --upgrade awscli && \
    apk --purge -v del py-pip && \
    rm /var/cache/apk/*

WORKDIR /app

# docker build -t cloudformation-dev .
# docker run -it --rm -v /path/to/cloudformation/files:/app cloudformation-dev /bin/bash