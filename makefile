it_test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from it_tests

it_test_down:
	docker-compose -f docker-compose.test.yml down

prod:
	docker-compose -f docker-compose.deploy.yml up --build --abort-on-container-exit --exit-code-from go_server_deploy

prod_down:
	docker-compose -f docker-compose.deploy.yml down