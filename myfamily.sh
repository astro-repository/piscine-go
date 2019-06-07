curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq ".[] | select( .id==70 ) | .connections.relatives" | sed -e 's/"//g'
