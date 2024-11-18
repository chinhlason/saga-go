DSN1="postgresql://root:root@localhost:5432/order?sslmode=disable"
DSN2="postgresql://root:root@localhost:5433/payment?sslmode=disable"
DSN3="postgresql://root:root@localhost:5434/inventory?sslmode=disable"
DSN4="postgresql://root:root@localhost:5435/delivery?sslmode=disable"

g-up:
	goose -dir ./order/migrations postgres $(DSN1) up
	goose -dir ./payment/migrations postgres $(DSN2) up
	goose -dir ./inventory/migrations postgres $(DSN3) up
	goose -dir ./delivery/migrations postgres $(DSN4) up

g-down:
	goose -dir ./order/migrations postgres $(DSN1) down
	goose -dir ./payment/migrations postgres $(DSN2) down
	goose -dir ./inventory/migrations postgres $(DSN3) down
	goose -dir ./delivery/migrations postgres $(DSN4) down