FROM debian

WORKDIR /opt/

#RUN apt-get update && apt-get install -y curl ca-certificates && apt-get clean && apt-get autoremove

COPY styxnode /bin/
COPY dockerconfig* /opt/

COPY host_key /opt/
COPY *.pem /opt/
#COPY ca.pem /usr/local/share/ca-certificates/

#RUN update-ca-certificates

ENV configfile ""

ENTRYPOINT [ "/bin/bash", "-c", "styxnode --overwrite-hostconfig --config ${configfile}" ]
