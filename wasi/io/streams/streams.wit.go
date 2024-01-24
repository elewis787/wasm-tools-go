// This file is generated by hand, as part of designing wit-bindgen-go. PLEASE EDIT.

// Package streams represents the interface "wasi:io/streams".
//
// WASI I/O is an I/O abstraction API which is currently focused on providing
// stream types.
//
// In the future, the component model is expected to add built-in stream types;
// when it does, they are expected to subsume this API.
package streams

import (
	"github.com/ydnar/wasm-tools-go/cm"
	ioerror "github.com/ydnar/wasm-tools-go/wasi/io/error"
	"github.com/ydnar/wasm-tools-go/wasi/io/poll"
)

type (
	Error    = ioerror.Error
	Pollable = poll.Pollable
)

// StreamError represents variant "wasi:io/streams.stream-error".
//
// An error for input-stream and output-stream operations.
type StreamError struct {
	v cm.Variant[bool, Error, Error]
}

// LastOperationFailed represents variant case "last-operation-failed(error)".
//
// The last operation (a write or flush) failed before completion.
//
// More information is available in the `error` payload.
func (self *StreamError) LastOperationFailed() (Error, bool) {
	return cm.Get[Error](&self.v, false)
}

// Closed represents variant case "closed".
//
// The stream is closed: no more input will be accepted by the
// stream. A closed output-stream will return this error on all
// future operations.
func (self *StreamError) Closed() bool {
	return self.v.Is(true)
}

// InputStream represents the resource "wasi:io/streams.input-stream".
//
// An input bytestream.
//
// `input-stream`s are *non-blocking* to the extent practical on underlying
// platforms. I/O operations always return promptly; if fewer bytes are
// promptly available than requested, they return the number of bytes promptly
// available, which could even be zero. To wait for data to be available,
// use the `subscribe` function to obtain a `pollable` which can be polled
// for using `wasi:io/poll`.
type InputStream cm.Resource

