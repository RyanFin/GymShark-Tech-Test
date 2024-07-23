# GymShark-Tech-Test
Golang Tech Test for Gymshark.

Software Engineering Challenge
Imagine for a moment that one of our product lines ships in various pack sizes:
	•	250 Items
	•	500 Items
	•	1000 Items
	•	2000 Items
	•	5000 Items
Our customers can order any number of these items through our website, but they will always only be given complete packs.
	•	Only whole packs can be sent. Packs cannot be broken open.
	•	Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
	•	Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.
So, for example:
Items ordered
Correct number of packs
Incorrect number of packs
1
1 x 250
1 x 500 – more items than necessary
250
1 x 250
1 x 500 – more items than necessary
251
1 x 500
2 x 250 – more packs than necessary
501
1 x 5001 x 250
1 x 1000 – more items than necessary3 x 250 – more packs than necessary
12001
2 x 50001 x 20001 x 250
3 x 5000 – more items than necessary

Write an application that can calculate the number of packs we need to ship to the customer.
The API must be written in Golang & be usable by a HTTP API (by whichever method you choose).
Optional: 
	•	Keep your application flexible so that pack sizes can be changed and added and removed without having to change the code.
	•	Create a UI to interact with your API
Please also send us your code via a publicly accessible git repository, GitHub or similar is fine, and deploy your application to an online environment so that we can access it and test your application out.
We look forward to receiving your application! Please return your completed solution to talent@gymshark.com by 1pm on Monday 29th July 2024. From here, we look forward to welcoming you into the office on Monday 5th August! 

## Development Notes
### Swagger

- [Swagger with Gin](https://santoshk.dev/posts/2022/how-to-integrate-swagger-ui-in-go-backend-gin-edition/)
  - docs/swagger.json and docs/swagger.yaml are the actual specification which you can upload to services like AWS API Gateway and similar services.
    Install gin-swagger

```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

- Generate swagger specs with: `$ swag init`
- Visit swagger docs at this URL: http://localhost:8080/docs/index.html

#### Requests
```
curl --location 'http://localhost:8080/calculate-packs/770'

curl --location 'http://localhost:8080/view-packsizes'

curl --location --request DELETE 'http://localhost:8080/remove-packsize?packsize=1000'

curl --location --request POST 'http://localhost:8080/add-packsize?packsize=800'

```

## Test Notes
```
--- FAIL: TestCalculatePacks (0.00s)
    --- FAIL: TestCalculatePacks/Order_size_just_higher_than_the_smallest_pack_size (0.00s)
        model_test.go:59: calculatePacks(251) = map[250:2]; want map[500:1]
FAIL
```
- Currently getting the wrong amount of packs for this test case. I should get one pack rather than 2 X 250 when the order size = 251.

Solution: Add logic to combine smaller packs into a larger one


