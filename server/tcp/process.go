package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type result struct {
	v []byte
	e error
}

func (s *Server) process(conn net.Conn) {
	r := bufio.NewReader(conn)
	resultCh := make(chan chan *result, 5000)
	defer close(resultCh)
	go reply(conn, resultCh)

	for {
		op, e := r.ReadByte()
		if e != nil {
			if e != io.EOF {
				log.Println("close connection due to error:", e)
			}
			return
		}
		switch op {
		case 'S':
			s.set(resultCh, r)
		case 'G':
			s.get(resultCh, r)
		case 'D':
			s.del(resultCh, r)
		default:
			log.Println("close connection due to invalid operation:", op)
			return
		}
	}
}

func reply(conn net.Conn, resultCh chan chan *result) {
	defer conn.Close()
	for {
		c, open := <-resultCh
		if !open {
			return
		}
		r := <-c
		e := sendResponse(r.v, r.e, conn)
		if e != nil {
			log.Println("close connection due to error:", e)
			return
		}
	}
}

func (s *Server) get(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c
	k, e := s.readKey(r)
	if e != nil {
		c <- &result{nil, e}
		return
	}
	go func() {
		v, e := s.Get(k)
		c <- &result{v, e}
	}()
}

func (s *Server) set(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c
	k, v, e := s.readKeyAndValue(r)
	if e != nil {
		c <- &result{nil, e}
		return
	}
	go func() {
		c <- &result{nil, s.Set(k, v)}
	}()
}

func (s *Server) del(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c
	k, e := s.readKey(r)
	if e != nil {
		c <- &result{nil, e}
		return
	}
	go func() {
		c <- &result{nil, s.Del(k)}
	}()
}
func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d ", len(errString)) + errString
		_, e := conn.Write([]byte(tmp))
		return e
	}
	vlen := fmt.Sprintf("%d ", len(value))
	_, e := conn.Write(append([]byte(vlen), value...))
	return e
}

func readLen(r *bufio.Reader) (int, error) {
	tmp, e := r.ReadString(' ')
	if e != nil {
		return 0, e
	}
	l, e := strconv.Atoi(strings.TrimSpace(tmp))
	if e != nil {
		return 0, e
	}
	return l, nil
}

func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}
	vlen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}
	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", nil, e
	}

	addr, ok := s.ShouldProcess(string(k))
	if !ok {
		return "", nil, errors.New("redirect " + addr)
	}

	v := make([]byte, vlen)
	_, e = io.ReadFull(r, v)
	if e != nil {
		return "", nil, e
	}
	return string(k), v, nil
}

func (s Server) readKey(r *bufio.Reader) (string, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", e
	}
	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", e
	}
	addr, ok := s.ShouldProcess(string(k))
	if !ok {
		return "", errors.New("redirect " + addr)
	}
	return string(k), nil
}
