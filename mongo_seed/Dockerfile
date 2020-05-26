FROM mongo

COPY init.json /init.json
CMD mongoimport --host mongodb --db weather --collection users --type json --file /init.json --jsonArray