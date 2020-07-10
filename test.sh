#/bin/bash

# unit test the code
go test || exit

sudo docker-compose up --build -d

# run test
test_output=$(sudo docker wait ecclesia_test_1)

# collect logs
sudo docker logs ecclesia_test_1 &> test.log

#results
echo
cat test.log
if [[ "$test_output" != "0" ]]
then
    echo "***************************************************"
    echo "***               TESTS FAILED                  ***"
    echo "***************************************************"
else
    echo "==================================================="
    echo "===               TESTS PASSED                  ==="
    echo "==================================================="
fi

# clean up
sudo docker-compose down
sudo docker ps -a