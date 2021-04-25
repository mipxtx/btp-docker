FROM ubuntu:20.04

RUN apt update
RUN apt install -y tzdata
RUN apt install -y php-fpm git
RUN mkdir /btp-ui
RUN git clone https://github.com/mipxtx/btp-ui.git /btp-ui
COPY . .
COPY www.conf /etc/php/7.4/fpm/pool.d/www.conf
RUN chmod +x ./run.sh
EXPOSE 9000
ENTRYPOINT ["./run.sh"]

