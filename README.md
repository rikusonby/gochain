docker build -t chain .

docker run -d --name chain -p 1323:1323 -v /Users/www/other/chain:/www chain
