// Package gumble is a client for the Mumble voice chat software.
//
// Getting started
//
//1. Create a new Config to hold your connection settings:
//
//        config := gumble.NewConfig()
//        config.Username = "gumble-test"
//        config.Address = "example.com:64738"
//
//2. Create a new Client:
//
//        client := gumble.NewClient(config)
//
//3. Implement EventListener (or use gumbleutil.Listener) and attach it to the client:
//
//        client.Attach(gumbleutil.Listener{
//          TextMessage: func(e *gumble.TextMessageEvent) {
//            fmt.Printf("Received text message: %s\n", e.Message)
//          },
//        })
//
//4. Connect to the server:
//
//        if err := client.Connect(); err != nil {
//          panic(err)
//        }
//
// Audio codecs
//
// Currently, only the Opus codec (https://www.opus-codec.org/) is supported
// for transmitting and receiving audio. It can be enabled by importing the
// following package for its side effect:
//  import (
//    _ "github.com/mikolan/gumble/opus"
//  )
//
// To ensure that gumble clients can always transmit and receive audio to and
// from your server, add the following line to your murmur configuration file:
//
//  opusthreshold=0
//
// Thread safety
//
// As an easy, general rule, a connecting/connected Client, and everything that
// is associated with it (e.g. Users, Channels), is thread-unsafe. Accessing or
// modifying such a Client should only be done from inside of an event listener
// or a Client.Do function.
package gumble
