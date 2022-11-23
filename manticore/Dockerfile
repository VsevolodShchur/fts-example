FROM manticoresearch/manticore:5.0.2

RUN apt-get update && apt-get install -y

COPY manticore.conf /etc/manticoresearch/manticore.conf
COPY wordforms.txt /etc/manticore/wordforms/wordforms.txt

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
