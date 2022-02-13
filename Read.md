## GoBazaar

# An e-commerce app that runs in terminal (backend)

# users can register themselves
# merchants can register
# merchant can upload any product, delete, and offer discount
# product categories
# search result
# users will have a wallet
# user can purchase any product , cancel the product before delivery date
# user-merchant transaction
# coupons and cash back
# user purchase chart
# customer support

## Bonus:
# recommendation engine
# delivery dates and ratings
# big billion sale 


## Road-map

# main.go (url that the user will hit)
# handlers (routing)
# controllers ( preparation )
# service ( real action )
# model
# database 



## hit the api through terminal
# add the user
curl http://localhost:8080/user/register \
     --include \
     --header "Content-Type: application/json"  \
     --request "POST" \
     --data '{"first_name": "Tarun1","last_name": "thakur1","email":"thakur.tarun@gmail.com","contact": "949494949","city":"delhi","wallet_balance":90.30}'


# user login 
curl http://localhost:8080/user/login \
     --include \
     --header "Content-Type: application/json"  \
     --request "POST" \
     --data '{"Id":1 ,"Pass":"Pass123"}'

# user logout
curl http://localhost:8080/user/logout

# user purchase
curl http://localhost:8080/user/purchase \
     --include \
     --header "Content-Type: application/json"  \
     --request "POST" \
     --data '{"id":}'


# add the merchant

curl http://localhost:8080/merchant/register \
     --include \
     --header "Content-Type: application/json"  \
     --request "POST" \
     --data '{"company_name": "Cipher","email":"thakur.tarun@gmail.com","merchant_address": "kjskjskjskjs","discount_offered":10.0}'

# merchant login
curl http://localhost:8080/merchant/login \
     --include \
     --header "Content-Type: application/json"  \
     --request "POST" \
     --data '{"Id":1 ,"Pass":"Pass123"}'


# merchant upload
curl http://localhost:8080/merchant/upload \
     --include \
     --header "Content-Type: application/json"  \
     --request "POST" \
     --data '{"Id":1 ,"Pass":"Pass123"}'     

     ProductID          int           `json:"product_id"`
	MerchantID         int           `json:"merchant_id"`
	Name               string        `json:"name"`
	ProductDescription string        `json:"product_description"`
	Price              float64       `json:"price"`
	Stock              int           `json:"stock"`
	DeliveryTime       time.Duration `json:"deliveryTime"`