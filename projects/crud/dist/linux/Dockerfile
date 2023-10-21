FROM nginx:alpine
RUN apk update
RUN apk add --no-cache --update libc6-compat
RUN apk add --no-cache tzdata
RUN apk add --no-cache supervisor
RUN apk add fontconfig freetype ttf-dejavu ttf-droid ttf-freefont ttf-liberation
RUN apk add openssl && rm -rf /var/cache/apk/*

RUN cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime
RUN echo "America/Sao_Paulo" > /etc/timezone
ENV TZ=America/Sao_Paulo

WORKDIR /home
ADD . /home
COPY supervisord.conf /etc/supervisord.conf
RUN mv /home/nginx.conf /etc/nginx/nginx.conf
RUN chmod +x /home/app

EXPOSE 80
CMD /usr/bin/supervisord -n
