FROM ysicing/debian

COPY image-syncer /root/image-syncer

COPY auth.json /root/auth.json

COPY habor-sync /root/habor-sync

COPY entrypoint.sh /root/entrypoint.sh

WORKDIR /root

RUN chmod +x /root/image-syncer && chmod +x /root/habor-sync  && chmod +x /root/entrypoint.sh

ENTRYPOINT [ "/root/entrypoint.sh" ]