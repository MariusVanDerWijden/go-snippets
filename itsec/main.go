package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"sync"

	"golang.org/x/sys/unix"
)

func main() {
	Spoof()
}

func ex4() {
	n := "344296855822"
	should := "70bff725118c3d777e1ca675d3d5779b"
	for i := 0; i < 99999999; i++ {
		str := fmt.Sprintf("%v %v", n, i)
		sum := fmt.Sprintf("%x", md5.Sum([]byte(str)))
		if sum == should {
			fmt.Println("SUM == Should")
			fmt.Println(i)
		}
	}
}

func aes() {

}

func replicate() {
	ownid := "a9809b7c7e"
	trunk := "trunk"
	n := "6942"
	salt := "3175039"
	str := fmt.Sprintf("%v%v%v%v", ownid, trunk, n, salt)
	fmt.Printf("Hello, playground: %x\n", sha256.Sum256([]byte(str)))
	fmt.Printf("salt: %v\n", salt)
	fmt.Println(str)
}

func salt() {
	ownid := "0909a9db49"
	trunk := "nav"
	n := "2711"
	salt := "1"
	replicate()

	search := "f74edd52d15b2574fb8e407f1e21f9a51baef3dd74beebf90312c58d15e3dce7"
	for i := 0; i < 10000000; i++ {
		salt = fmt.Sprintf("%d", i)
		str := fmt.Sprintf("%v%v%v%v", ownid, trunk, n, salt)
		if fmt.Sprintf("%x", sha256.Sum256([]byte(str))) == search {
			fmt.Printf("Hello, playground: %x", sha256.Sum256([]byte(str)))
			fmt.Printf("salt: %v", salt)
		}
	}
}

type iphdr struct {
	vhl   uint8
	tos   uint8
	iplen uint16
	id    uint16
	off   uint16
	ttl   uint8
	proto uint8
	csum  uint16
	src   [4]byte
	dst   [4]byte
}

type udphdr struct {
	src  uint16
	dst  uint16
	ulen uint16
	csum uint16
}

// pseudo header used for checksum calculation
type pseudohdr struct {
	ipsrc   [4]byte
	ipdst   [4]byte
	zero    uint8
	ipproto uint8
	plen    uint16
}

func checksum(buf []byte) uint16 {
	sum := uint32(0)

	for ; len(buf) >= 2; buf = buf[2:] {
		sum += uint32(buf[0])<<8 | uint32(buf[1])
	}
	if len(buf) > 0 {
		sum += uint32(buf[0]) << 8
	}
	for sum > 0xffff {
		sum = (sum >> 16) + (sum & 0xffff)
	}
	csum := ^uint16(sum)
	/*
	 * From RFC 768:
	 * If the computed checksum is zero, it is transmitted as all ones (the
	 * equivalent in one's complement arithmetic). An all zero transmitted
	 * checksum value means that the transmitter generated no checksum (for
	 * debugging or for higher level protocols that don't care).
	 */
	if csum == 0 {
		csum = 0xffff
	}
	return csum
}

func (h *iphdr) checksum() {
	h.csum = 0
	var b bytes.Buffer
	binary.Write(&b, binary.BigEndian, h)
	h.csum = checksum(b.Bytes())
}

func (u *udphdr) checksum(ip *iphdr, payload []byte) {
	u.csum = 0
	phdr := pseudohdr{
		ipsrc:   ip.src,
		ipdst:   ip.dst,
		zero:    0,
		ipproto: ip.proto,
		plen:    u.ulen,
	}
	var b bytes.Buffer
	binary.Write(&b, binary.BigEndian, &phdr)
	binary.Write(&b, binary.BigEndian, u)
	binary.Write(&b, binary.BigEndian, &payload)
	u.csum = checksum(b.Bytes())
}

func Spoof() {
	ipsrcstr := "2001:41b8:83f:4250::50"
	ipdststr := "130.83.186.170"
	udpsrc := uint(13338)
	udpdst := uint(13337)
	showcsum := false

	ipsrc := net.ParseIP(ipsrcstr)
	if ipsrc == nil {
		fmt.Fprintf(os.Stderr, "invalid source IP: %v\n", ipsrc)
		os.Exit(1)
	}
	ipdst := net.ParseIP(ipdststr)
	if ipdst == nil {
		fmt.Fprintf(os.Stderr, "invalid destination IP: %v\n", ipdst)
		os.Exit(1)
	}

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_RAW, unix.IPPROTO_RAW)

	if err != nil || fd < 0 {
		fmt.Fprintf(os.Stdout, "error creating a raw socket: %v\n", err)
		os.Exit(1)
	}

	err = unix.SetsockoptInt(fd, unix.IPPROTO_IP, unix.IP_HDRINCL, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error enabling IP_HDRINCL: %v\n", err)
		unix.Close(fd)
		os.Exit(1)
	}

	ip := iphdr{
		vhl:   0x45,
		tos:   0,
		id:    0x1234, // the kernel overwrites id if it is zero
		off:   0,
		ttl:   64,
		proto: unix.IPPROTO_UDP,
	}
	copy(ip.src[:], ipsrc.To4())
	copy(ip.dst[:], ipdst.To4())
	// iplen and csum set later

	udp := udphdr{
		src: uint16(udpsrc),
		dst: uint16(udpdst),
	}
	// ulen and csum set later

	// just use an empty IPv4 sockaddr for Sendto
	// the kernel will route the packet based on the IP header
	addr := unix.SockaddrInet4{}

	payload := []byte("mv94naqi")
	udplen := 8 + len(payload)
	totalLen := 20 + udplen
	if totalLen > 0xffff {
		fmt.Fprintf(os.Stderr, "message is too large to fit into a packet: %v > %v\n", totalLen, 0xffff)
	}

	// the kernel will overwrite the IP checksum, so this is included just for
	// completeness
	ip.iplen = uint16(totalLen)
	ip.checksum()

	// the kernel doesn't touch the UDP checksum, so we can either set it
	// correctly or leave it zero to indicate that we didn't use a checksum
	udp.ulen = uint16(udplen)
	udp.checksum(&ip, payload)
	if showcsum {
		fmt.Printf("ip checksum: 0x%x, udp checksum: 0x%x\n", ip.csum, udp.csum)
	}

	var b bytes.Buffer
	err = binary.Write(&b, binary.BigEndian, &ip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding the IP header: %v\n", err)
	}
	err = binary.Write(&b, binary.BigEndian, &udp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding the UDP header: %v\n", err)
	}
	err = binary.Write(&b, binary.BigEndian, &payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding the payload: %v\n", err)
	}
	bb := b.Bytes()

	/*
	* For some reason, the IP header's length field needs to be in host byte order
	* in OS X.
	 */
	if runtime.GOOS == "darwin" {
		bb[2], bb[3] = bb[3], bb[2]
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		buf := make([]byte, 8)
		nr, err := unix.Read(fd, buf)
		if err != nil {
			return
		}
		fmt.Printf("Read: %x : %v", buf, nr)
		wg.Done()
	}()

	err = unix.Sendto(fd, bb, 0, &addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending the packet: %v\n", err)
	}
	fmt.Printf("%v bytes were sent\n", len(bb))
	wg.Wait()
	err = unix.Close(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error closing the socket: %v\n", err)
		os.Exit(1)
	}
}

func RSA() {
	buf, err := ioutil.ReadFile("input.bin")
	if err != nil {
		fmt.Print("err reading file")
	}
	rsa_enc := buf[:16]
	
}
