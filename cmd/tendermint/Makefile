CC = gcc
GC = go run
INDEXES = 0 1 2 3

.PHONY: default local build run clean
default: build run
local: clean 
	$(GC) main.go init validator
	$(GC) main.go start --home="$$HOME/.tendermint" --proxy_app=kvstore
build: clean 
	$(GC) main.go testnet
	$(CC) build_testnet.c && ./a.out
	rm -f a.out
run:
	for i in $(INDEXES) ; do \
		gnome-terminal -- sh -c "./run_node$$i.sh; bash" ; \
	done
clean:
	rm -f a.out
	rm -f run_node*.sh
	rm -rf mytestnet
	rm -rf ~/.tendermint
