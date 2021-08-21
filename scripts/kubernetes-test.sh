#!/bin/bash 
MAX_ATTEMPTS=30
CURRENT_ATTEMPT=1
DELAY=${1:-0}
PORT=${2:-0}

echo "Starting up - Delay: ${DELAY} seconds"
sleep $DELAY

while [ $CURRENT_ATTEMPT -le $MAX_ATTEMPTS ]; do
    echo "Testing connection - Attempt: ${CURRENT_ATTEMPT} of ${MAX_ATTEMPTS}"
    RESULT=$(curl -o /dev/null -s -w "%{http_code}\n" http://34.79.255.98:${PORT}/)
    
    if [ $RESULT -ne '200' ]; then
     echo "Connection didn't return a status code of 200. Connection failed"
     exit 1
    fi

    CURRENT_ATTEMPT=$((CURRENT_ATTEMPT + 1))
    sleep 1s
done

echo "All connection attempts were successful"
exit 0

