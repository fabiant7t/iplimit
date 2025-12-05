package iplimit

import (
	"github.com/miekg/dns"
)

func ModResponseWriter(w dns.ResponseWriter, mods ...ResponseModifier) dns.ResponseWriter {
	return &modResponseWriter{
		ResponseWriter:    w,
		responseModifiers: mods,
	}
}

type modResponseWriter struct {
	dns.ResponseWriter
	responseModifiers []ResponseModifier
}

func (w *modResponseWriter) WriteMsg(res *dns.Msg) error {
	if res == nil {
		return w.ResponseWriter.WriteMsg(res)
	}
	for _, mod := range w.responseModifiers {
		if err := mod(res); err != nil {
			return err
		}
	}
	return w.ResponseWriter.WriteMsg(res)
}
