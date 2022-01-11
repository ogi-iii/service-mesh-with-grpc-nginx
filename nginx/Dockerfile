FROM nginx:1.17.5

ADD ./nginx/default.conf /etc/nginx/conf.d/default.conf
ADD ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ssl/ /etc/nginx/ssl/

VOLUME ["/var/cache/nginx"]

EXPOSE 50055

CMD ["nginx"]
