for i in $(seq 2000) # perfrom the inner command 5000 times.
do
    curl --location --request GET 'http://localhost:3000' > /dev/null & # send out a curl request, the & indicates not to wait for the response.
done

wait