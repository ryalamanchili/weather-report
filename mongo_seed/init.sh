#! /bin/bash

mongoimport --host mongodb --db weather --collection users --type json --file ./mongo_seed/init.json --jsonArray