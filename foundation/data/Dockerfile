FROM mongo:latest
COPY init_weather.json /init_weather.json
CMD mongoimport --host mongo --db weather --collection weather_conditions --drop --file /init_weather.json --jsonArray 