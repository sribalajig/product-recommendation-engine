1. Install Elasticsearch => https://www.elastic.co/guide/en/elasticsearch/reference/current/_installation.html. You can verify whether the installation works by executing the following command => "curl -i -XGET 'localhost:9200/'"
2. Import the product attribute data into elastic search. Navigate to the /data folder. You will see a file called setup.sh. Execute "sh setup.sh" in the command line. This will import data into elastic search.
3. Install Golang => https://golang.org/doc/install
4. Install beego which is a micro framework for Go.  Execute "go get github.com/astaxie/beego" in the command line.
5. Install the elastic library for Golang => Execute "go get gopkg.in/olivere/elastic.v5" in the command line.
6. Navigate to /recommendations.api. Execute "bee run".
7. The runs on port 8080. You can try out the following request from any HTTP client :

	http://localhost:8080/api/products/sku-7/recommendations

	You can replace sku-7 with any valid product id.
8. There is a test for the recommendation engine. Go to the root directory and execute "go test ./..." to run the test.