// Read represents the method "wasi:io/streams.input-stream#read".
//
// Perform a non-blocking read from the stream.
//
// This function returns a list of bytes containing the read data,
// when successful. The returned list will contain up to `len` bytes;
// it may return fewer than requested, but not more. The list is
// empty when no bytes are available for reading at this time. The
// pollable given by `subscribe` will be ready when more bytes are
// available.
//
// This function fails with a `stream-error` when the operation
// encounters an error, giving `last-operation-failed`, or when the
// stream is closed, giving `closed`.
//
// When the caller gives a `len` of 0, it represents a request to
// read 0 bytes. If the stream is still open, this call should
// succeed and return an empty list, or otherwise fail with `closed`.
//
// The `len` parameter is a `u64`, which could represent a list of u8 which
// is not possible to allocate in wasm32, or not desirable to allocate as
// as a return value by the callee. The callee may return a list of bytes
// less than `len` in size while more bytes are available for reading.
func (self InputStream) Read(len uint64) cm.OKSizedResult[cm.List[uint8], StreamError] {
	var ret cm.OKSizedResult[cm.List[uint8], StreamError]
	self.read(len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]input-stream.read
func (self InputStream) read(len uint64, ret *cm.OKSizedResult[cm.List[uint8], StreamError])

// BlockingRead represents the method "wasi:io/streams.input-stream.blocking-read".
//
// Read bytes from a stream, after blocking until at least one byte can
// be read. Except for blocking, behavior is identical to `read`.
func (self InputStream) BlockingRead(len uint64) cm.OKSizedResult[cm.List[uint8], StreamError] {
	var ret cm.OKSizedResult[cm.List[uint8], StreamError]
	self.blocking_read(len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]input-stream.blocking-read
func (self InputStream) blocking_read(len uint64, ret *cm.OKSizedResult[cm.List[uint8], StreamError])

// Skip represents the method "wasi:io/streams.input-stream#skip".
//
// Skip bytes from a stream. Returns number of bytes skipped.
//
// Behaves identical to `read`, except instead of returning a list
// of bytes, returns the number of bytes consumed from the stream.
func (self InputStream) Skip(len uint64) cm.OKSizedResult[uint64, StreamError] {
	var ret cm.OKSizedResult[uint64, StreamError]
	self.skip(len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]input-stream.skip
func (self InputStream) skip(len uint64, ret *cm.OKSizedResult[uint64, StreamError])

// BlockingSkip represents the method "wasi:io/streams.input-stream.blocking-skip".
//
// Skip bytes from a stream, after blocking until at least one byte
// can be skipped. Except for blocking behavior, identical to `skip`.
func (self InputStream) BlockingSkip(len uint64) cm.OKSizedResult[uint64, StreamError] {
	var ret cm.OKSizedResult[uint64, StreamError]
	self.blocking_skip(len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]input-stream.blocking-skip
func (self InputStream) blocking_skip(len uint64, ret *cm.OKSizedResult[uint64, StreamError])

// Subscribe represents the method "wasi:io/streams.input-stream.subscribe".
//
// Create a `pollable` which will resolve once either the specified stream
// has bytes available to read or the other end of the stream has been
// closed.
// The created `pollable` is a child resource of the `input-stream`.
// Implementations may trap if the `input-stream` is dropped before
// all derived `pollable`s created with this function are dropped.
func (self InputStream) Subscribe() Pollable {
	return self.subscribe()
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]input-stream.subscribe
func (self InputStream) subscribe() Pollable

// OutputStream represents the resource "wasi:io/streams.output-stream".
//
// An output bytestream.
//
// `output-stream`s are *non-blocking* to the extent practical on
// underlying platforms. Except where specified otherwise, I/O operations also
// always return promptly, after the number of bytes that can be written
// promptly, which could even be zero. To wait for the stream to be ready to
// accept data, the `subscribe` function to obtain a `pollable` which can be
// polled for using `wasi:io/poll`.
type OutputStream cm.Resource

// CheckWrite represents the method "wasi:io/streams.output-stream.check-write".
//
// Check readiness for writing. This function never blocks.
//
// Returns the number of bytes permitted for the next call to `write`,
// or an error. Calling `write` with more bytes than this function has
// permitted will trap.
//
// When this function returns 0 bytes, the `subscribe` pollable will
// become ready when this function will report at least 1 byte, or an
// error.
func (self OutputStream) CheckWrite() cm.OKSizedResult[uint64, StreamError] {
	var ret cm.OKSizedResult[uint64, StreamError]
	self.check_write(&ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.check-write
func (self OutputStream) check_write(ret *cm.OKSizedResult[uint64, StreamError])

// Write represents the method "wasi:io/streams.output-stream.write".
//
// Perform a write. This function never blocks.
//
// Precondition: check-write gave permit of Ok(n) and contents has a
// length of less than or equal to n. Otherwise, this function will trap.
//
// returns Err(closed) without writing if the stream has closed since
// the last call to check-write provided a permit.
func (self OutputStream) Write(contents cm.List[uint8]) cm.ErrResult[StreamError] {
	var ret cm.ErrResult[StreamError]
	self.write(contents, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.write
func (self OutputStream) write(contents cm.List[uint8], ret *cm.ErrResult[StreamError])

// BlockingWriteAndFlush represents the method "wasi:io/streams.output-stream.blocking-write-and-flush".
//
// Perform a write of up to 4096 bytes, and then flush the stream. Block
// until all of these operations are complete, or an error occurs.
//
// This is a convenience wrapper around the use of `check-write`,
// `subscribe`, `write`, and `flush`, and is implemented with the
// following pseudo-code:
//
// ```text
// let pollable = this.subscribe();
//
//	while !contents.is_empty() {
//	    // Wait for the stream to become writable
//	    pollable.block();
//	    let Ok(n) = this.check-write(); // eliding error handling
//	    let len = min(n, contents.len());
//	    let (chunk, rest) = contents.split_at(len);
//	    this.write(chunk  );            // eliding error handling
//	    contents = rest;
//	}
//
// this.flush();
// // Wait for completion of `flush`
// pollable.block();
// // Check for any errors that arose during `flush`
// let _ = this.check-write();         // eliding error handling
// ```
func (self OutputStream) BlockingWriteAndFlush(contents cm.List[uint8]) cm.ErrResult[StreamError] {
	var ret cm.ErrResult[StreamError]
	self.blocking_write_and_flush(contents, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.blocking-write-and-flush
func (self OutputStream) blocking_write_and_flush(contents cm.List[uint8], ret *cm.ErrResult[StreamError])

// Flush represents the method "wasi:io/streams.output-stream.flush".
//
// Request to flush buffered output. This function never blocks.
//
// This tells the output-stream that the caller intends any buffered
// output to be flushed. the output which is expected to be flushed
// is all that has been passed to `write` prior to this call.
//
// Upon calling this function, the `output-stream` will not accept any
// writes (`check-write` will return `ok(0)`) until the flush has
// completed. The `subscribe` pollable will become ready when the
// flush has completed and the stream can accept more writes.
func (self OutputStream) Flush() cm.ErrResult[StreamError] {
	var ret cm.ErrResult[StreamError]
	self.flush(&ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.flush
func (self OutputStream) flush(ret *cm.ErrResult[StreamError])

// BlockingFlush represents the method "wasi:io/streams.output-stream.blocking-flush".
//
// Request to flush buffered output, and block until flush completes
// and stream is ready for writing again.
func (self OutputStream) BlockingFlush() cm.ErrResult[StreamError] {
	var ret cm.ErrResult[StreamError]
	self.blocking_flush(&ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.blocking-flush
func (self OutputStream) blocking_flush(ret *cm.ErrResult[StreamError])

// Subscribe represents the method "wasi:io/streams.output-stream.subscribe".
//
// Create a `pollable` which will resolve once the output-stream
// is ready for more writing, or an error has occured. When this
// pollable is ready, `check-write` will return `ok(n)` with n>0, or an
// error.
//
// If the stream is closed, this pollable is always ready immediately.
//
// The created `pollable` is a child resource of the `output-stream`.
// Implementations may trap if the `output-stream` is dropped before
// all derived `pollable`s created with this function are dropped.
func (self OutputStream) Subscribe() Pollable {
	return self.subscribe()
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.subscribe
func (self OutputStream) subscribe() Pollable

// WriteZeroes represents the method "wasi:io/streams.output-stream.write-zeroes".
//
// Write zeroes to a stream.
//
// This should be used precisely like `write` with the exact same
// preconditions (must use check-write first), but instead of
// passing a list of bytes, you simply pass the number of zero-bytes
// that should be written.
func (self OutputStream) WriteZeroes(len uint64) cm.ErrResult[StreamError] {
	var ret cm.ErrResult[StreamError]
	self.write_zeroes(len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.write-zeroes
func (self OutputStream) write_zeroes(len uint64, ret *cm.ErrResult[StreamError])

// BlockingWriteZeroesAndFlush represents the method "wasi:io/streams.output-stream.blocking-write-zeroes-and-flush".
//
// Perform a write of up to 4096 zeroes, and then flush the stream.
// Block until all of these operations are complete, or an error
// occurs.
//
// This is a convenience wrapper around the use of `check-write`,
// `subscribe`, `write-zeroes`, and `flush`, and is implemented with
// the following pseudo-code:
//
// ```text
// let pollable = this.subscribe();
//
//	while num_zeroes != 0 {
//	    // Wait for the stream to become writable
//	    pollable.block();
//	    let Ok(n) = this.check-write(); // eliding error handling
//	    let len = min(n, num_zeroes);
//	    this.write-zeroes(len);         // eliding error handling
//	    num_zeroes -= len;
//	}
//
// this.flush();
// // Wait for completion of `flush`
// pollable.block();
// // Check for any errors that arose during `flush`
// let _ = this.check-write();         // eliding error handling
// ```
func (self OutputStream) BlockingWriteZeroesAndFlush(len uint64) cm.ErrResult[StreamError] {
	var ret cm.ErrResult[StreamError]
	self.blocking_write_zeroes_and_flush(len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.blocking-write-zeroes-and-flush
func (self OutputStream) blocking_write_zeroes_and_flush(len uint64, ret *cm.ErrResult[StreamError])

// Splice represents the method "wasi:io/streams.output-stream.splice".
//
// Read from one stream and write to another.
//
// The behavior of splice is equivelant to:
// 1. calling `check-write` on the `output-stream`
// 2. calling `read` on the `input-stream` with the smaller of the
// `check-write` permitted length and the `len` provided to `splice`
// 3. calling `write` on the `output-stream` with that read data.
//
// Any error reported by the call to `check-write`, `read`, or
// `write` ends the splice and reports that error.
//
// This function returns the number of bytes transferred; it may be less
// than `len`.
func (self OutputStream) Splice(src InputStream, len uint64) cm.OKSizedResult[uint64, StreamError] {
	var ret cm.OKSizedResult[uint64, StreamError]
	self.splice(src, len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.splice
func (self OutputStream) splice(src InputStream, len uint64, ret *cm.OKSizedResult[uint64, StreamError])

// BlockingSplice represents the method "wasi:io/streams.output-stream.blocking-splice".
//
// Read from one stream and write to another, with blocking.
//
// This is similar to `splice`, except that it blocks until the
// `output-stream` is ready for writing, and the `input-stream`
// is ready for reading, before performing the `splice`.
func (self OutputStream) BlockingSplice(src InputStream, len uint64) cm.OKSizedResult[uint64, StreamError] {
	var ret cm.OKSizedResult[uint64, StreamError]
	self.blocking_splice(src, len, &ret)
	return ret
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.blocking-splice
func (self OutputStream) blocking_splice(src InputStream, len uint64, ret *cm.OKSizedResult[uint64, StreamError])
