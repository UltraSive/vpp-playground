
//go:build ignore

#include "vmlinux.h"
// #include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
/*#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/ipv6.h>
#include <linux/tcp.h>
#include <linux/udp.h>*/

#define IP_PROTO_TCP 6
#define IP_PROTO_UDP 17
#define IP_PROTO_ICMP 1
#define IP6_PROTO_ICMPV6 58

// Define __constant_htons if not defined
#ifndef __constant_htons
#define __constant_htons(x) __builtin_bswap16(x)
#endif

// Define ETH_P_IP and ETH_P_IPV6 if not defined
#ifndef ETH_P_IP
#define ETH_P_IP 0x0800
#endif

#ifndef ETH_P_IPV6
#define ETH_P_IPV6 0x86DD
#endif

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __type(key, __u32); // 0
    __type(value, __u64);
    __uint(max_entries, 1);
} pkt_count SEC(".maps");

struct address {
    __u8 addr[16]; // Always 16 bytes to accommodate both IPv4 and IPv6
};

/*
enum src_usage {
    SOURCE_TUNNEL,
    SOURCE_BACKHAUL
};

struct source_info {
    enum src_usage usage;
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, sizeof(struct address));
    __type(value, sizeof(struct source_info));
    __uint(max_entries, 1024);
} source_map SEC(".maps");
*/

enum dst_usage {
    LOCAL_MACHINE,
    DEST_TUNNEL,
    FACILITY
};

struct destination_info {
    enum dst_usage usage;
    __u16 location_id;
    __u8 default_action;
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, sizeof(struct address));
    __type(value, sizeof(struct destination_info));
    __uint(max_entries, 1024);
} destination_map SEC(".maps");

struct tunnel {
    struct address local_ip;
    __u16 local_port;
    struct address remote_ip;
    __u16 remote_port;
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, sizeof(struct tunnel));
    __type(value, sizeof(struct tunnel));
    __uint(max_entries, 1024);
} tunnel_map SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_XSKMAP);
    __uint(max_entries, 64); // Maximum queues for XDP sockets
    __type(key, __u32);
    __type(value, __u32);
} xsks_map SEC(".maps");

// count_packets atomically increases a packet counter on every invocation.
SEC("xdp") 
int xdp_sock_prog(struct xdp_md *ctx) {
     /* START COUNT PACKET */
    __u32 key = 0;
    __u64 *count = bpf_map_lookup_elem(&pkt_count, &key);
    if (count)
    {
        __sync_fetch_and_add(count, 1);
    }
    /* END COUNT PACKET */
    
    void *data_end = (void *)(long)ctx->data_end;
    void *data = (void *)(long)ctx->data;

    struct ethhdr *eth = data;
    if ((void *)(eth + 1) > data_end)
        return XDP_PASS;

    __u64 now_ns = bpf_ktime_get_ns();

    __u32 rx_queue_index = ctx->rx_queue_index; // Get the RX queue index
    if (bpf_map_lookup_elem(&xsks_map, &rx_queue_index)) {
        return bpf_redirect_map(&xsks_map, ctx->rx_queue_index, XDP_REDIRECT);
    }

    return XDP_PASS; 
}

char __license[] SEC("license") = "Dual MIT/GPL";