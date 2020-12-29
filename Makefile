.PHONY:
gen:
	 GO111MODULE=off  swagger generate spec -o ./api/swagger/swagger.yaml --scan-models