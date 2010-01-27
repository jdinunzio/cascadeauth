# In a 386
#cc=8g
#ln=8l

# In a x64
cc=6g
ln=6l

cascadeauth: cascadeauth.go
	$(cc) -o cascadeauth.6 cascadeauth.go
	$(ln) -o cascadeauth   cascadeauth.6

test: test.go
	$(cc) -o test.6 test.go
	$(ln) -o test   test.6
	./test

run:
	./cascadeauth cascade.conf

clean:
	rm -f cascadeauth cascadeauth.6
