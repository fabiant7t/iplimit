# iplimit

## Name

*iplimit* - limits the number of IP addresses returned in A and AAAA DNS records.

## Description

The *iplimit* plugin is a CoreDNS middleware that restricts the number of IP addresses returned in DNS responses for A (IPv4) and AAAA (IPv6) record queries. This is useful for load balancing scenarios and reducing DNS response sizes.

When a DNS query returns multiple A or AAAA records, *iplimit* will trim the response to contain only the specified maximum number of IP addresses. The plugin operates as a response modifier, intercepting DNS responses from upstream plugins before they reach the client.

## Syntax

```
iplimit MAX_IPS
```

* **MAX_IPS** is a required integer specifying the maximum number of IP addresses to return in A and AAAA records. Must be a positive integer (1 or greater).

## Examples

### Basic Usage

Limit A and AAAA records to a max 5 IP addresses:

```corefile
.:53 {
	loadbalance round_robin
	iplimit 5
	forward . 1.1.1.1
}
```

## How It Works

The *iplimit* plugin works as a response modifier in the CoreDNS plugin chain:

1. It intercepts DNS queries and passes them to the next plugin in the chain
2. When a response comes back, it examines the answer section
3. For each A or AAAA record type, it limits the number of records to the configured maximum
4. The modified response is then returned to the client

The plugin preserves all other record types (CNAME, MX, TXT, etc.) and only affects A and AAAA records.

## License

This plugin is released under the MIT License. See the LICENSE file for details.

## Author

Fabian Topfstedt
