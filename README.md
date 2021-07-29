DONOTSHOUT
==========

... please. A simple chaotic randomized DNS server used to diagnose
problems with DNS resolvers. It is capable of adding a random delay,
randomly truncating and randomly dropping packets. It answers all
A and AAAA queries with the same IP address, both configurable.

## Quick Start

The quickest way is to simply run it right from the repository:

    $ sudo go run main.go # we bind to port 53 by default so we
                          # need root privs.

To drop 10% of all packets, truncate 5% and do a 
variable delay of between 1-10 seconds, answering with
the IPv4 Address of 1.1.1.1:

    $ sudo env DROP_PERCENT=10 TRUNCATE_PERCENT=5 MIN_JITTER=1000 \
        MAX_JITTER=10000 IPV4_ADDRESS=1.1.1.1 go run main.go

The latest version is also available as a docker image:

    $ docker run --rm --name donotshout -ti calyptia/donotshout:latest

## Configuration

Configuration is done either via environment variables or
a .env file. This is a full list of all configuration options:

  * HOST: Host address to listen on.
  * PORT: Port to listen on.
  * PROTOCOL: Protocol to listen on, either UDP or TCP.
  * DROP_PERCENT: the percentage of packets to drop.
  * TRUNC_PERCENT: the percentage of packets to truncate.
  * MIN_JITTER: the bottom range of packet delay in ms.
  * MAX_JITTER: the upper range of packet delay in ms.
  * IPV4_ADDRESS: the IPv4 address used to responde to A requests.
  * IPV6_ADDRESS: the IPv6 address used to responde to AAAA requests.
