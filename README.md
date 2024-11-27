# PoC: Bargain Hunter
This was done as a proof-of-concept as part of training to learn test-driven development in Go. The tool is supposed to keep track of the price of a product (via a URL passed to it) and notify the user whenever there is a bargain (e.g., a price drop). It was never completed, but concepts were learned.

# Usage

**`bargain https://www.newegg.com/lg-27gn950-b-27-uhd-144-hz-ultragear-nano-ips-black/p/N82E16824026052`**

```
LG 27" 144 Hz Nano IPS UHD Nano IPS Gaming Monitor AMD FreeSync Premium Pro and NVIDIA G-SYNC Compatible 3840 x 2160 (4K) 2xHDMI, DisplayPort, USB UltraGear 27GN950-B is not a good bargain today — price has been lower within the last month
```

# Errors

**`bargain https://example.com/bogus`**
```
Sorry, I don't recognise that as a product page I know how to extract the price from
```
