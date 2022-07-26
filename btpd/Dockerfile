FROM ubuntu:20.04
RUN apt-get update
RUN apt-get install -y tzdata
RUN apt-get install -y build-essential libsnappy-dev libbz2-dev liblz4-dev libzstd-dev g++-7  git cmake
RUN apt-get install -y curl zlib1g-dev wget
#RUN wget https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.tar.gz
#RUN tar -xf boost_1_76_0.tar.gz
#RUN cd boost_1_76_0 && ./bootstrap.sh
#RUN cd boost_1_76_0 && ./b2 link=shared runtime-link=shared threading=multi install
RUN mkdir /btpd
RUN git clone https://github.com/mambaru/btpd.git /btpd
RUN cd btpd && CXX=g++-7 make shared ARG=-j1
RUN mkdir -p /storage/btp
RUN mkdir -p /logs

COPY config /btpd
RUN chmod +x /btpd/run.sh

EXPOSE 37000 37001 37002 37003 37004 38001 38001/udp

ENTRYPOINT ["/btpd/run.sh"]

