FROM ubuntu:20.04
RUN apt update
RUN apt install -y tzdata
RUN apt install -y build-essential libsnappy-dev libbz2-dev liblz4-dev libzstd-dev g++-7  git cmake
RUN apt install -y curl zlib1g-dev
RUN mkdir /btpd
RUN git clone https://github.com/mambaru/btpd.git /btpd
RUN cd btpd && CXX=g++-7 make shared ARG=-j1
RUN mkdir -p /storage/btp
RUN mkdir -p /logs

COPY . /btpd
RUN chmod +x /btpd/run.sh


EXPOSE 37000 37001 37002 37003 37004 38001 38001/udp

ENTRYPOINT ["/btpd/run.sh"]

