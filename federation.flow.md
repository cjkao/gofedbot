# setup sequence
- setup coredns
- add `127.0.0.1` to `/etc/resolve.conf` before meteor start
- update `/etc/hosts` for browser
- start meteor 
  mongo port will be 3001 and 4001
   - meteor
   - meteor --port 4000

```
   play.cj 127.0.0.1
   example.cj 127.0.0.1
```

# coredns
2 domains, `play.cj` and `example.cj`


## Corefile

```

example.cj {
    file example.cj
    log
    errors
    reload
}
play.cj {
    file play.cj
    log
    errors
    reload
}


```

## play.cj

```ini
$ORIGIN play.cj.
@	3600 IN	SOA sns.dns.icann.org. noc.dns.icann.org. (
				2017042745 ; serial
				7200       ; refresh (2 hours)
				3600       ; retry (1 hour)
				1209600    ; expire (2 weeks)
				3600       ; minimum (1 hour)
				)

	3600 IN NS a.iana-servers.net.
	3600 IN NS b.iana-servers.net.

www     IN A     127.0.0.1
@       IN A     127.0.0.1
        IN AAAA  ::1
_rocketchat._http.play.cj. 1800 IN SRV 1 1 3000 play.cj.
rocketchat-public-key.play.cj. IN TXT 127.0.0.1 "-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiwtnfR2Qrlz0AXOdc0F+8JqQI5VyEe08cPQUwI8hGm9kUHUjF3XvebZaAL/Efpal7ivMRwzge7kYnChkzwTPjd415ek1N+EY6uzWnXw3hm1tMZEbnWnZvDwoEIbA77S/eJe1Ir7MDeJBn6qGB0s84vUunrHqDFV8oVtZIHpjHi+YD0hLRa" "XMn+B7sF1NYmjlq1x6Jz6iExnu0AcSWBoGqIERSxNKmG0l4OBTKguzKtIZTUIc2Wr/EHUYlYiwiBTOwVmpPqrsy883CLKd1dh1n1/oHFYD2CPKNFIlB9PzTAeGMwFrpiN6ZlrKT1CnUFOfWzpYkwBAtDtEQxnTNgGILQIDAQAB-----END PUBLIC KEY-----"
```

## example.cj

```ini
$ORIGIN example.cj.
@	3600 IN	SOA sns.dns.icann.org. noc.dns.icann.org. (
				2017042745 ; serial
				7200       ; refresh (2 hours)
				3600       ; retry (1 hour)
				1209600    ; expire (2 weeks)
				3600       ; minimum (1 hour)
				)

	3600 IN NS a.iana-servers.net.
	3600 IN NS b.iana-servers.net.
@       IN A     127.0.0.1
www     IN A     127.0.0.1
        IN AAAA  ::1
_rocketchat._http.example.cj. 1800 IN SRV 1 1 4000 example.cj.
rocketchat-public-key.example.cj. IN TXT 127.0.0.1 "-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiwtnfR2Qrlz0AXOdc0F+8JqQI5VyEe08cPQUwI8hGm9kUHUjF3XvebZaAL/Efpal7ivMRwzge7kYnChkzwTPjd415ek1N+EY6uzWnXw3hm1tMZEbnWnZvDwoEIbA77S/eJe1Ir7MDeJBn6qGB0s84vUunrHqDFV8oVtZIHpjHi+YD0hLRa" "XMn+B7sF1NYmjlq1x6Jz6iExnu0AcSWBoGqIERSxNKmG0l4OBTKguzKtIZTUIc2Wr/EHUYlYiwiBTOwVmpPqrsy883CLKd1dh1n1/oHFYD2CPKNFIlB9PzTAeGMwFrpiN6ZlrKT1CnUFOfWzpYkwBAtDtEQxnTNgGILQIDAQAB-----END PUBLIC KEY-----"
```