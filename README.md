System requirements

1. Install Elasticsearch => https://www.elastic.co/guide/en/elasticsearch/reference/current/_installation.html. Make sure elastic search is running => "curl -i -XGET 'localhost:9200/'"

2. Install Golang => https://golang.org/doc/install. This project is implemented in Golang.

3. Install npm (node package manager).

Running the application

1. Navigate to /setup. Run "sh setup.sh". This will import data to elastic search and install the Go dependencies. You will be prompted for your password.

2. Navigate to /recommendations.api. Execute "bee run". Bee is a micro-framework for Golang => https://beego.me/docs/install/

HTTP Requests to test the engine

1. The recommendation service runs on port 8080. You can try out the following request from any HTTP client =>

	http://localhost:8080/api/products/sku-7/recommendations

	You can replace sku-7 with any valid product id.

2. There is a test for the recommendation engine. Go to the root directory and execute "go test ./..." to run the test.