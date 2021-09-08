docker run -d --restart=always \
--name godbpostgis \
-p 5432:5432 \
-e "POSTGRES_USER=admin" \
-e "POSTGRES_PASS=admin" \
-e "POSTGRES_MULTIPLE_EXTENSIONS=postgis,hstore,postgis_topology,ogr_fdw,uuid-ossp" \
-e "POSTGRES_DB=godb" \
-e "ALLOW_IP_RANGE=0.0.0.0/0" \
-e "PG_USE_SSL=false" \
kartoza/postgis:latest